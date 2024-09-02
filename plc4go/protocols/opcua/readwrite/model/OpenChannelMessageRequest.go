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

// OpenChannelMessageRequest is the corresponding interface of OpenChannelMessageRequest
type OpenChannelMessageRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	OpenChannelMessage
	// GetSecureChannelId returns SecureChannelId (property field)
	GetSecureChannelId() int32
	// GetEndpoint returns Endpoint (property field)
	GetEndpoint() PascalString
	// GetSenderCertificate returns SenderCertificate (property field)
	GetSenderCertificate() PascalByteString
	// GetReceiverCertificateThumbprint returns ReceiverCertificateThumbprint (property field)
	GetReceiverCertificateThumbprint() PascalByteString
}

// OpenChannelMessageRequestExactly can be used when we want exactly this type and not a type which fulfills OpenChannelMessageRequest.
// This is useful for switch cases.
type OpenChannelMessageRequestExactly interface {
	OpenChannelMessageRequest
	isOpenChannelMessageRequest() bool
}

// _OpenChannelMessageRequest is the data-structure of this message
type _OpenChannelMessageRequest struct {
	OpenChannelMessageContract
	SecureChannelId               int32
	Endpoint                      PascalString
	SenderCertificate             PascalByteString
	ReceiverCertificateThumbprint PascalByteString
}

var _ OpenChannelMessageRequest = (*_OpenChannelMessageRequest)(nil)
var _ OpenChannelMessageRequirements = (*_OpenChannelMessageRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpenChannelMessageRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpenChannelMessageRequest) GetParent() OpenChannelMessageContract {
	return m.OpenChannelMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpenChannelMessageRequest) GetSecureChannelId() int32 {
	return m.SecureChannelId
}

func (m *_OpenChannelMessageRequest) GetEndpoint() PascalString {
	return m.Endpoint
}

func (m *_OpenChannelMessageRequest) GetSenderCertificate() PascalByteString {
	return m.SenderCertificate
}

func (m *_OpenChannelMessageRequest) GetReceiverCertificateThumbprint() PascalByteString {
	return m.ReceiverCertificateThumbprint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpenChannelMessageRequest factory function for _OpenChannelMessageRequest
func NewOpenChannelMessageRequest(secureChannelId int32, endpoint PascalString, senderCertificate PascalByteString, receiverCertificateThumbprint PascalByteString) *_OpenChannelMessageRequest {
	_result := &_OpenChannelMessageRequest{
		OpenChannelMessageContract:    NewOpenChannelMessage(),
		SecureChannelId:               secureChannelId,
		Endpoint:                      endpoint,
		SenderCertificate:             senderCertificate,
		ReceiverCertificateThumbprint: receiverCertificateThumbprint,
	}
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpenChannelMessageRequest(structType any) OpenChannelMessageRequest {
	if casted, ok := structType.(OpenChannelMessageRequest); ok {
		return casted
	}
	if casted, ok := structType.(*OpenChannelMessageRequest); ok {
		return *casted
	}
	return nil
}

func (m *_OpenChannelMessageRequest) GetTypeName() string {
	return "OpenChannelMessageRequest"
}

func (m *_OpenChannelMessageRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.OpenChannelMessageContract.(*_OpenChannelMessage).getLengthInBits(ctx))

	// Simple field (secureChannelId)
	lengthInBits += 32

	// Simple field (endpoint)
	lengthInBits += m.Endpoint.GetLengthInBits(ctx)

	// Simple field (senderCertificate)
	lengthInBits += m.SenderCertificate.GetLengthInBits(ctx)

	// Simple field (receiverCertificateThumbprint)
	lengthInBits += m.ReceiverCertificateThumbprint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpenChannelMessageRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpenChannelMessageRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_OpenChannelMessage, response bool) (__openChannelMessageRequest OpenChannelMessageRequest, err error) {
	m.OpenChannelMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpenChannelMessageRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpenChannelMessageRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	secureChannelId, err := ReadSimpleField(ctx, "secureChannelId", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'secureChannelId' field"))
	}
	m.SecureChannelId = secureChannelId

	endpoint, err := ReadSimpleField[PascalString](ctx, "endpoint", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endpoint' field"))
	}
	m.Endpoint = endpoint

	senderCertificate, err := ReadSimpleField[PascalByteString](ctx, "senderCertificate", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'senderCertificate' field"))
	}
	m.SenderCertificate = senderCertificate

	receiverCertificateThumbprint, err := ReadSimpleField[PascalByteString](ctx, "receiverCertificateThumbprint", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'receiverCertificateThumbprint' field"))
	}
	m.ReceiverCertificateThumbprint = receiverCertificateThumbprint

	if closeErr := readBuffer.CloseContext("OpenChannelMessageRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpenChannelMessageRequest")
	}

	return m, nil
}

func (m *_OpenChannelMessageRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpenChannelMessageRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpenChannelMessageRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpenChannelMessageRequest")
		}

		if err := WriteSimpleField[int32](ctx, "secureChannelId", m.GetSecureChannelId(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'secureChannelId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "endpoint", m.GetEndpoint(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'endpoint' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "senderCertificate", m.GetSenderCertificate(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'senderCertificate' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "receiverCertificateThumbprint", m.GetReceiverCertificateThumbprint(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'receiverCertificateThumbprint' field")
		}

		if popErr := writeBuffer.PopContext("OpenChannelMessageRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpenChannelMessageRequest")
		}
		return nil
	}
	return m.OpenChannelMessageContract.(*_OpenChannelMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpenChannelMessageRequest) isOpenChannelMessageRequest() bool {
	return true
}

func (m *_OpenChannelMessageRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
