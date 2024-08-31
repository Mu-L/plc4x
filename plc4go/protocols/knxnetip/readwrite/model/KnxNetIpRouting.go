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

// KnxNetIpRouting is the corresponding interface of KnxNetIpRouting
type KnxNetIpRouting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetIpRoutingExactly can be used when we want exactly this type and not a type which fulfills KnxNetIpRouting.
// This is useful for switch cases.
type KnxNetIpRoutingExactly interface {
	KnxNetIpRouting
	isKnxNetIpRouting() bool
}

// _KnxNetIpRouting is the data-structure of this message
type _KnxNetIpRouting struct {
	*_ServiceId
	Version uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpRouting) GetServiceType() uint8 {
	return 0x05
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpRouting) InitializeParent(parent ServiceId) {}

func (m *_KnxNetIpRouting) GetParent() ServiceId {
	return m._ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpRouting) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpRouting factory function for _KnxNetIpRouting
func NewKnxNetIpRouting(version uint8) *_KnxNetIpRouting {
	_result := &_KnxNetIpRouting{
		Version:    version,
		_ServiceId: NewServiceId(),
	}
	_result._ServiceId._ServiceIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetIpRouting(structType any) KnxNetIpRouting {
	if casted, ok := structType.(KnxNetIpRouting); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpRouting); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpRouting) GetTypeName() string {
	return "KnxNetIpRouting"
}

func (m *_KnxNetIpRouting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpRouting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func KnxNetIpRoutingParse(ctx context.Context, theBytes []byte) (KnxNetIpRouting, error) {
	return KnxNetIpRoutingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func KnxNetIpRoutingParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (KnxNetIpRouting, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (KnxNetIpRouting, error) {
		return KnxNetIpRoutingParseWithBuffer(ctx, readBuffer)
	}
}

func KnxNetIpRoutingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (KnxNetIpRouting, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("KnxNetIpRouting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpRouting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}

	if closeErr := readBuffer.CloseContext("KnxNetIpRouting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpRouting")
	}

	// Create a partially initialized instance
	_child := &_KnxNetIpRouting{
		_ServiceId: &_ServiceId{},
		Version:    version,
	}
	_child._ServiceId._ServiceIdChildRequirements = _child
	return _child, nil
}

func (m *_KnxNetIpRouting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetIpRouting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpRouting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpRouting")
		}

		// Simple Field (version)
		version := uint8(m.GetVersion())
		_versionErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("version", 8, uint8((version)))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpRouting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpRouting")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetIpRouting) isKnxNetIpRouting() bool {
	return true
}

func (m *_KnxNetIpRouting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
