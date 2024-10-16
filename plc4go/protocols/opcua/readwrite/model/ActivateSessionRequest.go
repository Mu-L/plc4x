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

// ActivateSessionRequest is the corresponding interface of ActivateSessionRequest
type ActivateSessionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetClientSignature returns ClientSignature (property field)
	GetClientSignature() ExtensionObjectDefinition
	// GetNoOfClientSoftwareCertificates returns NoOfClientSoftwareCertificates (property field)
	GetNoOfClientSoftwareCertificates() int32
	// GetClientSoftwareCertificates returns ClientSoftwareCertificates (property field)
	GetClientSoftwareCertificates() []ExtensionObjectDefinition
	// GetNoOfLocaleIds returns NoOfLocaleIds (property field)
	GetNoOfLocaleIds() int32
	// GetLocaleIds returns LocaleIds (property field)
	GetLocaleIds() []PascalString
	// GetUserIdentityToken returns UserIdentityToken (property field)
	GetUserIdentityToken() ExtensionObject
	// GetUserTokenSignature returns UserTokenSignature (property field)
	GetUserTokenSignature() ExtensionObjectDefinition
	// IsActivateSessionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsActivateSessionRequest()
	// CreateBuilder creates a ActivateSessionRequestBuilder
	CreateActivateSessionRequestBuilder() ActivateSessionRequestBuilder
}

// _ActivateSessionRequest is the data-structure of this message
type _ActivateSessionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader                  ExtensionObjectDefinition
	ClientSignature                ExtensionObjectDefinition
	NoOfClientSoftwareCertificates int32
	ClientSoftwareCertificates     []ExtensionObjectDefinition
	NoOfLocaleIds                  int32
	LocaleIds                      []PascalString
	UserIdentityToken              ExtensionObject
	UserTokenSignature             ExtensionObjectDefinition
}

var _ ActivateSessionRequest = (*_ActivateSessionRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ActivateSessionRequest)(nil)

// NewActivateSessionRequest factory function for _ActivateSessionRequest
func NewActivateSessionRequest(requestHeader ExtensionObjectDefinition, clientSignature ExtensionObjectDefinition, noOfClientSoftwareCertificates int32, clientSoftwareCertificates []ExtensionObjectDefinition, noOfLocaleIds int32, localeIds []PascalString, userIdentityToken ExtensionObject, userTokenSignature ExtensionObjectDefinition) *_ActivateSessionRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for ActivateSessionRequest must not be nil")
	}
	if clientSignature == nil {
		panic("clientSignature of type ExtensionObjectDefinition for ActivateSessionRequest must not be nil")
	}
	if userIdentityToken == nil {
		panic("userIdentityToken of type ExtensionObject for ActivateSessionRequest must not be nil")
	}
	if userTokenSignature == nil {
		panic("userTokenSignature of type ExtensionObjectDefinition for ActivateSessionRequest must not be nil")
	}
	_result := &_ActivateSessionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		ClientSignature:                   clientSignature,
		NoOfClientSoftwareCertificates:    noOfClientSoftwareCertificates,
		ClientSoftwareCertificates:        clientSoftwareCertificates,
		NoOfLocaleIds:                     noOfLocaleIds,
		LocaleIds:                         localeIds,
		UserIdentityToken:                 userIdentityToken,
		UserTokenSignature:                userTokenSignature,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ActivateSessionRequestBuilder is a builder for ActivateSessionRequest
type ActivateSessionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, clientSignature ExtensionObjectDefinition, noOfClientSoftwareCertificates int32, clientSoftwareCertificates []ExtensionObjectDefinition, noOfLocaleIds int32, localeIds []PascalString, userIdentityToken ExtensionObject, userTokenSignature ExtensionObjectDefinition) ActivateSessionRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) ActivateSessionRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ActivateSessionRequestBuilder
	// WithClientSignature adds ClientSignature (property field)
	WithClientSignature(ExtensionObjectDefinition) ActivateSessionRequestBuilder
	// WithClientSignatureBuilder adds ClientSignature (property field) which is build by the builder
	WithClientSignatureBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ActivateSessionRequestBuilder
	// WithNoOfClientSoftwareCertificates adds NoOfClientSoftwareCertificates (property field)
	WithNoOfClientSoftwareCertificates(int32) ActivateSessionRequestBuilder
	// WithClientSoftwareCertificates adds ClientSoftwareCertificates (property field)
	WithClientSoftwareCertificates(...ExtensionObjectDefinition) ActivateSessionRequestBuilder
	// WithNoOfLocaleIds adds NoOfLocaleIds (property field)
	WithNoOfLocaleIds(int32) ActivateSessionRequestBuilder
	// WithLocaleIds adds LocaleIds (property field)
	WithLocaleIds(...PascalString) ActivateSessionRequestBuilder
	// WithUserIdentityToken adds UserIdentityToken (property field)
	WithUserIdentityToken(ExtensionObject) ActivateSessionRequestBuilder
	// WithUserIdentityTokenBuilder adds UserIdentityToken (property field) which is build by the builder
	WithUserIdentityTokenBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) ActivateSessionRequestBuilder
	// WithUserTokenSignature adds UserTokenSignature (property field)
	WithUserTokenSignature(ExtensionObjectDefinition) ActivateSessionRequestBuilder
	// WithUserTokenSignatureBuilder adds UserTokenSignature (property field) which is build by the builder
	WithUserTokenSignatureBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ActivateSessionRequestBuilder
	// Build builds the ActivateSessionRequest or returns an error if something is wrong
	Build() (ActivateSessionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ActivateSessionRequest
}

