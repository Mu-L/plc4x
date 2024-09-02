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

// OpcuaNodeIdServicesVariableId is an enum
type OpcuaNodeIdServicesVariableId int32

type IOpcuaNodeIdServicesVariableId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableId_IdType_EnumStrings OpcuaNodeIdServicesVariableId = 7591
)

var OpcuaNodeIdServicesVariableIdValues []OpcuaNodeIdServicesVariableId

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableIdValues = []OpcuaNodeIdServicesVariableId{
		OpcuaNodeIdServicesVariableId_IdType_EnumStrings,
	}
}

func OpcuaNodeIdServicesVariableIdByValue(value int32) (enum OpcuaNodeIdServicesVariableId, ok bool) {
	switch value {
	case 7591:
		return OpcuaNodeIdServicesVariableId_IdType_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableIdByName(value string) (enum OpcuaNodeIdServicesVariableId, ok bool) {
	switch value {
	case "IdType_EnumStrings":
		return OpcuaNodeIdServicesVariableId_IdType_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableIdKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableIdValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableId(structType any) OpcuaNodeIdServicesVariableId {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableId {
		if sOpcuaNodeIdServicesVariableId, ok := typ.(OpcuaNodeIdServicesVariableId); ok {
			return sOpcuaNodeIdServicesVariableId
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableId) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableIdParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableId, error) {
	return OpcuaNodeIdServicesVariableIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableId, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableId", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableId")
	}
	if enum, ok := OpcuaNodeIdServicesVariableIdByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableId")
		return OpcuaNodeIdServicesVariableId(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableId", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableId) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableId) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableId_IdType_EnumStrings:
		return "IdType_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableId) String() string {
	return e.PLC4XEnumName()
}
