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

package test_constructed_data

import (
	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip"
)

type SequenceEqualityRequirements interface {
	GetSequenceElements() []Element
}

type SequenceEquality struct {
	_requirements SequenceEqualityRequirements
}

func NewSequenceEquality(requirements SequenceEqualityRequirements) *SequenceEquality {
	if requirements == nil {
		panic("requirements cannot be nil")
	}
	return &SequenceEquality{_requirements: requirements}
}

func (s *SequenceEquality) Equals(other any) bool {
	for _, element := range s._requirements.GetSequenceElements() {
		if !element.IsOptional() && true {
			panic("what??")
		}
	}
	return true
}

type EmptySequence struct {
	*Sequence
	*SequenceEquality
}

func NewEmptySequence(kwargs KWArgs) (*EmptySequence, error) {
	e := &EmptySequence{}
	var err error
	e.Sequence, err = NewSequence(kwargs, WithSequenceContract(e))
	if err != nil {
		return nil, errors.Wrap(err, "could not create sequence")
	}
	e.SequenceEquality = NewSequenceEquality(e)
	return e, nil
}

func (e *EmptySequence) SetSequence(sequence *Sequence) {
	e.Sequence = sequence
}

type SimpleSequence struct {
	*Sequence
	*SequenceEquality

	sequenceElements []Element
}

func NewSimpleSequence(kwargs KWArgs) (*SimpleSequence, error) {
	s := &SimpleSequence{
		sequenceElements: []Element{
			NewElement("hydrogen", func(args Args, _ KWArgs) (interface{ Encode(Arg) error }, error) {
				var arg any
				if len(args) == 1 {
					arg = args[0]
				}
				boolean, err := NewBoolean(arg)
				return boolean, err
			}),
		},
	}
	var err error
	s.Sequence, err = NewSequence(kwargs, WithSequenceContract(s))
	if err != nil {
		return nil, errors.Wrap(err, "could not create sequence")
	}
	s.SequenceEquality = NewSequenceEquality(s)
	return s, nil
}

func (e *SimpleSequence) SetSequence(sequence *Sequence) {
	e.Sequence = sequence
}

func (e *SimpleSequence) GetSequenceElements() []Element {
	return e.sequenceElements
}

type CompoundSequence1 struct {
	*Sequence
	*SequenceEquality
}

type CompoundSequence2 struct {
	*Sequence
	*SequenceEquality
}
