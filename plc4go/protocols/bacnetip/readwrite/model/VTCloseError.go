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

// VTCloseError is the corresponding interface of VTCloseError
type VTCloseError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetListOfVtSessionIdentifiers returns ListOfVtSessionIdentifiers (property field)
	GetListOfVtSessionIdentifiers() VTCloseErrorListOfVTSessionIdentifiers
}

// VTCloseErrorExactly can be used when we want exactly this type and not a type which fulfills VTCloseError.
// This is useful for switch cases.
type VTCloseErrorExactly interface {
	VTCloseError
	isVTCloseError() bool
}

// _VTCloseError is the data-structure of this message
type _VTCloseError struct {
	BACnetErrorContract
	ErrorType                  ErrorEnclosed
	ListOfVtSessionIdentifiers VTCloseErrorListOfVTSessionIdentifiers
}

var _ VTCloseError = (*_VTCloseError)(nil)
var _ BACnetErrorRequirements = (*_VTCloseError)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VTCloseError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_CLOSE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VTCloseError) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VTCloseError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_VTCloseError) GetListOfVtSessionIdentifiers() VTCloseErrorListOfVTSessionIdentifiers {
	return m.ListOfVtSessionIdentifiers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewVTCloseError factory function for _VTCloseError
func NewVTCloseError(errorType ErrorEnclosed, listOfVtSessionIdentifiers VTCloseErrorListOfVTSessionIdentifiers) *_VTCloseError {
	_result := &_VTCloseError{
		BACnetErrorContract:        NewBACnetError(),
		ErrorType:                  errorType,
		ListOfVtSessionIdentifiers: listOfVtSessionIdentifiers,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastVTCloseError(structType any) VTCloseError {
	if casted, ok := structType.(VTCloseError); ok {
		return casted
	}
	if casted, ok := structType.(*VTCloseError); ok {
		return *casted
	}
	return nil
}

func (m *_VTCloseError) GetTypeName() string {
	return "VTCloseError"
}

func (m *_VTCloseError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Optional Field (listOfVtSessionIdentifiers)
	if m.ListOfVtSessionIdentifiers != nil {
		lengthInBits += m.ListOfVtSessionIdentifiers.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_VTCloseError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VTCloseError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__vTCloseError VTCloseError, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VTCloseError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VTCloseError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	var listOfVtSessionIdentifiers VTCloseErrorListOfVTSessionIdentifiers
	_listOfVtSessionIdentifiers, err := ReadOptionalField[VTCloseErrorListOfVTSessionIdentifiers](ctx, "listOfVtSessionIdentifiers", ReadComplex[VTCloseErrorListOfVTSessionIdentifiers](VTCloseErrorListOfVTSessionIdentifiersParseWithBufferProducer((uint8)(uint8(1))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfVtSessionIdentifiers' field"))
	}
	if _listOfVtSessionIdentifiers != nil {
		listOfVtSessionIdentifiers = *_listOfVtSessionIdentifiers
		m.ListOfVtSessionIdentifiers = listOfVtSessionIdentifiers
	}

	if closeErr := readBuffer.CloseContext("VTCloseError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VTCloseError")
	}

	return m, nil
}

func (m *_VTCloseError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VTCloseError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VTCloseError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VTCloseError")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorType' field")
		}

		if err := WriteOptionalField[VTCloseErrorListOfVTSessionIdentifiers](ctx, "listOfVtSessionIdentifiers", GetRef(m.GetListOfVtSessionIdentifiers()), WriteComplex[VTCloseErrorListOfVTSessionIdentifiers](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfVtSessionIdentifiers' field")
		}

		if popErr := writeBuffer.PopContext("VTCloseError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VTCloseError")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VTCloseError) isVTCloseError() bool {
	return true
}

func (m *_VTCloseError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
