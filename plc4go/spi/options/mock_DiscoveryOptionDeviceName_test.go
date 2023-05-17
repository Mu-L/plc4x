/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.27.1. DO NOT EDIT.

package options

import mock "github.com/stretchr/testify/mock"

// MockDiscoveryOptionDeviceName is an autogenerated mock type for the DiscoveryOptionDeviceName type
type MockDiscoveryOptionDeviceName struct {
	mock.Mock
}

type MockDiscoveryOptionDeviceName_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDiscoveryOptionDeviceName) EXPECT() *MockDiscoveryOptionDeviceName_Expecter {
	return &MockDiscoveryOptionDeviceName_Expecter{mock: &_m.Mock}
}

// GetDeviceName provides a mock function with given fields:
func (_m *MockDiscoveryOptionDeviceName) GetDeviceName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDiscoveryOptionDeviceName_GetDeviceName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDeviceName'
type MockDiscoveryOptionDeviceName_GetDeviceName_Call struct {
	*mock.Call
}

// GetDeviceName is a helper method to define mock.On call
func (_e *MockDiscoveryOptionDeviceName_Expecter) GetDeviceName() *MockDiscoveryOptionDeviceName_GetDeviceName_Call {
	return &MockDiscoveryOptionDeviceName_GetDeviceName_Call{Call: _e.mock.On("GetDeviceName")}
}

func (_c *MockDiscoveryOptionDeviceName_GetDeviceName_Call) Run(run func()) *MockDiscoveryOptionDeviceName_GetDeviceName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDiscoveryOptionDeviceName_GetDeviceName_Call) Return(_a0 string) *MockDiscoveryOptionDeviceName_GetDeviceName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDiscoveryOptionDeviceName_GetDeviceName_Call) RunAndReturn(run func() string) *MockDiscoveryOptionDeviceName_GetDeviceName_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockDiscoveryOptionDeviceName interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDiscoveryOptionDeviceName creates a new instance of MockDiscoveryOptionDeviceName. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDiscoveryOptionDeviceName(t mockConstructorTestingTNewMockDiscoveryOptionDeviceName) *MockDiscoveryOptionDeviceName {
	mock := &MockDiscoveryOptionDeviceName{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
