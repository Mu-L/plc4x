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

// BACnetPropertyStatesLiftCarDriveStatus is the corresponding interface of BACnetPropertyStatesLiftCarDriveStatus
type BACnetPropertyStatesLiftCarDriveStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLiftCarDriveStatus returns LiftCarDriveStatus (property field)
	GetLiftCarDriveStatus() BACnetLiftCarDriveStatusTagged
	// IsBACnetPropertyStatesLiftCarDriveStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLiftCarDriveStatus()
}

// _BACnetPropertyStatesLiftCarDriveStatus is the data-structure of this message
type _BACnetPropertyStatesLiftCarDriveStatus struct {
	BACnetPropertyStatesContract
	LiftCarDriveStatus BACnetLiftCarDriveStatusTagged
}

var _ BACnetPropertyStatesLiftCarDriveStatus = (*_BACnetPropertyStatesLiftCarDriveStatus)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLiftCarDriveStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetLiftCarDriveStatus() BACnetLiftCarDriveStatusTagged {
	return m.LiftCarDriveStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLiftCarDriveStatus factory function for _BACnetPropertyStatesLiftCarDriveStatus
func NewBACnetPropertyStatesLiftCarDriveStatus(liftCarDriveStatus BACnetLiftCarDriveStatusTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesLiftCarDriveStatus {
	_result := &_BACnetPropertyStatesLiftCarDriveStatus{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LiftCarDriveStatus:           liftCarDriveStatus,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftCarDriveStatus(structType any) BACnetPropertyStatesLiftCarDriveStatus {
	if casted, ok := structType.(BACnetPropertyStatesLiftCarDriveStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftCarDriveStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetTypeName() string {
	return "BACnetPropertyStatesLiftCarDriveStatus"
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (liftCarDriveStatus)
	lengthInBits += m.LiftCarDriveStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLiftCarDriveStatus BACnetPropertyStatesLiftCarDriveStatus, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftCarDriveStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftCarDriveStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	liftCarDriveStatus, err := ReadSimpleField[BACnetLiftCarDriveStatusTagged](ctx, "liftCarDriveStatus", ReadComplex[BACnetLiftCarDriveStatusTagged](BACnetLiftCarDriveStatusTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'liftCarDriveStatus' field"))
	}
	m.LiftCarDriveStatus = liftCarDriveStatus

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftCarDriveStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftCarDriveStatus")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftCarDriveStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftCarDriveStatus")
		}

		if err := WriteSimpleField[BACnetLiftCarDriveStatusTagged](ctx, "liftCarDriveStatus", m.GetLiftCarDriveStatus(), WriteComplex[BACnetLiftCarDriveStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'liftCarDriveStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftCarDriveStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftCarDriveStatus")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) IsBACnetPropertyStatesLiftCarDriveStatus() {}

func (m *_BACnetPropertyStatesLiftCarDriveStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
