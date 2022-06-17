// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	intstr "k8s.io/apimachinery/pkg/util/intstr"

	v1 "k8s.io/api/core/v1"

	nexusApi "github.com/epam/edp-nexus-operator/v2/pkg/apis/edp/v1"

	keycloakApi "github.com/epam/edp-keycloak-operator/pkg/apis/v1/v1"
)

// PlatformService is an autogenerated mock type for the PlatformService type
type PlatformService struct {
	mock.Mock
}

// AddKeycloakProxyToDeployConf provides a mock function with given fields: instance, args
func (_m *PlatformService) AddKeycloakProxyToDeployConf(instance nexusApi.Nexus, args []string) error {
	ret := _m.Called(instance, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(nexusApi.Nexus, []string) error); ok {
		r0 = rf(instance, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddPortToService provides a mock function with given fields: instance, newPortSpec
func (_m *PlatformService) AddPortToService(instance nexusApi.Nexus, newPortSpec v1.ServicePort) error {
	ret := _m.Called(instance, newPortSpec)

	var r0 error
	if rf, ok := ret.Get(0).(func(nexusApi.Nexus, v1.ServicePort) error); ok {
		r0 = rf(instance, newPortSpec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateConfigMapFromFile provides a mock function with given fields: instance, configMapName, filePath
func (_m *PlatformService) CreateConfigMapFromFile(instance nexusApi.Nexus, configMapName string, filePath string) error {
	ret := _m.Called(instance, configMapName, filePath)

	var r0 error
	if rf, ok := ret.Get(0).(func(nexusApi.Nexus, string, string) error); ok {
		r0 = rf(instance, configMapName, filePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateEDPComponentIfNotExist provides a mock function with given fields: instance, url, icon
func (_m *PlatformService) CreateEDPComponentIfNotExist(instance nexusApi.Nexus, url string, icon string) error {
	ret := _m.Called(instance, url, icon)

	var r0 error
	if rf, ok := ret.Get(0).(func(nexusApi.Nexus, string, string) error); ok {
		r0 = rf(instance, url, icon)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateJenkinsServiceAccount provides a mock function with given fields: namespace, secretName
func (_m *PlatformService) CreateJenkinsServiceAccount(namespace string, secretName string) error {
	ret := _m.Called(namespace, secretName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(namespace, secretName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateKeycloakClient provides a mock function with given fields: kc
func (_m *PlatformService) CreateKeycloakClient(kc *keycloakApi.KeycloakClient) error {
	ret := _m.Called(kc)

	var r0 error
	if rf, ok := ret.Get(0).(func(*keycloakApi.KeycloakClient) error); ok {
		r0 = rf(kc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSecret provides a mock function with given fields: instance, name, data
func (_m *PlatformService) CreateSecret(instance nexusApi.Nexus, name string, data map[string][]byte) error {
	ret := _m.Called(instance, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(nexusApi.Nexus, string, map[string][]byte) error); ok {
		r0 = rf(instance, name, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetConfigMapData provides a mock function with given fields: namespace, name
func (_m *PlatformService) GetConfigMapData(namespace string, name string) (map[string]string, error) {
	ret := _m.Called(namespace, name)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, string) map[string]string); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExternalUrl provides a mock function with given fields: namespace, name
func (_m *PlatformService) GetExternalUrl(namespace string, name string) (string, string, string, error) {
	ret := _m.Called(namespace, name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(namespace, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(string, string) string); ok {
		r2 = rf(namespace, name)
	} else {
		r2 = ret.Get(2).(string)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string, string) error); ok {
		r3 = rf(namespace, name)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetKeycloakClient provides a mock function with given fields: name, namespace
func (_m *PlatformService) GetKeycloakClient(name string, namespace string) (keycloakApi.KeycloakClient, error) {
	ret := _m.Called(name, namespace)

	var r0 keycloakApi.KeycloakClient
	if rf, ok := ret.Get(0).(func(string, string) keycloakApi.KeycloakClient); ok {
		r0 = rf(name, namespace)
	} else {
		r0 = ret.Get(0).(keycloakApi.KeycloakClient)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecret provides a mock function with given fields: namespace, name
func (_m *PlatformService) GetSecret(namespace string, name string) (*v1.Secret, error) {
	ret := _m.Called(namespace, name)

	var r0 *v1.Secret
	if rf, ok := ret.Get(0).(func(string, string) *v1.Secret); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Secret)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecretData provides a mock function with given fields: namespace, name
func (_m *PlatformService) GetSecretData(namespace string, name string) (map[string][]byte, error) {
	ret := _m.Called(namespace, name)

	var r0 map[string][]byte
	if rf, ok := ret.Get(0).(func(string, string) map[string][]byte); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceByCr provides a mock function with given fields: name, namespace
func (_m *PlatformService) GetServiceByCr(name string, namespace string) (*v1.Service, error) {
	ret := _m.Called(name, namespace)

	var r0 *v1.Service
	if rf, ok := ret.Get(0).(func(string, string) *v1.Service); ok {
		r0 = rf(name, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsDeploymentReady provides a mock function with given fields: instance
func (_m *PlatformService) IsDeploymentReady(instance nexusApi.Nexus) (*bool, error) {
	ret := _m.Called(instance)

	var r0 *bool
	if rf, ok := ret.Get(0).(func(nexusApi.Nexus) *bool); ok {
		r0 = rf(instance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(nexusApi.Nexus) error); ok {
		r1 = rf(instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExternalTargetPath provides a mock function with given fields: instance, targetPort
func (_m *PlatformService) UpdateExternalTargetPath(instance nexusApi.Nexus, targetPort intstr.IntOrString) error {
	ret := _m.Called(instance, targetPort)

	var r0 error
	if rf, ok := ret.Get(0).(func(nexusApi.Nexus, intstr.IntOrString) error); ok {
		r0 = rf(instance, targetPort)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSecret provides a mock function with given fields: secret
func (_m *PlatformService) UpdateSecret(secret *v1.Secret) error {
	ret := _m.Called(secret)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.Secret) error); ok {
		r0 = rf(secret)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
