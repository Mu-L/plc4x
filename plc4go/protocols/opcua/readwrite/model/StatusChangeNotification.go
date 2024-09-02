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

// StatusChangeNotification is the corresponding interface of StatusChangeNotification
type StatusChangeNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatus returns Status (property field)
	GetStatus() StatusCode
	// GetDiagnosticInfo returns DiagnosticInfo (property field)
	GetDiagnosticInfo() DiagnosticInfo
}

// StatusChangeNotificationExactly can be used when we want exactly this type and not a type which fulfills StatusChangeNotification.
// This is useful for switch cases.
type StatusChangeNotificationExactly interface {
	StatusChangeNotification
	isStatusChangeNotification() bool
}

// _StatusChangeNotification is the data-structure of this message
type _StatusChangeNotification struct {
	ExtensionObjectDefinitionContract
	Status         StatusCode
	DiagnosticInfo DiagnosticInfo
}

var _ StatusChangeNotification = (*_StatusChangeNotification)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_StatusChangeNotification)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_StatusChangeNotification) GetIdentifier() string {
	return "820"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusChangeNotification) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusChangeNotification) GetStatus() StatusCode {
	return m.Status
}

func (m *_StatusChangeNotification) GetDiagnosticInfo() DiagnosticInfo {
	return m.DiagnosticInfo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusChangeNotification factory function for _StatusChangeNotification
func NewStatusChangeNotification(status StatusCode, diagnosticInfo DiagnosticInfo) *_StatusChangeNotification {
	_result := &_StatusChangeNotification{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Status:                            status,
		DiagnosticInfo:                    diagnosticInfo,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastStatusChangeNotification(structType any) StatusChangeNotification {
	if casted, ok := structType.(StatusChangeNotification); ok {
		return casted
	}
	if casted, ok := structType.(*StatusChangeNotification); ok {
		return *casted
	}
	return nil
}

func (m *_StatusChangeNotification) GetTypeName() string {
	return "StatusChangeNotification"
}

func (m *_StatusChangeNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (notificationLength)
	lengthInBits += 32

	// Simple field (status)
	lengthInBits += m.Status.GetLengthInBits(ctx)

	// Simple field (diagnosticInfo)
	lengthInBits += m.DiagnosticInfo.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_StatusChangeNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_StatusChangeNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__statusChangeNotification StatusChangeNotification, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusChangeNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusChangeNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	notificationLength, err := ReadImplicitField[int32](ctx, "notificationLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationLength' field"))
	}
	_ = notificationLength

	status, err := ReadSimpleField[StatusCode](ctx, "status", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	diagnosticInfo, err := ReadSimpleField[DiagnosticInfo](ctx, "diagnosticInfo", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfo' field"))
	}
	m.DiagnosticInfo = diagnosticInfo

	if closeErr := readBuffer.CloseContext("StatusChangeNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusChangeNotification")
	}

	return m, nil
}

func (m *_StatusChangeNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusChangeNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusChangeNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusChangeNotification")
		}
		notificationLength := int32(int32(m.GetLengthInBytes(ctx)))
		if err := WriteImplicitField(ctx, "notificationLength", notificationLength, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationLength' field")
		}

		if err := WriteSimpleField[StatusCode](ctx, "status", m.GetStatus(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[DiagnosticInfo](ctx, "diagnosticInfo", m.GetDiagnosticInfo(), WriteComplex[DiagnosticInfo](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfo' field")
		}

		if popErr := writeBuffer.PopContext("StatusChangeNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusChangeNotification")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_StatusChangeNotification) isStatusChangeNotification() bool {
	return true
}

func (m *_StatusChangeNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
