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

// BACnetLogRecord is the corresponding interface of BACnetLogRecord
type BACnetLogRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetDateTimeEnclosed
	// GetLogDatum returns LogDatum (property field)
	GetLogDatum() BACnetLogRecordLogDatum
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// IsBACnetLogRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogRecord()
}

// _BACnetLogRecord is the data-structure of this message
type _BACnetLogRecord struct {
	Timestamp   BACnetDateTimeEnclosed
	LogDatum    BACnetLogRecordLogDatum
	StatusFlags BACnetStatusFlagsTagged
}

var _ BACnetLogRecord = (*_BACnetLogRecord)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecord) GetTimestamp() BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *_BACnetLogRecord) GetLogDatum() BACnetLogRecordLogDatum {
	return m.LogDatum
}

func (m *_BACnetLogRecord) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecord factory function for _BACnetLogRecord
func NewBACnetLogRecord(timestamp BACnetDateTimeEnclosed, logDatum BACnetLogRecordLogDatum, statusFlags BACnetStatusFlagsTagged) *_BACnetLogRecord {
	return &_BACnetLogRecord{Timestamp: timestamp, LogDatum: logDatum, StatusFlags: statusFlags}
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecord(structType any) BACnetLogRecord {
	if casted, ok := structType.(BACnetLogRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecord) GetTypeName() string {
	return "BACnetLogRecord"
}

func (m *_BACnetLogRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (logDatum)
	lengthInBits += m.LogDatum.GetLengthInBits(ctx)

	// Optional Field (statusFlags)
	if m.StatusFlags != nil {
		lengthInBits += m.StatusFlags.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetLogRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogRecordParse(ctx context.Context, theBytes []byte) (BACnetLogRecord, error) {
	return BACnetLogRecordParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLogRecordParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogRecord, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogRecord, error) {
		return BACnetLogRecordParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetLogRecordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogRecord, error) {
	v, err := (&_BACnetLogRecord{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetLogRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetLogRecord BACnetLogRecord, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timestamp, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "timestamp", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	logDatum, err := ReadSimpleField[BACnetLogRecordLogDatum](ctx, "logDatum", ReadComplex[BACnetLogRecordLogDatum](BACnetLogRecordLogDatumParseWithBufferProducer[BACnetLogRecordLogDatum]((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logDatum' field"))
	}
	m.LogDatum = logDatum

	var statusFlags BACnetStatusFlagsTagged
	_statusFlags, err := ReadOptionalField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	if _statusFlags != nil {
		statusFlags = *_statusFlags
		m.StatusFlags = statusFlags
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecord")
	}

	return m, nil
}

func (m *_BACnetLogRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLogRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogRecord")
	}

	if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timestamp' field")
	}

	if err := WriteSimpleField[BACnetLogRecordLogDatum](ctx, "logDatum", m.GetLogDatum(), WriteComplex[BACnetLogRecordLogDatum](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'logDatum' field")
	}

	if err := WriteOptionalField[BACnetStatusFlagsTagged](ctx, "statusFlags", GetRef(m.GetStatusFlags()), WriteComplex[BACnetStatusFlagsTagged](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'statusFlags' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogRecord")
	}
	return nil
}

func (m *_BACnetLogRecord) IsBACnetLogRecord() {}

func (m *_BACnetLogRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
