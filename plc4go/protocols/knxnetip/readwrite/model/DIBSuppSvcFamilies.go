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

// DIBSuppSvcFamilies is the corresponding interface of DIBSuppSvcFamilies
type DIBSuppSvcFamilies interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDescriptionType returns DescriptionType (property field)
	GetDescriptionType() uint8
	// GetServiceIds returns ServiceIds (property field)
	GetServiceIds() []ServiceId
}

// DIBSuppSvcFamiliesExactly can be used when we want exactly this type and not a type which fulfills DIBSuppSvcFamilies.
// This is useful for switch cases.
type DIBSuppSvcFamiliesExactly interface {
	DIBSuppSvcFamilies
	isDIBSuppSvcFamilies() bool
}

// _DIBSuppSvcFamilies is the data-structure of this message
type _DIBSuppSvcFamilies struct {
	DescriptionType uint8
	ServiceIds      []ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DIBSuppSvcFamilies) GetDescriptionType() uint8 {
	return m.DescriptionType
}

func (m *_DIBSuppSvcFamilies) GetServiceIds() []ServiceId {
	return m.ServiceIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDIBSuppSvcFamilies factory function for _DIBSuppSvcFamilies
func NewDIBSuppSvcFamilies(descriptionType uint8, serviceIds []ServiceId) *_DIBSuppSvcFamilies {
	return &_DIBSuppSvcFamilies{DescriptionType: descriptionType, ServiceIds: serviceIds}
}

// Deprecated: use the interface for direct cast
func CastDIBSuppSvcFamilies(structType any) DIBSuppSvcFamilies {
	if casted, ok := structType.(DIBSuppSvcFamilies); ok {
		return casted
	}
	if casted, ok := structType.(*DIBSuppSvcFamilies); ok {
		return *casted
	}
	return nil
}

func (m *_DIBSuppSvcFamilies) GetTypeName() string {
	return "DIBSuppSvcFamilies"
}

func (m *_DIBSuppSvcFamilies) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (descriptionType)
	lengthInBits += 8

	// Array field
	if len(m.ServiceIds) > 0 {
		for _, element := range m.ServiceIds {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_DIBSuppSvcFamilies) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DIBSuppSvcFamiliesParse(ctx context.Context, theBytes []byte) (DIBSuppSvcFamilies, error) {
	return DIBSuppSvcFamiliesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DIBSuppSvcFamiliesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DIBSuppSvcFamilies, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DIBSuppSvcFamilies"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DIBSuppSvcFamilies")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field of DIBSuppSvcFamilies")
	}

	// Simple Field (descriptionType)
	_descriptionType, _descriptionTypeErr := /*TODO: migrate me*/ readBuffer.ReadUint8("descriptionType", 8)
	if _descriptionTypeErr != nil {
		return nil, errors.Wrap(_descriptionTypeErr, "Error parsing 'descriptionType' field of DIBSuppSvcFamilies")
	}
	descriptionType := _descriptionType

	serviceIds, err := ReadLengthArrayField[ServiceId](ctx, "serviceIds", ReadComplex[ServiceId](ServiceIdParseWithBuffer, readBuffer), int(int32(structureLength)-int32(int32(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceIds' field"))
	}

	if closeErr := readBuffer.CloseContext("DIBSuppSvcFamilies"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DIBSuppSvcFamilies")
	}

	// Create the instance
	return &_DIBSuppSvcFamilies{
		DescriptionType: descriptionType,
		ServiceIds:      serviceIds,
	}, nil
}

func (m *_DIBSuppSvcFamilies) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DIBSuppSvcFamilies) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DIBSuppSvcFamilies"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DIBSuppSvcFamilies")
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	_structureLengthErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("structureLength", 8, uint8((structureLength)))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (descriptionType)
	descriptionType := uint8(m.GetDescriptionType())
	_descriptionTypeErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("descriptionType", 8, uint8((descriptionType)))
	if _descriptionTypeErr != nil {
		return errors.Wrap(_descriptionTypeErr, "Error serializing 'descriptionType' field")
	}

	// Array Field (serviceIds)
	if pushErr := writeBuffer.PushContext("serviceIds", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for serviceIds")
	}
	for _curItem, _element := range m.GetServiceIds() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetServiceIds()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'serviceIds' field")
		}
	}
	if popErr := writeBuffer.PopContext("serviceIds", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for serviceIds")
	}

	if popErr := writeBuffer.PopContext("DIBSuppSvcFamilies"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DIBSuppSvcFamilies")
	}
	return nil
}

func (m *_DIBSuppSvcFamilies) isDIBSuppSvcFamilies() bool {
	return true
}

func (m *_DIBSuppSvcFamilies) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
