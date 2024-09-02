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

// BACnetConfirmedServiceRequestConfirmedPrivateTransfer is the corresponding interface of BACnetConfirmedServiceRequestConfirmedPrivateTransfer
type BACnetConfirmedServiceRequestConfirmedPrivateTransfer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() BACnetContextTagUnsignedInteger
	// GetServiceParameters returns ServiceParameters (property field)
	GetServiceParameters() BACnetConstructedData
	// IsBACnetConfirmedServiceRequestConfirmedPrivateTransfer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedPrivateTransfer()
}

// _BACnetConfirmedServiceRequestConfirmedPrivateTransfer is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedPrivateTransfer struct {
	BACnetConfirmedServiceRequestContract
	VendorId          BACnetVendorIdTagged
	ServiceNumber     BACnetContextTagUnsignedInteger
	ServiceParameters BACnetConstructedData
}

var _ BACnetConfirmedServiceRequestConfirmedPrivateTransfer = (*_BACnetConfirmedServiceRequestConfirmedPrivateTransfer)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestConfirmedPrivateTransfer)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetServiceNumber() BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetServiceParameters() BACnetConstructedData {
	return m.ServiceParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedPrivateTransfer factory function for _BACnetConfirmedServiceRequestConfirmedPrivateTransfer
func NewBACnetConfirmedServiceRequestConfirmedPrivateTransfer(vendorId BACnetVendorIdTagged, serviceNumber BACnetContextTagUnsignedInteger, serviceParameters BACnetConstructedData, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer {
	_result := &_BACnetConfirmedServiceRequestConfirmedPrivateTransfer{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		VendorId:                              vendorId,
		ServiceNumber:                         serviceNumber,
		ServiceParameters:                     serviceParameters,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedPrivateTransfer(structType any) BACnetConfirmedServiceRequestConfirmedPrivateTransfer {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedPrivateTransfer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedPrivateTransfer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedPrivateTransfer"
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (serviceNumber)
	lengthInBits += m.ServiceNumber.GetLengthInBits(ctx)

	// Optional Field (serviceParameters)
	if m.ServiceParameters != nil {
		lengthInBits += m.ServiceParameters.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestConfirmedPrivateTransfer BACnetConfirmedServiceRequestConfirmedPrivateTransfer, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedPrivateTransfer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedPrivateTransfer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	serviceNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "serviceNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceNumber' field"))
	}
	m.ServiceNumber = serviceNumber

	var serviceParameters BACnetConstructedData
	_serviceParameters, err := ReadOptionalField[BACnetConstructedData](ctx, "serviceParameters", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(2)), (BACnetObjectType)(BACnetObjectType_VENDOR_PROPRIETARY_VALUE), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceParameters' field"))
	}
	if _serviceParameters != nil {
		serviceParameters = *_serviceParameters
		m.ServiceParameters = serviceParameters
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedPrivateTransfer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedPrivateTransfer")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedPrivateTransfer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedPrivateTransfer")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorId", m.GetVendorId(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "serviceNumber", m.GetServiceNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceNumber' field")
		}

		if err := WriteOptionalField[BACnetConstructedData](ctx, "serviceParameters", GetRef(m.GetServiceParameters()), WriteComplex[BACnetConstructedData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'serviceParameters' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedPrivateTransfer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedPrivateTransfer")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) IsBACnetConfirmedServiceRequestConfirmedPrivateTransfer() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
