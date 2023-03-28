package main

import (
	"flag"
	"os"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	//+kubebuilder:scaffold:imports
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	buildInfo "github.com/epam/edp-common/pkg/config"
	edpCompApi "github.com/epam/edp-component-operator/api/v1"
	jenkinsApi "github.com/epam/edp-jenkins-operator/v2/pkg/apis/v2/v1"
	keycloakApi "github.com/epam/edp-keycloak-operator/api/v1"
	nexusApiV1 "github.com/epam/edp-nexus-operator/v2/api/v1"
	nexusApiV1Alpha1 "github.com/epam/edp-nexus-operator/v2/api/v1alpha1"
	"github.com/epam/edp-nexus-operator/v2/controllers/helper"
	"github.com/epam/edp-nexus-operator/v2/controllers/nexus"
	"github.com/epam/edp-nexus-operator/v2/controllers/user"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

const nexusOperatorLock = "edp-nexus-operator-lock"

func main() {
	var (
		metricsAddr          string
		enableLeaderElection bool
		probeAddr            string
	)

	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", helper.RunningInCluster(),
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")

	mode, err := helper.GetDebugMode()
	if err != nil {
		setupLog.Error(err, "unable to get debug mode value")
		os.Exit(1)
	}

	opts := zap.Options{
		Development: mode,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(nexusApiV1Alpha1.AddToScheme(scheme))
	utilruntime.Must(nexusApiV1.AddToScheme(scheme))
	utilruntime.Must(edpCompApi.AddToScheme(scheme))
	utilruntime.Must(jenkinsApi.AddToScheme(scheme))
	utilruntime.Must(keycloakApi.AddToScheme(scheme))

	v := buildInfo.Get()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	setupLog.Info("Starting the Nexus Operator",
		"version", v.Version,
		"git-commit", v.GitCommit,
		"git-tag", v.GitTag,
		"build-date", v.BuildDate,
		"go-version", v.Go,
		"go-client", v.KubectlVersion,
		"platform", v.Platform,
	)

	ns, err := helper.GetWatchNamespace()
	if err != nil {
		setupLog.Error(err, "unable to get watch namespace")
		os.Exit(1)
	}

	cfg := ctrl.GetConfigOrDie()

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		HealthProbeBindAddress: probeAddr,
		Port:                   9443,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       nexusOperatorLock,
		MapperProvider: func(c *rest.Config) (meta.RESTMapper, error) {
			return apiutil.NewDynamicRESTMapper(cfg)
		},
		Namespace: ns,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	platformType := helper.GetPlatformTypeEnv()

	nexusCtrl, err := nexus.NewReconcileNexus(mgr.GetClient(), mgr.GetScheme(), setupLog.WithName("nexus-ctrl"))
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "nexus")
		os.Exit(1)
	}

	if err = nexusCtrl.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "nexus")
		os.Exit(1)
	}

	userCtrl, err := user.NewReconcile(mgr.GetClient(), mgr.GetScheme(), setupLog.WithName("user-ctrl"), platformType)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "user")
		os.Exit(1)
	}

	if err = userCtrl.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "user")
		os.Exit(1)
	}

	if err = mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}

	if err = mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")

	if err = mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}