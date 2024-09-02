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

// BACnetProcessIdSelectionNull is the corresponding interface of BACnetProcessIdSelectionNull
type BACnetProcessIdSelectionNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetProcessIdSelection
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
}

// BACnetProcessIdSelectionNullExactly can be used when we want exactly this type and not a type which fulfills BACnetProcessIdSelectionNull.
// This is useful for switch cases.
type BACnetProcessIdSelectionNullExactly interface {
	BACnetProcessIdSelectionNull
	isBACnetProcessIdSelectionNull() bool
}

// _BACnetProcessIdSelectionNull is the data-structure of this message
type _BACnetProcessIdSelectionNull struct {
	BACnetProcessIdSelectionContract
	NullValue BACnetApplicationTagNull
}

var _ BACnetProcessIdSelectionNull = (*_BACnetProcessIdSelectionNull)(nil)
var _ BACnetProcessIdSelectionRequirements = (*_BACnetProcessIdSelectionNull)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetProcessIdSelectionNull) GetParent() BACnetProcessIdSelectionContract {
	return m.BACnetProcessIdSelectionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProcessIdSelectionNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProcessIdSelectionNull factory function for _BACnetProcessIdSelectionNull
func NewBACnetProcessIdSelectionNull(nullValue BACnetApplicationTagNull, peekedTagHeader BACnetTagHeader) *_BACnetProcessIdSelectionNull {
	_result := &_BACnetProcessIdSelectionNull{
		BACnetProcessIdSelectionContract: NewBACnetProcessIdSelection(peekedTagHeader),
		NullValue:                        nullValue,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetProcessIdSelectionNull(structType any) BACnetProcessIdSelectionNull {
	if casted, ok := structType.(BACnetProcessIdSelectionNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProcessIdSelectionNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProcessIdSelectionNull) GetTypeName() string {
	return "BACnetProcessIdSelectionNull"
}

func (m *_BACnetProcessIdSelectionNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetProcessIdSelectionNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetProcessIdSelectionNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetProcessIdSelection) (__bACnetProcessIdSelectionNull BACnetProcessIdSelectionNull, err error) {
	m.BACnetProcessIdSelectionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProcessIdSelectionNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProcessIdSelectionNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetProcessIdSelectionNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProcessIdSelectionNull")
	}

	return m, nil
}

func (m *_BACnetProcessIdSelectionNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProcessIdSelectionNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetProcessIdSelectionNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetProcessIdSelectionNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetProcessIdSelectionNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetProcessIdSelectionNull")
		}
		return nil
	}
	return m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetProcessIdSelectionNull) isBACnetProcessIdSelectionNull() bool {
	return true
}

func (m *_BACnetProcessIdSelectionNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
