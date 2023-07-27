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

// PublishResponse is the corresponding interface of PublishResponse
type PublishResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetNoOfAvailableSequenceNumbers returns NoOfAvailableSequenceNumbers (property field)
	GetNoOfAvailableSequenceNumbers() int32
	// GetAvailableSequenceNumbers returns AvailableSequenceNumbers (property field)
	GetAvailableSequenceNumbers() []uint32
	// GetMoreNotifications returns MoreNotifications (property field)
	GetMoreNotifications() bool
	// GetNotificationMessage returns NotificationMessage (property field)
	GetNotificationMessage() ExtensionObjectDefinition
	// GetNoOfResults returns NoOfResults (property field)
	GetNoOfResults() int32
	// GetResults returns Results (property field)
	GetResults() []StatusCode
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
}

// PublishResponseExactly can be used when we want exactly this type and not a type which fulfills PublishResponse.
// This is useful for switch cases.
type PublishResponseExactly interface {
	PublishResponse
	isPublishResponse() bool
}

// _PublishResponse is the data-structure of this message
type _PublishResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader               ExtensionObjectDefinition
	SubscriptionId               uint32
	NoOfAvailableSequenceNumbers int32
	AvailableSequenceNumbers     []uint32
	MoreNotifications            bool
	NotificationMessage          ExtensionObjectDefinition
	NoOfResults                  int32
	Results                      []StatusCode
	NoOfDiagnosticInfos          int32
	DiagnosticInfos              []DiagnosticInfo
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PublishResponse) GetIdentifier() string {
	return "829"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PublishResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PublishResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PublishResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_PublishResponse) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_PublishResponse) GetNoOfAvailableSequenceNumbers() int32 {
	return m.NoOfAvailableSequenceNumbers
}

func (m *_PublishResponse) GetAvailableSequenceNumbers() []uint32 {
	return m.AvailableSequenceNumbers
}

func (m *_PublishResponse) GetMoreNotifications() bool {
	return m.MoreNotifications
}

func (m *_PublishResponse) GetNotificationMessage() ExtensionObjectDefinition {
	return m.NotificationMessage
}

func (m *_PublishResponse) GetNoOfResults() int32 {
	return m.NoOfResults
}

func (m *_PublishResponse) GetResults() []StatusCode {
	return m.Results
}

