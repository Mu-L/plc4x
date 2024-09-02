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

// CreateSessionResponse is the corresponding interface of CreateSessionResponse
type CreateSessionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetSessionId returns SessionId (property field)
	GetSessionId() NodeId
	// GetAuthenticationToken returns AuthenticationToken (property field)
	GetAuthenticationToken() NodeId
	// GetRevisedSessionTimeout returns RevisedSessionTimeout (property field)
	GetRevisedSessionTimeout() float64
	// GetServerNonce returns ServerNonce (property field)
	GetServerNonce() PascalByteString
	// GetServerCertificate returns ServerCertificate (property field)
	GetServerCertificate() PascalByteString
	// GetNoOfServerEndpoints returns NoOfServerEndpoints (property field)
	GetNoOfServerEndpoints() int32
	// GetServerEndpoints returns ServerEndpoints (property field)
	GetServerEndpoints() []ExtensionObjectDefinition
	// GetNoOfServerSoftwareCertificates returns NoOfServerSoftwareCertificates (property field)
	GetNoOfServerSoftwareCertificates() int32
	// GetServerSoftwareCertificates returns ServerSoftwareCertificates (property field)
	GetServerSoftwareCertificates() []ExtensionObjectDefinition
	// GetServerSignature returns ServerSignature (property field)
	GetServerSignature() ExtensionObjectDefinition
	// GetMaxRequestMessageSize returns MaxRequestMessageSize (property field)
	GetMaxRequestMessageSize() uint32
	// IsCreateSessionResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCreateSessionResponse()
}

// _CreateSessionResponse is the data-structure of this message
type _CreateSessionResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader                 ExtensionObjectDefinition
	SessionId                      NodeId
	AuthenticationToken            NodeId
	RevisedSessionTimeout          float64
	ServerNonce                    PascalByteString
	ServerCertificate              PascalByteString
	NoOfServerEndpoints            int32
	ServerEndpoints                []ExtensionObjectDefinition
	NoOfServerSoftwareCertificates int32
	ServerSoftwareCertificates     []ExtensionObjectDefinition
	ServerSignature                ExtensionObjectDefinition
	MaxRequestMessageSize          uint32
}

var _ CreateSessionResponse = (*_CreateSessionResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CreateSessionResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateSessionResponse) GetIdentifier() string {
	return "464"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateSessionResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateSessionResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_CreateSessionResponse) GetSessionId() NodeId {
	return m.SessionId
}

func (m *_CreateSessionResponse) GetAuthenticationToken() NodeId {
	return m.AuthenticationToken
}

func (m *_CreateSessionResponse) GetRevisedSessionTimeout() float64 {
	return m.RevisedSessionTimeout
}

func (m *_CreateSessionResponse) GetServerNonce() PascalByteString {
	return m.ServerNonce
}

func (m *_CreateSessionResponse) GetServerCertificate() PascalByteString {
	return m.ServerCertificate
}

func (m *_CreateSessionResponse) GetNoOfServerEndpoints() int32 {
	return m.NoOfServerEndpoints
}

func (m *_CreateSessionResponse) GetServerEndpoints() []ExtensionObjectDefinition {
	return m.ServerEndpoints
}

func (m *_CreateSessionResponse) GetNoOfServerSoftwareCertificates() int32 {
	return m.NoOfServerSoftwareCertificates
}

func (m *_CreateSessionResponse) GetServerSoftwareCertificates() []ExtensionObjectDefinition {
	return m.ServerSoftwareCertificates
}

func (m *_CreateSessionResponse) GetServerSignature() ExtensionObjectDefinition {
	return m.ServerSignature
}

