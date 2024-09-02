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

// ChangeListAddError is the corresponding interface of ChangeListAddError
type ChangeListAddError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedElementNumber returns FirstFailedElementNumber (property field)
	GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger
	// IsChangeListAddError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsChangeListAddError()
}

// _ChangeListAddError is the data-structure of this message
type _ChangeListAddError struct {
	BACnetErrorContract
	ErrorType                ErrorEnclosed
	FirstFailedElementNumber BACnetContextTagUnsignedInteger
}

var _ ChangeListAddError = (*_ChangeListAddError)(nil)
var _ BACnetErrorRequirements = (*_ChangeListAddError)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ChangeListAddError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ChangeListAddError) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ChangeListAddError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_ChangeListAddError) GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger {
	return m.FirstFailedElementNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewChangeListAddError factory function for _ChangeListAddError
func NewChangeListAddError(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) *_ChangeListAddError {
	_result := &_ChangeListAddError{
		BACnetErrorContract:      NewBACnetError(),
		ErrorType:                errorType,
		FirstFailedElementNumber: firstFailedElementNumber,
	}
	_result.BACnetErrorContract.(*_BACnetError)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastChangeListAddError(structType any) ChangeListAddError {
	if casted, ok := structType.(ChangeListAddError); ok {
		return casted
	}
	if casted, ok := structType.(*ChangeListAddError); ok {
		return *casted
	}
	return nil
}

func (m *_ChangeListAddError) GetTypeName() string {
	return "ChangeListAddError"
}

func (m *_ChangeListAddError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (firstFailedElementNumber)
	lengthInBits += m.FirstFailedElementNumber.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ChangeListAddError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ChangeListAddError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__changeListAddError ChangeListAddError, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ChangeListAddError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ChangeListAddError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	firstFailedElementNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "firstFailedElementNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'firstFailedElementNumber' field"))
	}
	m.FirstFailedElementNumber = firstFailedElementNumber

	if closeErr := readBuffer.CloseContext("ChangeListAddError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ChangeListAddError")
	}

	return m, nil
}

func (m *_ChangeListAddError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ChangeListAddError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ChangeListAddError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ChangeListAddError")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorType' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "firstFailedElementNumber", m.GetFirstFailedElementNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'firstFailedElementNumber' field")
		}

		if popErr := writeBuffer.PopContext("ChangeListAddError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ChangeListAddError")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ChangeListAddError) IsChangeListAddError() {}

func (m *_ChangeListAddError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
