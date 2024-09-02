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

// APDUUnconfirmedRequest is the corresponding interface of APDUUnconfirmedRequest
type APDUUnconfirmedRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	APDU
	// GetServiceRequest returns ServiceRequest (property field)
	GetServiceRequest() BACnetUnconfirmedServiceRequest
	// IsAPDUUnconfirmedRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAPDUUnconfirmedRequest()
}

// _APDUUnconfirmedRequest is the data-structure of this message
type _APDUUnconfirmedRequest struct {
	APDUContract
	ServiceRequest BACnetUnconfirmedServiceRequest
	// Reserved Fields
	reservedField0 *uint8
}

var _ APDUUnconfirmedRequest = (*_APDUUnconfirmedRequest)(nil)
var _ APDURequirements = (*_APDUUnconfirmedRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUUnconfirmedRequest) GetApduType() ApduType {
	return ApduType_UNCONFIRMED_REQUEST_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUUnconfirmedRequest) GetParent() APDUContract {
	return m.APDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUUnconfirmedRequest) GetServiceRequest() BACnetUnconfirmedServiceRequest {
	return m.ServiceRequest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUUnconfirmedRequest factory function for _APDUUnconfirmedRequest
func NewAPDUUnconfirmedRequest(serviceRequest BACnetUnconfirmedServiceRequest, apduLength uint16) *_APDUUnconfirmedRequest {
	_result := &_APDUUnconfirmedRequest{
		APDUContract:   NewAPDU(apduLength),
		ServiceRequest: serviceRequest,
	}
	_result.APDUContract.(*_APDU)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUUnconfirmedRequest(structType any) APDUUnconfirmedRequest {
	if casted, ok := structType.(APDUUnconfirmedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*APDUUnconfirmedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_APDUUnconfirmedRequest) GetTypeName() string {
	return "APDUUnconfirmedRequest"
}

func (m *_APDUUnconfirmedRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.APDUContract.(*_APDU).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (serviceRequest)
	lengthInBits += m.ServiceRequest.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_APDUUnconfirmedRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_APDUUnconfirmedRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_APDU, apduLength uint16) (__aPDUUnconfirmedRequest APDUUnconfirmedRequest, err error) {
	m.APDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUUnconfirmedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUUnconfirmedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(4)), uint8(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	serviceRequest, err := ReadSimpleField[BACnetUnconfirmedServiceRequest](ctx, "serviceRequest", ReadComplex[BACnetUnconfirmedServiceRequest](BACnetUnconfirmedServiceRequestParseWithBufferProducer[BACnetUnconfirmedServiceRequest]((uint16)(uint16(apduLength)-uint16(uint16(1)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceRequest' field"))
	}
	m.ServiceRequest = serviceRequest

	if closeErr := readBuffer.CloseContext("APDUUnconfirmedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUUnconfirmedRequest")
	}

	return m, nil
}

func (m *_APDUUnconfirmedRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUUnconfirmedRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUUnconfirmedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUUnconfirmedRequest")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[BACnetUnconfirmedServiceRequest](ctx, "serviceRequest", m.GetServiceRequest(), WriteComplex[BACnetUnconfirmedServiceRequest](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceRequest' field")
		}

		if popErr := writeBuffer.PopContext("APDUUnconfirmedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUUnconfirmedRequest")
		}
		return nil
	}
	return m.APDUContract.(*_APDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUUnconfirmedRequest) IsAPDUUnconfirmedRequest() {}

func (m *_APDUUnconfirmedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
