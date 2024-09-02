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

// OpcuaNodeIdServicesVariableClose is an enum
type OpcuaNodeIdServicesVariableClose int32

type IOpcuaNodeIdServicesVariableClose interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableClose_CloseMethodType_InputArguments           OpcuaNodeIdServicesVariableClose = 11742
	OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_OutputArguments OpcuaNodeIdServicesVariableClose = 12517
	OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_InputArguments  OpcuaNodeIdServicesVariableClose = 12704
	OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_InputArguments  OpcuaNodeIdServicesVariableClose = 15801
	OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_OutputArguments OpcuaNodeIdServicesVariableClose = 15802
)

var OpcuaNodeIdServicesVariableCloseValues []OpcuaNodeIdServicesVariableClose

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableCloseValues = []OpcuaNodeIdServicesVariableClose{
		OpcuaNodeIdServicesVariableClose_CloseMethodType_InputArguments,
		OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_OutputArguments,
		OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_InputArguments,
		OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_InputArguments,
		OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_OutputArguments,
	}
}

func OpcuaNodeIdServicesVariableCloseByValue(value int32) (enum OpcuaNodeIdServicesVariableClose, ok bool) {
	switch value {
	case 11742:
		return OpcuaNodeIdServicesVariableClose_CloseMethodType_InputArguments, true
	case 12517:
		return OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_OutputArguments, true
	case 12704:
		return OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_InputArguments, true
	case 15801:
		return OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_InputArguments, true
	case 15802:
		return OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableCloseByName(value string) (enum OpcuaNodeIdServicesVariableClose, ok bool) {
	switch value {
	case "CloseMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableClose_CloseMethodType_InputArguments, true
	case "CloseAndUpdateMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_OutputArguments, true
	case "CloseAndUpdateMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_InputArguments, true
	case "CloseAndCommitMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_InputArguments, true
	case "CloseAndCommitMethodType_OutputArguments":
		return OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_OutputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableCloseKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableCloseValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableClose(structType any) OpcuaNodeIdServicesVariableClose {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableClose {
		if sOpcuaNodeIdServicesVariableClose, ok := typ.(OpcuaNodeIdServicesVariableClose); ok {
			return sOpcuaNodeIdServicesVariableClose
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableClose) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableClose) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableCloseParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableClose, error) {
	return OpcuaNodeIdServicesVariableCloseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableCloseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableClose, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableClose", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableClose")
	}
	if enum, ok := OpcuaNodeIdServicesVariableCloseByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableClose")
		return OpcuaNodeIdServicesVariableClose(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableClose) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableClose) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableClose", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableClose) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableClose) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableClose_CloseMethodType_InputArguments:
		return "CloseMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_OutputArguments:
		return "CloseAndUpdateMethodType_OutputArguments"
	case OpcuaNodeIdServicesVariableClose_CloseAndUpdateMethodType_InputArguments:
		return "CloseAndUpdateMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_InputArguments:
		return "CloseAndCommitMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableClose_CloseAndCommitMethodType_OutputArguments:
		return "CloseAndCommitMethodType_OutputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableClose) String() string {
	return e.PLC4XEnumName()
}
