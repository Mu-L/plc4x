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

// BACnetUnconfirmedServiceRequestTimeSynchronization is the corresponding interface of BACnetUnconfirmedServiceRequestTimeSynchronization
type BACnetUnconfirmedServiceRequestTimeSynchronization interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetSynchronizedDate returns SynchronizedDate (property field)
	GetSynchronizedDate() BACnetApplicationTagDate
	// GetSynchronizedTime returns SynchronizedTime (property field)
	GetSynchronizedTime() BACnetApplicationTagTime
}

// BACnetUnconfirmedServiceRequestTimeSynchronizationExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestTimeSynchronization.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestTimeSynchronizationExactly interface {
	BACnetUnconfirmedServiceRequestTimeSynchronization
	isBACnetUnconfirmedServiceRequestTimeSynchronization() bool
}

// _BACnetUnconfirmedServiceRequestTimeSynchronization is the data-structure of this message
type _BACnetUnconfirmedServiceRequestTimeSynchronization struct {
	BACnetUnconfirmedServiceRequestContract
	SynchronizedDate BACnetApplicationTagDate
	SynchronizedTime BACnetApplicationTagTime
}

var _ BACnetUnconfirmedServiceRequestTimeSynchronization = (*_BACnetUnconfirmedServiceRequestTimeSynchronization)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestTimeSynchronization)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetSynchronizedDate() BACnetApplicationTagDate {
	return m.SynchronizedDate
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetSynchronizedTime() BACnetApplicationTagTime {
	return m.SynchronizedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestTimeSynchronization factory function for _BACnetUnconfirmedServiceRequestTimeSynchronization
func NewBACnetUnconfirmedServiceRequestTimeSynchronization(synchronizedDate BACnetApplicationTagDate, synchronizedTime BACnetApplicationTagTime, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestTimeSynchronization {
	_result := &_BACnetUnconfirmedServiceRequestTimeSynchronization{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		SynchronizedDate:                        synchronizedDate,
		SynchronizedTime:                        synchronizedTime,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestTimeSynchronization(structType any) BACnetUnconfirmedServiceRequestTimeSynchronization {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestTimeSynchronization); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestTimeSynchronization); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestTimeSynchronization"
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (synchronizedDate)
	lengthInBits += m.SynchronizedDate.GetLengthInBits(ctx)

	// Simple field (synchronizedTime)
	lengthInBits += m.SynchronizedTime.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestTimeSynchronization BACnetUnconfirmedServiceRequestTimeSynchronization, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestTimeSynchronization")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	synchronizedDate, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "synchronizedDate", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'synchronizedDate' field"))
	}
	m.SynchronizedDate = synchronizedDate

	synchronizedTime, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "synchronizedTime", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'synchronizedTime' field"))
	}
	m.SynchronizedTime = synchronizedTime

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestTimeSynchronization")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestTimeSynchronization")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "synchronizedDate", m.GetSynchronizedDate(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'synchronizedDate' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "synchronizedTime", m.GetSynchronizedTime(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'synchronizedTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestTimeSynchronization"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestTimeSynchronization")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) isBACnetUnconfirmedServiceRequestTimeSynchronization() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestTimeSynchronization) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
