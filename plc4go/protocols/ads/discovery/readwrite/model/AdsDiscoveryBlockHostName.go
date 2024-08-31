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

// AdsDiscoveryBlockHostName is the corresponding interface of AdsDiscoveryBlockHostName
type AdsDiscoveryBlockHostName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsDiscoveryBlock
	// GetHostName returns HostName (property field)
	GetHostName() AmsString
}

// AdsDiscoveryBlockHostNameExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryBlockHostName.
// This is useful for switch cases.
type AdsDiscoveryBlockHostNameExactly interface {
	AdsDiscoveryBlockHostName
	isAdsDiscoveryBlockHostName() bool
}

// _AdsDiscoveryBlockHostName is the data-structure of this message
type _AdsDiscoveryBlockHostName struct {
	*_AdsDiscoveryBlock
	HostName AmsString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockHostName) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_HOST_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockHostName) InitializeParent(parent AdsDiscoveryBlock) {}

func (m *_AdsDiscoveryBlockHostName) GetParent() AdsDiscoveryBlock {
	return m._AdsDiscoveryBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockHostName) GetHostName() AmsString {
	return m.HostName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscoveryBlockHostName factory function for _AdsDiscoveryBlockHostName
func NewAdsDiscoveryBlockHostName(hostName AmsString) *_AdsDiscoveryBlockHostName {
	_result := &_AdsDiscoveryBlockHostName{
		HostName:           hostName,
		_AdsDiscoveryBlock: NewAdsDiscoveryBlock(),
	}
	_result._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockHostName(structType any) AdsDiscoveryBlockHostName {
	if casted, ok := structType.(AdsDiscoveryBlockHostName); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockHostName); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockHostName) GetTypeName() string {
	return "AdsDiscoveryBlockHostName"
}

func (m *_AdsDiscoveryBlockHostName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (hostName)
	lengthInBits += m.HostName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AdsDiscoveryBlockHostName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryBlockHostNameParse(ctx context.Context, theBytes []byte) (AdsDiscoveryBlockHostName, error) {
	return AdsDiscoveryBlockHostNameParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryBlockHostNameParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockHostName, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockHostName, error) {
		return AdsDiscoveryBlockHostNameParseWithBuffer(ctx, readBuffer)
	}
}

func AdsDiscoveryBlockHostNameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockHostName, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockHostName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockHostName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hostName, err := ReadSimpleField[AmsString](ctx, "hostName", ReadComplex[AmsString](AmsStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hostName' field"))
	}

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockHostName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockHostName")
	}

	// Create a partially initialized instance
	_child := &_AdsDiscoveryBlockHostName{
		_AdsDiscoveryBlock: &_AdsDiscoveryBlock{},
		HostName:           hostName,
	}
	_child._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _child
	return _child, nil
}

func (m *_AdsDiscoveryBlockHostName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockHostName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockHostName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockHostName")
		}

		// Simple Field (hostName)
		if pushErr := writeBuffer.PushContext("hostName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for hostName")
		}
		_hostNameErr := writeBuffer.WriteSerializable(ctx, m.GetHostName())
		if popErr := writeBuffer.PopContext("hostName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for hostName")
		}
		if _hostNameErr != nil {
			return errors.Wrap(_hostNameErr, "Error serializing 'hostName' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockHostName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockHostName")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockHostName) isAdsDiscoveryBlockHostName() bool {
	return true
}

func (m *_AdsDiscoveryBlockHostName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
