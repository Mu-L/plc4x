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

package transports

import (
	bufio "bufio"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockDefaultBufferedTransportInstance is an autogenerated mock type for the DefaultBufferedTransportInstance type
type MockDefaultBufferedTransportInstance struct {
	mock.Mock
}

type MockDefaultBufferedTransportInstance_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDefaultBufferedTransportInstance) EXPECT() *MockDefaultBufferedTransportInstance_Expecter {
	return &MockDefaultBufferedTransportInstance_Expecter{mock: &_m.Mock}
}

// ConnectWithContext provides a mock function with given fields: ctx
func (_m *MockDefaultBufferedTransportInstance) ConnectWithContext(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultBufferedTransportInstance_ConnectWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectWithContext'
type MockDefaultBufferedTransportInstance_ConnectWithContext_Call struct {
	*mock.Call
}

// ConnectWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDefaultBufferedTransportInstance_Expecter) ConnectWithContext(ctx interface{}) *MockDefaultBufferedTransportInstance_ConnectWithContext_Call {
	return &MockDefaultBufferedTransportInstance_ConnectWithContext_Call{Call: _e.mock.On("ConnectWithContext", ctx)}
}

func (_c *MockDefaultBufferedTransportInstance_ConnectWithContext_Call) Run(run func(ctx context.Context)) *MockDefaultBufferedTransportInstance_ConnectWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_ConnectWithContext_Call) Return(_a0 error) *MockDefaultBufferedTransportInstance_ConnectWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_ConnectWithContext_Call) RunAndReturn(run func(context.Context) error) *MockDefaultBufferedTransportInstance_ConnectWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// FillBuffer provides a mock function with given fields: until
func (_m *MockDefaultBufferedTransportInstance) FillBuffer(until func(uint, byte, *bufio.Reader) bool) error {
	ret := _m.Called(until)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(uint, byte, *bufio.Reader) bool) error); ok {
		r0 = rf(until)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDefaultBufferedTransportInstance_FillBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FillBuffer'
type MockDefaultBufferedTransportInstance_FillBuffer_Call struct {
	*mock.Call
}

// FillBuffer is a helper method to define mock.On call
//   - until func(uint , byte , *bufio.Reader) bool
func (_e *MockDefaultBufferedTransportInstance_Expecter) FillBuffer(until interface{}) *MockDefaultBufferedTransportInstance_FillBuffer_Call {
	return &MockDefaultBufferedTransportInstance_FillBuffer_Call{Call: _e.mock.On("FillBuffer", until)}
}

func (_c *MockDefaultBufferedTransportInstance_FillBuffer_Call) Run(run func(until func(uint, byte, *bufio.Reader) bool)) *MockDefaultBufferedTransportInstance_FillBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(uint, byte, *bufio.Reader) bool))
	})
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_FillBuffer_Call) Return(_a0 error) *MockDefaultBufferedTransportInstance_FillBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_FillBuffer_Call) RunAndReturn(run func(func(uint, byte, *bufio.Reader) bool) error) *MockDefaultBufferedTransportInstance_FillBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// GetNumBytesAvailableInBuffer provides a mock function with given fields:
func (_m *MockDefaultBufferedTransportInstance) GetNumBytesAvailableInBuffer() (uint32, error) {
	ret := _m.Called()

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint32, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNumBytesAvailableInBuffer'
type MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call struct {
	*mock.Call
}

// GetNumBytesAvailableInBuffer is a helper method to define mock.On call
func (_e *MockDefaultBufferedTransportInstance_Expecter) GetNumBytesAvailableInBuffer() *MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call {
	return &MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call{Call: _e.mock.On("GetNumBytesAvailableInBuffer")}
}

func (_c *MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call) Run(run func()) *MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call) Return(_a0 uint32, _a1 error) *MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call) RunAndReturn(run func() (uint32, error)) *MockDefaultBufferedTransportInstance_GetNumBytesAvailableInBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// PeekReadableBytes provides a mock function with given fields: numBytes
func (_m *MockDefaultBufferedTransportInstance) PeekReadableBytes(numBytes uint32) ([]byte, error) {
	ret := _m.Called(numBytes)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32) ([]byte, error)); ok {
		return rf(numBytes)
	}
	if rf, ok := ret.Get(0).(func(uint32) []byte); ok {
		r0 = rf(numBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(numBytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDefaultBufferedTransportInstance_PeekReadableBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PeekReadableBytes'
type MockDefaultBufferedTransportInstance_PeekReadableBytes_Call struct {
	*mock.Call
}

// PeekReadableBytes is a helper method to define mock.On call
//   - numBytes uint32
func (_e *MockDefaultBufferedTransportInstance_Expecter) PeekReadableBytes(numBytes interface{}) *MockDefaultBufferedTransportInstance_PeekReadableBytes_Call {
	return &MockDefaultBufferedTransportInstance_PeekReadableBytes_Call{Call: _e.mock.On("PeekReadableBytes", numBytes)}
}

func (_c *MockDefaultBufferedTransportInstance_PeekReadableBytes_Call) Run(run func(numBytes uint32)) *MockDefaultBufferedTransportInstance_PeekReadableBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_PeekReadableBytes_Call) Return(_a0 []byte, _a1 error) *MockDefaultBufferedTransportInstance_PeekReadableBytes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_PeekReadableBytes_Call) RunAndReturn(run func(uint32) ([]byte, error)) *MockDefaultBufferedTransportInstance_PeekReadableBytes_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: numBytes
func (_m *MockDefaultBufferedTransportInstance) Read(numBytes uint32) ([]byte, error) {
	ret := _m.Called(numBytes)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32) ([]byte, error)); ok {
		return rf(numBytes)
	}
	if rf, ok := ret.Get(0).(func(uint32) []byte); ok {
		r0 = rf(numBytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(numBytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDefaultBufferedTransportInstance_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockDefaultBufferedTransportInstance_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - numBytes uint32
func (_e *MockDefaultBufferedTransportInstance_Expecter) Read(numBytes interface{}) *MockDefaultBufferedTransportInstance_Read_Call {
	return &MockDefaultBufferedTransportInstance_Read_Call{Call: _e.mock.On("Read", numBytes)}
}

func (_c *MockDefaultBufferedTransportInstance_Read_Call) Run(run func(numBytes uint32)) *MockDefaultBufferedTransportInstance_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint32))
	})
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_Read_Call) Return(_a0 []byte, _a1 error) *MockDefaultBufferedTransportInstance_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDefaultBufferedTransportInstance_Read_Call) RunAndReturn(run func(uint32) ([]byte, error)) *MockDefaultBufferedTransportInstance_Read_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockDefaultBufferedTransportInstance interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDefaultBufferedTransportInstance creates a new instance of MockDefaultBufferedTransportInstance. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDefaultBufferedTransportInstance(t mockConstructorTestingTNewMockDefaultBufferedTransportInstance) *MockDefaultBufferedTransportInstance {
	mock := &MockDefaultBufferedTransportInstance{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
