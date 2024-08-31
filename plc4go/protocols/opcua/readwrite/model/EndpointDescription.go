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

// EndpointDescription is the corresponding interface of EndpointDescription
type EndpointDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetEndpointUrl returns EndpointUrl (property field)
	GetEndpointUrl() PascalString
	// GetServer returns Server (property field)
	GetServer() ExtensionObjectDefinition
	// GetServerCertificate returns ServerCertificate (property field)
	GetServerCertificate() PascalByteString
	// GetSecurityMode returns SecurityMode (property field)
	GetSecurityMode() MessageSecurityMode
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
	// GetNoOfUserIdentityTokens returns NoOfUserIdentityTokens (property field)
	GetNoOfUserIdentityTokens() int32
	// GetUserIdentityTokens returns UserIdentityTokens (property field)
	GetUserIdentityTokens() []ExtensionObjectDefinition
	// GetTransportProfileUri returns TransportProfileUri (property field)
	GetTransportProfileUri() PascalString
	// GetSecurityLevel returns SecurityLevel (property field)
	GetSecurityLevel() uint8
}

// EndpointDescriptionExactly can be used when we want exactly this type and not a type which fulfills EndpointDescription.
// This is useful for switch cases.
type EndpointDescriptionExactly interface {
	EndpointDescription
	isEndpointDescription() bool
}

// _EndpointDescription is the data-structure of this message
type _EndpointDescription struct {
	*_ExtensionObjectDefinition
	EndpointUrl            PascalString
	Server                 ExtensionObjectDefinition
	ServerCertificate      PascalByteString
	SecurityMode           MessageSecurityMode
	SecurityPolicyUri      PascalString
	NoOfUserIdentityTokens int32
	UserIdentityTokens     []ExtensionObjectDefinition
	TransportProfileUri    PascalString
	SecurityLevel          uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EndpointDescription) GetIdentifier() string {
	return "314"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EndpointDescription) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_EndpointDescription) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EndpointDescription) GetEndpointUrl() PascalString {
	return m.EndpointUrl
}

func (m *_EndpointDescription) GetServer() ExtensionObjectDefinition {
	return m.Server
}

func (m *_EndpointDescription) GetServerCertificate() PascalByteString {
	return m.ServerCertificate
}

func (m *_EndpointDescription) GetSecurityMode() MessageSecurityMode {
	return m.SecurityMode
}

func (m *_EndpointDescription) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

func (m *_EndpointDescription) GetNoOfUserIdentityTokens() int32 {
	return m.NoOfUserIdentityTokens
}

func (m *_EndpointDescription) GetUserIdentityTokens() []ExtensionObjectDefinition {
	return m.UserIdentityTokens
}

func (m *_EndpointDescription) GetTransportProfileUri() PascalString {
	return m.TransportProfileUri
}

