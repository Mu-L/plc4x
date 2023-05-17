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

package interceptors

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	spi "github.com/apache/plc4x/plc4go/spi"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// mockWriteRequestFactory is an autogenerated mock type for the writeRequestFactory type
type mockWriteRequestFactory struct {
	mock.Mock
}

type mockWriteRequestFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *mockWriteRequestFactory) EXPECT() *mockWriteRequestFactory_Expecter {
	return &mockWriteRequestFactory_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: tags, tagNames, _a2, writer, writeRequestInterceptor
func (_m *mockWriteRequestFactory) Execute(tags map[string]model.PlcTag, tagNames []string, _a2 map[string]values.PlcValue, writer spi.PlcWriter, writeRequestInterceptor WriteRequestInterceptor) model.PlcWriteRequest {
	ret := _m.Called(tags, tagNames, _a2, writer, writeRequestInterceptor)

	var r0 model.PlcWriteRequest
	if rf, ok := ret.Get(0).(func(map[string]model.PlcTag, []string, map[string]values.PlcValue, spi.PlcWriter, WriteRequestInterceptor) model.PlcWriteRequest); ok {
		r0 = rf(tags, tagNames, _a2, writer, writeRequestInterceptor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcWriteRequest)
		}
	}

	return r0
}

// mockWriteRequestFactory_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type mockWriteRequestFactory_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - tags map[string]model.PlcTag
//   - tagNames []string
//   - _a2 map[string]values.PlcValue
//   - writer spi.PlcWriter
//   - writeRequestInterceptor WriteRequestInterceptor
func (_e *mockWriteRequestFactory_Expecter) Execute(tags interface{}, tagNames interface{}, _a2 interface{}, writer interface{}, writeRequestInterceptor interface{}) *mockWriteRequestFactory_Execute_Call {
	return &mockWriteRequestFactory_Execute_Call{Call: _e.mock.On("Execute", tags, tagNames, _a2, writer, writeRequestInterceptor)}
}

func (_c *mockWriteRequestFactory_Execute_Call) Run(run func(tags map[string]model.PlcTag, tagNames []string, _a2 map[string]values.PlcValue, writer spi.PlcWriter, writeRequestInterceptor WriteRequestInterceptor)) *mockWriteRequestFactory_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]model.PlcTag), args[1].([]string), args[2].(map[string]values.PlcValue), args[3].(spi.PlcWriter), args[4].(WriteRequestInterceptor))
	})
	return _c
}

func (_c *mockWriteRequestFactory_Execute_Call) Return(_a0 model.PlcWriteRequest) *mockWriteRequestFactory_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockWriteRequestFactory_Execute_Call) RunAndReturn(run func(map[string]model.PlcTag, []string, map[string]values.PlcValue, spi.PlcWriter, WriteRequestInterceptor) model.PlcWriteRequest) *mockWriteRequestFactory_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTnewMockWriteRequestFactory interface {
	mock.TestingT
	Cleanup(func())
}

// newMockWriteRequestFactory creates a new instance of mockWriteRequestFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockWriteRequestFactory(t mockConstructorTestingTnewMockWriteRequestFactory) *mockWriteRequestFactory {
	mock := &mockWriteRequestFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
