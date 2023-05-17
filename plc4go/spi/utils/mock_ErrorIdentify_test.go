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

package utils

import mock "github.com/stretchr/testify/mock"

// MockErrorIdentify is an autogenerated mock type for the ErrorIdentify type
type MockErrorIdentify struct {
	mock.Mock
}

type MockErrorIdentify_Expecter struct {
	mock *mock.Mock
}

func (_m *MockErrorIdentify) EXPECT() *MockErrorIdentify_Expecter {
	return &MockErrorIdentify_Expecter{mock: &_m.Mock}
}

// Is provides a mock function with given fields: target
func (_m *MockErrorIdentify) Is(target error) bool {
	ret := _m.Called(target)

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(target)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockErrorIdentify_Is_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Is'
type MockErrorIdentify_Is_Call struct {
	*mock.Call
}

// Is is a helper method to define mock.On call
//   - target error
func (_e *MockErrorIdentify_Expecter) Is(target interface{}) *MockErrorIdentify_Is_Call {
	return &MockErrorIdentify_Is_Call{Call: _e.mock.On("Is", target)}
}

func (_c *MockErrorIdentify_Is_Call) Run(run func(target error)) *MockErrorIdentify_Is_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockErrorIdentify_Is_Call) Return(_a0 bool) *MockErrorIdentify_Is_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockErrorIdentify_Is_Call) RunAndReturn(run func(error) bool) *MockErrorIdentify_Is_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockErrorIdentify interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockErrorIdentify creates a new instance of MockErrorIdentify. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockErrorIdentify(t mockConstructorTestingTNewMockErrorIdentify) *MockErrorIdentify {
	mock := &MockErrorIdentify{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
