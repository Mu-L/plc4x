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

// ApduData is the corresponding interface of ApduData
type ApduData interface {
	ApduDataContract
	ApduDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// ApduDataContract provides a set of functions which can be overwritten by a sub struct
type ApduDataContract interface {
	// GetDataLength() returns a parser argument
	GetDataLength() uint8
}

// ApduDataRequirements provides a set of functions which need to be implemented by a sub struct
type ApduDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetApciType returns ApciType (discriminator field)
	GetApciType() uint8
}

// ApduDataExactly can be used when we want exactly this type and not a type which fulfills ApduData.
// This is useful for switch cases.
type ApduDataExactly interface {
	ApduData
	isApduData() bool
}

// _ApduData is the data-structure of this message
type _ApduData struct {
	_SubType ApduData

	// Arguments.
	DataLength uint8
}

var _ ApduDataContract = (*_ApduData)(nil)

type ApduDataChild interface {
	utils.Serializable

	GetParent() *ApduData

	GetTypeName() string
	ApduData
}

// NewApduData factory function for _ApduData
func NewApduData(dataLength uint8) *_ApduData {
	return &_ApduData{DataLength: dataLength}
}

// Deprecated: use the interface for direct cast
func CastApduData(structType any) ApduData {
	if casted, ok := structType.(ApduData); ok {
		return casted
	}
	if casted, ok := structType.(*ApduData); ok {
		return *casted
	}
	return nil
}

func (m *_ApduData) GetTypeName() string {
	return "ApduData"
}

func (m *_ApduData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (apciType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ApduData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ApduDataParse[T ApduData](ctx context.Context, theBytes []byte, dataLength uint8) (T, error) {
	return ApduDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), dataLength)
}

func ApduDataParseWithBufferProducer[T ApduData](dataLength uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ApduDataParseWithBuffer[T](ctx, readBuffer, dataLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func ApduDataParseWithBuffer[T ApduData](ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (T, error) {
	v, err := (&_ApduData{}).parse(ctx, readBuffer, dataLength)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_ApduData) parse(ctx context.Context, readBuffer utils.ReadBuffer, dataLength uint8) (__apduData ApduData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	apciType, err := ReadDiscriminatorField[uint8](ctx, "apciType", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'apciType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ApduData
	switch {
	case apciType == 0x0: // ApduDataGroupValueRead
		if _child, err = (&_ApduDataGroupValueRead{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataGroupValueRead for type-switch of ApduData")
		}
	case apciType == 0x1: // ApduDataGroupValueResponse
		if _child, err = (&_ApduDataGroupValueResponse{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataGroupValueResponse for type-switch of ApduData")
		}
	case apciType == 0x2: // ApduDataGroupValueWrite
		if _child, err = (&_ApduDataGroupValueWrite{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataGroupValueWrite for type-switch of ApduData")
		}
	case apciType == 0x3: // ApduDataIndividualAddressWrite
		if _child, err = (&_ApduDataIndividualAddressWrite{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataIndividualAddressWrite for type-switch of ApduData")
		}
	case apciType == 0x4: // ApduDataIndividualAddressRead
		if _child, err = (&_ApduDataIndividualAddressRead{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataIndividualAddressRead for type-switch of ApduData")
		}
	case apciType == 0x5: // ApduDataIndividualAddressResponse
		if _child, err = (&_ApduDataIndividualAddressResponse{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataIndividualAddressResponse for type-switch of ApduData")
		}
	case apciType == 0x6: // ApduDataAdcRead
		if _child, err = (&_ApduDataAdcRead{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataAdcRead for type-switch of ApduData")
		}
	case apciType == 0x7: // ApduDataAdcResponse
		if _child, err = (&_ApduDataAdcResponse{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataAdcResponse for type-switch of ApduData")
		}
	case apciType == 0x8: // ApduDataMemoryRead
		if _child, err = (&_ApduDataMemoryRead{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataMemoryRead for type-switch of ApduData")
		}
	case apciType == 0x9: // ApduDataMemoryResponse
		if _child, err = (&_ApduDataMemoryResponse{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataMemoryResponse for type-switch of ApduData")
		}
	case apciType == 0xA: // ApduDataMemoryWrite
		if _child, err = (&_ApduDataMemoryWrite{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataMemoryWrite for type-switch of ApduData")
		}
	case apciType == 0xB: // ApduDataUserMessage
		if _child, err = (&_ApduDataUserMessage{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataUserMessage for type-switch of ApduData")
		}
	case apciType == 0xC: // ApduDataDeviceDescriptorRead
		if _child, err = (&_ApduDataDeviceDescriptorRead{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataDeviceDescriptorRead for type-switch of ApduData")
		}
	case apciType == 0xD: // ApduDataDeviceDescriptorResponse
		if _child, err = (&_ApduDataDeviceDescriptorResponse{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataDeviceDescriptorResponse for type-switch of ApduData")
		}
	case apciType == 0xE: // ApduDataRestart
		if _child, err = (&_ApduDataRestart{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataRestart for type-switch of ApduData")
		}
	case apciType == 0xF: // ApduDataOther
		if _child, err = (&_ApduDataOther{}).parse(ctx, readBuffer, m, dataLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ApduDataOther for type-switch of ApduData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [apciType=%v]", apciType)
	}

	if closeErr := readBuffer.CloseContext("ApduData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduData")
	}

	return _child, nil
}

func (pm *_ApduData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ApduData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ApduData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApduData")
	}

	if err := WriteDiscriminatorField(ctx, "apciType", m.GetApciType(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
		return errors.Wrap(err, "Error serializing 'apciType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApduData")
	}
	return nil
}

////
// Arguments Getter

func (m *_ApduData) GetDataLength() uint8 {
	return m.DataLength
}

//
////

func (m *_ApduData) isApduData() bool {
	return true
}
