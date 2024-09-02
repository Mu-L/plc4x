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

// SALDataLighting is the corresponding interface of SALDataLighting
type SALDataLighting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetLightingData returns LightingData (property field)
	GetLightingData() LightingData
	// IsSALDataLighting is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataLighting()
}

// _SALDataLighting is the data-structure of this message
type _SALDataLighting struct {
	SALDataContract
	LightingData LightingData
}

var _ SALDataLighting = (*_SALDataLighting)(nil)
var _ SALDataRequirements = (*_SALDataLighting)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataLighting) GetApplicationId() ApplicationId {
	return ApplicationId_LIGHTING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataLighting) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataLighting) GetLightingData() LightingData {
	return m.LightingData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataLighting factory function for _SALDataLighting
func NewSALDataLighting(lightingData LightingData, salData SALData) *_SALDataLighting {
	_result := &_SALDataLighting{
		SALDataContract: NewSALData(salData),
		LightingData:    lightingData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataLighting(structType any) SALDataLighting {
	if casted, ok := structType.(SALDataLighting); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataLighting); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataLighting) GetTypeName() string {
	return "SALDataLighting"
}

func (m *_SALDataLighting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (lightingData)
	lengthInBits += m.LightingData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataLighting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataLighting) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataLighting SALDataLighting, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataLighting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataLighting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightingData, err := ReadSimpleField[LightingData](ctx, "lightingData", ReadComplex[LightingData](LightingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightingData' field"))
	}
	m.LightingData = lightingData

	if closeErr := readBuffer.CloseContext("SALDataLighting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataLighting")
	}

	return m, nil
}

func (m *_SALDataLighting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataLighting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataLighting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataLighting")
		}

		if err := WriteSimpleField[LightingData](ctx, "lightingData", m.GetLightingData(), WriteComplex[LightingData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightingData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataLighting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataLighting")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataLighting) IsSALDataLighting() {}

func (m *_SALDataLighting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
