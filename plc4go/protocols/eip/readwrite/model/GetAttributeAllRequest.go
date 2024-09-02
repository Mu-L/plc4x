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

// GetAttributeAllRequest is the corresponding interface of GetAttributeAllRequest
type GetAttributeAllRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() PathSegment
	// GetInstanceSegment returns InstanceSegment (property field)
	GetInstanceSegment() PathSegment
	// IsGetAttributeAllRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGetAttributeAllRequest()
}

// _GetAttributeAllRequest is the data-structure of this message
type _GetAttributeAllRequest struct {
	CipServiceContract
	ClassSegment    PathSegment
	InstanceSegment PathSegment
}

var _ GetAttributeAllRequest = (*_GetAttributeAllRequest)(nil)
var _ CipServiceRequirements = (*_GetAttributeAllRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeAllRequest) GetService() uint8 {
	return 0x01
}

func (m *_GetAttributeAllRequest) GetResponse() bool {
	return bool(false)
}

func (m *_GetAttributeAllRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeAllRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GetAttributeAllRequest) GetClassSegment() PathSegment {
	return m.ClassSegment
}

func (m *_GetAttributeAllRequest) GetInstanceSegment() PathSegment {
	return m.InstanceSegment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGetAttributeAllRequest factory function for _GetAttributeAllRequest
func NewGetAttributeAllRequest(classSegment PathSegment, instanceSegment PathSegment, serviceLen uint16) *_GetAttributeAllRequest {
	_result := &_GetAttributeAllRequest{
		CipServiceContract: NewCipService(serviceLen),
		ClassSegment:       classSegment,
		InstanceSegment:    instanceSegment,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastGetAttributeAllRequest(structType any) GetAttributeAllRequest {
	if casted, ok := structType.(GetAttributeAllRequest); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeAllRequest); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeAllRequest) GetTypeName() string {
	return "GetAttributeAllRequest"
}

func (m *_GetAttributeAllRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Simple field (classSegment)
	lengthInBits += m.ClassSegment.GetLengthInBits(ctx)

	// Simple field (instanceSegment)
	lengthInBits += m.InstanceSegment.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_GetAttributeAllRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetAttributeAllRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__getAttributeAllRequest GetAttributeAllRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeAllRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeAllRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadImplicitField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	classSegment, err := ReadSimpleField[PathSegment](ctx, "classSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'classSegment' field"))
	}
	m.ClassSegment = classSegment

	instanceSegment, err := ReadSimpleField[PathSegment](ctx, "instanceSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instanceSegment' field"))
	}
	m.InstanceSegment = instanceSegment

	if closeErr := readBuffer.CloseContext("GetAttributeAllRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeAllRequest")
	}

	return m, nil
}

func (m *_GetAttributeAllRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeAllRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeAllRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeAllRequest")
		}
		requestPathSize := uint8(uint8((uint8(m.GetClassSegment().GetLengthInBytes(ctx)) + uint8(m.GetInstanceSegment().GetLengthInBytes(ctx)))) / uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "requestPathSize", requestPathSize, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "classSegment", m.GetClassSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'classSegment' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "instanceSegment", m.GetInstanceSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'instanceSegment' field")
		}

		if popErr := writeBuffer.PopContext("GetAttributeAllRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeAllRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeAllRequest) IsGetAttributeAllRequest() {}

func (m *_GetAttributeAllRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