// NewActivateSessionRequestBuilder() creates a ActivateSessionRequestBuilder
func NewActivateSessionRequestBuilder() ActivateSessionRequestBuilder {
	return &_ActivateSessionRequestBuilder{_ActivateSessionRequest: new(_ActivateSessionRequest)}
}

type _ActivateSessionRequestBuilder struct {
	*_ActivateSessionRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ActivateSessionRequestBuilder) = (*_ActivateSessionRequestBuilder)(nil)

func (b *_ActivateSessionRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ActivateSessionRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, clientSignature ExtensionObjectDefinition, noOfClientSoftwareCertificates int32, clientSoftwareCertificates []ExtensionObjectDefinition, noOfLocaleIds int32, localeIds []PascalString, userIdentityToken ExtensionObject, userTokenSignature ExtensionObjectDefinition) ActivateSessionRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithClientSignature(clientSignature).WithNoOfClientSoftwareCertificates(noOfClientSoftwareCertificates).WithClientSoftwareCertificates(clientSoftwareCertificates...).WithNoOfLocaleIds(noOfLocaleIds).WithLocaleIds(localeIds...).WithUserIdentityToken(userIdentityToken).WithUserTokenSignature(userTokenSignature)
}

func (b *_ActivateSessionRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) ActivateSessionRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_ActivateSessionRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ActivateSessionRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_ActivateSessionRequestBuilder) WithClientSignature(clientSignature ExtensionObjectDefinition) ActivateSessionRequestBuilder {
	b.ClientSignature = clientSignature
	return b
}

func (b *_ActivateSessionRequestBuilder) WithClientSignatureBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ActivateSessionRequestBuilder {
	builder := builderSupplier(b.ClientSignature.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.ClientSignature, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_ActivateSessionRequestBuilder) WithNoOfClientSoftwareCertificates(noOfClientSoftwareCertificates int32) ActivateSessionRequestBuilder {
	b.NoOfClientSoftwareCertificates = noOfClientSoftwareCertificates
	return b
}

func (b *_ActivateSessionRequestBuilder) WithClientSoftwareCertificates(clientSoftwareCertificates ...ExtensionObjectDefinition) ActivateSessionRequestBuilder {
	b.ClientSoftwareCertificates = clientSoftwareCertificates
	return b
}

func (b *_ActivateSessionRequestBuilder) WithNoOfLocaleIds(noOfLocaleIds int32) ActivateSessionRequestBuilder {
	b.NoOfLocaleIds = noOfLocaleIds
	return b
}

func (b *_ActivateSessionRequestBuilder) WithLocaleIds(localeIds ...PascalString) ActivateSessionRequestBuilder {
	b.LocaleIds = localeIds
	return b
}

func (b *_ActivateSessionRequestBuilder) WithUserIdentityToken(userIdentityToken ExtensionObject) ActivateSessionRequestBuilder {
	b.UserIdentityToken = userIdentityToken
	return b
}

func (b *_ActivateSessionRequestBuilder) WithUserIdentityTokenBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) ActivateSessionRequestBuilder {
	builder := builderSupplier(b.UserIdentityToken.CreateExtensionObjectBuilder())
	var err error
	b.UserIdentityToken, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_ActivateSessionRequestBuilder) WithUserTokenSignature(userTokenSignature ExtensionObjectDefinition) ActivateSessionRequestBuilder {
	b.UserTokenSignature = userTokenSignature
	return b
}

func (b *_ActivateSessionRequestBuilder) WithUserTokenSignatureBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) ActivateSessionRequestBuilder {
	builder := builderSupplier(b.UserTokenSignature.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.UserTokenSignature, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_ActivateSessionRequestBuilder) Build() (ActivateSessionRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.ClientSignature == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'clientSignature' not set"))
	}
	if b.UserIdentityToken == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'userIdentityToken' not set"))
	}
	if b.UserTokenSignature == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'userTokenSignature' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ActivateSessionRequest.deepCopy(), nil
}

