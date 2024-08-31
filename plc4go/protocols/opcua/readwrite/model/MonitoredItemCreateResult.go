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

// MonitoredItemCreateResult is the corresponding interface of MonitoredItemCreateResult
type MonitoredItemCreateResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetMonitoredItemId returns MonitoredItemId (property field)
	GetMonitoredItemId() uint32
	// GetRevisedSamplingInterval returns RevisedSamplingInterval (property field)
	GetRevisedSamplingInterval() float64
	// GetRevisedQueueSize returns RevisedQueueSize (property field)
	GetRevisedQueueSize() uint32
	// GetFilterResult returns FilterResult (property field)
	GetFilterResult() ExtensionObject
}

// MonitoredItemCreateResultExactly can be used when we want exactly this type and not a type which fulfills MonitoredItemCreateResult.
// This is useful for switch cases.
type MonitoredItemCreateResultExactly interface {
	MonitoredItemCreateResult
	isMonitoredItemCreateResult() bool
}

// _MonitoredItemCreateResult is the data-structure of this message
type _MonitoredItemCreateResult struct {
	*_ExtensionObjectDefinition
	StatusCode              StatusCode
	MonitoredItemId         uint32
	RevisedSamplingInterval float64
	RevisedQueueSize        uint32
	FilterResult            ExtensionObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemCreateResult) GetIdentifier() string {
	return "748"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemCreateResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_MonitoredItemCreateResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemCreateResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_MonitoredItemCreateResult) GetMonitoredItemId() uint32 {
	return m.MonitoredItemId
}

func (m *_MonitoredItemCreateResult) GetRevisedSamplingInterval() float64 {
	return m.RevisedSamplingInterval
}

func (m *_MonitoredItemCreateResult) GetRevisedQueueSize() uint32 {
	return m.RevisedQueueSize
}

func (m *_MonitoredItemCreateResult) GetFilterResult() ExtensionObject {
	return m.FilterResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredItemCreateResult factory function for _MonitoredItemCreateResult
func NewMonitoredItemCreateResult(statusCode StatusCode, monitoredItemId uint32, revisedSamplingInterval float64, revisedQueueSize uint32, filterResult ExtensionObject) *_MonitoredItemCreateResult {
	_result := &_MonitoredItemCreateResult{
		StatusCode:                 statusCode,
		MonitoredItemId:            monitoredItemId,
		RevisedSamplingInterval:    revisedSamplingInterval,
		RevisedQueueSize:           revisedQueueSize,
		FilterResult:               filterResult,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredItemCreateResult(structType any) MonitoredItemCreateResult {
	if casted, ok := structType.(MonitoredItemCreateResult); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemCreateResult); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemCreateResult) GetTypeName() string {
	return "MonitoredItemCreateResult"
}

func (m *_MonitoredItemCreateResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (monitoredItemId)
	lengthInBits += 32

	// Simple field (revisedSamplingInterval)
	lengthInBits += 64

	// Simple field (revisedQueueSize)
	lengthInBits += 32

	// Simple field (filterResult)
	lengthInBits += m.FilterResult.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemCreateResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredItemCreateResultParse(ctx context.Context, theBytes []byte, identifier string) (MonitoredItemCreateResult, error) {
	return MonitoredItemCreateResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func MonitoredItemCreateResultParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoredItemCreateResult, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoredItemCreateResult, error) {
		return MonitoredItemCreateResultParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func MonitoredItemCreateResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (MonitoredItemCreateResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MonitoredItemCreateResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemCreateResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}

	monitoredItemId, err := ReadSimpleField(ctx, "monitoredItemId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredItemId' field"))
	}

	revisedSamplingInterval, err := ReadSimpleField(ctx, "revisedSamplingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisedSamplingInterval' field"))
	}

	revisedQueueSize, err := ReadSimpleField(ctx, "revisedQueueSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisedQueueSize' field"))
	}

	filterResult, err := ReadSimpleField[ExtensionObject](ctx, "filterResult", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'filterResult' field"))
	}

	if closeErr := readBuffer.CloseContext("MonitoredItemCreateResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemCreateResult")
	}

	// Create a partially initialized instance
	_child := &_MonitoredItemCreateResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StatusCode:                 statusCode,
		MonitoredItemId:            monitoredItemId,
		RevisedSamplingInterval:    revisedSamplingInterval,
		RevisedQueueSize:           revisedQueueSize,
		FilterResult:               filterResult,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredItemCreateResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemCreateResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemCreateResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemCreateResult")
		}

		// Simple Field (statusCode)
		if pushErr := writeBuffer.PushContext("statusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusCode")
		}
		_statusCodeErr := writeBuffer.WriteSerializable(ctx, m.GetStatusCode())
		if popErr := writeBuffer.PopContext("statusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusCode")
		}
		if _statusCodeErr != nil {
			return errors.Wrap(_statusCodeErr, "Error serializing 'statusCode' field")
		}

		// Simple Field (monitoredItemId)
		monitoredItemId := uint32(m.GetMonitoredItemId())
		_monitoredItemIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("monitoredItemId", 32, uint32((monitoredItemId)))
		if _monitoredItemIdErr != nil {
			return errors.Wrap(_monitoredItemIdErr, "Error serializing 'monitoredItemId' field")
		}

		// Simple Field (revisedSamplingInterval)
		revisedSamplingInterval := float64(m.GetRevisedSamplingInterval())
		_revisedSamplingIntervalErr := /*TODO: migrate me*/ writeBuffer.WriteFloat64("revisedSamplingInterval", 64, (revisedSamplingInterval))
		if _revisedSamplingIntervalErr != nil {
			return errors.Wrap(_revisedSamplingIntervalErr, "Error serializing 'revisedSamplingInterval' field")
		}

		// Simple Field (revisedQueueSize)
		revisedQueueSize := uint32(m.GetRevisedQueueSize())
		_revisedQueueSizeErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("revisedQueueSize", 32, uint32((revisedQueueSize)))
		if _revisedQueueSizeErr != nil {
			return errors.Wrap(_revisedQueueSizeErr, "Error serializing 'revisedQueueSize' field")
		}

		// Simple Field (filterResult)
		if pushErr := writeBuffer.PushContext("filterResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for filterResult")
		}
		_filterResultErr := writeBuffer.WriteSerializable(ctx, m.GetFilterResult())
		if popErr := writeBuffer.PopContext("filterResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for filterResult")
		}
		if _filterResultErr != nil {
			return errors.Wrap(_filterResultErr, "Error serializing 'filterResult' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemCreateResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemCreateResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemCreateResult) isMonitoredItemCreateResult() bool {
	return true
}

func (m *_MonitoredItemCreateResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
