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

// CipConnectionManagerRequest is the corresponding interface of CipConnectionManagerRequest
type CipConnectionManagerRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() PathSegment
	// GetInstanceSegment returns InstanceSegment (property field)
	GetInstanceSegment() PathSegment
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// GetTickTime returns TickTime (property field)
	GetTickTime() uint8
	// GetTimeoutTicks returns TimeoutTicks (property field)
	GetTimeoutTicks() uint8
	// GetOtConnectionId returns OtConnectionId (property field)
	GetOtConnectionId() uint32
	// GetToConnectionId returns ToConnectionId (property field)
	GetToConnectionId() uint32
	// GetConnectionSerialNumber returns ConnectionSerialNumber (property field)
	GetConnectionSerialNumber() uint16
	// GetOriginatorVendorId returns OriginatorVendorId (property field)
	GetOriginatorVendorId() uint16
	// GetOriginatorSerialNumber returns OriginatorSerialNumber (property field)
	GetOriginatorSerialNumber() uint32
	// GetTimeoutMultiplier returns TimeoutMultiplier (property field)
	GetTimeoutMultiplier() uint8
	// GetOtRpi returns OtRpi (property field)
	GetOtRpi() uint32
	// GetOtConnectionParameters returns OtConnectionParameters (property field)
	GetOtConnectionParameters() NetworkConnectionParameters
	// GetToRpi returns ToRpi (property field)
	GetToRpi() uint32
	// GetToConnectionParameters returns ToConnectionParameters (property field)
	GetToConnectionParameters() NetworkConnectionParameters
	// GetTransportType returns TransportType (property field)
	GetTransportType() TransportType
	// GetConnectionPathSize returns ConnectionPathSize (property field)
	GetConnectionPathSize() uint8
	// GetConnectionPaths returns ConnectionPaths (property field)
	GetConnectionPaths() []PathSegment
}

// CipConnectionManagerRequestExactly can be used when we want exactly this type and not a type which fulfills CipConnectionManagerRequest.
// This is useful for switch cases.
type CipConnectionManagerRequestExactly interface {
	CipConnectionManagerRequest
	isCipConnectionManagerRequest() bool
}

// _CipConnectionManagerRequest is the data-structure of this message
type _CipConnectionManagerRequest struct {
	*_CipService
	ClassSegment           PathSegment
	InstanceSegment        PathSegment
	Priority               uint8
	TickTime               uint8
	TimeoutTicks           uint8
	OtConnectionId         uint32
	ToConnectionId         uint32
	ConnectionSerialNumber uint16
	OriginatorVendorId     uint16
	OriginatorSerialNumber uint32
	TimeoutMultiplier      uint8
	OtRpi                  uint32
	OtConnectionParameters NetworkConnectionParameters
	ToRpi                  uint32
	ToConnectionParameters NetworkConnectionParameters
	TransportType          TransportType
	ConnectionPathSize     uint8
	ConnectionPaths        []PathSegment
	// Reserved Fields
	reservedField0 *uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectionManagerRequest) GetService() uint8 {
	return 0x5B
}

func (m *_CipConnectionManagerRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipConnectionManagerRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectionManagerRequest) InitializeParent(parent CipService) {}

func (m *_CipConnectionManagerRequest) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectionManagerRequest) GetClassSegment() PathSegment {
	return m.ClassSegment
}

func (m *_CipConnectionManagerRequest) GetInstanceSegment() PathSegment {
	return m.InstanceSegment
}

func (m *_CipConnectionManagerRequest) GetPriority() uint8 {
	return m.Priority
}

func (m *_CipConnectionManagerRequest) GetTickTime() uint8 {
	return m.TickTime
}

func (m *_CipConnectionManagerRequest) GetTimeoutTicks() uint8 {
	return m.TimeoutTicks
}

func (m *_CipConnectionManagerRequest) GetOtConnectionId() uint32 {
	return m.OtConnectionId
}

