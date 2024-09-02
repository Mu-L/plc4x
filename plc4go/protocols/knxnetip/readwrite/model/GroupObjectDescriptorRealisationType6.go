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

// GroupObjectDescriptorRealisationType6 is the corresponding interface of GroupObjectDescriptorRealisationType6
type GroupObjectDescriptorRealisationType6 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsGroupObjectDescriptorRealisationType6 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGroupObjectDescriptorRealisationType6()
}

// _GroupObjectDescriptorRealisationType6 is the data-structure of this message
type _GroupObjectDescriptorRealisationType6 struct {
}

var _ GroupObjectDescriptorRealisationType6 = (*_GroupObjectDescriptorRealisationType6)(nil)

// NewGroupObjectDescriptorRealisationType6 factory function for _GroupObjectDescriptorRealisationType6
func NewGroupObjectDescriptorRealisationType6() *_GroupObjectDescriptorRealisationType6 {
	return &_GroupObjectDescriptorRealisationType6{}
}

// Deprecated: use the interface for direct cast
func CastGroupObjectDescriptorRealisationType6(structType any) GroupObjectDescriptorRealisationType6 {
	if casted, ok := structType.(GroupObjectDescriptorRealisationType6); ok {
		return casted
	}
	if casted, ok := structType.(*GroupObjectDescriptorRealisationType6); ok {
		return *casted
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType6) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType6"
}

func (m *_GroupObjectDescriptorRealisationType6) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_GroupObjectDescriptorRealisationType6) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GroupObjectDescriptorRealisationType6Parse(ctx context.Context, theBytes []byte) (GroupObjectDescriptorRealisationType6, error) {
	return GroupObjectDescriptorRealisationType6ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func GroupObjectDescriptorRealisationType6ParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType6, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType6, error) {
		return GroupObjectDescriptorRealisationType6ParseWithBuffer(ctx, readBuffer)
	}
}

func GroupObjectDescriptorRealisationType6ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType6, error) {
	v, err := (&_GroupObjectDescriptorRealisationType6{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_GroupObjectDescriptorRealisationType6) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__groupObjectDescriptorRealisationType6 GroupObjectDescriptorRealisationType6, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationType6"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GroupObjectDescriptorRealisationType6")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationType6"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GroupObjectDescriptorRealisationType6")
	}

	return m, nil
}

func (m *_GroupObjectDescriptorRealisationType6) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GroupObjectDescriptorRealisationType6) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationType6"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GroupObjectDescriptorRealisationType6")
	}

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationType6"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GroupObjectDescriptorRealisationType6")
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType6) IsGroupObjectDescriptorRealisationType6() {}

func (m *_GroupObjectDescriptorRealisationType6) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
