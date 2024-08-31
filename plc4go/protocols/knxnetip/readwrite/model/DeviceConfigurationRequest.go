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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DeviceConfigurationRequest is the corresponding interface of DeviceConfigurationRequest
type DeviceConfigurationRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetDeviceConfigurationRequestDataBlock returns DeviceConfigurationRequestDataBlock (property field)
	GetDeviceConfigurationRequestDataBlock() DeviceConfigurationRequestDataBlock
	// GetCemi returns Cemi (property field)
	GetCemi() CEMI
}

// DeviceConfigurationRequestExactly can be used when we want exactly this type and not a type which fulfills DeviceConfigurationRequest.
// This is useful for switch cases.
type DeviceConfigurationRequestExactly interface {
	DeviceConfigurationRequest
	isDeviceConfigurationRequest() bool
}

// _DeviceConfigurationRequest is the data-structure of this message
type _DeviceConfigurationRequest struct {
	*_KnxNetIpMessage
	DeviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock
	Cemi                                CEMI

	// Arguments.
	TotalLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeviceConfigurationRequest) GetMsgType() uint16 {
	return 0x0310
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeviceConfigurationRequest) InitializeParent(parent KnxNetIpMessage) {}

func (m *_DeviceConfigurationRequest) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationRequest) GetDeviceConfigurationRequestDataBlock() DeviceConfigurationRequestDataBlock {
	return m.DeviceConfigurationRequestDataBlock
}

func (m *_DeviceConfigurationRequest) GetCemi() CEMI {
	return m.Cemi
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeviceConfigurationRequest factory function for _DeviceConfigurationRequest
func NewDeviceConfigurationRequest(deviceConfigurationRequestDataBlock DeviceConfigurationRequestDataBlock, cemi CEMI, totalLength uint16) *_DeviceConfigurationRequest {
	_result := &_DeviceConfigurationRequest{
		DeviceConfigurationRequestDataBlock: deviceConfigurationRequestDataBlock,
		Cemi:                                cemi,
		_KnxNetIpMessage:                    NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationRequest(structType any) DeviceConfigurationRequest {
	if casted, ok := structType.(DeviceConfigurationRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationRequest) GetTypeName() string {
	return "DeviceConfigurationRequest"
}

func (m *_DeviceConfigurationRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (deviceConfigurationRequestDataBlock)
	lengthInBits += m.DeviceConfigurationRequestDataBlock.GetLengthInBits(ctx)

	// Simple field (cemi)
	lengthInBits += m.Cemi.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DeviceConfigurationRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeviceConfigurationRequestParse(ctx context.Context, theBytes []byte, totalLength uint16) (DeviceConfigurationRequest, error) {
	return DeviceConfigurationRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), totalLength)
}

func DeviceConfigurationRequestParseWithBufferProducer(totalLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceConfigurationRequest, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceConfigurationRequest, error) {
		return DeviceConfigurationRequestParseWithBuffer(ctx, readBuffer, totalLength)
	}
}

func DeviceConfigurationRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, totalLength uint16) (DeviceConfigurationRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DeviceConfigurationRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceConfigurationRequestDataBlock, err := ReadSimpleField[DeviceConfigurationRequestDataBlock](ctx, "deviceConfigurationRequestDataBlock", ReadComplex[DeviceConfigurationRequestDataBlock](DeviceConfigurationRequestDataBlockParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceConfigurationRequestDataBlock' field"))
	}

	cemi, err := ReadSimpleField[CEMI](ctx, "cemi", ReadComplex[CEMI](CEMIParseWithBufferProducer[CEMI]((uint16)(uint16(totalLength)-uint16((uint16(uint16(6))+uint16(deviceConfigurationRequestDataBlock.GetLengthInBytes(ctx)))))), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cemi' field"))
	}

	if closeErr := readBuffer.CloseContext("DeviceConfigurationRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationRequest")
	}

	// Create a partially initialized instance
	_child := &_DeviceConfigurationRequest{
		_KnxNetIpMessage:                    &_KnxNetIpMessage{},
		DeviceConfigurationRequestDataBlock: deviceConfigurationRequestDataBlock,
		Cemi:                                cemi,
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_DeviceConfigurationRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceConfigurationRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeviceConfigurationRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationRequest")
		}

		// Simple Field (deviceConfigurationRequestDataBlock)
		if pushErr := writeBuffer.PushContext("deviceConfigurationRequestDataBlock"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceConfigurationRequestDataBlock")
		}
		_deviceConfigurationRequestDataBlockErr := writeBuffer.WriteSerializable(ctx, m.GetDeviceConfigurationRequestDataBlock())
		if popErr := writeBuffer.PopContext("deviceConfigurationRequestDataBlock"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceConfigurationRequestDataBlock")
		}
		if _deviceConfigurationRequestDataBlockErr != nil {
			return errors.Wrap(_deviceConfigurationRequestDataBlockErr, "Error serializing 'deviceConfigurationRequestDataBlock' field")
		}

		// Simple Field (cemi)
		if pushErr := writeBuffer.PushContext("cemi"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for cemi")
		}
		_cemiErr := writeBuffer.WriteSerializable(ctx, m.GetCemi())
		if popErr := writeBuffer.PopContext("cemi"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for cemi")
		}
		if _cemiErr != nil {
			return errors.Wrap(_cemiErr, "Error serializing 'cemi' field")
		}

		if popErr := writeBuffer.PopContext("DeviceConfigurationRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeviceConfigurationRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_DeviceConfigurationRequest) GetTotalLength() uint16 {
	return m.TotalLength
}

//
////

func (m *_DeviceConfigurationRequest) isDeviceConfigurationRequest() bool {
	return true
}

func (m *_DeviceConfigurationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
