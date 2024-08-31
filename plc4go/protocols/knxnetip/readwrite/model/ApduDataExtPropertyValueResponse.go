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

// ApduDataExtPropertyValueResponse is the corresponding interface of ApduDataExtPropertyValueResponse
type ApduDataExtPropertyValueResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetCount returns Count (property field)
	GetCount() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint16
	// GetData returns Data (property field)
	GetData() []byte
}

// ApduDataExtPropertyValueResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtPropertyValueResponse.
// This is useful for switch cases.
type ApduDataExtPropertyValueResponseExactly interface {
	ApduDataExtPropertyValueResponse
	isApduDataExtPropertyValueResponse() bool
}

// _ApduDataExtPropertyValueResponse is the data-structure of this message
type _ApduDataExtPropertyValueResponse struct {
	*_ApduDataExt
	ObjectIndex uint8
	PropertyId  uint8
	Count       uint8
	Index       uint16
	Data        []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtPropertyValueResponse) GetExtApciType() uint8 {
	return 0x16
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtPropertyValueResponse) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtPropertyValueResponse) GetParent() ApduDataExt {
	return m._ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtPropertyValueResponse) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *_ApduDataExtPropertyValueResponse) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_ApduDataExtPropertyValueResponse) GetCount() uint8 {
	return m.Count
}

func (m *_ApduDataExtPropertyValueResponse) GetIndex() uint16 {
	return m.Index
}

func (m *_ApduDataExtPropertyValueResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtPropertyValueResponse factory function for _ApduDataExtPropertyValueResponse
func NewApduDataExtPropertyValueResponse(objectIndex uint8, propertyId uint8, count uint8, index uint16, data []byte, length uint8) *_ApduDataExtPropertyValueResponse {
	_result := &_ApduDataExtPropertyValueResponse{
		ObjectIndex:  objectIndex,
		PropertyId:   propertyId,
		Count:        count,
		Index:        index,
		Data:         data,
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtPropertyValueResponse(structType any) ApduDataExtPropertyValueResponse {
	if casted, ok := structType.(ApduDataExtPropertyValueResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyValueResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtPropertyValueResponse) GetTypeName() string {
	return "ApduDataExtPropertyValueResponse"
}

func (m *_ApduDataExtPropertyValueResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 4

	// Simple field (index)
	lengthInBits += 12

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ApduDataExtPropertyValueResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtPropertyValueResponseParse(ctx context.Context, theBytes []byte, length uint8) (ApduDataExtPropertyValueResponse, error) {
	return ApduDataExtPropertyValueResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtPropertyValueResponseParseWithBufferProducer(length uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (ApduDataExtPropertyValueResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ApduDataExtPropertyValueResponse, error) {
		return ApduDataExtPropertyValueResponseParseWithBuffer(ctx, readBuffer, length)
	}
}

func ApduDataExtPropertyValueResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtPropertyValueResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyValueResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyValueResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIndex, err := ReadSimpleField(ctx, "objectIndex", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIndex' field"))
	}

	propertyId, err := ReadSimpleField(ctx, "propertyId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyId' field"))
	}

	count, err := ReadSimpleField(ctx, "count", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'count' field"))
	}

	index, err := ReadSimpleField(ctx, "index", ReadUnsignedShort(readBuffer, uint8(12)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'index' field"))
	}

	data, err := readBuffer.ReadByteArray("data", int(int32(length)-int32(int32(5))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyValueResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyValueResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtPropertyValueResponse{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
		ObjectIndex: objectIndex,
		PropertyId:  propertyId,
		Count:       count,
		Index:       index,
		Data:        data,
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtPropertyValueResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtPropertyValueResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyValueResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyValueResponse")
		}

		// Simple Field (objectIndex)
		objectIndex := uint8(m.GetObjectIndex())
		_objectIndexErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("objectIndex", 8, uint8((objectIndex)))
		if _objectIndexErr != nil {
			return errors.Wrap(_objectIndexErr, "Error serializing 'objectIndex' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.GetPropertyId())
		_propertyIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("propertyId", 8, uint8((propertyId)))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (count)
		count := uint8(m.GetCount())
		_countErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("count", 4, uint8((count)))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		// Simple Field (index)
		index := uint16(m.GetIndex())
		_indexErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("index", 12, uint16((index)))
		if _indexErr != nil {
			return errors.Wrap(_indexErr, "Error serializing 'index' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyValueResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyValueResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtPropertyValueResponse) isApduDataExtPropertyValueResponse() bool {
	return true
}

func (m *_ApduDataExtPropertyValueResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
