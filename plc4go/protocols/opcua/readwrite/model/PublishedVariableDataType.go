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

// PublishedVariableDataType is the corresponding interface of PublishedVariableDataType
type PublishedVariableDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetPublishedVariable returns PublishedVariable (property field)
	GetPublishedVariable() NodeId
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetSamplingIntervalHint returns SamplingIntervalHint (property field)
	GetSamplingIntervalHint() float64
	// GetDeadbandType returns DeadbandType (property field)
	GetDeadbandType() uint32
	// GetDeadbandValue returns DeadbandValue (property field)
	GetDeadbandValue() float64
	// GetIndexRange returns IndexRange (property field)
	GetIndexRange() PascalString
	// GetSubstituteValue returns SubstituteValue (property field)
	GetSubstituteValue() Variant
	// GetNoOfMetaDataProperties returns NoOfMetaDataProperties (property field)
	GetNoOfMetaDataProperties() int32
	// GetMetaDataProperties returns MetaDataProperties (property field)
	GetMetaDataProperties() []QualifiedName
}

// PublishedVariableDataTypeExactly can be used when we want exactly this type and not a type which fulfills PublishedVariableDataType.
// This is useful for switch cases.
type PublishedVariableDataTypeExactly interface {
	PublishedVariableDataType
	isPublishedVariableDataType() bool
}

// _PublishedVariableDataType is the data-structure of this message
type _PublishedVariableDataType struct {
	*_ExtensionObjectDefinition
	PublishedVariable      NodeId
	AttributeId            uint32
	SamplingIntervalHint   float64
	DeadbandType           uint32
	DeadbandValue          float64
	IndexRange             PascalString
	SubstituteValue        Variant
	NoOfMetaDataProperties int32
	MetaDataProperties     []QualifiedName
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PublishedVariableDataType) GetIdentifier() string {
	return "14275"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PublishedVariableDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PublishedVariableDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PublishedVariableDataType) GetPublishedVariable() NodeId {
	return m.PublishedVariable
}

func (m *_PublishedVariableDataType) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_PublishedVariableDataType) GetSamplingIntervalHint() float64 {
	return m.SamplingIntervalHint
}

func (m *_PublishedVariableDataType) GetDeadbandType() uint32 {
	return m.DeadbandType
}

func (m *_PublishedVariableDataType) GetDeadbandValue() float64 {
	return m.DeadbandValue
}

func (m *_PublishedVariableDataType) GetIndexRange() PascalString {
	return m.IndexRange
}

func (m *_PublishedVariableDataType) GetSubstituteValue() Variant {
	return m.SubstituteValue
}

func (m *_PublishedVariableDataType) GetNoOfMetaDataProperties() int32 {
	return m.NoOfMetaDataProperties
}

