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

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// DataChangeTrigger is an enum
type DataChangeTrigger uint32

type IDataChangeTrigger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	DataChangeTrigger_dataChangeTriggerStatus               DataChangeTrigger = 0
	DataChangeTrigger_dataChangeTriggerStatusValue          DataChangeTrigger = 1
	DataChangeTrigger_dataChangeTriggerStatusValueTimestamp DataChangeTrigger = 2
)

var DataChangeTriggerValues []DataChangeTrigger

func init() {
	_ = errors.New
	DataChangeTriggerValues = []DataChangeTrigger{
		DataChangeTrigger_dataChangeTriggerStatus,
		DataChangeTrigger_dataChangeTriggerStatusValue,
		DataChangeTrigger_dataChangeTriggerStatusValueTimestamp,
	}
}

func DataChangeTriggerByValue(value uint32) (enum DataChangeTrigger, ok bool) {
	switch value {
	case 0:
		return DataChangeTrigger_dataChangeTriggerStatus, true
	case 1:
		return DataChangeTrigger_dataChangeTriggerStatusValue, true
	case 2:
		return DataChangeTrigger_dataChangeTriggerStatusValueTimestamp, true
	}
	return 0, false
}

func DataChangeTriggerByName(value string) (enum DataChangeTrigger, ok bool) {
	switch value {
	case "dataChangeTriggerStatus":
		return DataChangeTrigger_dataChangeTriggerStatus, true
	case "dataChangeTriggerStatusValue":
		return DataChangeTrigger_dataChangeTriggerStatusValue, true
	case "dataChangeTriggerStatusValueTimestamp":
		return DataChangeTrigger_dataChangeTriggerStatusValueTimestamp, true
	}
	return 0, false
}

func DataChangeTriggerKnows(value uint32) bool {
	for _, typeValue := range DataChangeTriggerValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDataChangeTrigger(structType any) DataChangeTrigger {
	castFunc := func(typ any) DataChangeTrigger {
		if sDataChangeTrigger, ok := typ.(DataChangeTrigger); ok {
			return sDataChangeTrigger
		}
		return 0
	}
	return castFunc(structType)
}

func (m DataChangeTrigger) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m DataChangeTrigger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DataChangeTriggerParse(ctx context.Context, theBytes []byte) (DataChangeTrigger, error) {
	return DataChangeTriggerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DataChangeTriggerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DataChangeTrigger, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("DataChangeTrigger", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DataChangeTrigger")
	}
	if enum, ok := DataChangeTriggerByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return DataChangeTrigger(val), nil
	} else {
		return enum, nil
	}
}

func (e DataChangeTrigger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DataChangeTrigger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("DataChangeTrigger", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DataChangeTrigger) PLC4XEnumName() string {
	switch e {
	case DataChangeTrigger_dataChangeTriggerStatus:
		return "dataChangeTriggerStatus"
	case DataChangeTrigger_dataChangeTriggerStatusValue:
		return "dataChangeTriggerStatusValue"
	case DataChangeTrigger_dataChangeTriggerStatusValueTimestamp:
		return "dataChangeTriggerStatusValueTimestamp"
	}
	return ""
}

func (e DataChangeTrigger) String() string {
	return e.PLC4XEnumName()
}
