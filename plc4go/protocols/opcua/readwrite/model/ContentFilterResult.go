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

// ContentFilterResult is the corresponding interface of ContentFilterResult
type ContentFilterResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfElementResults returns NoOfElementResults (property field)
	GetNoOfElementResults() int32
	// GetElementResults returns ElementResults (property field)
	GetElementResults() []ExtensionObjectDefinition
	// GetNoOfElementDiagnosticInfos returns NoOfElementDiagnosticInfos (property field)
	GetNoOfElementDiagnosticInfos() int32
	// GetElementDiagnosticInfos returns ElementDiagnosticInfos (property field)
	GetElementDiagnosticInfos() []DiagnosticInfo
}

// ContentFilterResultExactly can be used when we want exactly this type and not a type which fulfills ContentFilterResult.
// This is useful for switch cases.
type ContentFilterResultExactly interface {
	ContentFilterResult
	isContentFilterResult() bool
}

// _ContentFilterResult is the data-structure of this message
type _ContentFilterResult struct {
	*_ExtensionObjectDefinition
	NoOfElementResults         int32
	ElementResults             []ExtensionObjectDefinition
	NoOfElementDiagnosticInfos int32
	ElementDiagnosticInfos     []DiagnosticInfo
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ContentFilterResult) GetIdentifier() string {
	return "609"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ContentFilterResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ContentFilterResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ContentFilterResult) GetNoOfElementResults() int32 {
	return m.NoOfElementResults
}

func (m *_ContentFilterResult) GetElementResults() []ExtensionObjectDefinition {
	return m.ElementResults
}

func (m *_ContentFilterResult) GetNoOfElementDiagnosticInfos() int32 {
	return m.NoOfElementDiagnosticInfos
}

func (m *_ContentFilterResult) GetElementDiagnosticInfos() []DiagnosticInfo {
	return m.ElementDiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewContentFilterResult factory function for _ContentFilterResult
func NewContentFilterResult(noOfElementResults int32, elementResults []ExtensionObjectDefinition, noOfElementDiagnosticInfos int32, elementDiagnosticInfos []DiagnosticInfo) *_ContentFilterResult {
	_result := &_ContentFilterResult{
		NoOfElementResults:         noOfElementResults,
		ElementResults:             elementResults,
		NoOfElementDiagnosticInfos: noOfElementDiagnosticInfos,
		ElementDiagnosticInfos:     elementDiagnosticInfos,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastContentFilterResult(structType any) ContentFilterResult {
	if casted, ok := structType.(ContentFilterResult); ok {
		return casted
	}
	if casted, ok := structType.(*ContentFilterResult); ok {
		return *casted
	}
	return nil
}

func (m *_ContentFilterResult) GetTypeName() string {
	return "ContentFilterResult"
}

func (m *_ContentFilterResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noOfElementResults)
	lengthInBits += 32

	// Array field
	if len(m.ElementResults) > 0 {
		for _curItem, element := range m.ElementResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ElementResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfElementDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.ElementDiagnosticInfos) > 0 {
		for _curItem, element := range m.ElementDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ElementDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ContentFilterResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ContentFilterResultParse(ctx context.Context, theBytes []byte, identifier string) (ContentFilterResult, error) {
	return ContentFilterResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ContentFilterResultParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (ContentFilterResult, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ContentFilterResult, error) {
		return ContentFilterResultParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func ContentFilterResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ContentFilterResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ContentFilterResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ContentFilterResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfElementResults, err := ReadSimpleField(ctx, "noOfElementResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfElementResults' field"))
	}

	elementResults, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "elementResults", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("606")), readBuffer), uint64(noOfElementResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elementResults' field"))
	}

	noOfElementDiagnosticInfos, err := ReadSimpleField(ctx, "noOfElementDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfElementDiagnosticInfos' field"))
	}

	elementDiagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "elementDiagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfElementDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elementDiagnosticInfos' field"))
	}

	if closeErr := readBuffer.CloseContext("ContentFilterResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ContentFilterResult")
	}

	// Create a partially initialized instance
	_child := &_ContentFilterResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NoOfElementResults:         noOfElementResults,
		ElementResults:             elementResults,
		NoOfElementDiagnosticInfos: noOfElementDiagnosticInfos,
		ElementDiagnosticInfos:     elementDiagnosticInfos,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ContentFilterResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ContentFilterResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ContentFilterResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ContentFilterResult")
		}

		// Simple Field (noOfElementResults)
		noOfElementResults := int32(m.GetNoOfElementResults())
		_noOfElementResultsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfElementResults", 32, int32((noOfElementResults)))
		if _noOfElementResultsErr != nil {
			return errors.Wrap(_noOfElementResultsErr, "Error serializing 'noOfElementResults' field")
		}

		// Array Field (elementResults)
		if pushErr := writeBuffer.PushContext("elementResults", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for elementResults")
		}
		for _curItem, _element := range m.GetElementResults() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetElementResults()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'elementResults' field")
			}
		}
		if popErr := writeBuffer.PopContext("elementResults", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for elementResults")
		}

		// Simple Field (noOfElementDiagnosticInfos)
		noOfElementDiagnosticInfos := int32(m.GetNoOfElementDiagnosticInfos())
		_noOfElementDiagnosticInfosErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfElementDiagnosticInfos", 32, int32((noOfElementDiagnosticInfos)))
		if _noOfElementDiagnosticInfosErr != nil {
			return errors.Wrap(_noOfElementDiagnosticInfosErr, "Error serializing 'noOfElementDiagnosticInfos' field")
		}

		// Array Field (elementDiagnosticInfos)
		if pushErr := writeBuffer.PushContext("elementDiagnosticInfos", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for elementDiagnosticInfos")
		}
		for _curItem, _element := range m.GetElementDiagnosticInfos() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetElementDiagnosticInfos()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'elementDiagnosticInfos' field")
			}
		}
		if popErr := writeBuffer.PopContext("elementDiagnosticInfos", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for elementDiagnosticInfos")
		}

		if popErr := writeBuffer.PopContext("ContentFilterResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ContentFilterResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ContentFilterResult) isContentFilterResult() bool {
	return true
}

func (m *_ContentFilterResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
