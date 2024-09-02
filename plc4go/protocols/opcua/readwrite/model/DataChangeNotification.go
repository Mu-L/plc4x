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

// DataChangeNotification is the corresponding interface of DataChangeNotification
type DataChangeNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfMonitoredItems returns NoOfMonitoredItems (property field)
	GetNoOfMonitoredItems() int32
	// GetMonitoredItems returns MonitoredItems (property field)
	GetMonitoredItems() []ExtensionObjectDefinition
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
}

// DataChangeNotificationExactly can be used when we want exactly this type and not a type which fulfills DataChangeNotification.
// This is useful for switch cases.
type DataChangeNotificationExactly interface {
	DataChangeNotification
	isDataChangeNotification() bool
}

// _DataChangeNotification is the data-structure of this message
type _DataChangeNotification struct {
	ExtensionObjectDefinitionContract
	NoOfMonitoredItems  int32
	MonitoredItems      []ExtensionObjectDefinition
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
}

var _ DataChangeNotification = (*_DataChangeNotification)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataChangeNotification)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataChangeNotification) GetIdentifier() string {
	return "811"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataChangeNotification) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataChangeNotification) GetNoOfMonitoredItems() int32 {
	return m.NoOfMonitoredItems
}

func (m *_DataChangeNotification) GetMonitoredItems() []ExtensionObjectDefinition {
	return m.MonitoredItems
}

func (m *_DataChangeNotification) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_DataChangeNotification) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDataChangeNotification factory function for _DataChangeNotification
func NewDataChangeNotification(noOfMonitoredItems int32, monitoredItems []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_DataChangeNotification {
	_result := &_DataChangeNotification{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NoOfMonitoredItems:                noOfMonitoredItems,
		MonitoredItems:                    monitoredItems,
		NoOfDiagnosticInfos:               noOfDiagnosticInfos,
		DiagnosticInfos:                   diagnosticInfos,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastDataChangeNotification(structType any) DataChangeNotification {
	if casted, ok := structType.(DataChangeNotification); ok {
		return casted
	}
	if casted, ok := structType.(*DataChangeNotification); ok {
		return *casted
	}
	return nil
}

func (m *_DataChangeNotification) GetTypeName() string {
	return "DataChangeNotification"
}

func (m *_DataChangeNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (notificationLength)
	lengthInBits += 32

	// Simple field (noOfMonitoredItems)
	lengthInBits += 32

	// Array field
	if len(m.MonitoredItems) > 0 {
		for _curItem, element := range m.MonitoredItems {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MonitoredItems), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DataChangeNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataChangeNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__dataChangeNotification DataChangeNotification, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataChangeNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataChangeNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	notificationLength, err := ReadImplicitField[int32](ctx, "notificationLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationLength' field"))
	}
	_ = notificationLength

	noOfMonitoredItems, err := ReadSimpleField(ctx, "noOfMonitoredItems", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfMonitoredItems' field"))
	}
	m.NoOfMonitoredItems = noOfMonitoredItems

	monitoredItems, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "monitoredItems", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("808")), readBuffer), uint64(noOfMonitoredItems))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredItems' field"))
	}
	m.MonitoredItems = monitoredItems

	noOfDiagnosticInfos, err := ReadSimpleField(ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	m.NoOfDiagnosticInfos = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("DataChangeNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataChangeNotification")
	}

	return m, nil
}

func (m *_DataChangeNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataChangeNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataChangeNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataChangeNotification")
		}
		notificationLength := int32(int32(m.GetLengthInBytes(ctx)))
		if err := WriteImplicitField(ctx, "notificationLength", notificationLength, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationLength' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfMonitoredItems", m.GetNoOfMonitoredItems(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfMonitoredItems' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "monitoredItems", m.GetMonitoredItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredItems' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiagnosticInfos", m.GetNoOfDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("DataChangeNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataChangeNotification")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataChangeNotification) isDataChangeNotification() bool {
	return true
}

func (m *_DataChangeNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