func (m *_CipConnectionManagerRequest) GetToConnectionId() uint32 {
	return m.ToConnectionId
}

func (m *_CipConnectionManagerRequest) GetConnectionSerialNumber() uint16 {
	return m.ConnectionSerialNumber
}

func (m *_CipConnectionManagerRequest) GetOriginatorVendorId() uint16 {
	return m.OriginatorVendorId
}

func (m *_CipConnectionManagerRequest) GetOriginatorSerialNumber() uint32 {
	return m.OriginatorSerialNumber
}

func (m *_CipConnectionManagerRequest) GetTimeoutMultiplier() uint8 {
	return m.TimeoutMultiplier
}

func (m *_CipConnectionManagerRequest) GetOtRpi() uint32 {
	return m.OtRpi
}

func (m *_CipConnectionManagerRequest) GetOtConnectionParameters() NetworkConnectionParameters {
	return m.OtConnectionParameters
}

func (m *_CipConnectionManagerRequest) GetToRpi() uint32 {
	return m.ToRpi
}

func (m *_CipConnectionManagerRequest) GetToConnectionParameters() NetworkConnectionParameters {
	return m.ToConnectionParameters
}

func (m *_CipConnectionManagerRequest) GetTransportType() TransportType {
	return m.TransportType
}

func (m *_CipConnectionManagerRequest) GetConnectionPathSize() uint8 {
	return m.ConnectionPathSize
}

