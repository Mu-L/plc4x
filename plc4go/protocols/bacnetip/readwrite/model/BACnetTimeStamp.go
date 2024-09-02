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

// BACnetTimeStamp is the corresponding interface of BACnetTimeStamp
type BACnetTimeStamp interface {
	BACnetTimeStampContract
	BACnetTimeStampRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetTimeStampContract provides a set of functions which can be overwritten by a sub struct
type BACnetTimeStampContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetTimeStampRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetTimeStampRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// BACnetTimeStampExactly can be used when we want exactly this type and not a type which fulfills BACnetTimeStamp.
// This is useful for switch cases.
type BACnetTimeStampExactly interface {
	BACnetTimeStamp
	isBACnetTimeStamp() bool
}

// _BACnetTimeStamp is the data-structure of this message
type _BACnetTimeStamp struct {
	_SubType        BACnetTimeStamp
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetTimeStampContract = (*_BACnetTimeStamp)(nil)

type BACnetTimeStampChild interface {
	utils.Serializable

	GetParent() *BACnetTimeStamp

	GetTypeName() string
	BACnetTimeStamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStamp) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetTimeStamp) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimeStamp factory function for _BACnetTimeStamp
func NewBACnetTimeStamp(peekedTagHeader BACnetTagHeader) *_BACnetTimeStamp {
	return &_BACnetTimeStamp{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetTimeStamp(structType any) BACnetTimeStamp {
	if casted, ok := structType.(BACnetTimeStamp); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStamp); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStamp) GetTypeName() string {
	return "BACnetTimeStamp"
}

func (m *_BACnetTimeStamp) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTimeStamp) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetTimeStampParse[T BACnetTimeStamp](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetTimeStampParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTimeStampParseWithBufferProducer[T BACnetTimeStamp]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetTimeStampParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetTimeStampParseWithBuffer[T BACnetTimeStamp](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetTimeStamp{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetTimeStamp) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetTimeStamp BACnetTimeStamp, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetTimeStamp
	switch {
	case peekedTagNumber == uint8(0): // BACnetTimeStampTime
		if _child, err = (&_BACnetTimeStampTime{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimeStampTime for type-switch of BACnetTimeStamp")
		}
	case peekedTagNumber == uint8(1): // BACnetTimeStampSequence
		if _child, err = (&_BACnetTimeStampSequence{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimeStampSequence for type-switch of BACnetTimeStamp")
		}
	case peekedTagNumber == uint8(2): // BACnetTimeStampDateTime
		if _child, err = (&_BACnetTimeStampDateTime{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetTimeStampDateTime for type-switch of BACnetTimeStamp")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStamp")
	}

	return _child, nil
}

func (pm *_BACnetTimeStamp) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetTimeStamp, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTimeStamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimeStamp")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeStamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimeStamp")
	}
	return nil
}

func (m *_BACnetTimeStamp) isBACnetTimeStamp() bool {
	return true
}
