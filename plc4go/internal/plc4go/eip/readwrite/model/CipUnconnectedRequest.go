/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CipUnconnectedRequest_ROUTE uint16 = 0x0001

// The data-structure of this message
type CipUnconnectedRequest struct {
	UnconnectedService *CipService
	BackPlane          int8
	Slot               int8
	Parent             *CipService
}

// The corresponding interface
type ICipUnconnectedRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *CipUnconnectedRequest) Service() uint8 {
	return 0x52
}

func (m *CipUnconnectedRequest) InitializeParent(parent *CipService) {
}

func NewCipUnconnectedRequest(unconnectedService *CipService, backPlane int8, slot int8) *CipService {
	child := &CipUnconnectedRequest{
		UnconnectedService: unconnectedService,
		BackPlane:          backPlane,
		Slot:               slot,
		Parent:             NewCipService(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCipUnconnectedRequest(structType interface{}) *CipUnconnectedRequest {
	castFunc := func(typ interface{}) *CipUnconnectedRequest {
		if casted, ok := typ.(CipUnconnectedRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*CipUnconnectedRequest); ok {
			return casted
		}
		if casted, ok := typ.(CipService); ok {
			return CastCipUnconnectedRequest(casted.Child)
		}
		if casted, ok := typ.(*CipService); ok {
			return CastCipUnconnectedRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *CipUnconnectedRequest) GetTypeName() string {
	return "CipUnconnectedRequest"
}

func (m *CipUnconnectedRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *CipUnconnectedRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (messageSize)
	lengthInBits += 16

	// Simple field (unconnectedService)
	lengthInBits += m.UnconnectedService.LengthInBits()

	// Const Field (route)
	lengthInBits += 16

	// Simple field (backPlane)
	lengthInBits += 8

	// Simple field (slot)
	lengthInBits += 8

	return lengthInBits
}

func (m *CipUnconnectedRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func CipUnconnectedRequestParse(readBuffer utils.ReadBuffer) (*CipService, error) {
	if pullErr := readBuffer.PullContext("CipUnconnectedRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x02) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x02),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x20) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x20),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x06) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x06),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x24) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x24),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x01) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x01),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x9D05) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x9D05),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Implicit Field (messageSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	messageSize, _messageSizeErr := readBuffer.ReadUint16("messageSize", 16)
	_ = messageSize
	if _messageSizeErr != nil {
		return nil, errors.Wrap(_messageSizeErr, "Error parsing 'messageSize' field")
	}

	if pullErr := readBuffer.PullContext("unconnectedService"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (unconnectedService)
	unconnectedService, _unconnectedServiceErr := CipServiceParse(readBuffer, messageSize)
	if _unconnectedServiceErr != nil {
		return nil, errors.Wrap(_unconnectedServiceErr, "Error parsing 'unconnectedService' field")
	}
	if closeErr := readBuffer.CloseContext("unconnectedService"); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (route)
	route, _routeErr := readBuffer.ReadUint16("route", 16)
	if _routeErr != nil {
		return nil, errors.Wrap(_routeErr, "Error parsing 'route' field")
	}
	if route != CipUnconnectedRequest_ROUTE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CipUnconnectedRequest_ROUTE) + " but got " + fmt.Sprintf("%d", route))
	}

	// Simple Field (backPlane)
	backPlane, _backPlaneErr := readBuffer.ReadInt8("backPlane", 8)
	if _backPlaneErr != nil {
		return nil, errors.Wrap(_backPlaneErr, "Error parsing 'backPlane' field")
	}

	// Simple Field (slot)
	slot, _slotErr := readBuffer.ReadInt8("slot", 8)
	if _slotErr != nil {
		return nil, errors.Wrap(_slotErr, "Error parsing 'slot' field")
	}

	if closeErr := readBuffer.CloseContext("CipUnconnectedRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CipUnconnectedRequest{
		UnconnectedService: unconnectedService,
		BackPlane:          backPlane,
		Slot:               slot,
		Parent:             &CipService{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *CipUnconnectedRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipUnconnectedRequest"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x02))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x20))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x06))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x24))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x01))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint16("reserved", 16, uint16(0x9D05))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Implicit Field (messageSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		messageSize := uint16(uint16(uint16(uint16(m.LengthInBytes()))-uint16(uint16(10))) - uint16(uint16(4)))
		_messageSizeErr := writeBuffer.WriteUint16("messageSize", 16, (messageSize))
		if _messageSizeErr != nil {
			return errors.Wrap(_messageSizeErr, "Error serializing 'messageSize' field")
		}

		// Simple Field (unconnectedService)
		if pushErr := writeBuffer.PushContext("unconnectedService"); pushErr != nil {
			return pushErr
		}
		_unconnectedServiceErr := m.UnconnectedService.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("unconnectedService"); popErr != nil {
			return popErr
		}
		if _unconnectedServiceErr != nil {
			return errors.Wrap(_unconnectedServiceErr, "Error serializing 'unconnectedService' field")
		}

		// Const Field (route)
		_routeErr := writeBuffer.WriteUint16("route", 16, 0x0001)
		if _routeErr != nil {
			return errors.Wrap(_routeErr, "Error serializing 'route' field")
		}

		// Simple Field (backPlane)
		backPlane := int8(m.BackPlane)
		_backPlaneErr := writeBuffer.WriteInt8("backPlane", 8, (backPlane))
		if _backPlaneErr != nil {
			return errors.Wrap(_backPlaneErr, "Error serializing 'backPlane' field")
		}

		// Simple Field (slot)
		slot := int8(m.Slot)
		_slotErr := writeBuffer.WriteInt8("slot", 8, (slot))
		if _slotErr != nil {
			return errors.Wrap(_slotErr, "Error serializing 'slot' field")
		}

		if popErr := writeBuffer.PopContext("CipUnconnectedRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *CipUnconnectedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
