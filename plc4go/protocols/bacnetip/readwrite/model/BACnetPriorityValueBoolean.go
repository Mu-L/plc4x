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

// BACnetPriorityValueBoolean is the corresponding interface of BACnetPriorityValueBoolean
type BACnetPriorityValueBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
	// IsBACnetPriorityValueBoolean is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueBoolean()
}

// _BACnetPriorityValueBoolean is the data-structure of this message
type _BACnetPriorityValueBoolean struct {
	BACnetPriorityValueContract
	BooleanValue BACnetApplicationTagBoolean
}

var _ BACnetPriorityValueBoolean = (*_BACnetPriorityValueBoolean)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueBoolean)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueBoolean) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueBoolean factory function for _BACnetPriorityValueBoolean
func NewBACnetPriorityValueBoolean(booleanValue BACnetApplicationTagBoolean, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueBoolean {
	_result := &_BACnetPriorityValueBoolean{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		BooleanValue:                booleanValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueBoolean(structType any) BACnetPriorityValueBoolean {
	if casted, ok := structType.(BACnetPriorityValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueBoolean) GetTypeName() string {
	return "BACnetPriorityValueBoolean"
}

func (m *_BACnetPriorityValueBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueBoolean BACnetPriorityValueBoolean, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	booleanValue, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	m.BooleanValue = booleanValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueBoolean")
	}

	return m, nil
}

func (m *_BACnetPriorityValueBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueBoolean")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "booleanValue", m.GetBooleanValue(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueBoolean")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueBoolean) IsBACnetPriorityValueBoolean() {}

func (m *_BACnetPriorityValueBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
