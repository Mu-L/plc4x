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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsReadStateRequest is the corresponding interface of AdsReadStateRequest
type AdsReadStateRequest interface {
	utils.LengthAware
	utils.Serializable
	AdsData
}

// AdsReadStateRequestExactly can be used when we want exactly this type and not a type which fulfills AdsReadStateRequest.
// This is useful for switch cases.
type AdsReadStateRequestExactly interface {
	AdsReadStateRequest
	isAdsReadStateRequest() bool
}

// _AdsReadStateRequest is the data-structure of this message
type _AdsReadStateRequest struct {
	*_AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadStateRequest) GetCommandId() CommandId {
	return CommandId_ADS_READ_STATE
}

func (m *_AdsReadStateRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadStateRequest) InitializeParent(parent AdsData) {}

func (m *_AdsReadStateRequest) GetParent() AdsData {
	return m._AdsData
}

// NewAdsReadStateRequest factory function for _AdsReadStateRequest
func NewAdsReadStateRequest() *_AdsReadStateRequest {
	_result := &_AdsReadStateRequest{
		_AdsData: NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadStateRequest(structType interface{}) AdsReadStateRequest {
	if casted, ok := structType.(AdsReadStateRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadStateRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadStateRequest) GetTypeName() string {
	return "AdsReadStateRequest"
}

func (m *_AdsReadStateRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsReadStateRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_AdsReadStateRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadStateRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsReadStateRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadStateRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadStateRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsReadStateRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadStateRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsReadStateRequest{
		_AdsData: &_AdsData{},
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsReadStateRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadStateRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadStateRequest")
		}

		if popErr := writeBuffer.PopContext("AdsReadStateRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadStateRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsReadStateRequest) isAdsReadStateRequest() bool {
	return true
}

func (m *_AdsReadStateRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
