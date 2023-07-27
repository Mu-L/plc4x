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

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// Counter is the corresponding interface of Counter
type Counter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// CounterExactly can be used when we want exactly this type and not a type which fulfills Counter.
// This is useful for switch cases.
type CounterExactly interface {
	Counter
	isCounter() bool
}

// _Counter is the data-structure of this message
type _Counter struct {
}

// NewCounter factory function for _Counter
func NewCounter() *_Counter {
	return &_Counter{}
}

// Deprecated: use the interface for direct cast
func CastCounter(structType any) Counter {
	if casted, ok := structType.(Counter); ok {
		return casted
	}
	if casted, ok := structType.(*Counter); ok {
		return *casted
	}
	return nil
}

func (m *_Counter) GetTypeName() string {
	return "Counter"
}

func (m *_Counter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_Counter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CounterParse(ctx context.Context, theBytes []byte) (Counter, error) {
	return CounterParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CounterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Counter, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("Counter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Counter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Counter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Counter")
	}

	// Create the instance
	return &_Counter{}, nil
}

func (m *_Counter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Counter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Counter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Counter")
	}

	if popErr := writeBuffer.PopContext("Counter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Counter")
	}
	return nil
}

func (m *_Counter) isCounter() bool {
	return true
}

func (m *_Counter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
