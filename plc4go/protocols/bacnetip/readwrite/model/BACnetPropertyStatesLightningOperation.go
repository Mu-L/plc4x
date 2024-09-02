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

// BACnetPropertyStatesLightningOperation is the corresponding interface of BACnetPropertyStatesLightningOperation
type BACnetPropertyStatesLightningOperation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLightningOperation returns LightningOperation (property field)
	GetLightningOperation() BACnetLightingOperationTagged
}

// BACnetPropertyStatesLightningOperationExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLightningOperation.
// This is useful for switch cases.
type BACnetPropertyStatesLightningOperationExactly interface {
	BACnetPropertyStatesLightningOperation
	isBACnetPropertyStatesLightningOperation() bool
}

// _BACnetPropertyStatesLightningOperation is the data-structure of this message
type _BACnetPropertyStatesLightningOperation struct {
	BACnetPropertyStatesContract
	LightningOperation BACnetLightingOperationTagged
}

var _ BACnetPropertyStatesLightningOperation = (*_BACnetPropertyStatesLightningOperation)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLightningOperation)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLightningOperation) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLightningOperation) GetLightningOperation() BACnetLightingOperationTagged {
	return m.LightningOperation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLightningOperation factory function for _BACnetPropertyStatesLightningOperation
func NewBACnetPropertyStatesLightningOperation(lightningOperation BACnetLightingOperationTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLightningOperation {
	_result := &_BACnetPropertyStatesLightningOperation{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LightningOperation:           lightningOperation,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLightningOperation(structType any) BACnetPropertyStatesLightningOperation {
	if casted, ok := structType.(BACnetPropertyStatesLightningOperation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLightningOperation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLightningOperation) GetTypeName() string {
	return "BACnetPropertyStatesLightningOperation"
}

func (m *_BACnetPropertyStatesLightningOperation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lightningOperation)
	lengthInBits += m.LightningOperation.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLightningOperation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLightningOperation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLightningOperation BACnetPropertyStatesLightningOperation, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLightningOperation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLightningOperation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightningOperation, err := ReadSimpleField[BACnetLightingOperationTagged](ctx, "lightningOperation", ReadComplex[BACnetLightingOperationTagged](BACnetLightingOperationTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightningOperation' field"))
	}
	m.LightningOperation = lightningOperation

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLightningOperation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLightningOperation")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLightningOperation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLightningOperation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLightningOperation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLightningOperation")
		}

		if err := WriteSimpleField[BACnetLightingOperationTagged](ctx, "lightningOperation", m.GetLightningOperation(), WriteComplex[BACnetLightingOperationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightningOperation' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLightningOperation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLightningOperation")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLightningOperation) isBACnetPropertyStatesLightningOperation() bool {
	return true
}

func (m *_BACnetPropertyStatesLightningOperation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
