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
	"io"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPriorityArray is the corresponding interface of BACnetPriorityArray
type BACnetPriorityArray interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetData returns Data (property field)
	GetData() []BACnetPriorityValue
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetPriorityValue01 returns PriorityValue01 (virtual field)
	GetPriorityValue01() BACnetPriorityValue
	// GetPriorityValue02 returns PriorityValue02 (virtual field)
	GetPriorityValue02() BACnetPriorityValue
	// GetPriorityValue03 returns PriorityValue03 (virtual field)
	GetPriorityValue03() BACnetPriorityValue
	// GetPriorityValue04 returns PriorityValue04 (virtual field)
	GetPriorityValue04() BACnetPriorityValue
	// GetPriorityValue05 returns PriorityValue05 (virtual field)
	GetPriorityValue05() BACnetPriorityValue
	// GetPriorityValue06 returns PriorityValue06 (virtual field)
	GetPriorityValue06() BACnetPriorityValue
	// GetPriorityValue07 returns PriorityValue07 (virtual field)
	GetPriorityValue07() BACnetPriorityValue
	// GetPriorityValue08 returns PriorityValue08 (virtual field)
	GetPriorityValue08() BACnetPriorityValue
	// GetPriorityValue09 returns PriorityValue09 (virtual field)
	GetPriorityValue09() BACnetPriorityValue
	// GetPriorityValue10 returns PriorityValue10 (virtual field)
	GetPriorityValue10() BACnetPriorityValue
	// GetPriorityValue11 returns PriorityValue11 (virtual field)
	GetPriorityValue11() BACnetPriorityValue
	// GetPriorityValue12 returns PriorityValue12 (virtual field)
	GetPriorityValue12() BACnetPriorityValue
	// GetPriorityValue13 returns PriorityValue13 (virtual field)
	GetPriorityValue13() BACnetPriorityValue
	// GetPriorityValue14 returns PriorityValue14 (virtual field)
	GetPriorityValue14() BACnetPriorityValue
	// GetPriorityValue15 returns PriorityValue15 (virtual field)
	GetPriorityValue15() BACnetPriorityValue
	// GetPriorityValue16 returns PriorityValue16 (virtual field)
	GetPriorityValue16() BACnetPriorityValue
	// GetIsIndexedAccess returns IsIndexedAccess (virtual field)
	GetIsIndexedAccess() bool
	// GetIndexEntry returns IndexEntry (virtual field)
	GetIndexEntry() BACnetPriorityValue
}

// BACnetPriorityArrayExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityArray.
// This is useful for switch cases.
type BACnetPriorityArrayExactly interface {
	BACnetPriorityArray
	isBACnetPriorityArray() bool
}

// _BACnetPriorityArray is the data-structure of this message
type _BACnetPriorityArray struct {
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	Data                 []BACnetPriorityValue

	// Arguments.
	ObjectTypeArgument BACnetObjectType
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityArray) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetPriorityArray) GetData() []BACnetPriorityValue {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetPriorityArray) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

