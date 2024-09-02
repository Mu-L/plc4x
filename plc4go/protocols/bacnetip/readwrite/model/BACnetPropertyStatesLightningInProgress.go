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

// BACnetPropertyStatesLightningInProgress is the corresponding interface of BACnetPropertyStatesLightningInProgress
type BACnetPropertyStatesLightningInProgress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLightningInProgress returns LightningInProgress (property field)
	GetLightningInProgress() BACnetLightingInProgressTagged
	// IsBACnetPropertyStatesLightningInProgress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLightningInProgress()
}

// _BACnetPropertyStatesLightningInProgress is the data-structure of this message
type _BACnetPropertyStatesLightningInProgress struct {
	BACnetPropertyStatesContract
	LightningInProgress BACnetLightingInProgressTagged
}

var _ BACnetPropertyStatesLightningInProgress = (*_BACnetPropertyStatesLightningInProgress)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLightningInProgress)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLightningInProgress) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLightningInProgress) GetLightningInProgress() BACnetLightingInProgressTagged {
	return m.LightningInProgress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLightningInProgress factory function for _BACnetPropertyStatesLightningInProgress
func NewBACnetPropertyStatesLightningInProgress(lightningInProgress BACnetLightingInProgressTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLightningInProgress {
	_result := &_BACnetPropertyStatesLightningInProgress{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LightningInProgress:          lightningInProgress,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLightningInProgress(structType any) BACnetPropertyStatesLightningInProgress {
	if casted, ok := structType.(BACnetPropertyStatesLightningInProgress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLightningInProgress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLightningInProgress) GetTypeName() string {
	return "BACnetPropertyStatesLightningInProgress"
}

func (m *_BACnetPropertyStatesLightningInProgress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lightningInProgress)
	lengthInBits += m.LightningInProgress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLightningInProgress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLightningInProgress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLightningInProgress BACnetPropertyStatesLightningInProgress, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLightningInProgress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLightningInProgress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightningInProgress, err := ReadSimpleField[BACnetLightingInProgressTagged](ctx, "lightningInProgress", ReadComplex[BACnetLightingInProgressTagged](BACnetLightingInProgressTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightningInProgress' field"))
	}
	m.LightningInProgress = lightningInProgress

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLightningInProgress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLightningInProgress")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLightningInProgress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLightningInProgress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLightningInProgress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLightningInProgress")
		}

		if err := WriteSimpleField[BACnetLightingInProgressTagged](ctx, "lightningInProgress", m.GetLightningInProgress(), WriteComplex[BACnetLightingInProgressTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightningInProgress' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLightningInProgress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLightningInProgress")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLightningInProgress) IsBACnetPropertyStatesLightningInProgress() {}

func (m *_BACnetPropertyStatesLightningInProgress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
