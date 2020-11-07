//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type KNXAddress struct {
    MainGroup uint8
    MiddleGroup uint8
    SubGroup uint8
    IKNXAddress
}

// The corresponding interface
type IKNXAddress interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

func NewKNXAddress(mainGroup uint8, middleGroup uint8, subGroup uint8) *KNXAddress {
    return &KNXAddress{MainGroup: mainGroup, MiddleGroup: middleGroup, SubGroup: subGroup}
}

func CastKNXAddress(structType interface{}) KNXAddress {
    castFunc := func(typ interface{}) KNXAddress {
        if casted, ok := typ.(KNXAddress); ok {
            return casted
        }
        if casted, ok := typ.(*KNXAddress); ok {
            return *casted
        }
        return KNXAddress{}
    }
    return castFunc(structType)
}

func (m *KNXAddress) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (mainGroup)
    lengthInBits += 4

    // Simple field (middleGroup)
    lengthInBits += 4

    // Simple field (subGroup)
    lengthInBits += 8

    return lengthInBits
}

func (m *KNXAddress) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func KNXAddressParse(io *utils.ReadBuffer) (*KNXAddress, error) {

    // Simple Field (mainGroup)
    mainGroup, _mainGroupErr := io.ReadUint8(4)
    if _mainGroupErr != nil {
        return nil, errors.New("Error parsing 'mainGroup' field " + _mainGroupErr.Error())
    }

    // Simple Field (middleGroup)
    middleGroup, _middleGroupErr := io.ReadUint8(4)
    if _middleGroupErr != nil {
        return nil, errors.New("Error parsing 'middleGroup' field " + _middleGroupErr.Error())
    }

    // Simple Field (subGroup)
    subGroup, _subGroupErr := io.ReadUint8(8)
    if _subGroupErr != nil {
        return nil, errors.New("Error parsing 'subGroup' field " + _subGroupErr.Error())
    }

    // Create the instance
    return NewKNXAddress(mainGroup, middleGroup, subGroup), nil
}

func (m *KNXAddress) Serialize(io utils.WriteBuffer) error {

    // Simple Field (mainGroup)
    mainGroup := uint8(m.MainGroup)
    _mainGroupErr := io.WriteUint8(4, (mainGroup))
    if _mainGroupErr != nil {
        return errors.New("Error serializing 'mainGroup' field " + _mainGroupErr.Error())
    }

    // Simple Field (middleGroup)
    middleGroup := uint8(m.MiddleGroup)
    _middleGroupErr := io.WriteUint8(4, (middleGroup))
    if _middleGroupErr != nil {
        return errors.New("Error serializing 'middleGroup' field " + _middleGroupErr.Error())
    }

    // Simple Field (subGroup)
    subGroup := uint8(m.SubGroup)
    _subGroupErr := io.WriteUint8(8, (subGroup))
    if _subGroupErr != nil {
        return errors.New("Error serializing 'subGroup' field " + _subGroupErr.Error())
    }

    return nil
}

func (m *KNXAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    for {
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "mainGroup":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.MainGroup = data
            case "middleGroup":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.MiddleGroup = data
            case "subGroup":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.SubGroup = data
            }
        }
    }
}

func (m *KNXAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := "org.apache.plc4x.java.knxnetip.readwrite.KNXAddress"
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.MainGroup, xml.StartElement{Name: xml.Name{Local: "mainGroup"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.MiddleGroup, xml.StartElement{Name: xml.Name{Local: "middleGroup"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.SubGroup, xml.StartElement{Name: xml.Name{Local: "subGroup"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

