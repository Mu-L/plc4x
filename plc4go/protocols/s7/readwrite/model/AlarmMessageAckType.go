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

// AlarmMessageAckType is the corresponding interface of AlarmMessageAckType
type AlarmMessageAckType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects (property field)
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []AlarmMessageObjectAckType
}

// AlarmMessageAckTypeExactly can be used when we want exactly this type and not a type which fulfills AlarmMessageAckType.
// This is useful for switch cases.
type AlarmMessageAckTypeExactly interface {
	AlarmMessageAckType
	isAlarmMessageAckType() bool
}

// _AlarmMessageAckType is the data-structure of this message
type _AlarmMessageAckType struct {
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []AlarmMessageObjectAckType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageAckType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_AlarmMessageAckType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *_AlarmMessageAckType) GetMessageObjects() []AlarmMessageObjectAckType {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageAckType factory function for _AlarmMessageAckType
func NewAlarmMessageAckType(functionId uint8, numberOfObjects uint8, messageObjects []AlarmMessageObjectAckType) *_AlarmMessageAckType {
	return &_AlarmMessageAckType{FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

// Deprecated: use the interface for direct cast
func CastAlarmMessageAckType(structType any) AlarmMessageAckType {
	if casted, ok := structType.(AlarmMessageAckType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageAckType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageAckType) GetTypeName() string {
	return "AlarmMessageAckType"
}

func (m *_AlarmMessageAckType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for _curItem, element := range m.MessageObjects {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MessageObjects), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_AlarmMessageAckType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AlarmMessageAckTypeParse(ctx context.Context, theBytes []byte) (AlarmMessageAckType, error) {
	return AlarmMessageAckTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AlarmMessageAckTypeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckType, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckType, error) {
		return AlarmMessageAckTypeParseWithBuffer(ctx, readBuffer)
	}
}

func AlarmMessageAckTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AlarmMessageAckType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AlarmMessageAckType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageAckType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	functionId, err := ReadSimpleField(ctx, "functionId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionId' field"))
	}

	numberOfObjects, err := ReadSimpleField(ctx, "numberOfObjects", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfObjects' field"))
	}

	messageObjects, err := ReadCountArrayField[AlarmMessageObjectAckType](ctx, "messageObjects", ReadComplex[AlarmMessageObjectAckType](AlarmMessageObjectAckTypeParseWithBuffer, readBuffer), uint64(numberOfObjects))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageObjects' field"))
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageAckType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageAckType")
	}

	// Create the instance
	return &_AlarmMessageAckType{
		FunctionId:      functionId,
		NumberOfObjects: numberOfObjects,
		MessageObjects:  messageObjects,
	}, nil
}

func (m *_AlarmMessageAckType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AlarmMessageAckType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AlarmMessageAckType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageAckType")
	}

	// Simple Field (functionId)
	functionId := uint8(m.GetFunctionId())
	_functionIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("functionId", 8, uint8((functionId)))
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Simple Field (numberOfObjects)
	numberOfObjects := uint8(m.GetNumberOfObjects())
	_numberOfObjectsErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("numberOfObjects", 8, uint8((numberOfObjects)))
	if _numberOfObjectsErr != nil {
		return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
	}

	// Array Field (messageObjects)
	if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for messageObjects")
	}
	for _curItem, _element := range m.GetMessageObjects() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetMessageObjects()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
		}
	}
	if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for messageObjects")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageAckType")
	}
	return nil
}

func (m *_AlarmMessageAckType) isAlarmMessageAckType() bool {
	return true
}

func (m *_AlarmMessageAckType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