func (m *_PublishedVariableDataType) GetMetaDataProperties() []QualifiedName {
	return m.MetaDataProperties
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPublishedVariableDataType factory function for _PublishedVariableDataType
func NewPublishedVariableDataType(publishedVariable NodeId, attributeId uint32, samplingIntervalHint float64, deadbandType uint32, deadbandValue float64, indexRange PascalString, substituteValue Variant, noOfMetaDataProperties int32, metaDataProperties []QualifiedName) *_PublishedVariableDataType {
	_result := &_PublishedVariableDataType{
		PublishedVariable:          publishedVariable,
		AttributeId:                attributeId,
		SamplingIntervalHint:       samplingIntervalHint,
		DeadbandType:               deadbandType,
		DeadbandValue:              deadbandValue,
		IndexRange:                 indexRange,
		SubstituteValue:            substituteValue,
		NoOfMetaDataProperties:     noOfMetaDataProperties,
		MetaDataProperties:         metaDataProperties,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPublishedVariableDataType(structType any) PublishedVariableDataType {
	if casted, ok := structType.(PublishedVariableDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PublishedVariableDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PublishedVariableDataType) GetTypeName() string {
	return "PublishedVariableDataType"
}

func (m *_PublishedVariableDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (publishedVariable)
	lengthInBits += m.PublishedVariable.GetLengthInBits(ctx)

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (samplingIntervalHint)
	lengthInBits += 64

	// Simple field (deadbandType)
	lengthInBits += 32

	// Simple field (deadbandValue)
	lengthInBits += 64

	// Simple field (indexRange)
	lengthInBits += m.IndexRange.GetLengthInBits(ctx)

	// Simple field (substituteValue)
	lengthInBits += m.SubstituteValue.GetLengthInBits(ctx)

	// Simple field (noOfMetaDataProperties)
	lengthInBits += 32

	// Array field
	if len(m.MetaDataProperties) > 0 {
		for _curItem, element := range m.MetaDataProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MetaDataProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PublishedVariableDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PublishedVariableDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (PublishedVariableDataType, error) {
	return PublishedVariableDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PublishedVariableDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PublishedVariableDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PublishedVariableDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PublishedVariableDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (publishedVariable)
	if pullErr := readBuffer.PullContext("publishedVariable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for publishedVariable")
	}
	_publishedVariable, _publishedVariableErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _publishedVariableErr != nil {
		return nil, errors.Wrap(_publishedVariableErr, "Error parsing 'publishedVariable' field of PublishedVariableDataType")
	}
	publishedVariable := _publishedVariable.(NodeId)
	if closeErr := readBuffer.CloseContext("publishedVariable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for publishedVariable")
	}

	// Simple Field (attributeId)
	_attributeId, _attributeIdErr := /*TODO: migrate me*/ readBuffer.ReadUint32("attributeId", 32)
	if _attributeIdErr != nil {
		return nil, errors.Wrap(_attributeIdErr, "Error parsing 'attributeId' field of PublishedVariableDataType")
	}
	attributeId := _attributeId

	// Simple Field (samplingIntervalHint)
	_samplingIntervalHint, _samplingIntervalHintErr := /*TODO: migrate me*/ readBuffer.ReadFloat64("samplingIntervalHint", 64)
	if _samplingIntervalHintErr != nil {
		return nil, errors.Wrap(_samplingIntervalHintErr, "Error parsing 'samplingIntervalHint' field of PublishedVariableDataType")
	}
	samplingIntervalHint := _samplingIntervalHint

	// Simple Field (deadbandType)
	_deadbandType, _deadbandTypeErr := /*TODO: migrate me*/ readBuffer.ReadUint32("deadbandType", 32)
	if _deadbandTypeErr != nil {
		return nil, errors.Wrap(_deadbandTypeErr, "Error parsing 'deadbandType' field of PublishedVariableDataType")
	}
	deadbandType := _deadbandType

	// Simple Field (deadbandValue)
	_deadbandValue, _deadbandValueErr := /*TODO: migrate me*/ readBuffer.ReadFloat64("deadbandValue", 64)
	if _deadbandValueErr != nil {
		return nil, errors.Wrap(_deadbandValueErr, "Error parsing 'deadbandValue' field of PublishedVariableDataType")
	}
	deadbandValue := _deadbandValue

	// Simple Field (indexRange)
	if pullErr := readBuffer.PullContext("indexRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for indexRange")
	}
	_indexRange, _indexRangeErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _indexRangeErr != nil {
		return nil, errors.Wrap(_indexRangeErr, "Error parsing 'indexRange' field of PublishedVariableDataType")
	}
	indexRange := _indexRange.(PascalString)
	if closeErr := readBuffer.CloseContext("indexRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for indexRange")
	}

	// Simple Field (substituteValue)
	if pullErr := readBuffer.PullContext("substituteValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for substituteValue")
	}
	_substituteValue, _substituteValueErr := VariantParseWithBuffer(ctx, readBuffer)
	if _substituteValueErr != nil {
		return nil, errors.Wrap(_substituteValueErr, "Error parsing 'substituteValue' field of PublishedVariableDataType")
	}
	substituteValue := _substituteValue.(Variant)
	if closeErr := readBuffer.CloseContext("substituteValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for substituteValue")
	}

	// Simple Field (noOfMetaDataProperties)
	_noOfMetaDataProperties, _noOfMetaDataPropertiesErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfMetaDataProperties", 32)
	if _noOfMetaDataPropertiesErr != nil {
		return nil, errors.Wrap(_noOfMetaDataPropertiesErr, "Error parsing 'noOfMetaDataProperties' field of PublishedVariableDataType")
	}
	noOfMetaDataProperties := _noOfMetaDataProperties

	metaDataProperties, err := ReadCountArrayField[QualifiedName](ctx, "metaDataProperties", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer), uint64(noOfMetaDataProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'metaDataProperties' field"))
	}

	if closeErr := readBuffer.CloseContext("PublishedVariableDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PublishedVariableDataType")
	}

	// Create a partially initialized instance
	_child := &_PublishedVariableDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		PublishedVariable:          publishedVariable,
		AttributeId:                attributeId,
		SamplingIntervalHint:       samplingIntervalHint,
		DeadbandType:               deadbandType,
		DeadbandValue:              deadbandValue,
		IndexRange:                 indexRange,
		SubstituteValue:            substituteValue,
		NoOfMetaDataProperties:     noOfMetaDataProperties,
		MetaDataProperties:         metaDataProperties,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PublishedVariableDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PublishedVariableDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PublishedVariableDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PublishedVariableDataType")
		}

		// Simple Field (publishedVariable)
		if pushErr := writeBuffer.PushContext("publishedVariable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for publishedVariable")
		}
		_publishedVariableErr := writeBuffer.WriteSerializable(ctx, m.GetPublishedVariable())
		if popErr := writeBuffer.PopContext("publishedVariable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for publishedVariable")
		}
		if _publishedVariableErr != nil {
			return errors.Wrap(_publishedVariableErr, "Error serializing 'publishedVariable' field")
		}

		// Simple Field (attributeId)
		attributeId := uint32(m.GetAttributeId())
		_attributeIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("attributeId", 32, uint32((attributeId)))
		if _attributeIdErr != nil {
			return errors.Wrap(_attributeIdErr, "Error serializing 'attributeId' field")
		}

		// Simple Field (samplingIntervalHint)
		samplingIntervalHint := float64(m.GetSamplingIntervalHint())
		_samplingIntervalHintErr := /*TODO: migrate me*/ writeBuffer.WriteFloat64("samplingIntervalHint", 64, (samplingIntervalHint))
		if _samplingIntervalHintErr != nil {
			return errors.Wrap(_samplingIntervalHintErr, "Error serializing 'samplingIntervalHint' field")
		}

		// Simple Field (deadbandType)
		deadbandType := uint32(m.GetDeadbandType())
		_deadbandTypeErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("deadbandType", 32, uint32((deadbandType)))
		if _deadbandTypeErr != nil {
			return errors.Wrap(_deadbandTypeErr, "Error serializing 'deadbandType' field")
		}

		// Simple Field (deadbandValue)
		deadbandValue := float64(m.GetDeadbandValue())
		_deadbandValueErr := /*TODO: migrate me*/ writeBuffer.WriteFloat64("deadbandValue", 64, (deadbandValue))
		if _deadbandValueErr != nil {
			return errors.Wrap(_deadbandValueErr, "Error serializing 'deadbandValue' field")
		}

		// Simple Field (indexRange)
		if pushErr := writeBuffer.PushContext("indexRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for indexRange")
		}
		_indexRangeErr := writeBuffer.WriteSerializable(ctx, m.GetIndexRange())
		if popErr := writeBuffer.PopContext("indexRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for indexRange")
		}
		if _indexRangeErr != nil {
			return errors.Wrap(_indexRangeErr, "Error serializing 'indexRange' field")
		}

		// Simple Field (substituteValue)
		if pushErr := writeBuffer.PushContext("substituteValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for substituteValue")
		}
		_substituteValueErr := writeBuffer.WriteSerializable(ctx, m.GetSubstituteValue())
		if popErr := writeBuffer.PopContext("substituteValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for substituteValue")
		}
		if _substituteValueErr != nil {
			return errors.Wrap(_substituteValueErr, "Error serializing 'substituteValue' field")
		}

		// Simple Field (noOfMetaDataProperties)
		noOfMetaDataProperties := int32(m.GetNoOfMetaDataProperties())
		_noOfMetaDataPropertiesErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfMetaDataProperties", 32, int32((noOfMetaDataProperties)))
		if _noOfMetaDataPropertiesErr != nil {
			return errors.Wrap(_noOfMetaDataPropertiesErr, "Error serializing 'noOfMetaDataProperties' field")
		}

		// Array Field (metaDataProperties)
		if pushErr := writeBuffer.PushContext("metaDataProperties", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for metaDataProperties")
		}
		for _curItem, _element := range m.GetMetaDataProperties() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetMetaDataProperties()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'metaDataProperties' field")
			}
		}
		if popErr := writeBuffer.PopContext("metaDataProperties", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for metaDataProperties")
		}

		if popErr := writeBuffer.PopContext("PublishedVariableDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PublishedVariableDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PublishedVariableDataType) isPublishedVariableDataType() bool {
	return true
}

func (m *_PublishedVariableDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