func (m *_BACnetPriorityArray) GetPriorityValue01() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (0)), func() any { return CastBACnetPriorityValue(m.GetData()[0]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue02() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (1)), func() any { return CastBACnetPriorityValue(m.GetData()[1]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue03() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (2)), func() any { return CastBACnetPriorityValue(m.GetData()[2]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue04() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (3)), func() any { return CastBACnetPriorityValue(m.GetData()[3]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue05() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (4)), func() any { return CastBACnetPriorityValue(m.GetData()[4]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue06() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (5)), func() any { return CastBACnetPriorityValue(m.GetData()[5]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue07() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (6)), func() any { return CastBACnetPriorityValue(m.GetData()[6]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue08() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (7)), func() any { return CastBACnetPriorityValue(m.GetData()[7]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue09() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (8)), func() any { return CastBACnetPriorityValue(m.GetData()[8]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue10() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (9)), func() any { return CastBACnetPriorityValue(m.GetData()[9]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue11() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (10)), func() any { return CastBACnetPriorityValue(m.GetData()[10]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue12() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (11)), func() any { return CastBACnetPriorityValue(m.GetData()[11]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue13() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (12)), func() any { return CastBACnetPriorityValue(m.GetData()[12]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue14() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (13)), func() any { return CastBACnetPriorityValue(m.GetData()[13]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue15() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (14)), func() any { return CastBACnetPriorityValue(m.GetData()[14]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetPriorityValue16() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(CastBACnetPriorityValue(utils.InlineIf(bool((len(m.GetData())) > (15)), func() any { return CastBACnetPriorityValue(m.GetData()[15]) }, func() any { return CastBACnetPriorityValue(nil) })))
}

func (m *_BACnetPriorityArray) GetIsIndexedAccess() bool {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return bool(bool((len(m.GetData())) == (1)))
}

func (m *_BACnetPriorityArray) GetIndexEntry() BACnetPriorityValue {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetPriorityValue(m.GetPriorityValue01())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityArray factory function for _BACnetPriorityArray
func NewBACnetPriorityArray(numberOfDataElements BACnetApplicationTagUnsignedInteger, data []BACnetPriorityValue, objectTypeArgument BACnetObjectType, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetPriorityArray {
	return &_BACnetPriorityArray{NumberOfDataElements: numberOfDataElements, Data: data, ObjectTypeArgument: objectTypeArgument, TagNumber: tagNumber, ArrayIndexArgument: arrayIndexArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityArray(structType any) BACnetPriorityArray {
	if casted, ok := structType.(BACnetPriorityArray); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityArray); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityArray) GetTypeName() string {
	return "BACnetPriorityArray"
}

func (m *_BACnetPriorityArray) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPriorityArray) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPriorityArrayParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetPriorityArray, error) {
	return BACnetPriorityArrayParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument, tagNumber, arrayIndexArgument)
}

func BACnetPriorityArrayParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetPriorityArray, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPriorityArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityArray")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetPriorityArray")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	data, err := ReadTerminatedArrayField[BACnetPriorityValue](ctx, "data", ReadComplex[BACnetPriorityValue](func(ctx context.Context, buffer utils.ReadBuffer) (BACnetPriorityValue, error) {
		v, err := BACnetPriorityValueParseWithBuffer(ctx, readBuffer, (BACnetObjectType)(objectTypeArgument))
		if err != nil {
			return nil, err
		}
		return v.(BACnetPriorityValue), nil
	}, readBuffer), func() bool { return IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber) })
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}

	// Virtual field
	_priorityValue01 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (0)), func() any { return CastBACnetPriorityValue(data[0]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue01 := _priorityValue01
	_ = priorityValue01

	// Virtual field
	_priorityValue02 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (1)), func() any { return CastBACnetPriorityValue(data[1]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue02 := _priorityValue02
	_ = priorityValue02

	// Virtual field
	_priorityValue03 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (2)), func() any { return CastBACnetPriorityValue(data[2]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue03 := _priorityValue03
	_ = priorityValue03

	// Virtual field
	_priorityValue04 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (3)), func() any { return CastBACnetPriorityValue(data[3]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue04 := _priorityValue04
	_ = priorityValue04

	// Virtual field
	_priorityValue05 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (4)), func() any { return CastBACnetPriorityValue(data[4]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue05 := _priorityValue05
	_ = priorityValue05

	// Virtual field
	_priorityValue06 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (5)), func() any { return CastBACnetPriorityValue(data[5]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue06 := _priorityValue06
	_ = priorityValue06

	// Virtual field
	_priorityValue07 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (6)), func() any { return CastBACnetPriorityValue(data[6]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue07 := _priorityValue07
	_ = priorityValue07

	// Virtual field
	_priorityValue08 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (7)), func() any { return CastBACnetPriorityValue(data[7]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue08 := _priorityValue08
	_ = priorityValue08

	// Virtual field
	_priorityValue09 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (8)), func() any { return CastBACnetPriorityValue(data[8]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue09 := _priorityValue09
	_ = priorityValue09

	// Virtual field
	_priorityValue10 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (9)), func() any { return CastBACnetPriorityValue(data[9]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue10 := _priorityValue10
	_ = priorityValue10

	// Virtual field
	_priorityValue11 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (10)), func() any { return CastBACnetPriorityValue(data[10]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue11 := _priorityValue11
	_ = priorityValue11

	// Virtual field
	_priorityValue12 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (11)), func() any { return CastBACnetPriorityValue(data[11]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue12 := _priorityValue12
	_ = priorityValue12

	// Virtual field
	_priorityValue13 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (12)), func() any { return CastBACnetPriorityValue(data[12]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue13 := _priorityValue13
	_ = priorityValue13

	// Virtual field
	_priorityValue14 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (13)), func() any { return CastBACnetPriorityValue(data[13]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue14 := _priorityValue14
	_ = priorityValue14

	// Virtual field
	_priorityValue15 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (14)), func() any { return CastBACnetPriorityValue(data[14]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue15 := _priorityValue15
	_ = priorityValue15

	// Virtual field
	_priorityValue16 := CastBACnetPriorityValue(utils.InlineIf(bool((len(data)) > (15)), func() any { return CastBACnetPriorityValue(data[15]) }, func() any { return CastBACnetPriorityValue(nil) }))
	priorityValue16 := _priorityValue16
	_ = priorityValue16

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(data)) == (16)))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "Either indexed access or lenght 16 expected"})
	}

	// Virtual field
	_isIndexedAccess := bool((len(data)) == (1))
	isIndexedAccess := bool(_isIndexedAccess)
	_ = isIndexedAccess

	// Virtual field
	_indexEntry := priorityValue01
	indexEntry := _indexEntry
	_ = indexEntry

	if closeErr := readBuffer.CloseContext("BACnetPriorityArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityArray")
	}

	// Create the instance
	return &_BACnetPriorityArray{
		ObjectTypeArgument:   objectTypeArgument,
		TagNumber:            tagNumber,
		ArrayIndexArgument:   arrayIndexArgument,
		NumberOfDataElements: numberOfDataElements,
		Data:                 data,
	}, nil
}

func (m *_BACnetPriorityArray) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityArray) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPriorityArray"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPriorityArray")
	}
	// Virtual field
	zero := m.GetZero()
	_ = zero
	if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
		return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
	}

	// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if m.GetNumberOfDataElements() != nil {
		if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
		}
		numberOfDataElements = m.GetNumberOfDataElements()
		_numberOfDataElementsErr := writeBuffer.WriteSerializable(ctx, numberOfDataElements)
		if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for numberOfDataElements")
		}
		if _numberOfDataElementsErr != nil {
			return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
		}
	}

	// Array Field (data)
	if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for data")
	}
	for _curItem, _element := range m.GetData() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetData()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'data' field")
		}
	}
	if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for data")
	}
	// Virtual field
	priorityValue01 := m.GetPriorityValue01()
	_ = priorityValue01
	if _priorityValue01Err := writeBuffer.WriteVirtual(ctx, "priorityValue01", m.GetPriorityValue01()); _priorityValue01Err != nil {
		return errors.Wrap(_priorityValue01Err, "Error serializing 'priorityValue01' field")
	}
	// Virtual field
	priorityValue02 := m.GetPriorityValue02()
	_ = priorityValue02
	if _priorityValue02Err := writeBuffer.WriteVirtual(ctx, "priorityValue02", m.GetPriorityValue02()); _priorityValue02Err != nil {
		return errors.Wrap(_priorityValue02Err, "Error serializing 'priorityValue02' field")
	}
	// Virtual field
	priorityValue03 := m.GetPriorityValue03()
	_ = priorityValue03
	if _priorityValue03Err := writeBuffer.WriteVirtual(ctx, "priorityValue03", m.GetPriorityValue03()); _priorityValue03Err != nil {
		return errors.Wrap(_priorityValue03Err, "Error serializing 'priorityValue03' field")
	}
	// Virtual field
	priorityValue04 := m.GetPriorityValue04()
	_ = priorityValue04
	if _priorityValue04Err := writeBuffer.WriteVirtual(ctx, "priorityValue04", m.GetPriorityValue04()); _priorityValue04Err != nil {
		return errors.Wrap(_priorityValue04Err, "Error serializing 'priorityValue04' field")
	}
	// Virtual field
	priorityValue05 := m.GetPriorityValue05()
	_ = priorityValue05
	if _priorityValue05Err := writeBuffer.WriteVirtual(ctx, "priorityValue05", m.GetPriorityValue05()); _priorityValue05Err != nil {
		return errors.Wrap(_priorityValue05Err, "Error serializing 'priorityValue05' field")
	}
	// Virtual field
	priorityValue06 := m.GetPriorityValue06()
	_ = priorityValue06
	if _priorityValue06Err := writeBuffer.WriteVirtual(ctx, "priorityValue06", m.GetPriorityValue06()); _priorityValue06Err != nil {
		return errors.Wrap(_priorityValue06Err, "Error serializing 'priorityValue06' field")
	}
	// Virtual field
	priorityValue07 := m.GetPriorityValue07()
	_ = priorityValue07
	if _priorityValue07Err := writeBuffer.WriteVirtual(ctx, "priorityValue07", m.GetPriorityValue07()); _priorityValue07Err != nil {
		return errors.Wrap(_priorityValue07Err, "Error serializing 'priorityValue07' field")
	}
	// Virtual field
	priorityValue08 := m.GetPriorityValue08()
	_ = priorityValue08
	if _priorityValue08Err := writeBuffer.WriteVirtual(ctx, "priorityValue08", m.GetPriorityValue08()); _priorityValue08Err != nil {
		return errors.Wrap(_priorityValue08Err, "Error serializing 'priorityValue08' field")
	}
	// Virtual field
	priorityValue09 := m.GetPriorityValue09()
	_ = priorityValue09
	if _priorityValue09Err := writeBuffer.WriteVirtual(ctx, "priorityValue09", m.GetPriorityValue09()); _priorityValue09Err != nil {
		return errors.Wrap(_priorityValue09Err, "Error serializing 'priorityValue09' field")
	}
	// Virtual field
	priorityValue10 := m.GetPriorityValue10()
	_ = priorityValue10
	if _priorityValue10Err := writeBuffer.WriteVirtual(ctx, "priorityValue10", m.GetPriorityValue10()); _priorityValue10Err != nil {
		return errors.Wrap(_priorityValue10Err, "Error serializing 'priorityValue10' field")
	}
	// Virtual field
	priorityValue11 := m.GetPriorityValue11()
	_ = priorityValue11
	if _priorityValue11Err := writeBuffer.WriteVirtual(ctx, "priorityValue11", m.GetPriorityValue11()); _priorityValue11Err != nil {
		return errors.Wrap(_priorityValue11Err, "Error serializing 'priorityValue11' field")
	}
	// Virtual field
	priorityValue12 := m.GetPriorityValue12()
	_ = priorityValue12
	if _priorityValue12Err := writeBuffer.WriteVirtual(ctx, "priorityValue12", m.GetPriorityValue12()); _priorityValue12Err != nil {
		return errors.Wrap(_priorityValue12Err, "Error serializing 'priorityValue12' field")
	}
	// Virtual field
	priorityValue13 := m.GetPriorityValue13()
	_ = priorityValue13
	if _priorityValue13Err := writeBuffer.WriteVirtual(ctx, "priorityValue13", m.GetPriorityValue13()); _priorityValue13Err != nil {
		return errors.Wrap(_priorityValue13Err, "Error serializing 'priorityValue13' field")
	}
	// Virtual field
	priorityValue14 := m.GetPriorityValue14()
	_ = priorityValue14
	if _priorityValue14Err := writeBuffer.WriteVirtual(ctx, "priorityValue14", m.GetPriorityValue14()); _priorityValue14Err != nil {
		return errors.Wrap(_priorityValue14Err, "Error serializing 'priorityValue14' field")
	}
	// Virtual field
	priorityValue15 := m.GetPriorityValue15()
	_ = priorityValue15
	if _priorityValue15Err := writeBuffer.WriteVirtual(ctx, "priorityValue15", m.GetPriorityValue15()); _priorityValue15Err != nil {
		return errors.Wrap(_priorityValue15Err, "Error serializing 'priorityValue15' field")
	}
	// Virtual field
	priorityValue16 := m.GetPriorityValue16()
	_ = priorityValue16
	if _priorityValue16Err := writeBuffer.WriteVirtual(ctx, "priorityValue16", m.GetPriorityValue16()); _priorityValue16Err != nil {
		return errors.Wrap(_priorityValue16Err, "Error serializing 'priorityValue16' field")
	}
	// Virtual field
	isIndexedAccess := m.GetIsIndexedAccess()
	_ = isIndexedAccess
	if _isIndexedAccessErr := writeBuffer.WriteVirtual(ctx, "isIndexedAccess", m.GetIsIndexedAccess()); _isIndexedAccessErr != nil {
		return errors.Wrap(_isIndexedAccessErr, "Error serializing 'isIndexedAccess' field")
	}
	// Virtual field
	indexEntry := m.GetIndexEntry()
	_ = indexEntry
	if _indexEntryErr := writeBuffer.WriteVirtual(ctx, "indexEntry", m.GetIndexEntry()); _indexEntryErr != nil {
		return errors.Wrap(_indexEntryErr, "Error serializing 'indexEntry' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPriorityArray"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPriorityArray")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPriorityArray) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}
func (m *_BACnetPriorityArray) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetPriorityArray) GetArrayIndexArgument() BACnetTagPayloadUnsignedInteger {
	return m.ArrayIndexArgument
}

//
////

func (m *_BACnetPriorityArray) isBACnetPriorityArray() bool {
	return true
}

func (m *_BACnetPriorityArray) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
