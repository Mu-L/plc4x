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

// ApplicationDescription is the corresponding interface of ApplicationDescription
type ApplicationDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetApplicationUri returns ApplicationUri (property field)
	GetApplicationUri() PascalString
	// GetProductUri returns ProductUri (property field)
	GetProductUri() PascalString
	// GetApplicationName returns ApplicationName (property field)
	GetApplicationName() LocalizedText
	// GetApplicationType returns ApplicationType (property field)
	GetApplicationType() ApplicationType
	// GetGatewayServerUri returns GatewayServerUri (property field)
	GetGatewayServerUri() PascalString
	// GetDiscoveryProfileUri returns DiscoveryProfileUri (property field)
	GetDiscoveryProfileUri() PascalString
	// GetNoOfDiscoveryUrls returns NoOfDiscoveryUrls (property field)
	GetNoOfDiscoveryUrls() int32
	// GetDiscoveryUrls returns DiscoveryUrls (property field)
	GetDiscoveryUrls() []PascalString
	// IsApplicationDescription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApplicationDescription()
}

// _ApplicationDescription is the data-structure of this message
type _ApplicationDescription struct {
	ExtensionObjectDefinitionContract
	ApplicationUri      PascalString
	ProductUri          PascalString
	ApplicationName     LocalizedText
	ApplicationType     ApplicationType
	GatewayServerUri    PascalString
	DiscoveryProfileUri PascalString
	NoOfDiscoveryUrls   int32
	DiscoveryUrls       []PascalString
}

var _ ApplicationDescription = (*_ApplicationDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ApplicationDescription)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApplicationDescription) GetIdentifier() string {
	return "310"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApplicationDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApplicationDescription) GetApplicationUri() PascalString {
	return m.ApplicationUri
}

func (m *_ApplicationDescription) GetProductUri() PascalString {
	return m.ProductUri
}

func (m *_ApplicationDescription) GetApplicationName() LocalizedText {
	return m.ApplicationName
}

func (m *_ApplicationDescription) GetApplicationType() ApplicationType {
	return m.ApplicationType
}

func (m *_ApplicationDescription) GetGatewayServerUri() PascalString {
	return m.GatewayServerUri
}

func (m *_ApplicationDescription) GetDiscoveryProfileUri() PascalString {
	return m.DiscoveryProfileUri
}

func (m *_ApplicationDescription) GetNoOfDiscoveryUrls() int32 {
	return m.NoOfDiscoveryUrls
}

func (m *_ApplicationDescription) GetDiscoveryUrls() []PascalString {
	return m.DiscoveryUrls
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApplicationDescription factory function for _ApplicationDescription
func NewApplicationDescription(applicationUri PascalString, productUri PascalString, applicationName LocalizedText, applicationType ApplicationType, gatewayServerUri PascalString, discoveryProfileUri PascalString, noOfDiscoveryUrls int32, discoveryUrls []PascalString) *_ApplicationDescription {
	_result := &_ApplicationDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ApplicationUri:                    applicationUri,
		ProductUri:                        productUri,
		ApplicationName:                   applicationName,
		ApplicationType:                   applicationType,
		GatewayServerUri:                  gatewayServerUri,
		DiscoveryProfileUri:               discoveryProfileUri,
		NoOfDiscoveryUrls:                 noOfDiscoveryUrls,
		DiscoveryUrls:                     discoveryUrls,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApplicationDescription(structType any) ApplicationDescription {
	if casted, ok := structType.(ApplicationDescription); ok {
		return casted
	}
	if casted, ok := structType.(*ApplicationDescription); ok {
		return *casted
	}
	return nil
}

func (m *_ApplicationDescription) GetTypeName() string {
	return "ApplicationDescription"
}

func (m *_ApplicationDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (applicationUri)
	lengthInBits += m.ApplicationUri.GetLengthInBits(ctx)

	// Simple field (productUri)
	lengthInBits += m.ProductUri.GetLengthInBits(ctx)

	// Simple field (applicationName)
	lengthInBits += m.ApplicationName.GetLengthInBits(ctx)

	// Simple field (applicationType)
	lengthInBits += 32

	// Simple field (gatewayServerUri)
	lengthInBits += m.GatewayServerUri.GetLengthInBits(ctx)

	// Simple field (discoveryProfileUri)
	lengthInBits += m.DiscoveryProfileUri.GetLengthInBits(ctx)

	// Simple field (noOfDiscoveryUrls)
	lengthInBits += 32

	// Array field
	if len(m.DiscoveryUrls) > 0 {
		for _curItem, element := range m.DiscoveryUrls {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiscoveryUrls), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ApplicationDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApplicationDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__applicationDescription ApplicationDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApplicationDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApplicationDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	applicationUri, err := ReadSimpleField[PascalString](ctx, "applicationUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'applicationUri' field"))
	}
	m.ApplicationUri = applicationUri

	productUri, err := ReadSimpleField[PascalString](ctx, "productUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'productUri' field"))
	}
	m.ProductUri = productUri

	applicationName, err := ReadSimpleField[LocalizedText](ctx, "applicationName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'applicationName' field"))
	}
	m.ApplicationName = applicationName

	applicationType, err := ReadEnumField[ApplicationType](ctx, "applicationType", "ApplicationType", ReadEnum(ApplicationTypeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'applicationType' field"))
	}
	m.ApplicationType = applicationType

	gatewayServerUri, err := ReadSimpleField[PascalString](ctx, "gatewayServerUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'gatewayServerUri' field"))
	}
	m.GatewayServerUri = gatewayServerUri

	discoveryProfileUri, err := ReadSimpleField[PascalString](ctx, "discoveryProfileUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'discoveryProfileUri' field"))
	}
	m.DiscoveryProfileUri = discoveryProfileUri

	noOfDiscoveryUrls, err := ReadSimpleField(ctx, "noOfDiscoveryUrls", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiscoveryUrls' field"))
	}
	m.NoOfDiscoveryUrls = noOfDiscoveryUrls

	discoveryUrls, err := ReadCountArrayField[PascalString](ctx, "discoveryUrls", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfDiscoveryUrls))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'discoveryUrls' field"))
	}
	m.DiscoveryUrls = discoveryUrls

	if closeErr := readBuffer.CloseContext("ApplicationDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApplicationDescription")
	}

	return m, nil
}

func (m *_ApplicationDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApplicationDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApplicationDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApplicationDescription")
		}

		if err := WriteSimpleField[PascalString](ctx, "applicationUri", m.GetApplicationUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'applicationUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "productUri", m.GetProductUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'productUri' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "applicationName", m.GetApplicationName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'applicationName' field")
		}

		if err := WriteSimpleEnumField[ApplicationType](ctx, "applicationType", "ApplicationType", m.GetApplicationType(), WriteEnum[ApplicationType, uint32](ApplicationType.GetValue, ApplicationType.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'applicationType' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "gatewayServerUri", m.GetGatewayServerUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'gatewayServerUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "discoveryProfileUri", m.GetDiscoveryProfileUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'discoveryProfileUri' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiscoveryUrls", m.GetNoOfDiscoveryUrls(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiscoveryUrls' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "discoveryUrls", m.GetDiscoveryUrls(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'discoveryUrls' field")
		}

		if popErr := writeBuffer.PopContext("ApplicationDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApplicationDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApplicationDescription) IsApplicationDescription() {}

func (m *_ApplicationDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
