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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SysexCommandSysexRealtime is the corresponding interface of SysexCommandSysexRealtime
type SysexCommandSysexRealtime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandSysexRealtimeExactly can be used when we want exactly this type and not a type which fulfills SysexCommandSysexRealtime.
// This is useful for switch cases.
type SysexCommandSysexRealtimeExactly interface {
	SysexCommandSysexRealtime
	isSysexCommandSysexRealtime() bool
}

// _SysexCommandSysexRealtime is the data-structure of this message
type _SysexCommandSysexRealtime struct {
	*_SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandSysexRealtime) GetCommandType() uint8 {
	return 0x7F
}

func (m *_SysexCommandSysexRealtime) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandSysexRealtime) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandSysexRealtime) GetParent() SysexCommand {
	return m._SysexCommand
}

// NewSysexCommandSysexRealtime factory function for _SysexCommandSysexRealtime
func NewSysexCommandSysexRealtime() *_SysexCommandSysexRealtime {
	_result := &_SysexCommandSysexRealtime{
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandSysexRealtime(structType any) SysexCommandSysexRealtime {
	if casted, ok := structType.(SysexCommandSysexRealtime); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandSysexRealtime); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandSysexRealtime) GetTypeName() string {
	return "SysexCommandSysexRealtime"
}

func (m *_SysexCommandSysexRealtime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandSysexRealtime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SysexCommandSysexRealtimeParse(ctx context.Context, theBytes []byte, response bool) (SysexCommandSysexRealtime, error) {
	return SysexCommandSysexRealtimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func SysexCommandSysexRealtimeParseWithBufferProducer(response bool) func(ctx context.Context, readBuffer utils.ReadBuffer) (SysexCommandSysexRealtime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SysexCommandSysexRealtime, error) {
		return SysexCommandSysexRealtimeParseWithBuffer(ctx, readBuffer, response)
	}
}

func SysexCommandSysexRealtimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (SysexCommandSysexRealtime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SysexCommandSysexRealtime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandSysexRealtime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSysexRealtime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandSysexRealtime")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandSysexRealtime{
		_SysexCommand: &_SysexCommand{},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandSysexRealtime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandSysexRealtime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSysexRealtime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandSysexRealtime")
		}

		if popErr := writeBuffer.PopContext("SysexCommandSysexRealtime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandSysexRealtime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandSysexRealtime) isSysexCommandSysexRealtime() bool {
	return true
}

func (m *_SysexCommandSysexRealtime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