func (m *_PublishResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_PublishResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPublishResponse factory function for _PublishResponse
func NewPublishResponse(responseHeader ExtensionObjectDefinition, subscriptionId uint32, noOfAvailableSequenceNumbers int32, availableSequenceNumbers []uint32, moreNotifications bool, notificationMessage ExtensionObjectDefinition, noOfResults int32, results []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_PublishResponse {
	_result := &_PublishResponse{
		ResponseHeader:               responseHeader,
		SubscriptionId:               subscriptionId,
		NoOfAvailableSequenceNumbers: noOfAvailableSequenceNumbers,
		AvailableSequenceNumbers:     availableSequenceNumbers,
		MoreNotifications:            moreNotifications,
		NotificationMessage:          notificationMessage,
		NoOfResults:                  noOfResults,
		Results:                      results,
		NoOfDiagnosticInfos:          noOfDiagnosticInfos,
		DiagnosticInfos:              diagnosticInfos,
		_ExtensionObjectDefinition:   NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPublishResponse(structType any) PublishResponse {
	if casted, ok := structType.(PublishResponse); ok {
		return casted
	}
	if casted, ok := structType.(*PublishResponse); ok {
		return *casted
	}
	return nil
}

func (m *_PublishResponse) GetTypeName() string {
	return "PublishResponse"
}

func (m *_PublishResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (noOfAvailableSequenceNumbers)
	lengthInBits += 32

	// Array field
	if len(m.AvailableSequenceNumbers) > 0 {
		lengthInBits += 32 * uint16(len(m.AvailableSequenceNumbers))
	}

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (moreNotifications)
	lengthInBits += 1

	// Simple field (notificationMessage)
	lengthInBits += m.NotificationMessage.GetLengthInBits(ctx)

	// Simple field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PublishResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PublishResponseParse(ctx context.Context, theBytes []byte, identifier string) (PublishResponse, error) {
	return PublishResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PublishResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PublishResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PublishResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PublishResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of PublishResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (subscriptionId)
	_subscriptionId, _subscriptionIdErr := readBuffer.ReadUint32("subscriptionId", 32)
	if _subscriptionIdErr != nil {
		return nil, errors.Wrap(_subscriptionIdErr, "Error parsing 'subscriptionId' field of PublishResponse")
	}
	subscriptionId := _subscriptionId

	// Simple Field (noOfAvailableSequenceNumbers)
	_noOfAvailableSequenceNumbers, _noOfAvailableSequenceNumbersErr := readBuffer.ReadInt32("noOfAvailableSequenceNumbers", 32)
	if _noOfAvailableSequenceNumbersErr != nil {
		return nil, errors.Wrap(_noOfAvailableSequenceNumbersErr, "Error parsing 'noOfAvailableSequenceNumbers' field of PublishResponse")
	}
	noOfAvailableSequenceNumbers := _noOfAvailableSequenceNumbers

	// Array field (availableSequenceNumbers)
	if pullErr := readBuffer.PullContext("availableSequenceNumbers", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for availableSequenceNumbers")
	}
	// Count array
	availableSequenceNumbers := make([]uint32, noOfAvailableSequenceNumbers)
	// This happens when the size is set conditional to 0
	if len(availableSequenceNumbers) == 0 {
		availableSequenceNumbers = nil
	}
	{
		_numItems := uint16(noOfAvailableSequenceNumbers)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint32("", 32)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'availableSequenceNumbers' field of PublishResponse")
			}
			availableSequenceNumbers[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("availableSequenceNumbers", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for availableSequenceNumbers")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of PublishResponse")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (moreNotifications)
	_moreNotifications, _moreNotificationsErr := readBuffer.ReadBit("moreNotifications")
	if _moreNotificationsErr != nil {
		return nil, errors.Wrap(_moreNotificationsErr, "Error parsing 'moreNotifications' field of PublishResponse")
	}
	moreNotifications := _moreNotifications

	// Simple Field (notificationMessage)
	if pullErr := readBuffer.PullContext("notificationMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for notificationMessage")
	}
	_notificationMessage, _notificationMessageErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("805"))
	if _notificationMessageErr != nil {
		return nil, errors.Wrap(_notificationMessageErr, "Error parsing 'notificationMessage' field of PublishResponse")
	}
	notificationMessage := _notificationMessage.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("notificationMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for notificationMessage")
	}

	// Simple Field (noOfResults)
	_noOfResults, _noOfResultsErr := readBuffer.ReadInt32("noOfResults", 32)
	if _noOfResultsErr != nil {
		return nil, errors.Wrap(_noOfResultsErr, "Error parsing 'noOfResults' field of PublishResponse")
	}
	noOfResults := _noOfResults

	// Array field (results)
	if pullErr := readBuffer.PullContext("results", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for results")
	}
	// Count array
	results := make([]StatusCode, noOfResults)
	// This happens when the size is set conditional to 0
	if len(results) == 0 {
		results = nil
	}
	{
		_numItems := uint16(noOfResults)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := StatusCodeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'results' field of PublishResponse")
			}
			results[_curItem] = _item.(StatusCode)
		}
	}
	if closeErr := readBuffer.CloseContext("results", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for results")
	}

	// Simple Field (noOfDiagnosticInfos)
	_noOfDiagnosticInfos, _noOfDiagnosticInfosErr := readBuffer.ReadInt32("noOfDiagnosticInfos", 32)
	if _noOfDiagnosticInfosErr != nil {
		return nil, errors.Wrap(_noOfDiagnosticInfosErr, "Error parsing 'noOfDiagnosticInfos' field of PublishResponse")
	}
	noOfDiagnosticInfos := _noOfDiagnosticInfos

	// Array field (diagnosticInfos)
	if pullErr := readBuffer.PullContext("diagnosticInfos", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for diagnosticInfos")
	}
	// Count array
	diagnosticInfos := make([]DiagnosticInfo, noOfDiagnosticInfos)
	// This happens when the size is set conditional to 0
	if len(diagnosticInfos) == 0 {
		diagnosticInfos = nil
	}
	{
		_numItems := uint16(noOfDiagnosticInfos)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := DiagnosticInfoParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'diagnosticInfos' field of PublishResponse")
			}
			diagnosticInfos[_curItem] = _item.(DiagnosticInfo)
		}
	}
	if closeErr := readBuffer.CloseContext("diagnosticInfos", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for diagnosticInfos")
	}

	if closeErr := readBuffer.CloseContext("PublishResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PublishResponse")
	}

	// Create a partially initialized instance
	_child := &_PublishResponse{
		_ExtensionObjectDefinition:   &_ExtensionObjectDefinition{},
		ResponseHeader:               responseHeader,
		SubscriptionId:               subscriptionId,
		NoOfAvailableSequenceNumbers: noOfAvailableSequenceNumbers,
		AvailableSequenceNumbers:     availableSequenceNumbers,
		MoreNotifications:            moreNotifications,
		NotificationMessage:          notificationMessage,
		NoOfResults:                  noOfResults,
		Results:                      results,
		NoOfDiagnosticInfos:          noOfDiagnosticInfos,
		DiagnosticInfos:              diagnosticInfos,
		reservedField0:               reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PublishResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PublishResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PublishResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PublishResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (subscriptionId)
		subscriptionId := uint32(m.GetSubscriptionId())
		_subscriptionIdErr := writeBuffer.WriteUint32("subscriptionId", 32, (subscriptionId))
		if _subscriptionIdErr != nil {
			return errors.Wrap(_subscriptionIdErr, "Error serializing 'subscriptionId' field")
		}

		// Simple Field (noOfAvailableSequenceNumbers)
		noOfAvailableSequenceNumbers := int32(m.GetNoOfAvailableSequenceNumbers())
		_noOfAvailableSequenceNumbersErr := writeBuffer.WriteInt32("noOfAvailableSequenceNumbers", 32, (noOfAvailableSequenceNumbers))
		if _noOfAvailableSequenceNumbersErr != nil {
			return errors.Wrap(_noOfAvailableSequenceNumbersErr, "Error serializing 'noOfAvailableSequenceNumbers' field")
		}

		// Array Field (availableSequenceNumbers)
		if pushErr := writeBuffer.PushContext("availableSequenceNumbers", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for availableSequenceNumbers")
		}
		for _curItem, _element := range m.GetAvailableSequenceNumbers() {
			_ = _curItem
			_elementErr := writeBuffer.WriteUint32("", 32, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'availableSequenceNumbers' field")
			}
		}
		if popErr := writeBuffer.PopContext("availableSequenceNumbers", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for availableSequenceNumbers")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (moreNotifications)
		moreNotifications := bool(m.GetMoreNotifications())
		_moreNotificationsErr := writeBuffer.WriteBit("moreNotifications", (moreNotifications))
		if _moreNotificationsErr != nil {
			return errors.Wrap(_moreNotificationsErr, "Error serializing 'moreNotifications' field")
		}

		// Simple Field (notificationMessage)
		if pushErr := writeBuffer.PushContext("notificationMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for notificationMessage")
		}
		_notificationMessageErr := writeBuffer.WriteSerializable(ctx, m.GetNotificationMessage())
		if popErr := writeBuffer.PopContext("notificationMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for notificationMessage")
		}
		if _notificationMessageErr != nil {
			return errors.Wrap(_notificationMessageErr, "Error serializing 'notificationMessage' field")
		}

		// Simple Field (noOfResults)
		noOfResults := int32(m.GetNoOfResults())
		_noOfResultsErr := writeBuffer.WriteInt32("noOfResults", 32, (noOfResults))
		if _noOfResultsErr != nil {
			return errors.Wrap(_noOfResultsErr, "Error serializing 'noOfResults' field")
		}

		// Array Field (results)
		if pushErr := writeBuffer.PushContext("results", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for results")
		}
		for _curItem, _element := range m.GetResults() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetResults()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'results' field")
			}
		}
		if popErr := writeBuffer.PopContext("results", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for results")
		}

		// Simple Field (noOfDiagnosticInfos)
		noOfDiagnosticInfos := int32(m.GetNoOfDiagnosticInfos())
		_noOfDiagnosticInfosErr := writeBuffer.WriteInt32("noOfDiagnosticInfos", 32, (noOfDiagnosticInfos))
		if _noOfDiagnosticInfosErr != nil {
			return errors.Wrap(_noOfDiagnosticInfosErr, "Error serializing 'noOfDiagnosticInfos' field")
		}

		// Array Field (diagnosticInfos)
		if pushErr := writeBuffer.PushContext("diagnosticInfos", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for diagnosticInfos")
		}
		for _curItem, _element := range m.GetDiagnosticInfos() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetDiagnosticInfos()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'diagnosticInfos' field")
			}
		}
		if popErr := writeBuffer.PopContext("diagnosticInfos", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for diagnosticInfos")
		}

		if popErr := writeBuffer.PopContext("PublishResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PublishResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PublishResponse) isPublishResponse() bool {
	return true
}

func (m *_PublishResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
