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

// BACnetServiceAckGetAlarmSummary is the corresponding interface of BACnetServiceAckGetAlarmSummary
type BACnetServiceAckGetAlarmSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetEventState returns EventState (property field)
	GetEventState() BACnetEventStateTagged
	// GetAcknowledgedTransitions returns AcknowledgedTransitions (property field)
	GetAcknowledgedTransitions() BACnetEventTransitionBitsTagged
}

// BACnetServiceAckGetAlarmSummaryExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckGetAlarmSummary.
// This is useful for switch cases.
type BACnetServiceAckGetAlarmSummaryExactly interface {
	BACnetServiceAckGetAlarmSummary
	isBACnetServiceAckGetAlarmSummary() bool
}

// _BACnetServiceAckGetAlarmSummary is the data-structure of this message
type _BACnetServiceAckGetAlarmSummary struct {
	BACnetServiceAckContract
	ObjectIdentifier        BACnetApplicationTagObjectIdentifier
	EventState              BACnetEventStateTagged
	AcknowledgedTransitions BACnetEventTransitionBitsTagged
}

var _ BACnetServiceAckGetAlarmSummary = (*_BACnetServiceAckGetAlarmSummary)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckGetAlarmSummary)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckGetAlarmSummary) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckGetAlarmSummary) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckGetAlarmSummary) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetServiceAckGetAlarmSummary) GetEventState() BACnetEventStateTagged {
	return m.EventState
}

func (m *_BACnetServiceAckGetAlarmSummary) GetAcknowledgedTransitions() BACnetEventTransitionBitsTagged {
	return m.AcknowledgedTransitions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckGetAlarmSummary factory function for _BACnetServiceAckGetAlarmSummary
func NewBACnetServiceAckGetAlarmSummary(objectIdentifier BACnetApplicationTagObjectIdentifier, eventState BACnetEventStateTagged, acknowledgedTransitions BACnetEventTransitionBitsTagged, serviceAckLength uint32) *_BACnetServiceAckGetAlarmSummary {
	_result := &_BACnetServiceAckGetAlarmSummary{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		ObjectIdentifier:         objectIdentifier,
		EventState:               eventState,
		AcknowledgedTransitions:  acknowledgedTransitions,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckGetAlarmSummary(structType any) BACnetServiceAckGetAlarmSummary {
	if casted, ok := structType.(BACnetServiceAckGetAlarmSummary); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckGetAlarmSummary); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckGetAlarmSummary) GetTypeName() string {
	return "BACnetServiceAckGetAlarmSummary"
}

func (m *_BACnetServiceAckGetAlarmSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits(ctx)

	// Simple field (acknowledgedTransitions)
	lengthInBits += m.AcknowledgedTransitions.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckGetAlarmSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckGetAlarmSummary) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckGetAlarmSummary BACnetServiceAckGetAlarmSummary, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckGetAlarmSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckGetAlarmSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	eventState, err := ReadSimpleField[BACnetEventStateTagged](ctx, "eventState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventState' field"))
	}
	m.EventState = eventState

	acknowledgedTransitions, err := ReadSimpleField[BACnetEventTransitionBitsTagged](ctx, "acknowledgedTransitions", ReadComplex[BACnetEventTransitionBitsTagged](BACnetEventTransitionBitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'acknowledgedTransitions' field"))
	}
	m.AcknowledgedTransitions = acknowledgedTransitions

	if closeErr := readBuffer.CloseContext("BACnetServiceAckGetAlarmSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckGetAlarmSummary")
	}

	return m, nil
}

func (m *_BACnetServiceAckGetAlarmSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckGetAlarmSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckGetAlarmSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckGetAlarmSummary")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetEventStateTagged](ctx, "eventState", m.GetEventState(), WriteComplex[BACnetEventStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventState' field")
		}

		if err := WriteSimpleField[BACnetEventTransitionBitsTagged](ctx, "acknowledgedTransitions", m.GetAcknowledgedTransitions(), WriteComplex[BACnetEventTransitionBitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'acknowledgedTransitions' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckGetAlarmSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckGetAlarmSummary")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckGetAlarmSummary) isBACnetServiceAckGetAlarmSummary() bool {
	return true
}

func (m *_BACnetServiceAckGetAlarmSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
