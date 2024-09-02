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

// BACnetPropertyStates is the corresponding interface of BACnetPropertyStates
type BACnetPropertyStates interface {
	BACnetPropertyStatesContract
	BACnetPropertyStatesRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetPropertyStatesContract provides a set of functions which can be overwritten by a sub struct
type BACnetPropertyStatesContract interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetPropertyStatesRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetPropertyStatesRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// BACnetPropertyStatesExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStates.
// This is useful for switch cases.
type BACnetPropertyStatesExactly interface {
	BACnetPropertyStates
	isBACnetPropertyStates() bool
}

// _BACnetPropertyStates is the data-structure of this message
type _BACnetPropertyStates struct {
	_SubType        BACnetPropertyStates
	PeekedTagHeader BACnetTagHeader
}

var _ BACnetPropertyStatesContract = (*_BACnetPropertyStates)(nil)

type BACnetPropertyStatesChild interface {
	utils.Serializable

	GetParent() *BACnetPropertyStates

	GetTypeName() string
	BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStates) GetPeekedTagHeader() BACnetTagHeader {
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

func (pm *_BACnetPropertyStates) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStates factory function for _BACnetPropertyStates
func NewBACnetPropertyStates(peekedTagHeader BACnetTagHeader) *_BACnetPropertyStates {
	return &_BACnetPropertyStates{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStates(structType any) BACnetPropertyStates {
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStates) GetTypeName() string {
	return "BACnetPropertyStates"
}

func (m *_BACnetPropertyStates) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPropertyStates) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesParse[T BACnetPropertyStates](ctx context.Context, theBytes []byte) (T, error) {
	return BACnetPropertyStatesParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetPropertyStatesParseWithBufferProducer[T BACnetPropertyStates]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetPropertyStatesParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetPropertyStatesParseWithBuffer[T BACnetPropertyStates](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_BACnetPropertyStates{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetPropertyStates) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetPropertyStates BACnetPropertyStates, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStates"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStates")
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
	var _child BACnetPropertyStates
	switch {
	case peekedTagNumber == uint8(0): // BACnetPropertyStatesBoolean
		if _child, err = (&_BACnetPropertyStatesBoolean{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBoolean for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(1): // BACnetPropertyStatesBinaryValue
		if _child, err = (&_BACnetPropertyStatesBinaryValue{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBinaryValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(2): // BACnetPropertyStatesEventType
		if _child, err = (&_BACnetPropertyStatesEventType{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesEventType for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(3): // BACnetPropertyStatesPolarity
		if _child, err = (&_BACnetPropertyStatesPolarity{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesPolarity for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(4): // BACnetPropertyStatesProgramChange
		if _child, err = (&_BACnetPropertyStatesProgramChange{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesProgramChange for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(5): // BACnetPropertyStatesProgramChange
		if _child, err = (&_BACnetPropertyStatesProgramChange{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesProgramChange for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(6): // BACnetPropertyStatesReasonForHalt
		if _child, err = (&_BACnetPropertyStatesReasonForHalt{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesReasonForHalt for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(7): // BACnetPropertyStatesReliability
		if _child, err = (&_BACnetPropertyStatesReliability{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesReliability for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(8): // BACnetPropertyStatesState
		if _child, err = (&_BACnetPropertyStatesState{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(9): // BACnetPropertyStatesSystemStatus
		if _child, err = (&_BACnetPropertyStatesSystemStatus{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesSystemStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(10): // BACnetPropertyStatesUnits
		if _child, err = (&_BACnetPropertyStatesUnits{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesUnits for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(11): // BACnetPropertyStatesExtendedValue
		if _child, err = (&_BACnetPropertyStatesExtendedValue{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesExtendedValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(12): // BACnetPropertyStatesLifeSafetyMode
		if _child, err = (&_BACnetPropertyStatesLifeSafetyMode{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLifeSafetyMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(13): // BACnetPropertyStatesLifeSafetyState
		if _child, err = (&_BACnetPropertyStatesLifeSafetyState{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLifeSafetyState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(14): // BACnetPropertyStatesRestartReason
		if _child, err = (&_BACnetPropertyStatesRestartReason{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesRestartReason for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(15): // BACnetPropertyStatesDoorAlarmState
		if _child, err = (&_BACnetPropertyStatesDoorAlarmState{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesDoorAlarmState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(16): // BACnetPropertyStatesAction
		if _child, err = (&_BACnetPropertyStatesAction{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAction for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(17): // BACnetPropertyStatesDoorSecuredStatus
		if _child, err = (&_BACnetPropertyStatesDoorSecuredStatus{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesDoorSecuredStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(18): // BACnetPropertyStatesDoorStatus
		if _child, err = (&_BACnetPropertyStatesDoorStatus{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesDoorStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(19): // BACnetPropertyStatesDoorValue
		if _child, err = (&_BACnetPropertyStatesDoorValue{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesDoorValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(20): // BACnetPropertyStatesFileAccessMethod
		if _child, err = (&_BACnetPropertyStatesFileAccessMethod{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesFileAccessMethod for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(21): // BACnetPropertyStatesLockStatus
		if _child, err = (&_BACnetPropertyStatesLockStatus{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLockStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(22): // BACnetPropertyStatesLifeSafetyOperations
		if _child, err = (&_BACnetPropertyStatesLifeSafetyOperations{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLifeSafetyOperations for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(23): // BACnetPropertyStatesMaintenance
		if _child, err = (&_BACnetPropertyStatesMaintenance{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesMaintenance for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(24): // BACnetPropertyStatesNodeType
		if _child, err = (&_BACnetPropertyStatesNodeType{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNodeType for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(25): // BACnetPropertyStatesNotifyType
		if _child, err = (&_BACnetPropertyStatesNotifyType{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNotifyType for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(26): // BACnetPropertyStatesSecurityLevel
		if _child, err = (&_BACnetPropertyStatesSecurityLevel{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesSecurityLevel for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(27): // BACnetPropertyStatesShedState
		if _child, err = (&_BACnetPropertyStatesShedState{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesShedState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(28): // BACnetPropertyStatesSilencedState
		if _child, err = (&_BACnetPropertyStatesSilencedState{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesSilencedState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(30): // BACnetPropertyStatesAccessEvent
		if _child, err = (&_BACnetPropertyStatesAccessEvent{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAccessEvent for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(31): // BACnetPropertyStatesZoneOccupanyState
		if _child, err = (&_BACnetPropertyStatesZoneOccupanyState{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesZoneOccupanyState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(32): // BACnetPropertyStatesAccessCredentialDisableReason
		if _child, err = (&_BACnetPropertyStatesAccessCredentialDisableReason{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAccessCredentialDisableReason for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(33): // BACnetPropertyStatesAccessCredentialDisable
		if _child, err = (&_BACnetPropertyStatesAccessCredentialDisable{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAccessCredentialDisable for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(34): // BACnetPropertyStatesAuthenticationStatus
		if _child, err = (&_BACnetPropertyStatesAuthenticationStatus{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesAuthenticationStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(36): // BACnetPropertyStatesBackupState
		if _child, err = (&_BACnetPropertyStatesBackupState{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBackupState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(37): // BACnetPropertyStatesWriteStatus
		if _child, err = (&_BACnetPropertyStatesWriteStatus{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesWriteStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(38): // BACnetPropertyStatesLightningInProgress
		if _child, err = (&_BACnetPropertyStatesLightningInProgress{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLightningInProgress for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(39): // BACnetPropertyStatesLightningOperation
		if _child, err = (&_BACnetPropertyStatesLightningOperation{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLightningOperation for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(40): // BACnetPropertyStatesLightningTransition
		if _child, err = (&_BACnetPropertyStatesLightningTransition{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLightningTransition for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(41): // BACnetPropertyStatesIntegerValue
		if _child, err = (&_BACnetPropertyStatesIntegerValue{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesIntegerValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(42): // BACnetPropertyStatesBinaryLightningValue
		if _child, err = (&_BACnetPropertyStatesBinaryLightningValue{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBinaryLightningValue for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(43): // BACnetPropertyStatesTimerState
		if _child, err = (&_BACnetPropertyStatesTimerState{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesTimerState for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(44): // BACnetPropertyStatesTimerTransition
		if _child, err = (&_BACnetPropertyStatesTimerTransition{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesTimerTransition for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(45): // BACnetPropertyStatesBacnetIpMode
		if _child, err = (&_BACnetPropertyStatesBacnetIpMode{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesBacnetIpMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(46): // BACnetPropertyStatesNetworkPortCommand
		if _child, err = (&_BACnetPropertyStatesNetworkPortCommand{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNetworkPortCommand for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(47): // BACnetPropertyStatesNetworkType
		if _child, err = (&_BACnetPropertyStatesNetworkType{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNetworkType for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(48): // BACnetPropertyStatesNetworkNumberQuality
		if _child, err = (&_BACnetPropertyStatesNetworkNumberQuality{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesNetworkNumberQuality for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(49): // BACnetPropertyStatesEscalatorOperationDirection
		if _child, err = (&_BACnetPropertyStatesEscalatorOperationDirection{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesEscalatorOperationDirection for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(50): // BACnetPropertyStatesEscalatorFault
		if _child, err = (&_BACnetPropertyStatesEscalatorFault{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesEscalatorFault for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(51): // BACnetPropertyStatesEscalatorMode
		if _child, err = (&_BACnetPropertyStatesEscalatorMode{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesEscalatorMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(52): // BACnetPropertyStatesLiftCarDirection
		if _child, err = (&_BACnetPropertyStatesLiftCarDirection{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftCarDirection for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(53): // BACnetPropertyStatesLiftCarDoorCommand
		if _child, err = (&_BACnetPropertyStatesLiftCarDoorCommand{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftCarDoorCommand for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(54): // BACnetPropertyStatesLiftCarDriveStatus
		if _child, err = (&_BACnetPropertyStatesLiftCarDriveStatus{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftCarDriveStatus for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(55): // BACnetPropertyStatesLiftCarMode
		if _child, err = (&_BACnetPropertyStatesLiftCarMode{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftCarMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(56): // BACnetPropertyStatesLiftGroupMode
		if _child, err = (&_BACnetPropertyStatesLiftGroupMode{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftGroupMode for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(57): // BACnetPropertyStatesLiftFault
		if _child, err = (&_BACnetPropertyStatesLiftFault{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesLiftFault for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(58): // BACnetPropertyStatesProtocolLevel
		if _child, err = (&_BACnetPropertyStatesProtocolLevel{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesProtocolLevel for type-switch of BACnetPropertyStates")
		}
	case peekedTagNumber == uint8(63): // BACnetPropertyStatesExtendedValue
		if _child, err = (&_BACnetPropertyStatesExtendedValue{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStatesExtendedValue for type-switch of BACnetPropertyStates")
		}
	case true: // BACnetPropertyStateActionUnknown
		if _child, err = (&_BACnetPropertyStateActionUnknown{}).parse(ctx, readBuffer, m, peekedTagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetPropertyStateActionUnknown for type-switch of BACnetPropertyStates")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStates"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStates")
	}

	return _child, nil
}

func (pm *_BACnetPropertyStates) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetPropertyStates, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyStates"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStates")
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

	if popErr := writeBuffer.PopContext("BACnetPropertyStates"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyStates")
	}
	return nil
}

func (m *_BACnetPropertyStates) isBACnetPropertyStates() bool {
	return true
}
