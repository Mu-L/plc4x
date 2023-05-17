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

package testutils

import mock "github.com/stretchr/testify/mock"

// MockXmlParser is an autogenerated mock type for the XmlParser type
type MockXmlParser struct {
	mock.Mock
}

type MockXmlParser_Expecter struct {
	mock *mock.Mock
}

func (_m *MockXmlParser) EXPECT() *MockXmlParser_Expecter {
	return &MockXmlParser_Expecter{mock: &_m.Mock}
}

// Parse provides a mock function with given fields: typeName, xmlString, parserArguments
func (_m *MockXmlParser) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
	_va := make([]interface{}, len(parserArguments))
	for _i := range parserArguments {
		_va[_i] = parserArguments[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, typeName, xmlString)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, ...string) (interface{}, error)); ok {
		return rf(typeName, xmlString, parserArguments...)
	}
	if rf, ok := ret.Get(0).(func(string, string, ...string) interface{}); ok {
		r0 = rf(typeName, xmlString, parserArguments...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, ...string) error); ok {
		r1 = rf(typeName, xmlString, parserArguments...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockXmlParser_Parse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Parse'
type MockXmlParser_Parse_Call struct {
	*mock.Call
}

// Parse is a helper method to define mock.On call
//   - typeName string
//   - xmlString string
//   - parserArguments ...string
func (_e *MockXmlParser_Expecter) Parse(typeName interface{}, xmlString interface{}, parserArguments ...interface{}) *MockXmlParser_Parse_Call {
	return &MockXmlParser_Parse_Call{Call: _e.mock.On("Parse",
		append([]interface{}{typeName, xmlString}, parserArguments...)...)}
}

func (_c *MockXmlParser_Parse_Call) Run(run func(typeName string, xmlString string, parserArguments ...string)) *MockXmlParser_Parse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(string), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockXmlParser_Parse_Call) Return(_a0 interface{}, _a1 error) *MockXmlParser_Parse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockXmlParser_Parse_Call) RunAndReturn(run func(string, string, ...string) (interface{}, error)) *MockXmlParser_Parse_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockXmlParser interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockXmlParser creates a new instance of MockXmlParser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockXmlParser(t mockConstructorTestingTNewMockXmlParser) *MockXmlParser {
	mock := &MockXmlParser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
