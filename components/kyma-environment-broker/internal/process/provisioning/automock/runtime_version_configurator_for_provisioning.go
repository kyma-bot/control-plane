// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	internal "github.com/kyma-project/control-plane/components/kyma-environment-broker/internal"
	mock "github.com/stretchr/testify/mock"
)

// RuntimeVersionConfiguratorForProvisioning is an autogenerated mock type for the RuntimeVersionConfiguratorForProvisioning type
type RuntimeVersionConfiguratorForProvisioning struct {
	mock.Mock
}

// ForProvisioning provides a mock function with given fields: op
func (_m *RuntimeVersionConfiguratorForProvisioning) ForProvisioning(op internal.ProvisioningOperation) (*internal.RuntimeVersionData, error) {
	ret := _m.Called(op)

	var r0 *internal.RuntimeVersionData
	if rf, ok := ret.Get(0).(func(internal.ProvisioningOperation) *internal.RuntimeVersionData); ok {
		r0 = rf(op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.RuntimeVersionData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.ProvisioningOperation) error); ok {
		r1 = rf(op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