func (m *_EndpointDescription) GetSecurityLevel() uint8 {
	return m.SecurityLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEndpointDescription factory function for _EndpointDescription
func NewEndpointDescription(endpointUrl PascalString, server ExtensionObjectDefinition, serverCertificate PascalByteString, securityMode MessageSecurityMode, securityPolicyUri PascalString, noOfUserIdentityTokens int32, userIdentityTokens []ExtensionObjectDefinition, transportProfileUri PascalString, securityLevel uint8) *_EndpointDescription {
	_result := &_EndpointDescription{
		EndpointUrl:                endpointUrl,
		Server:                     server,
		ServerCertificate:          serverCertificate,
		SecurityMode:               securityMode,
		SecurityPolicyUri:          securityPolicyUri,
		NoOfUserIdentityTokens:     noOfUserIdentityTokens,
		UserIdentityTokens:         userIdentityTokens,
		TransportProfileUri:        transportProfileUri,
		SecurityLevel:              securityLevel,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEndpointDescription(structType any) EndpointDescription {
	if casted, ok := structType.(EndpointDescription); ok {
		return casted
	}
	if casted, ok := structType.(*EndpointDescription); ok {
		return *casted
	}
	return nil
}

func (m *_EndpointDescription) GetTypeName() string {
	return "EndpointDescription"
}

func (m *_EndpointDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (endpointUrl)
	lengthInBits += m.EndpointUrl.GetLengthInBits(ctx)

	// Simple field (server)
	lengthInBits += m.Server.GetLengthInBits(ctx)

	// Simple field (serverCertificate)
	lengthInBits += m.ServerCertificate.GetLengthInBits(ctx)

	// Simple field (securityMode)
	lengthInBits += 32

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	// Simple field (noOfUserIdentityTokens)
	lengthInBits += 32

	// Array field
	if len(m.UserIdentityTokens) > 0 {
		for _curItem, element := range m.UserIdentityTokens {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.UserIdentityTokens), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (transportProfileUri)
	lengthInBits += m.TransportProfileUri.GetLengthInBits(ctx)

	// Simple field (securityLevel)
	lengthInBits += 8

	return lengthInBits
}

func (m *_EndpointDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EndpointDescriptionParse(ctx context.Context, theBytes []byte, identifier string) (EndpointDescription, error) {
	return EndpointDescriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func EndpointDescriptionParseWithBufferProducer(identifier string) func(ctx context.Context, readBuffer utils.ReadBuffer) (EndpointDescription, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (EndpointDescription, error) {
		return EndpointDescriptionParseWithBuffer(ctx, readBuffer, identifier)
	}
}

func EndpointDescriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (EndpointDescription, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("EndpointDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EndpointDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	endpointUrl, err := ReadSimpleField[PascalString](ctx, "endpointUrl", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endpointUrl' field"))
	}

	server, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "server", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("310")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'server' field"))
	}

	serverCertificate, err := ReadSimpleField[PascalByteString](ctx, "serverCertificate", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverCertificate' field"))
	}

	securityMode, err := ReadEnumField[MessageSecurityMode](ctx, "securityMode", "MessageSecurityMode", ReadEnum(MessageSecurityModeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityMode' field"))
	}

	securityPolicyUri, err := ReadSimpleField[PascalString](ctx, "securityPolicyUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityPolicyUri' field"))
	}

	noOfUserIdentityTokens, err := ReadSimpleField(ctx, "noOfUserIdentityTokens", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfUserIdentityTokens' field"))
	}

	userIdentityTokens, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "userIdentityTokens", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("306")), readBuffer), uint64(noOfUserIdentityTokens))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userIdentityTokens' field"))
	}

	transportProfileUri, err := ReadSimpleField[PascalString](ctx, "transportProfileUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportProfileUri' field"))
	}

	securityLevel, err := ReadSimpleField(ctx, "securityLevel", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityLevel' field"))
	}

	if closeErr := readBuffer.CloseContext("EndpointDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EndpointDescription")
	}

	// Create a partially initialized instance
	_child := &_EndpointDescription{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		EndpointUrl:                endpointUrl,
		Server:                     server,
		ServerCertificate:          serverCertificate,
		SecurityMode:               securityMode,
		SecurityPolicyUri:          securityPolicyUri,
		NoOfUserIdentityTokens:     noOfUserIdentityTokens,
		UserIdentityTokens:         userIdentityTokens,
		TransportProfileUri:        transportProfileUri,
		SecurityLevel:              securityLevel,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_EndpointDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EndpointDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EndpointDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EndpointDescription")
		}

		// Simple Field (endpointUrl)
		if pushErr := writeBuffer.PushContext("endpointUrl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for endpointUrl")
		}
		_endpointUrlErr := writeBuffer.WriteSerializable(ctx, m.GetEndpointUrl())
		if popErr := writeBuffer.PopContext("endpointUrl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for endpointUrl")
		}
		if _endpointUrlErr != nil {
			return errors.Wrap(_endpointUrlErr, "Error serializing 'endpointUrl' field")
		}

		// Simple Field (server)
		if pushErr := writeBuffer.PushContext("server"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for server")
		}
		_serverErr := writeBuffer.WriteSerializable(ctx, m.GetServer())
		if popErr := writeBuffer.PopContext("server"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for server")
		}
		if _serverErr != nil {
			return errors.Wrap(_serverErr, "Error serializing 'server' field")
		}

		// Simple Field (serverCertificate)
		if pushErr := writeBuffer.PushContext("serverCertificate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverCertificate")
		}
		_serverCertificateErr := writeBuffer.WriteSerializable(ctx, m.GetServerCertificate())
		if popErr := writeBuffer.PopContext("serverCertificate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverCertificate")
		}
		if _serverCertificateErr != nil {
			return errors.Wrap(_serverCertificateErr, "Error serializing 'serverCertificate' field")
		}

		// Simple Field (securityMode)
		if pushErr := writeBuffer.PushContext("securityMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityMode")
		}
		_securityModeErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityMode())
		if popErr := writeBuffer.PopContext("securityMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityMode")
		}
		if _securityModeErr != nil {
			return errors.Wrap(_securityModeErr, "Error serializing 'securityMode' field")
		}

		// Simple Field (securityPolicyUri)
		if pushErr := writeBuffer.PushContext("securityPolicyUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityPolicyUri")
		}
		_securityPolicyUriErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityPolicyUri())
		if popErr := writeBuffer.PopContext("securityPolicyUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityPolicyUri")
		}
		if _securityPolicyUriErr != nil {
			return errors.Wrap(_securityPolicyUriErr, "Error serializing 'securityPolicyUri' field")
		}

		// Simple Field (noOfUserIdentityTokens)
		noOfUserIdentityTokens := int32(m.GetNoOfUserIdentityTokens())
		_noOfUserIdentityTokensErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfUserIdentityTokens", 32, int32((noOfUserIdentityTokens)))
		if _noOfUserIdentityTokensErr != nil {
			return errors.Wrap(_noOfUserIdentityTokensErr, "Error serializing 'noOfUserIdentityTokens' field")
		}

		// Array Field (userIdentityTokens)
		if pushErr := writeBuffer.PushContext("userIdentityTokens", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userIdentityTokens")
		}
		for _curItem, _element := range m.GetUserIdentityTokens() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetUserIdentityTokens()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'userIdentityTokens' field")
			}
		}
		if popErr := writeBuffer.PopContext("userIdentityTokens", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userIdentityTokens")
		}

		// Simple Field (transportProfileUri)
		if pushErr := writeBuffer.PushContext("transportProfileUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transportProfileUri")
		}
		_transportProfileUriErr := writeBuffer.WriteSerializable(ctx, m.GetTransportProfileUri())
		if popErr := writeBuffer.PopContext("transportProfileUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transportProfileUri")
		}
		if _transportProfileUriErr != nil {
			return errors.Wrap(_transportProfileUriErr, "Error serializing 'transportProfileUri' field")
		}

		// Simple Field (securityLevel)
		securityLevel := uint8(m.GetSecurityLevel())
		_securityLevelErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("securityLevel", 8, uint8((securityLevel)))
		if _securityLevelErr != nil {
			return errors.Wrap(_securityLevelErr, "Error serializing 'securityLevel' field")
		}

		if popErr := writeBuffer.PopContext("EndpointDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EndpointDescription")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EndpointDescription) isEndpointDescription() bool {
	return true
}

func (m *_EndpointDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
