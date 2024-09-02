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

// BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
type BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() BACnetContextTagUnsignedInteger
	// GetServiceParameters returns ServiceParameters (property field)
	GetServiceParameters() BACnetConstructedData
}

// BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferExactly interface {
	BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
	isBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer() bool
}

// _BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer struct {
	BACnetUnconfirmedServiceRequestContract
	VendorId          BACnetVendorIdTagged
	ServiceNumber     BACnetContextTagUnsignedInteger
	ServiceParameters BACnetConstructedData
}

var _ BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer = (*_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UNCONFIRMED_PRIVATE_TRANSFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetServiceNumber() BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetServiceParameters() BACnetConstructedData {
	return m.ServiceParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer factory function for _BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
func NewBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(vendorId BACnetVendorIdTagged, serviceNumber BACnetContextTagUnsignedInteger, serviceParameters BACnetConstructedData, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
	_result := &_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		VendorId:                                vendorId,
		ServiceNumber:                           serviceNumber,
		ServiceParameters:                       serviceParameters,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(structType any) BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

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

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")
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

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")
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

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) isBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
