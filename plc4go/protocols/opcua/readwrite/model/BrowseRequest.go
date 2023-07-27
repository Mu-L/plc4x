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

// BrowseRequest is the corresponding interface of BrowseRequest
type BrowseRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetView returns View (property field)
	GetView() ExtensionObjectDefinition
	// GetRequestedMaxReferencesPerNode returns RequestedMaxReferencesPerNode (property field)
	GetRequestedMaxReferencesPerNode() uint32
	// GetNoOfNodesToBrowse returns NoOfNodesToBrowse (property field)
	GetNoOfNodesToBrowse() int32
	// GetNodesToBrowse returns NodesToBrowse (property field)
	GetNodesToBrowse() []ExtensionObjectDefinition
}

// BrowseRequestExactly can be used when we want exactly this type and not a type which fulfills BrowseRequest.
// This is useful for switch cases.
type BrowseRequestExactly interface {
	BrowseRequest
	isBrowseRequest() bool
}

// _BrowseRequest is the data-structure of this message
type _BrowseRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader                 ExtensionObjectDefinition
	View                          ExtensionObjectDefinition
	RequestedMaxReferencesPerNode uint32
	NoOfNodesToBrowse             int32
	NodesToBrowse                 []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowseRequest) GetIdentifier() string {
	return "527"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowseRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_BrowseRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowseRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_BrowseRequest) GetView() ExtensionObjectDefinition {
	return m.View
}

func (m *_BrowseRequest) GetRequestedMaxReferencesPerNode() uint32 {
	return m.RequestedMaxReferencesPerNode
}

func (m *_BrowseRequest) GetNoOfNodesToBrowse() int32 {
	return m.NoOfNodesToBrowse
}

