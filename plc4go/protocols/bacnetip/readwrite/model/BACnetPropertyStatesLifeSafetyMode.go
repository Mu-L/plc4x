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

// BACnetPropertyStatesLifeSafetyMode is the corresponding interface of BACnetPropertyStatesLifeSafetyMode
type BACnetPropertyStatesLifeSafetyMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLifeSafetyMode returns LifeSafetyMode (property field)
	GetLifeSafetyMode() BACnetLifeSafetyModeTagged
	// IsBACnetPropertyStatesLifeSafetyMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLifeSafetyMode()
}

// _BACnetPropertyStatesLifeSafetyMode is the data-structure of this message
type _BACnetPropertyStatesLifeSafetyMode struct {
	BACnetPropertyStatesContract
	LifeSafetyMode BACnetLifeSafetyModeTagged
}

var _ BACnetPropertyStatesLifeSafetyMode = (*_BACnetPropertyStatesLifeSafetyMode)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLifeSafetyMode)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLifeSafetyMode) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLifeSafetyMode() BACnetLifeSafetyModeTagged {
	return m.LifeSafetyMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLifeSafetyMode factory function for _BACnetPropertyStatesLifeSafetyMode
func NewBACnetPropertyStatesLifeSafetyMode(lifeSafetyMode BACnetLifeSafetyModeTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLifeSafetyMode {
	_result := &_BACnetPropertyStatesLifeSafetyMode{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LifeSafetyMode:               lifeSafetyMode,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLifeSafetyMode(structType any) BACnetPropertyStatesLifeSafetyMode {
	if casted, ok := structType.(BACnetPropertyStatesLifeSafetyMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLifeSafetyMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetTypeName() string {
	return "BACnetPropertyStatesLifeSafetyMode"
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lifeSafetyMode)
	lengthInBits += m.LifeSafetyMode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLifeSafetyMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLifeSafetyMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLifeSafetyMode BACnetPropertyStatesLifeSafetyMode, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLifeSafetyMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLifeSafetyMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lifeSafetyMode, err := ReadSimpleField[BACnetLifeSafetyModeTagged](ctx, "lifeSafetyMode", ReadComplex[BACnetLifeSafetyModeTagged](BACnetLifeSafetyModeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifeSafetyMode' field"))
	}
	m.LifeSafetyMode = lifeSafetyMode

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLifeSafetyMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLifeSafetyMode")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLifeSafetyMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLifeSafetyMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLifeSafetyMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLifeSafetyMode")
		}

		if err := WriteSimpleField[BACnetLifeSafetyModeTagged](ctx, "lifeSafetyMode", m.GetLifeSafetyMode(), WriteComplex[BACnetLifeSafetyModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lifeSafetyMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLifeSafetyMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLifeSafetyMode")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLifeSafetyMode) IsBACnetPropertyStatesLifeSafetyMode() {}

func (m *_BACnetPropertyStatesLifeSafetyMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