func (b *_ActivateSessionRequestBuilder) MustBuild() ActivateSessionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ActivateSessionRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ActivateSessionRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ActivateSessionRequestBuilder) DeepCopy() any {
	_copy := b.CreateActivateSessionRequestBuilder().(*_ActivateSessionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateActivateSessionRequestBuilder creates a ActivateSessionRequestBuilder
func (b *_ActivateSessionRequest) CreateActivateSessionRequestBuilder() ActivateSessionRequestBuilder {
	if b == nil {
		return NewActivateSessionRequestBuilder()
	}
	return &_ActivateSessionRequestBuilder{_ActivateSessionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ActivateSessionRequest) GetIdentifier() string {
	return "467"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ActivateSessionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ActivateSessionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_ActivateSessionRequest) GetClientSignature() ExtensionObjectDefinition {
	return m.ClientSignature
}

func (m *_ActivateSessionRequest) GetNoOfClientSoftwareCertificates() int32 {
	return m.NoOfClientSoftwareCertificates
}

func (m *_ActivateSessionRequest) GetClientSoftwareCertificates() []ExtensionObjectDefinition {
	return m.ClientSoftwareCertificates
}

func (m *_ActivateSessionRequest) GetNoOfLocaleIds() int32 {
	return m.NoOfLocaleIds
}

func (m *_ActivateSessionRequest) GetLocaleIds() []PascalString {
	return m.LocaleIds
}

func (m *_ActivateSessionRequest) GetUserIdentityToken() ExtensionObject {
	return m.UserIdentityToken
}

func (m *_ActivateSessionRequest) GetUserTokenSignature() ExtensionObjectDefinition {
	return m.UserTokenSignature
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastActivateSessionRequest(structType any) ActivateSessionRequest {
	if casted, ok := structType.(ActivateSessionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ActivateSessionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ActivateSessionRequest) GetTypeName() string {
	return "ActivateSessionRequest"
}

func (m *_ActivateSessionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (clientSignature)
	lengthInBits += m.ClientSignature.GetLengthInBits(ctx)

	// Simple field (noOfClientSoftwareCertificates)
	lengthInBits += 32

	// Array field
	if len(m.ClientSoftwareCertificates) > 0 {
		for _curItem, element := range m.ClientSoftwareCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ClientSoftwareCertificates), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfLocaleIds)
	lengthInBits += 32

	// Array field
	if len(m.LocaleIds) > 0 {
		for _curItem, element := range m.LocaleIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LocaleIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (userIdentityToken)
	lengthInBits += m.UserIdentityToken.GetLengthInBits(ctx)

	// Simple field (userTokenSignature)
	lengthInBits += m.UserTokenSignature.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ActivateSessionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ActivateSessionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__activateSessionRequest ActivateSessionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ActivateSessionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ActivateSessionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	clientSignature, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "clientSignature", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("458")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientSignature' field"))
	}
	m.ClientSignature = clientSignature

	noOfClientSoftwareCertificates, err := ReadSimpleField(ctx, "noOfClientSoftwareCertificates", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfClientSoftwareCertificates' field"))
	}
	m.NoOfClientSoftwareCertificates = noOfClientSoftwareCertificates

	clientSoftwareCertificates, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "clientSoftwareCertificates", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("346")), readBuffer), uint64(noOfClientSoftwareCertificates))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientSoftwareCertificates' field"))
	}
	m.ClientSoftwareCertificates = clientSoftwareCertificates

	noOfLocaleIds, err := ReadSimpleField(ctx, "noOfLocaleIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLocaleIds' field"))
	}
	m.NoOfLocaleIds = noOfLocaleIds

	localeIds, err := ReadCountArrayField[PascalString](ctx, "localeIds", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfLocaleIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localeIds' field"))
	}
	m.LocaleIds = localeIds

	userIdentityToken, err := ReadSimpleField[ExtensionObject](ctx, "userIdentityToken", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userIdentityToken' field"))
	}
	m.UserIdentityToken = userIdentityToken

	userTokenSignature, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "userTokenSignature", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("458")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userTokenSignature' field"))
	}
	m.UserTokenSignature = userTokenSignature

	if closeErr := readBuffer.CloseContext("ActivateSessionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ActivateSessionRequest")
	}

	return m, nil
}

func (m *_ActivateSessionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ActivateSessionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ActivateSessionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ActivateSessionRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "clientSignature", m.GetClientSignature(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientSignature' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfClientSoftwareCertificates", m.GetNoOfClientSoftwareCertificates(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfClientSoftwareCertificates' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "clientSoftwareCertificates", m.GetClientSoftwareCertificates(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'clientSoftwareCertificates' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfLocaleIds", m.GetNoOfLocaleIds(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLocaleIds' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "localeIds", m.GetLocaleIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'localeIds' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "userIdentityToken", m.GetUserIdentityToken(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userIdentityToken' field")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "userTokenSignature", m.GetUserTokenSignature(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userTokenSignature' field")
		}

		if popErr := writeBuffer.PopContext("ActivateSessionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ActivateSessionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ActivateSessionRequest) IsActivateSessionRequest() {}

func (m *_ActivateSessionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ActivateSessionRequest) deepCopy() *_ActivateSessionRequest {
	if m == nil {
		return nil
	}
	_ActivateSessionRequestCopy := &_ActivateSessionRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.ClientSignature.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfClientSoftwareCertificates,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.ClientSoftwareCertificates),
		m.NoOfLocaleIds,
		utils.DeepCopySlice[PascalString, PascalString](m.LocaleIds),
		m.UserIdentityToken.DeepCopy().(ExtensionObject),
		m.UserTokenSignature.DeepCopy().(ExtensionObjectDefinition),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ActivateSessionRequestCopy
}

func (m *_ActivateSessionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
