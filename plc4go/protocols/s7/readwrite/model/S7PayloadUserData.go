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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// S7PayloadUserData is the corresponding interface of S7PayloadUserData
type S7PayloadUserData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Payload
	// GetItems returns Items (property field)
	GetItems() []S7PayloadUserDataItem
}

// S7PayloadUserDataExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserData.
// This is useful for switch cases.
type S7PayloadUserDataExactly interface {
	S7PayloadUserData
	isS7PayloadUserData() bool
}

// _S7PayloadUserData is the data-structure of this message
type _S7PayloadUserData struct {
	*_S7Payload
	Items []S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserData) GetParameterParameterType() uint8 {
	return 0x00
}

func (m *_S7PayloadUserData) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserData) InitializeParent(parent S7Payload) {}

func (m *_S7PayloadUserData) GetParent() S7Payload {
	return m._S7Payload
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserData) GetItems() []S7PayloadUserDataItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserData factory function for _S7PayloadUserData
func NewS7PayloadUserData(items []S7PayloadUserDataItem, parameter S7Parameter) *_S7PayloadUserData {
	_result := &_S7PayloadUserData{
		Items:      items,
		_S7Payload: NewS7Payload(parameter),
	}
	_result._S7Payload._S7PayloadChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserData(structType any) S7PayloadUserData {
	if casted, ok := structType.(S7PayloadUserData); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserData); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserData) GetTypeName() string {
	return "S7PayloadUserData"
}

func (m *_S7PayloadUserData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadUserData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataParse(ctx context.Context, theBytes []byte, messageType uint8, parameter S7Parameter) (S7PayloadUserData, error) {
	return S7PayloadUserDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), messageType, parameter)
}

func S7PayloadUserDataParseWithBufferProducer(messageType uint8, parameter S7Parameter) func(ctx context.Context, readBuffer utils.ReadBuffer) (S7PayloadUserData, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7PayloadUserData, error) {
		return S7PayloadUserDataParseWithBuffer(ctx, readBuffer, messageType, parameter)
	}
}

func S7PayloadUserDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8, parameter S7Parameter) (S7PayloadUserData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadUserData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	items, err := ReadCountArrayField[S7PayloadUserDataItem](ctx, "items", ReadComplex[S7PayloadUserDataItem](S7PayloadUserDataItemParseWithBufferProducer[S7PayloadUserDataItem]((uint8)(CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).GetItems()[0]).GetCpuFunctionGroup()), (uint8)(CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).GetItems()[0]).GetCpuFunctionType()), (uint8)(CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).GetItems()[0]).GetCpuSubfunction())), readBuffer), uint64(int32(len(CastS7ParameterUserData(parameter).GetItems()))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserData")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserData{
		_S7Payload: &_S7Payload{
			Parameter: parameter,
		},
		Items: items,
	}
	_child._S7Payload._S7PayloadChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserData")
		}

		// Array Field (items)
		if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for items")
		}
		for _curItem, _element := range m.GetItems() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetItems()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'items' field")
			}
		}
		if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for items")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserData")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserData) isS7PayloadUserData() bool {
	return true
}

func (m *_S7PayloadUserData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
