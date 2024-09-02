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

// CBusPointToPointToMultiPointCommand is the corresponding interface of CBusPointToPointToMultiPointCommand
type CBusPointToPointToMultiPointCommand interface {
	CBusPointToPointToMultiPointCommandContract
	CBusPointToPointToMultiPointCommandRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// CBusPointToPointToMultiPointCommandContract provides a set of functions which can be overwritten by a sub struct
type CBusPointToPointToMultiPointCommandContract interface {
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetNetworkRoute returns NetworkRoute (property field)
	GetNetworkRoute() NetworkRoute
	// GetPeekedApplication returns PeekedApplication (property field)
	GetPeekedApplication() byte
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
}

// CBusPointToPointToMultiPointCommandRequirements provides a set of functions which need to be implemented by a sub struct
type CBusPointToPointToMultiPointCommandRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedApplication returns PeekedApplication (discriminator field)
	GetPeekedApplication() byte
}

// CBusPointToPointToMultiPointCommandExactly can be used when we want exactly this type and not a type which fulfills CBusPointToPointToMultiPointCommand.
// This is useful for switch cases.
type CBusPointToPointToMultiPointCommandExactly interface {
	CBusPointToPointToMultiPointCommand
	isCBusPointToPointToMultiPointCommand() bool
}

// _CBusPointToPointToMultiPointCommand is the data-structure of this message
type _CBusPointToPointToMultiPointCommand struct {
	_SubType          CBusPointToPointToMultiPointCommand
	BridgeAddress     BridgeAddress
	NetworkRoute      NetworkRoute
	PeekedApplication byte

	// Arguments.
	CBusOptions CBusOptions
}

var _ CBusPointToPointToMultiPointCommandContract = (*_CBusPointToPointToMultiPointCommand)(nil)

type CBusPointToPointToMultiPointCommandChild interface {
	utils.Serializable

	GetParent() *CBusPointToPointToMultiPointCommand

	GetTypeName() string
	CBusPointToPointToMultiPointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToPointToMultiPointCommand) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_CBusPointToPointToMultiPointCommand) GetNetworkRoute() NetworkRoute {
	return m.NetworkRoute
}

func (m *_CBusPointToPointToMultiPointCommand) GetPeekedApplication() byte {
	return m.PeekedApplication
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToPointToMultiPointCommand factory function for _CBusPointToPointToMultiPointCommand
func NewCBusPointToPointToMultiPointCommand(bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte, cBusOptions CBusOptions) *_CBusPointToPointToMultiPointCommand {
	return &_CBusPointToPointToMultiPointCommand{BridgeAddress: bridgeAddress, NetworkRoute: networkRoute, PeekedApplication: peekedApplication, CBusOptions: cBusOptions}
}

// Deprecated: use the interface for direct cast
func CastCBusPointToPointToMultiPointCommand(structType any) CBusPointToPointToMultiPointCommand {
	if casted, ok := structType.(CBusPointToPointToMultiPointCommand); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToPointToMultiPointCommand); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToPointToMultiPointCommand) GetTypeName() string {
	return "CBusPointToPointToMultiPointCommand"
}

func (m *_CBusPointToPointToMultiPointCommand) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (bridgeAddress)
	lengthInBits += m.BridgeAddress.GetLengthInBits(ctx)

	// Simple field (networkRoute)
	lengthInBits += m.NetworkRoute.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToPointToMultiPointCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CBusPointToPointToMultiPointCommandParse[T CBusPointToPointToMultiPointCommand](ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (T, error) {
	return CBusPointToPointToMultiPointCommandParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusPointToPointToMultiPointCommandParseWithBufferProducer[T CBusPointToPointToMultiPointCommand](cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CBusPointToPointToMultiPointCommandParseWithBuffer[T](ctx, readBuffer, cBusOptions)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func CBusPointToPointToMultiPointCommandParseWithBuffer[T CBusPointToPointToMultiPointCommand](ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (T, error) {
	v, err := (&_CBusPointToPointToMultiPointCommand{}).parse(ctx, readBuffer, cBusOptions)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_CBusPointToPointToMultiPointCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (__cBusPointToPointToMultiPointCommand CBusPointToPointToMultiPointCommand, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToPointToMultiPointCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToPointToMultiPointCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bridgeAddress, err := ReadSimpleField[BridgeAddress](ctx, "bridgeAddress", ReadComplex[BridgeAddress](BridgeAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeAddress' field"))
	}
	m.BridgeAddress = bridgeAddress

	networkRoute, err := ReadSimpleField[NetworkRoute](ctx, "networkRoute", ReadComplex[NetworkRoute](NetworkRouteParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkRoute' field"))
	}
	m.NetworkRoute = networkRoute

	peekedApplication, err := ReadPeekField[byte](ctx, "peekedApplication", ReadByte(readBuffer, 8), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedApplication' field"))
	}
	m.PeekedApplication = peekedApplication

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CBusPointToPointToMultiPointCommand
	switch {
	case peekedApplication == 0xFF: // CBusPointToPointToMultiPointCommandStatus
		if _child, err = (&_CBusPointToPointToMultiPointCommandStatus{}).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusPointToPointToMultiPointCommandStatus for type-switch of CBusPointToPointToMultiPointCommand")
		}
	case 0 == 0: // CBusPointToPointToMultiPointCommandNormal
		if _child, err = (&_CBusPointToPointToMultiPointCommandNormal{}).parse(ctx, readBuffer, m, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusPointToPointToMultiPointCommandNormal for type-switch of CBusPointToPointToMultiPointCommand")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedApplication=%v]", peekedApplication)
	}

	if closeErr := readBuffer.CloseContext("CBusPointToPointToMultiPointCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToPointToMultiPointCommand")
	}

	return _child, nil
}

func (pm *_CBusPointToPointToMultiPointCommand) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusPointToPointToMultiPointCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CBusPointToPointToMultiPointCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusPointToPointToMultiPointCommand")
	}

	if err := WriteSimpleField[BridgeAddress](ctx, "bridgeAddress", m.GetBridgeAddress(), WriteComplex[BridgeAddress](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'bridgeAddress' field")
	}

	if err := WriteSimpleField[NetworkRoute](ctx, "networkRoute", m.GetNetworkRoute(), WriteComplex[NetworkRoute](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'networkRoute' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusPointToPointToMultiPointCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusPointToPointToMultiPointCommand")
	}
	return nil
}

////
// Arguments Getter

func (m *_CBusPointToPointToMultiPointCommand) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_CBusPointToPointToMultiPointCommand) isCBusPointToPointToMultiPointCommand() bool {
	return true
}
