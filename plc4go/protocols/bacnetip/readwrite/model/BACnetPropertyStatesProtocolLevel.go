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

// BACnetPropertyStatesProtocolLevel is the corresponding interface of BACnetPropertyStatesProtocolLevel
type BACnetPropertyStatesProtocolLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetProtocolLevel returns ProtocolLevel (property field)
	GetProtocolLevel() BACnetProtocolLevelTagged
	// IsBACnetPropertyStatesProtocolLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesProtocolLevel()
}

// _BACnetPropertyStatesProtocolLevel is the data-structure of this message
type _BACnetPropertyStatesProtocolLevel struct {
	BACnetPropertyStatesContract
	ProtocolLevel BACnetProtocolLevelTagged
}

var _ BACnetPropertyStatesProtocolLevel = (*_BACnetPropertyStatesProtocolLevel)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesProtocolLevel)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesProtocolLevel) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesProtocolLevel) GetProtocolLevel() BACnetProtocolLevelTagged {
	return m.ProtocolLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesProtocolLevel factory function for _BACnetPropertyStatesProtocolLevel
func NewBACnetPropertyStatesProtocolLevel(protocolLevel BACnetProtocolLevelTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesProtocolLevel {
	_result := &_BACnetPropertyStatesProtocolLevel{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		ProtocolLevel:                protocolLevel,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesProtocolLevel(structType any) BACnetPropertyStatesProtocolLevel {
	if casted, ok := structType.(BACnetPropertyStatesProtocolLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesProtocolLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesProtocolLevel) GetTypeName() string {
	return "BACnetPropertyStatesProtocolLevel"
}

func (m *_BACnetPropertyStatesProtocolLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (protocolLevel)
	lengthInBits += m.ProtocolLevel.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesProtocolLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesProtocolLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesProtocolLevel BACnetPropertyStatesProtocolLevel, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesProtocolLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesProtocolLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolLevel, err := ReadSimpleField[BACnetProtocolLevelTagged](ctx, "protocolLevel", ReadComplex[BACnetProtocolLevelTagged](BACnetProtocolLevelTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolLevel' field"))
	}
	m.ProtocolLevel = protocolLevel

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesProtocolLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesProtocolLevel")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesProtocolLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesProtocolLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesProtocolLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesProtocolLevel")
		}

		if err := WriteSimpleField[BACnetProtocolLevelTagged](ctx, "protocolLevel", m.GetProtocolLevel(), WriteComplex[BACnetProtocolLevelTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolLevel' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesProtocolLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesProtocolLevel")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesProtocolLevel) IsBACnetPropertyStatesProtocolLevel() {}

func (m *_BACnetPropertyStatesProtocolLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
