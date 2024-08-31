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

// AdsReadStateResponse is the corresponding interface of AdsReadStateResponse
type AdsReadStateResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetAdsState returns AdsState (property field)
	GetAdsState() uint16
	// GetDeviceState returns DeviceState (property field)
	GetDeviceState() uint16
}

// AdsReadStateResponseExactly can be used when we want exactly this type and not a type which fulfills AdsReadStateResponse.
// This is useful for switch cases.
type AdsReadStateResponseExactly interface {
	AdsReadStateResponse
	isAdsReadStateResponse() bool
}

// _AdsReadStateResponse is the data-structure of this message
type _AdsReadStateResponse struct {
	*_AmsPacket
	Result      ReturnCode
	AdsState    uint16
	DeviceState uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadStateResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ_STATE
}

func (m *_AdsReadStateResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadStateResponse) InitializeParent(parent AmsPacket, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) {
	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsReadStateResponse) GetParent() AmsPacket {
	return m._AmsPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadStateResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsReadStateResponse) GetAdsState() uint16 {
	return m.AdsState
}

func (m *_AdsReadStateResponse) GetDeviceState() uint16 {
	return m.DeviceState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadStateResponse factory function for _AdsReadStateResponse
func NewAdsReadStateResponse(result ReturnCode, adsState uint16, deviceState uint16, targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsReadStateResponse {
	_result := &_AdsReadStateResponse{
		Result:      result,
		AdsState:    adsState,
		DeviceState: deviceState,
		_AmsPacket:  NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadStateResponse(structType any) AdsReadStateResponse {
	if casted, ok := structType.(AdsReadStateResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadStateResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadStateResponse) GetTypeName() string {
	return "AdsReadStateResponse"
}

func (m *_AdsReadStateResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (result)
	lengthInBits += 32

	// Simple field (adsState)
	lengthInBits += 16

	// Simple field (deviceState)
	lengthInBits += 16

	return lengthInBits
}

func (m *_AdsReadStateResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsReadStateResponseParse(ctx context.Context, theBytes []byte) (AdsReadStateResponse, error) {
	return AdsReadStateResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsReadStateResponseParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsReadStateResponse, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsReadStateResponse, error) {
		return AdsReadStateResponseParseWithBuffer(ctx, readBuffer)
	}
}

func AdsReadStateResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsReadStateResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdsReadStateResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadStateResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	result, err := ReadEnumField[ReturnCode](ctx, "result", "ReturnCode", ReadEnum(ReturnCodeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'result' field"))
	}

	adsState, err := ReadSimpleField(ctx, "adsState", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adsState' field"))
	}

	deviceState, err := ReadSimpleField(ctx, "deviceState", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceState' field"))
	}

	if closeErr := readBuffer.CloseContext("AdsReadStateResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadStateResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsReadStateResponse{
		_AmsPacket:  &_AmsPacket{},
		Result:      result,
		AdsState:    adsState,
		DeviceState: deviceState,
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsReadStateResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadStateResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadStateResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadStateResponse")
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for result")
		}
		_resultErr := writeBuffer.WriteSerializable(ctx, m.GetResult())
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for result")
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (adsState)
		adsState := uint16(m.GetAdsState())
		_adsStateErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("adsState", 16, uint16((adsState)))
		if _adsStateErr != nil {
			return errors.Wrap(_adsStateErr, "Error serializing 'adsState' field")
		}

		// Simple Field (deviceState)
		deviceState := uint16(m.GetDeviceState())
		_deviceStateErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("deviceState", 16, uint16((deviceState)))
		if _deviceStateErr != nil {
			return errors.Wrap(_deviceStateErr, "Error serializing 'deviceState' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadStateResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadStateResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsReadStateResponse) isAdsReadStateResponse() bool {
	return true
}

func (m *_AdsReadStateResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
