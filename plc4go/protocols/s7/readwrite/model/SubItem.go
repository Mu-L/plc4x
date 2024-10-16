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

// SubItem is the corresponding interface of SubItem
type SubItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetBytesToRead returns BytesToRead (property field)
	GetBytesToRead() uint8
	// GetDbNumber returns DbNumber (property field)
	GetDbNumber() uint16
	// GetStartAddress returns StartAddress (property field)
	GetStartAddress() uint16
	// IsSubItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSubItem()
	// CreateBuilder creates a SubItemBuilder
	CreateSubItemBuilder() SubItemBuilder
}

// _SubItem is the data-structure of this message
type _SubItem struct {
	BytesToRead  uint8
	DbNumber     uint16
	StartAddress uint16
}

var _ SubItem = (*_SubItem)(nil)

// NewSubItem factory function for _SubItem
func NewSubItem(bytesToRead uint8, dbNumber uint16, startAddress uint16) *_SubItem {
	return &_SubItem{BytesToRead: bytesToRead, DbNumber: dbNumber, StartAddress: startAddress}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SubItemBuilder is a builder for SubItem
type SubItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bytesToRead uint8, dbNumber uint16, startAddress uint16) SubItemBuilder
	// WithBytesToRead adds BytesToRead (property field)
	WithBytesToRead(uint8) SubItemBuilder
	// WithDbNumber adds DbNumber (property field)
	WithDbNumber(uint16) SubItemBuilder
	// WithStartAddress adds StartAddress (property field)
	WithStartAddress(uint16) SubItemBuilder
	// Build builds the SubItem or returns an error if something is wrong
	Build() (SubItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SubItem
}

// NewSubItemBuilder() creates a SubItemBuilder
func NewSubItemBuilder() SubItemBuilder {
	return &_SubItemBuilder{_SubItem: new(_SubItem)}
}

type _SubItemBuilder struct {
	*_SubItem

	err *utils.MultiError
}

var _ (SubItemBuilder) = (*_SubItemBuilder)(nil)

func (b *_SubItemBuilder) WithMandatoryFields(bytesToRead uint8, dbNumber uint16, startAddress uint16) SubItemBuilder {
	return b.WithBytesToRead(bytesToRead).WithDbNumber(dbNumber).WithStartAddress(startAddress)
}

func (b *_SubItemBuilder) WithBytesToRead(bytesToRead uint8) SubItemBuilder {
	b.BytesToRead = bytesToRead
	return b
}

func (b *_SubItemBuilder) WithDbNumber(dbNumber uint16) SubItemBuilder {
	b.DbNumber = dbNumber
	return b
}

func (b *_SubItemBuilder) WithStartAddress(startAddress uint16) SubItemBuilder {
	b.StartAddress = startAddress
	return b
}

func (b *_SubItemBuilder) Build() (SubItem, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SubItem.deepCopy(), nil
}

func (b *_SubItemBuilder) MustBuild() SubItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SubItemBuilder) DeepCopy() any {
	_copy := b.CreateSubItemBuilder().(*_SubItemBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSubItemBuilder creates a SubItemBuilder
func (b *_SubItem) CreateSubItemBuilder() SubItemBuilder {
	if b == nil {
		return NewSubItemBuilder()
	}
	return &_SubItemBuilder{_SubItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SubItem) GetBytesToRead() uint8 {
	return m.BytesToRead
}

func (m *_SubItem) GetDbNumber() uint16 {
	return m.DbNumber
}

func (m *_SubItem) GetStartAddress() uint16 {
	return m.StartAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSubItem(structType any) SubItem {
	if casted, ok := structType.(SubItem); ok {
		return casted
	}
	if casted, ok := structType.(*SubItem); ok {
		return *casted
	}
	return nil
}

func (m *_SubItem) GetTypeName() string {
	return "SubItem"
}

func (m *_SubItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (bytesToRead)
	lengthInBits += 8

	// Simple field (dbNumber)
	lengthInBits += 16

	// Simple field (startAddress)
	lengthInBits += 16

	return lengthInBits
}

func (m *_SubItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SubItemParse(ctx context.Context, theBytes []byte) (SubItem, error) {
	return SubItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SubItemParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SubItem, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SubItem, error) {
		return SubItemParseWithBuffer(ctx, readBuffer)
	}
}

func SubItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SubItem, error) {
	v, err := (&_SubItem{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_SubItem) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__subItem SubItem, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bytesToRead, err := ReadSimpleField(ctx, "bytesToRead", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bytesToRead' field"))
	}
	m.BytesToRead = bytesToRead

	dbNumber, err := ReadSimpleField(ctx, "dbNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dbNumber' field"))
	}
	m.DbNumber = dbNumber

	startAddress, err := ReadSimpleField(ctx, "startAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startAddress' field"))
	}
	m.StartAddress = startAddress

	if closeErr := readBuffer.CloseContext("SubItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubItem")
	}

	return m, nil
}

func (m *_SubItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SubItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SubItem")
	}

	if err := WriteSimpleField[uint8](ctx, "bytesToRead", m.GetBytesToRead(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'bytesToRead' field")
	}

	if err := WriteSimpleField[uint16](ctx, "dbNumber", m.GetDbNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'dbNumber' field")
	}

	if err := WriteSimpleField[uint16](ctx, "startAddress", m.GetStartAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'startAddress' field")
	}

	if popErr := writeBuffer.PopContext("SubItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SubItem")
	}
	return nil
}

func (m *_SubItem) IsSubItem() {}

func (m *_SubItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SubItem) deepCopy() *_SubItem {
	if m == nil {
		return nil
	}
	_SubItemCopy := &_SubItem{
		m.BytesToRead,
		m.DbNumber,
		m.StartAddress,
	}
	return _SubItemCopy
}

func (m *_SubItem) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