func (m *_CipConnectionManagerRequest) GetConnectionPaths() []PathSegment {
	return m.ConnectionPaths
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipConnectionManagerRequest factory function for _CipConnectionManagerRequest
func NewCipConnectionManagerRequest(classSegment PathSegment, instanceSegment PathSegment, priority uint8, tickTime uint8, timeoutTicks uint8, otConnectionId uint32, toConnectionId uint32, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, timeoutMultiplier uint8, otRpi uint32, otConnectionParameters NetworkConnectionParameters, toRpi uint32, toConnectionParameters NetworkConnectionParameters, transportType TransportType, connectionPathSize uint8, connectionPaths []PathSegment, serviceLen uint16) *_CipConnectionManagerRequest {
	_result := &_CipConnectionManagerRequest{
		ClassSegment:           classSegment,
		InstanceSegment:        instanceSegment,
		Priority:               priority,
		TickTime:               tickTime,
		TimeoutTicks:           timeoutTicks,
		OtConnectionId:         otConnectionId,
		ToConnectionId:         toConnectionId,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		TimeoutMultiplier:      timeoutMultiplier,
		OtRpi:                  otRpi,
		OtConnectionParameters: otConnectionParameters,
		ToRpi:                  toRpi,
		ToConnectionParameters: toConnectionParameters,
		TransportType:          transportType,
		ConnectionPathSize:     connectionPathSize,
		ConnectionPaths:        connectionPaths,
		_CipService:            NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipConnectionManagerRequest(structType any) CipConnectionManagerRequest {
	if casted, ok := structType.(CipConnectionManagerRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectionManagerRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectionManagerRequest) GetTypeName() string {
	return "CipConnectionManagerRequest"
}

func (m *_CipConnectionManagerRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Simple field (classSegment)
	lengthInBits += m.ClassSegment.GetLengthInBits(ctx)

	// Simple field (instanceSegment)
	lengthInBits += m.InstanceSegment.GetLengthInBits(ctx)

	// Simple field (priority)
	lengthInBits += 4

	// Simple field (tickTime)
	lengthInBits += 4

	// Simple field (timeoutTicks)
	lengthInBits += 8

	// Simple field (otConnectionId)
	lengthInBits += 32

	// Simple field (toConnectionId)
	lengthInBits += 32

	// Simple field (connectionSerialNumber)
	lengthInBits += 16

	// Simple field (originatorVendorId)
	lengthInBits += 16

	// Simple field (originatorSerialNumber)
	lengthInBits += 32

	// Simple field (timeoutMultiplier)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 24

	// Simple field (otRpi)
	lengthInBits += 32

	// Simple field (otConnectionParameters)
	lengthInBits += m.OtConnectionParameters.GetLengthInBits(ctx)

	// Simple field (toRpi)
	lengthInBits += 32

	// Simple field (toConnectionParameters)
	lengthInBits += m.ToConnectionParameters.GetLengthInBits(ctx)

	// Simple field (transportType)
	lengthInBits += m.TransportType.GetLengthInBits(ctx)

	// Simple field (connectionPathSize)
	lengthInBits += 8

	// Array field
	if len(m.ConnectionPaths) > 0 {
		for _, element := range m.ConnectionPaths {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_CipConnectionManagerRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipConnectionManagerRequestParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (CipConnectionManagerRequest, error) {
	return CipConnectionManagerRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipConnectionManagerRequestParseWithBufferProducer(connected bool, serviceLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (CipConnectionManagerRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CipConnectionManagerRequest, error) {
		return CipConnectionManagerRequestParseWithBuffer(ctx, readBuffer, connected, serviceLen)
	}
}

func CipConnectionManagerRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (CipConnectionManagerRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CipConnectionManagerRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectionManagerRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadImplicitField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	classSegment, err := ReadSimpleField[PathSegment](ctx, "classSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'classSegment' field"))
	}

	instanceSegment, err := ReadSimpleField[PathSegment](ctx, "instanceSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instanceSegment' field"))
	}

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}

	tickTime, err := ReadSimpleField(ctx, "tickTime", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tickTime' field"))
	}

	timeoutTicks, err := ReadSimpleField(ctx, "timeoutTicks", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeoutTicks' field"))
	}

	otConnectionId, err := ReadSimpleField(ctx, "otConnectionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'otConnectionId' field"))
	}

	toConnectionId, err := ReadSimpleField(ctx, "toConnectionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toConnectionId' field"))
	}

	connectionSerialNumber, err := ReadSimpleField(ctx, "connectionSerialNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSerialNumber' field"))
	}

	originatorVendorId, err := ReadSimpleField(ctx, "originatorVendorId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorVendorId' field"))
	}

	originatorSerialNumber, err := ReadSimpleField(ctx, "originatorSerialNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorSerialNumber' field"))
	}

	timeoutMultiplier, err := ReadSimpleField(ctx, "timeoutMultiplier", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeoutMultiplier' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedInt(readBuffer, uint8(24)), uint32(0x000000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}

	otRpi, err := ReadSimpleField(ctx, "otRpi", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'otRpi' field"))
	}

	otConnectionParameters, err := ReadSimpleField[NetworkConnectionParameters](ctx, "otConnectionParameters", ReadComplex[NetworkConnectionParameters](NetworkConnectionParametersParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'otConnectionParameters' field"))
	}

	toRpi, err := ReadSimpleField(ctx, "toRpi", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toRpi' field"))
	}

	toConnectionParameters, err := ReadSimpleField[NetworkConnectionParameters](ctx, "toConnectionParameters", ReadComplex[NetworkConnectionParameters](NetworkConnectionParametersParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toConnectionParameters' field"))
	}

	transportType, err := ReadSimpleField[TransportType](ctx, "transportType", ReadComplex[TransportType](TransportTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportType' field"))
	}

	connectionPathSize, err := ReadSimpleField(ctx, "connectionPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionPathSize' field"))
	}

	connectionPaths, err := ReadTerminatedArrayField[PathSegment](ctx, "connectionPaths", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer), NoMorePathSegments(ctx, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionPaths' field"))
	}

	if closeErr := readBuffer.CloseContext("CipConnectionManagerRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectionManagerRequest")
	}

	// Create a partially initialized instance
	_child := &_CipConnectionManagerRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		ClassSegment:           classSegment,
		InstanceSegment:        instanceSegment,
		Priority:               priority,
		TickTime:               tickTime,
		TimeoutTicks:           timeoutTicks,
		OtConnectionId:         otConnectionId,
		ToConnectionId:         toConnectionId,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		TimeoutMultiplier:      timeoutMultiplier,
		OtRpi:                  otRpi,
		OtConnectionParameters: otConnectionParameters,
		ToRpi:                  toRpi,
		ToConnectionParameters: toConnectionParameters,
		TransportType:          transportType,
		ConnectionPathSize:     connectionPathSize,
		ConnectionPaths:        connectionPaths,
		reservedField0:         reservedField0,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipConnectionManagerRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectionManagerRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectionManagerRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectionManagerRequest")
		}

		// Implicit Field (requestPathSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		requestPathSize := uint8(uint8((uint8(m.GetClassSegment().GetLengthInBytes(ctx)) + uint8(m.GetInstanceSegment().GetLengthInBytes(ctx)))) / uint8(uint8(2)))
		_requestPathSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("requestPathSize", 8, uint8((requestPathSize)))
		if _requestPathSizeErr != nil {
			return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
		}

		// Simple Field (classSegment)
		if pushErr := writeBuffer.PushContext("classSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for classSegment")
		}
		_classSegmentErr := writeBuffer.WriteSerializable(ctx, m.GetClassSegment())
		if popErr := writeBuffer.PopContext("classSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for classSegment")
		}
		if _classSegmentErr != nil {
			return errors.Wrap(_classSegmentErr, "Error serializing 'classSegment' field")
		}

		// Simple Field (instanceSegment)
		if pushErr := writeBuffer.PushContext("instanceSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for instanceSegment")
		}
		_instanceSegmentErr := writeBuffer.WriteSerializable(ctx, m.GetInstanceSegment())
		if popErr := writeBuffer.PopContext("instanceSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for instanceSegment")
		}
		if _instanceSegmentErr != nil {
			return errors.Wrap(_instanceSegmentErr, "Error serializing 'instanceSegment' field")
		}

		// Simple Field (priority)
		priority := uint8(m.GetPriority())
		_priorityErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("priority", 4, uint8((priority)))
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}

		// Simple Field (tickTime)
		tickTime := uint8(m.GetTickTime())
		_tickTimeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("tickTime", 4, uint8((tickTime)))
		if _tickTimeErr != nil {
			return errors.Wrap(_tickTimeErr, "Error serializing 'tickTime' field")
		}

		// Simple Field (timeoutTicks)
		timeoutTicks := uint8(m.GetTimeoutTicks())
		_timeoutTicksErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("timeoutTicks", 8, uint8((timeoutTicks)))
		if _timeoutTicksErr != nil {
			return errors.Wrap(_timeoutTicksErr, "Error serializing 'timeoutTicks' field")
		}

		// Simple Field (otConnectionId)
		otConnectionId := uint32(m.GetOtConnectionId())
		_otConnectionIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("otConnectionId", 32, uint32((otConnectionId)))
		if _otConnectionIdErr != nil {
			return errors.Wrap(_otConnectionIdErr, "Error serializing 'otConnectionId' field")
		}

		// Simple Field (toConnectionId)
		toConnectionId := uint32(m.GetToConnectionId())
		_toConnectionIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("toConnectionId", 32, uint32((toConnectionId)))
		if _toConnectionIdErr != nil {
			return errors.Wrap(_toConnectionIdErr, "Error serializing 'toConnectionId' field")
		}

		// Simple Field (connectionSerialNumber)
		connectionSerialNumber := uint16(m.GetConnectionSerialNumber())
		_connectionSerialNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("connectionSerialNumber", 16, uint16((connectionSerialNumber)))
		if _connectionSerialNumberErr != nil {
			return errors.Wrap(_connectionSerialNumberErr, "Error serializing 'connectionSerialNumber' field")
		}

		// Simple Field (originatorVendorId)
		originatorVendorId := uint16(m.GetOriginatorVendorId())
		_originatorVendorIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("originatorVendorId", 16, uint16((originatorVendorId)))
		if _originatorVendorIdErr != nil {
			return errors.Wrap(_originatorVendorIdErr, "Error serializing 'originatorVendorId' field")
		}

		// Simple Field (originatorSerialNumber)
		originatorSerialNumber := uint32(m.GetOriginatorSerialNumber())
		_originatorSerialNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("originatorSerialNumber", 32, uint32((originatorSerialNumber)))
		if _originatorSerialNumberErr != nil {
			return errors.Wrap(_originatorSerialNumberErr, "Error serializing 'originatorSerialNumber' field")
		}

		// Simple Field (timeoutMultiplier)
		timeoutMultiplier := uint8(m.GetTimeoutMultiplier())
		_timeoutMultiplierErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("timeoutMultiplier", 8, uint8((timeoutMultiplier)))
		if _timeoutMultiplierErr != nil {
			return errors.Wrap(_timeoutMultiplierErr, "Error serializing 'timeoutMultiplier' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint32 = uint32(0x000000)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint32(0x000000),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint32("reserved", 24, uint32(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (otRpi)
		otRpi := uint32(m.GetOtRpi())
		_otRpiErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("otRpi", 32, uint32((otRpi)))
		if _otRpiErr != nil {
			return errors.Wrap(_otRpiErr, "Error serializing 'otRpi' field")
		}

		// Simple Field (otConnectionParameters)
		if pushErr := writeBuffer.PushContext("otConnectionParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for otConnectionParameters")
		}
		_otConnectionParametersErr := writeBuffer.WriteSerializable(ctx, m.GetOtConnectionParameters())
		if popErr := writeBuffer.PopContext("otConnectionParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for otConnectionParameters")
		}
		if _otConnectionParametersErr != nil {
			return errors.Wrap(_otConnectionParametersErr, "Error serializing 'otConnectionParameters' field")
		}

		// Simple Field (toRpi)
		toRpi := uint32(m.GetToRpi())
		_toRpiErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("toRpi", 32, uint32((toRpi)))
		if _toRpiErr != nil {
			return errors.Wrap(_toRpiErr, "Error serializing 'toRpi' field")
		}

		// Simple Field (toConnectionParameters)
		if pushErr := writeBuffer.PushContext("toConnectionParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for toConnectionParameters")
		}
		_toConnectionParametersErr := writeBuffer.WriteSerializable(ctx, m.GetToConnectionParameters())
		if popErr := writeBuffer.PopContext("toConnectionParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for toConnectionParameters")
		}
		if _toConnectionParametersErr != nil {
			return errors.Wrap(_toConnectionParametersErr, "Error serializing 'toConnectionParameters' field")
		}

		// Simple Field (transportType)
		if pushErr := writeBuffer.PushContext("transportType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transportType")
		}
		_transportTypeErr := writeBuffer.WriteSerializable(ctx, m.GetTransportType())
		if popErr := writeBuffer.PopContext("transportType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transportType")
		}
		if _transportTypeErr != nil {
			return errors.Wrap(_transportTypeErr, "Error serializing 'transportType' field")
		}

		// Simple Field (connectionPathSize)
		connectionPathSize := uint8(m.GetConnectionPathSize())
		_connectionPathSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("connectionPathSize", 8, uint8((connectionPathSize)))
		if _connectionPathSizeErr != nil {
			return errors.Wrap(_connectionPathSizeErr, "Error serializing 'connectionPathSize' field")
		}

		// Array Field (connectionPaths)
		if pushErr := writeBuffer.PushContext("connectionPaths", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for connectionPaths")
		}
		for _curItem, _element := range m.GetConnectionPaths() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetConnectionPaths()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'connectionPaths' field")
			}
		}
		if popErr := writeBuffer.PopContext("connectionPaths", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for connectionPaths")
		}

		if popErr := writeBuffer.PopContext("CipConnectionManagerRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectionManagerRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectionManagerRequest) isCipConnectionManagerRequest() bool {
	return true
}

func (m *_CipConnectionManagerRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
