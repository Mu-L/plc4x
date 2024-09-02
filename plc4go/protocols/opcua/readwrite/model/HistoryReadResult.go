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

// HistoryReadResult is the corresponding interface of HistoryReadResult
type HistoryReadResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetContinuationPoint returns ContinuationPoint (property field)
	GetContinuationPoint() PascalByteString
	// GetHistoryData returns HistoryData (property field)
	GetHistoryData() ExtensionObject
}

// HistoryReadResultExactly can be used when we want exactly this type and not a type which fulfills HistoryReadResult.
// This is useful for switch cases.
type HistoryReadResultExactly interface {
	HistoryReadResult
	isHistoryReadResult() bool
}

// _HistoryReadResult is the data-structure of this message
type _HistoryReadResult struct {
	ExtensionObjectDefinitionContract
	StatusCode        StatusCode
	ContinuationPoint PascalByteString
	HistoryData       ExtensionObject
}

var _ HistoryReadResult = (*_HistoryReadResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryReadResult)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryReadResult) GetIdentifier() string {
	return "640"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryReadResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryReadResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_HistoryReadResult) GetContinuationPoint() PascalByteString {
	return m.ContinuationPoint
}

func (m *_HistoryReadResult) GetHistoryData() ExtensionObject {
	return m.HistoryData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHistoryReadResult factory function for _HistoryReadResult
func NewHistoryReadResult(statusCode StatusCode, continuationPoint PascalByteString, historyData ExtensionObject) *_HistoryReadResult {
	_result := &_HistoryReadResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		ContinuationPoint:                 continuationPoint,
		HistoryData:                       historyData,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastHistoryReadResult(structType any) HistoryReadResult {
	if casted, ok := structType.(HistoryReadResult); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryReadResult); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryReadResult) GetTypeName() string {
	return "HistoryReadResult"
}

func (m *_HistoryReadResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (continuationPoint)
	lengthInBits += m.ContinuationPoint.GetLengthInBits(ctx)

	// Simple field (historyData)
	lengthInBits += m.HistoryData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_HistoryReadResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryReadResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__historyReadResult HistoryReadResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryReadResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryReadResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	continuationPoint, err := ReadSimpleField[PascalByteString](ctx, "continuationPoint", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'continuationPoint' field"))
	}
	m.ContinuationPoint = continuationPoint

	historyData, err := ReadSimpleField[ExtensionObject](ctx, "historyData", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'historyData' field"))
	}
	m.HistoryData = historyData

	if closeErr := readBuffer.CloseContext("HistoryReadResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryReadResult")
	}

	return m, nil
}

func (m *_HistoryReadResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryReadResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryReadResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryReadResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "continuationPoint", m.GetContinuationPoint(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'continuationPoint' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "historyData", m.GetHistoryData(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'historyData' field")
		}

		if popErr := writeBuffer.PopContext("HistoryReadResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryReadResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryReadResult) isHistoryReadResult() bool {
	return true
}

func (m *_HistoryReadResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
