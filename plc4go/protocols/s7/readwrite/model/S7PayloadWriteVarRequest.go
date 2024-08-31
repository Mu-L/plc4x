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

// S7PayloadWriteVarRequest is the corresponding interface of S7PayloadWriteVarRequest
type S7PayloadWriteVarRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7Payload
	// GetItems returns Items (property field)
	GetItems() []S7VarPayloadDataItem
}

// S7PayloadWriteVarRequestExactly can be used when we want exactly this type and not a type which fulfills S7PayloadWriteVarRequest.
// This is useful for switch cases.
type S7PayloadWriteVarRequestExactly interface {
	S7PayloadWriteVarRequest
	isS7PayloadWriteVarRequest() bool
}

// _S7PayloadWriteVarRequest is the data-structure of this message
type _S7PayloadWriteVarRequest struct {
	*_S7Payload
	Items []S7VarPayloadDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadWriteVarRequest) GetParameterParameterType() uint8 {
	return 0x05
}

func (m *_S7PayloadWriteVarRequest) GetMessageType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadWriteVarRequest) InitializeParent(parent S7Payload) {}

func (m *_S7PayloadWriteVarRequest) GetParent() S7Payload {
	return m._S7Payload
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadWriteVarRequest) GetItems() []S7VarPayloadDataItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadWriteVarRequest factory function for _S7PayloadWriteVarRequest
func NewS7PayloadWriteVarRequest(items []S7VarPayloadDataItem, parameter S7Parameter) *_S7PayloadWriteVarRequest {
	_result := &_S7PayloadWriteVarRequest{
		Items:      items,
		_S7Payload: NewS7Payload(parameter),
	}
	_result._S7Payload._S7PayloadChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadWriteVarRequest(structType any) S7PayloadWriteVarRequest {
	if casted, ok := structType.(S7PayloadWriteVarRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadWriteVarRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadWriteVarRequest) GetTypeName() string {
	return "S7PayloadWriteVarRequest"
}

func (m *_S7PayloadWriteVarRequest) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_S7PayloadWriteVarRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadWriteVarRequestParse(ctx context.Context, theBytes []byte, messageType uint8, parameter S7Parameter) (S7PayloadWriteVarRequest, error) {
	return S7PayloadWriteVarRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), messageType, parameter)
}

func S7PayloadWriteVarRequestParseWithBufferProducer(messageType uint8, parameter S7Parameter) func(ctx context.Context, readBuffer utils.ReadBuffer) (S7PayloadWriteVarRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (S7PayloadWriteVarRequest, error) {
		return S7PayloadWriteVarRequestParseWithBuffer(ctx, readBuffer, messageType, parameter)
	}
}

func S7PayloadWriteVarRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, messageType uint8, parameter S7Parameter) (S7PayloadWriteVarRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadWriteVarRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadWriteVarRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	items, err := ReadCountArrayField[S7VarPayloadDataItem](ctx, "items", ReadComplex[S7VarPayloadDataItem](S7VarPayloadDataItemParseWithBuffer, readBuffer), uint64(int32(len(CastS7ParameterWriteVarRequest(parameter).GetItems()))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}

	if closeErr := readBuffer.CloseContext("S7PayloadWriteVarRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadWriteVarRequest")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadWriteVarRequest{
		_S7Payload: &_S7Payload{
			Parameter: parameter,
		},
		Items: items,
	}
	_child._S7Payload._S7PayloadChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadWriteVarRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadWriteVarRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadWriteVarRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadWriteVarRequest")
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

		if popErr := writeBuffer.PopContext("S7PayloadWriteVarRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadWriteVarRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadWriteVarRequest) isS7PayloadWriteVarRequest() bool {
	return true
}

func (m *_S7PayloadWriteVarRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