func (m *_BrowseRequest) GetNodesToBrowse() []ExtensionObjectDefinition {
	return m.NodesToBrowse
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBrowseRequest factory function for _BrowseRequest
func NewBrowseRequest(requestHeader ExtensionObjectDefinition, view ExtensionObjectDefinition, requestedMaxReferencesPerNode uint32, noOfNodesToBrowse int32, nodesToBrowse []ExtensionObjectDefinition) *_BrowseRequest {
	_result := &_BrowseRequest{
		RequestHeader:                 requestHeader,
		View:                          view,
		RequestedMaxReferencesPerNode: requestedMaxReferencesPerNode,
		NoOfNodesToBrowse:             noOfNodesToBrowse,
		NodesToBrowse:                 nodesToBrowse,
		_ExtensionObjectDefinition:    NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBrowseRequest(structType any) BrowseRequest {
	if casted, ok := structType.(BrowseRequest); ok {
		return casted
	}
	if casted, ok := structType.(*BrowseRequest); ok {
		return *casted
	}
	return nil
}

func (m *_BrowseRequest) GetTypeName() string {
	return "BrowseRequest"
}

func (m *_BrowseRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (view)
	lengthInBits += m.View.GetLengthInBits(ctx)

	// Simple field (requestedMaxReferencesPerNode)
	lengthInBits += 32

	// Simple field (noOfNodesToBrowse)
	lengthInBits += 32

	// Array field
	if len(m.NodesToBrowse) > 0 {
		for _curItem, element := range m.NodesToBrowse {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToBrowse), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_BrowseRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrowseRequestParse(ctx context.Context, theBytes []byte, identifier string) (BrowseRequest, error) {
	return BrowseRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func BrowseRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (BrowseRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BrowseRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowseRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of BrowseRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (view)
	if pullErr := readBuffer.PullContext("view"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for view")
	}
	_view, _viewErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("513"))
	if _viewErr != nil {
		return nil, errors.Wrap(_viewErr, "Error parsing 'view' field of BrowseRequest")
	}
	view := _view.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("view"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for view")
	}

	// Simple Field (requestedMaxReferencesPerNode)
	_requestedMaxReferencesPerNode, _requestedMaxReferencesPerNodeErr := readBuffer.ReadUint32("requestedMaxReferencesPerNode", 32)
	if _requestedMaxReferencesPerNodeErr != nil {
		return nil, errors.Wrap(_requestedMaxReferencesPerNodeErr, "Error parsing 'requestedMaxReferencesPerNode' field of BrowseRequest")
	}
	requestedMaxReferencesPerNode := _requestedMaxReferencesPerNode

	// Simple Field (noOfNodesToBrowse)
	_noOfNodesToBrowse, _noOfNodesToBrowseErr := readBuffer.ReadInt32("noOfNodesToBrowse", 32)
	if _noOfNodesToBrowseErr != nil {
		return nil, errors.Wrap(_noOfNodesToBrowseErr, "Error parsing 'noOfNodesToBrowse' field of BrowseRequest")
	}
	noOfNodesToBrowse := _noOfNodesToBrowse

	// Array field (nodesToBrowse)
	if pullErr := readBuffer.PullContext("nodesToBrowse", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodesToBrowse")
	}
	// Count array
	nodesToBrowse := make([]ExtensionObjectDefinition, noOfNodesToBrowse)
	// This happens when the size is set conditional to 0
	if len(nodesToBrowse) == 0 {
		nodesToBrowse = nil
	}
	{
		_numItems := uint16(noOfNodesToBrowse)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "516")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'nodesToBrowse' field of BrowseRequest")
			}
			nodesToBrowse[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("nodesToBrowse", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodesToBrowse")
	}

	if closeErr := readBuffer.CloseContext("BrowseRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowseRequest")
	}

	// Create a partially initialized instance
	_child := &_BrowseRequest{
		_ExtensionObjectDefinition:    &_ExtensionObjectDefinition{},
		RequestHeader:                 requestHeader,
		View:                          view,
		RequestedMaxReferencesPerNode: requestedMaxReferencesPerNode,
		NoOfNodesToBrowse:             noOfNodesToBrowse,
		NodesToBrowse:                 nodesToBrowse,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_BrowseRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowseRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowseRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowseRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (view)
		if pushErr := writeBuffer.PushContext("view"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for view")
		}
		_viewErr := writeBuffer.WriteSerializable(ctx, m.GetView())
		if popErr := writeBuffer.PopContext("view"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for view")
		}
		if _viewErr != nil {
			return errors.Wrap(_viewErr, "Error serializing 'view' field")
		}

		// Simple Field (requestedMaxReferencesPerNode)
		requestedMaxReferencesPerNode := uint32(m.GetRequestedMaxReferencesPerNode())
		_requestedMaxReferencesPerNodeErr := writeBuffer.WriteUint32("requestedMaxReferencesPerNode", 32, (requestedMaxReferencesPerNode))
		if _requestedMaxReferencesPerNodeErr != nil {
			return errors.Wrap(_requestedMaxReferencesPerNodeErr, "Error serializing 'requestedMaxReferencesPerNode' field")
		}

		// Simple Field (noOfNodesToBrowse)
		noOfNodesToBrowse := int32(m.GetNoOfNodesToBrowse())
		_noOfNodesToBrowseErr := writeBuffer.WriteInt32("noOfNodesToBrowse", 32, (noOfNodesToBrowse))
		if _noOfNodesToBrowseErr != nil {
			return errors.Wrap(_noOfNodesToBrowseErr, "Error serializing 'noOfNodesToBrowse' field")
		}

		// Array Field (nodesToBrowse)
		if pushErr := writeBuffer.PushContext("nodesToBrowse", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodesToBrowse")
		}
		for _curItem, _element := range m.GetNodesToBrowse() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNodesToBrowse()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'nodesToBrowse' field")
			}
		}
		if popErr := writeBuffer.PopContext("nodesToBrowse", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodesToBrowse")
		}

		if popErr := writeBuffer.PopContext("BrowseRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowseRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowseRequest) isBrowseRequest() bool {
	return true
}

func (m *_BrowseRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
