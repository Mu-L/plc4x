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

// CALReply is the corresponding interface of CALReply
type CALReply interface {
	CALReplyContract
	CALReplyRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsCALReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALReply()
}

// CALReplyContract provides a set of functions which can be overwritten by a sub struct
type CALReplyContract interface {
	// GetCalType returns CalType (property field)
	GetCalType() byte
	// GetCalData returns CalData (property field)
	GetCalData() CALData
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
	// GetRequestContext() returns a parser argument
	GetRequestContext() RequestContext
	// IsCALReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALReply()
}

// CALReplyRequirements provides a set of functions which need to be implemented by a sub struct
type CALReplyRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCalType returns CalType (discriminator field)
	GetCalType() byte
}

// _CALReply is the data-structure of this message
type _CALReply struct {
	_SubType CALReply
	CalType  byte
	CalData  CALData

	// Arguments.
	CBusOptions    CBusOptions
	RequestContext RequestContext
}

var _ CALReplyContract = (*_CALReply)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALReply) GetCalType() byte {
	return m.CalType
}

func (m *_CALReply) GetCalData() CALData {
	return m.CalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALReply factory function for _CALReply
func NewCALReply(calType byte, calData CALData, cBusOptions CBusOptions, requestContext RequestContext) *_CALReply {
	return &_CALReply{CalType: calType, CalData: calData, CBusOptions: cBusOptions, RequestContext: requestContext}
}

// Deprecated: use the interface for direct cast
func CastCALReply(structType any) CALReply {
	if casted, ok := structType.(CALReply); ok {
		return casted
	}
	if casted, ok := structType.(*CALReply); ok {
		return *casted
	}
	return nil
}

func (m *_CALReply) GetTypeName() string {
	return "CALReply"
}

func (m *_CALReply) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (calData)
	lengthInBits += m.CalData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CALReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CALReplyParse[T CALReply](ctx context.Context, theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (T, error) {
	return CALReplyParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func CALReplyParseWithBufferProducer[T CALReply](cBusOptions CBusOptions, requestContext RequestContext) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CALReplyParseWithBuffer[T](ctx, readBuffer, cBusOptions, requestContext)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func CALReplyParseWithBuffer[T CALReply](ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (T, error) {
	v, err := (&_CALReply{CBusOptions: cBusOptions, RequestContext: requestContext}).parse(ctx, readBuffer, cBusOptions, requestContext)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_CALReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (__cALReply CALReply, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	calType, err := ReadPeekField[byte](ctx, "calType", ReadByte(readBuffer, 8), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'calType' field"))
	}
	m.CalType = calType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CALReply
	switch {
	case calType == 0x86: // CALReplyLong
		if _child, err = (&_CALReplyLong{}).parse(ctx, readBuffer, m, cBusOptions, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALReplyLong for type-switch of CALReply")
		}
	case true: // CALReplyShort
		if _child, err = (&_CALReplyShort{}).parse(ctx, readBuffer, m, cBusOptions, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CALReplyShort for type-switch of CALReply")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [calType=%v]", calType)
	}

	calData, err := ReadSimpleField[CALData](ctx, "calData", ReadComplex[CALData](CALDataParseWithBufferProducer[CALData]((RequestContext)(requestContext)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'calData' field"))
	}
	m.CalData = calData

	if closeErr := readBuffer.CloseContext("CALReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALReply")
	}

	return _child, nil
}

func (pm *_CALReply) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CALReply, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CALReply"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CALReply")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[CALData](ctx, "calData", m.GetCalData(), WriteComplex[CALData](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'calData' field")
	}

	if popErr := writeBuffer.PopContext("CALReply"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CALReply")
	}
	return nil
}

////
// Arguments Getter

func (m *_CALReply) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}
func (m *_CALReply) GetRequestContext() RequestContext {
	return m.RequestContext
}

//
////

func (m *_CALReply) IsCALReply() {}
