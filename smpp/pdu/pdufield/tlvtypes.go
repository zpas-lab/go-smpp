// Copyright 2015 go-smpp authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pdufield

import "fmt"


type (
// TLVType is the Tag Length Value.
	TLVType uint16

// MessageState is sent by SMSC to ESME to indicate the final message state for an SMS DeliveryReport
	MessageStateValue byte
)


const (
	TLVReceiptedMessageId = TLVType(0x001e)
	TLVMessagePayload = TLVType(0x0424)
	TLVMessageState = TLVType(0x0427)
)

var tlvNames = map[TLVType]string{
	TLVReceiptedMessageId: "receipted_message_id",
	TLVMessagePayload:     "message_payload",
	TLVMessageState:       "message_state",
}

func (t TLVType) String() string {
	name, ok := tlvNames[t]
	if !ok {
		return fmt.Sprintf("0x%X", t)
	}
	return name
}


const (
	ENROUTE = MessageStateValue(0x01)
	DELIVERED = MessageStateValue(0x02)
	EXPIRED = MessageStateValue(0x03)
	DELETED = MessageStateValue(0x04)
	UNDELIVERABLE = MessageStateValue(0x05)
	ACCEPTED = MessageStateValue(0x06)
	UNKNOWN = MessageStateValue(0x07)
	REJECTED = MessageStateValue(0x08)
)

var messageStateNames = map[MessageStateValue]string{
	ENROUTE:       "ENROUTE",
	DELIVERED:     "DELIVERED",
	EXPIRED:       "EXPIRED",
	DELETED:       "DELETED",
	UNDELIVERABLE: "UNDELIVERABLE",
	ACCEPTED:      "ACCEPTED",
	UNKNOWN:       "UNKNOWN",
	REJECTED:      "REJECTED",
}

func (ms MessageStateValue) String() string {
	name, ok := messageStateNames[ms]
	if !ok {
		return fmt.Sprintf("%x", ms)
	}
	return name
}
