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

package bacnetip

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"

	"github.com/pkg/errors"

	"github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
)

type Unsigned16 struct {
	*Atomic[uint16]
	*CommonMath
}

func NewUnsigned16(arg Arg) (*Unsigned16, error) {
	i := &Unsigned16{}
	i.Atomic = NewAtomic[uint16](i)

	if arg == nil {
		return i, nil
	}
	switch arg := arg.(type) {
	case *tag:
		err := i.Decode(arg)
		if err != nil {
			return nil, errors.Wrap(err, "error decoding")
		}
		return i, nil
	case uint16:
		i.value = arg
	case uint:
		if !i.IsValid(arg) {
			return nil, errors.Errorf("invalid Unsigned16: %d", arg)
		}
		i.value = uint16(arg)
	case int:
		if !i.IsValid(arg) {
			return nil, errors.Errorf("invalid Unsigned16: %d", arg)
		}
		i.value = uint16(arg)
	case *Unsigned16:
		i.value = arg.value
	default:
		return nil, errors.Errorf("invalid constructor datatype: %T", arg)
	}

	return i, nil
}

func (i *Unsigned16) Encode(arg Arg) error {
	tag, ok := arg.(Tag)
	if !ok {
		return errors.Errorf("%T is not a Tag", arg)
	}
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, uint32(i.value))

	// reduce the value to the smallest number of bytes
	for len(data) > 1 && data[0] == 0 {
		data = data[1:]
	}

	tag.setAppData(uint(model.BACnetDataType_UNSIGNED_INTEGER), data)
	return nil
}

func (i *Unsigned16) Decode(arg Arg) error {
	tag, ok := arg.(Tag)
	if !ok {
		return errors.Errorf("%T is not a Tag", arg)
	}
	if tag.GetTagClass() != model.TagClass_APPLICATION_TAGS || tag.GetTagNumber() != uint(model.BACnetDataType_UNSIGNED_INTEGER) {
		return errors.New("Unsigned16 application tag required")
	}
	if len(tag.GetTagData()) == 0 {
		return errors.New("invalid tag length")
	}

	tagData := tag.GetTagData()

	// get the data
	rslt := uint16(0)
	for _, c := range tagData {
		rslt = (rslt << 8) + uint16(c)
	}

	// save the result
	i.value = rslt
	return nil
}

func (i *Unsigned16) IsValid(arg any) bool {
	switch arg := arg.(type) {
	case string:
		_, err := strconv.Atoi(arg)
		return err == nil
	case int:
		return arg > 0 && arg < 65536
	case int16:
		return arg > 0
	case uint:
		return arg <= math.MaxUint32
	case uint16:
		return true
	default:
		return false
	}
}

func (i *Unsigned16) String() string {
	return fmt.Sprintf("Unsigned16(%d)", i.value)
}
