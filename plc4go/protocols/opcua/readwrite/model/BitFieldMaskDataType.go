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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BitFieldMaskDataType is the corresponding interface of BitFieldMaskDataType
type BitFieldMaskDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BitFieldMaskDataTypeExactly can be used when we want exactly this type and not a type which fulfills BitFieldMaskDataType.
// This is useful for switch cases.
type BitFieldMaskDataTypeExactly interface {
	BitFieldMaskDataType
	isBitFieldMaskDataType() bool
}

// _BitFieldMaskDataType is the data-structure of this message
type _BitFieldMaskDataType struct {
}

// NewBitFieldMaskDataType factory function for _BitFieldMaskDataType
func NewBitFieldMaskDataType() *_BitFieldMaskDataType {
	return &_BitFieldMaskDataType{}
}

// Deprecated: use the interface for direct cast
func CastBitFieldMaskDataType(structType any) BitFieldMaskDataType {
	if casted, ok := structType.(BitFieldMaskDataType); ok {
		return casted
	}
	if casted, ok := structType.(*BitFieldMaskDataType); ok {
		return *casted
	}
	return nil
}

func (m *_BitFieldMaskDataType) GetTypeName() string {
	return "BitFieldMaskDataType"
}

func (m *_BitFieldMaskDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_BitFieldMaskDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BitFieldMaskDataTypeParse(ctx context.Context, theBytes []byte) (BitFieldMaskDataType, error) {
	return BitFieldMaskDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BitFieldMaskDataTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BitFieldMaskDataType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BitFieldMaskDataType, error) {
		return BitFieldMaskDataTypeParseWithBuffer(ctx, readBuffer)
	}
}

func BitFieldMaskDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BitFieldMaskDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BitFieldMaskDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BitFieldMaskDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BitFieldMaskDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BitFieldMaskDataType")
	}

	// Create the instance
	return &_BitFieldMaskDataType{}, nil
}

func (m *_BitFieldMaskDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BitFieldMaskDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BitFieldMaskDataType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BitFieldMaskDataType")
	}

	if popErr := writeBuffer.PopContext("BitFieldMaskDataType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BitFieldMaskDataType")
	}
	return nil
}

func (m *_BitFieldMaskDataType) isBitFieldMaskDataType() bool {
	return true
}

func (m *_BitFieldMaskDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