func (m *_CreateSessionResponse) GetMaxRequestMessageSize() uint32 {
	return m.MaxRequestMessageSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCreateSessionResponse factory function for _CreateSessionResponse
func NewCreateSessionResponse(responseHeader ExtensionObjectDefinition, sessionId NodeId, authenticationToken NodeId, revisedSessionTimeout float64, serverNonce PascalByteString, serverCertificate PascalByteString, noOfServerEndpoints int32, serverEndpoints []ExtensionObjectDefinition, noOfServerSoftwareCertificates int32, serverSoftwareCertificates []ExtensionObjectDefinition, serverSignature ExtensionObjectDefinition, maxRequestMessageSize uint32) *_CreateSessionResponse {
	_result := &_CreateSessionResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		SessionId:                         sessionId,
		AuthenticationToken:               authenticationToken,
		RevisedSessionTimeout:             revisedSessionTimeout,
		ServerNonce:                       serverNonce,
		ServerCertificate:                 serverCertificate,
		NoOfServerEndpoints:               noOfServerEndpoints,
		ServerEndpoints:                   serverEndpoints,
		NoOfServerSoftwareCertificates:    noOfServerSoftwareCertificates,
		ServerSoftwareCertificates:        serverSoftwareCertificates,
		ServerSignature:                   serverSignature,
		MaxRequestMessageSize:             maxRequestMessageSize,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCreateSessionResponse(structType any) CreateSessionResponse {
	if casted, ok := structType.(CreateSessionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CreateSessionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CreateSessionResponse) GetTypeName() string {
	return "CreateSessionResponse"
}

func (m *_CreateSessionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (sessionId)
	lengthInBits += m.SessionId.GetLengthInBits(ctx)

	// Simple field (authenticationToken)
	lengthInBits += m.AuthenticationToken.GetLengthInBits(ctx)

	// Simple field (revisedSessionTimeout)
	lengthInBits += 64

	// Simple field (serverNonce)
	lengthInBits += m.ServerNonce.GetLengthInBits(ctx)

	// Simple field (serverCertificate)
	lengthInBits += m.ServerCertificate.GetLengthInBits(ctx)

	// Simple field (noOfServerEndpoints)
	lengthInBits += 32

	// Array field
	if len(m.ServerEndpoints) > 0 {
		for _curItem, element := range m.ServerEndpoints {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerEndpoints), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfServerSoftwareCertificates)
	lengthInBits += 32

	// Array field
	if len(m.ServerSoftwareCertificates) > 0 {
		for _curItem, element := range m.ServerSoftwareCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerSoftwareCertificates), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (serverSignature)
	lengthInBits += m.ServerSignature.GetLengthInBits(ctx)

	// Simple field (maxRequestMessageSize)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CreateSessionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CreateSessionResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__createSessionResponse CreateSessionResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CreateSessionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateSessionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	sessionId, err := ReadSimpleField[NodeId](ctx, "sessionId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sessionId' field"))
	}
	m.SessionId = sessionId

	authenticationToken, err := ReadSimpleField[NodeId](ctx, "authenticationToken", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationToken' field"))
	}
	m.AuthenticationToken = authenticationToken

	revisedSessionTimeout, err := ReadSimpleField(ctx, "revisedSessionTimeout", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisedSessionTimeout' field"))
	}
	m.RevisedSessionTimeout = revisedSessionTimeout

	serverNonce, err := ReadSimpleField[PascalByteString](ctx, "serverNonce", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverNonce' field"))
	}
	m.ServerNonce = serverNonce

	serverCertificate, err := ReadSimpleField[PascalByteString](ctx, "serverCertificate", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverCertificate' field"))
	}
	m.ServerCertificate = serverCertificate

	noOfServerEndpoints, err := ReadSimpleField(ctx, "noOfServerEndpoints", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServerEndpoints' field"))
	}
	m.NoOfServerEndpoints = noOfServerEndpoints

	serverEndpoints, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "serverEndpoints", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("314")), readBuffer), uint64(noOfServerEndpoints))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverEndpoints' field"))
	}
	m.ServerEndpoints = serverEndpoints

	noOfServerSoftwareCertificates, err := ReadSimpleField(ctx, "noOfServerSoftwareCertificates", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServerSoftwareCertificates' field"))
	}
	m.NoOfServerSoftwareCertificates = noOfServerSoftwareCertificates

	serverSoftwareCertificates, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "serverSoftwareCertificates", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("346")), readBuffer), uint64(noOfServerSoftwareCertificates))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverSoftwareCertificates' field"))
	}
	m.ServerSoftwareCertificates = serverSoftwareCertificates

	serverSignature, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "serverSignature", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("458")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverSignature' field"))
	}
	m.ServerSignature = serverSignature

	maxRequestMessageSize, err := ReadSimpleField(ctx, "maxRequestMessageSize", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxRequestMessageSize' field"))
	}
	m.MaxRequestMessageSize = maxRequestMessageSize

	if closeErr := readBuffer.CloseContext("CreateSessionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateSessionResponse")
	}

	return m, nil
}

func (m *_CreateSessionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateSessionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateSessionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateSessionResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "sessionId", m.GetSessionId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sessionId' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "authenticationToken", m.GetAuthenticationToken(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationToken' field")
		}

		if err := WriteSimpleField[float64](ctx, "revisedSessionTimeout", m.GetRevisedSessionTimeout(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'revisedSessionTimeout' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "serverNonce", m.GetServerNonce(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverNonce' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "serverCertificate", m.GetServerCertificate(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverCertificate' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfServerEndpoints", m.GetNoOfServerEndpoints(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServerEndpoints' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "serverEndpoints", m.GetServerEndpoints(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'serverEndpoints' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfServerSoftwareCertificates", m.GetNoOfServerSoftwareCertificates(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServerSoftwareCertificates' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "serverSoftwareCertificates", m.GetServerSoftwareCertificates(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'serverSoftwareCertificates' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "serverSignature", m.GetServerSignature(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverSignature' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxRequestMessageSize", m.GetMaxRequestMessageSize(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxRequestMessageSize' field")
		}

		if popErr := writeBuffer.PopContext("CreateSessionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateSessionResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateSessionResponse) IsCreateSessionResponse() {}

func (m *_CreateSessionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
