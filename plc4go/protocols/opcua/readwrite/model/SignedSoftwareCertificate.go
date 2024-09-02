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

// SignedSoftwareCertificate is the corresponding interface of SignedSoftwareCertificate
type SignedSoftwareCertificate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetCertificateData returns CertificateData (property field)
	GetCertificateData() PascalByteString
	// GetSignature returns Signature (property field)
	GetSignature() PascalByteString
}

// SignedSoftwareCertificateExactly can be used when we want exactly this type and not a type which fulfills SignedSoftwareCertificate.
// This is useful for switch cases.
type SignedSoftwareCertificateExactly interface {
	SignedSoftwareCertificate
	isSignedSoftwareCertificate() bool
}

// _SignedSoftwareCertificate is the data-structure of this message
type _SignedSoftwareCertificate struct {
	ExtensionObjectDefinitionContract
	CertificateData PascalByteString
	Signature       PascalByteString
}

var _ SignedSoftwareCertificate = (*_SignedSoftwareCertificate)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SignedSoftwareCertificate)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SignedSoftwareCertificate) GetIdentifier() string {
	return "346"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SignedSoftwareCertificate) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SignedSoftwareCertificate) GetCertificateData() PascalByteString {
	return m.CertificateData
}

func (m *_SignedSoftwareCertificate) GetSignature() PascalByteString {
	return m.Signature
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSignedSoftwareCertificate factory function for _SignedSoftwareCertificate
func NewSignedSoftwareCertificate(certificateData PascalByteString, signature PascalByteString) *_SignedSoftwareCertificate {
	_result := &_SignedSoftwareCertificate{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		CertificateData:                   certificateData,
		Signature:                         signature,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastSignedSoftwareCertificate(structType any) SignedSoftwareCertificate {
	if casted, ok := structType.(SignedSoftwareCertificate); ok {
		return casted
	}
	if casted, ok := structType.(*SignedSoftwareCertificate); ok {
		return *casted
	}
	return nil
}

func (m *_SignedSoftwareCertificate) GetTypeName() string {
	return "SignedSoftwareCertificate"
}

func (m *_SignedSoftwareCertificate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (certificateData)
	lengthInBits += m.CertificateData.GetLengthInBits(ctx)

	// Simple field (signature)
	lengthInBits += m.Signature.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SignedSoftwareCertificate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SignedSoftwareCertificate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__signedSoftwareCertificate SignedSoftwareCertificate, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SignedSoftwareCertificate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SignedSoftwareCertificate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	certificateData, err := ReadSimpleField[PascalByteString](ctx, "certificateData", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'certificateData' field"))
	}
	m.CertificateData = certificateData

	signature, err := ReadSimpleField[PascalByteString](ctx, "signature", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'signature' field"))
	}
	m.Signature = signature

	if closeErr := readBuffer.CloseContext("SignedSoftwareCertificate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SignedSoftwareCertificate")
	}

	return m, nil
}

func (m *_SignedSoftwareCertificate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SignedSoftwareCertificate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SignedSoftwareCertificate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SignedSoftwareCertificate")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "certificateData", m.GetCertificateData(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'certificateData' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "signature", m.GetSignature(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'signature' field")
		}

		if popErr := writeBuffer.PopContext("SignedSoftwareCertificate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SignedSoftwareCertificate")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SignedSoftwareCertificate) isSignedSoftwareCertificate() bool {
	return true
}

func (m *_SignedSoftwareCertificate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
