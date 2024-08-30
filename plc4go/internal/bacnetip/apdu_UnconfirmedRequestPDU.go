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

package bacnetip

import (
	"github.com/pkg/errors"

	readWriteModel "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
)

type UnconfirmedRequestPDU struct {
	*___APDU

	// args
	argChoice *readWriteModel.BACnetUnconfirmedServiceChoice

	pduType readWriteModel.ApduType
}

func NewUnconfirmedRequestPDU(opts ...func(*UnconfirmedRequestPDU)) (*UnconfirmedRequestPDU, error) {
	u := &UnconfirmedRequestPDU{
		pduType: readWriteModel.ApduType_UNCONFIRMED_REQUEST_PDU,
	}
	for _, opt := range opts {
		opt(u)
	}
	apdu, err := new_APDU()
	if err != nil {
		return nil, errors.Wrap(err, "error creating _APDU")
	}
	u.___APDU = apdu.(*___APDU)

	u.apduType = &u.pduType
	u.apduService = (*uint8)(u.argChoice)

	return u, nil
}

func WithUnconfirmedRequestPDU(choice readWriteModel.BACnetUnconfirmedServiceChoice) func(*UnconfirmedRequestPDU) {
	return func(u *UnconfirmedRequestPDU) {
		u.argChoice = &choice
	}
}