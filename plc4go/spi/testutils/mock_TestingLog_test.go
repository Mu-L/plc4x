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

// Code generated by mockery v2.42.2. DO NOT EDIT.

package testutils

import mock "github.com/stretchr/testify/mock"

// MockTestingLog is an autogenerated mock type for the TestingLog type
type MockTestingLog struct {
	mock.Mock
}

type MockTestingLog_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTestingLog) EXPECT() *MockTestingLog_Expecter {
	return &MockTestingLog_Expecter{mock: &_m.Mock}
}

// Helper provides a mock function with given fields:
func (_m *MockTestingLog) Helper() {
	_m.Called()
}

// MockTestingLog_Helper_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Helper'
type MockTestingLog_Helper_Call struct {
	*mock.Call
}

// Helper is a helper method to define mock.On call
func (_e *MockTestingLog_Expecter) Helper() *MockTestingLog_Helper_Call {
	return &MockTestingLog_Helper_Call{Call: _e.mock.On("Helper")}
}

func (_c *MockTestingLog_Helper_Call) Run(run func()) *MockTestingLog_Helper_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestingLog_Helper_Call) Return() *MockTestingLog_Helper_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTestingLog_Helper_Call) RunAndReturn(run func()) *MockTestingLog_Helper_Call {
	_c.Call.Return(run)
	return _c
}

// Log provides a mock function with given fields: args
func (_m *MockTestingLog) Log(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockTestingLog_Log_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Log'
type MockTestingLog_Log_Call struct {
	*mock.Call
}

// Log is a helper method to define mock.On call
//   - args ...interface{}
func (_e *MockTestingLog_Expecter) Log(args ...interface{}) *MockTestingLog_Log_Call {
	return &MockTestingLog_Log_Call{Call: _e.mock.On("Log",
		append([]interface{}{}, args...)...)}
}

func (_c *MockTestingLog_Log_Call) Run(run func(args ...interface{})) *MockTestingLog_Log_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockTestingLog_Log_Call) Return() *MockTestingLog_Log_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTestingLog_Log_Call) RunAndReturn(run func(...interface{})) *MockTestingLog_Log_Call {
	_c.Call.Return(run)
	return _c
}

// Logf provides a mock function with given fields: format, args
func (_m *MockTestingLog) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// MockTestingLog_Logf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Logf'
type MockTestingLog_Logf_Call struct {
	*mock.Call
}

// Logf is a helper method to define mock.On call
//   - format string
//   - args ...interface{}
func (_e *MockTestingLog_Expecter) Logf(format interface{}, args ...interface{}) *MockTestingLog_Logf_Call {
	return &MockTestingLog_Logf_Call{Call: _e.mock.On("Logf",
		append([]interface{}{format}, args...)...)}
}

func (_c *MockTestingLog_Logf_Call) Run(run func(format string, args ...interface{})) *MockTestingLog_Logf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockTestingLog_Logf_Call) Return() *MockTestingLog_Logf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTestingLog_Logf_Call) RunAndReturn(run func(string, ...interface{})) *MockTestingLog_Logf_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTestingLog creates a new instance of MockTestingLog. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTestingLog(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTestingLog {
	mock := &MockTestingLog{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}