package messages

import (
	"fmt"
)

// Below are the status for the official price message defining if the
// official price message is for the open or close.
type OFFICIAL_PRICE IEX_BYTE

const (
	OFFICIAL_PRICE_OPENING OFFICIAL_PRICE = 0x51
	OFFICIAL_PRICE_CLOSING OFFICIAL_PRICE = 0x4d
)

// OfficialPriceMessage is a message defining the official open or closing price
// of a security
type OfficialPriceMessage struct {
	// MessageType is the type
	MessageType IEX_BYTE

	// PriceType represents if the price is for opening or closing official
	// prices.
	PriceType OFFICIAL_PRICE

	// Timestamp the time of the update event as set by teh IEX Trading System
	// logic.
	Timestamp IEX_TIMESTAMP

	// RawSymbol is the symbol that this operation hault message refers to
	RawSymbol [8]byte

	// OfficialPrice is the official opening or closing price.
	OfficialPrice IEX_PRICE
}

// Symbol returns RawSymbol as string type
func (r *OfficialPriceMessage) Symbol() string {
	return string(r.RawSymbol[:])
}

// String converts this packet into a human readable string containing
// the import information.
func (r *OfficialPriceMessage) String() (string, error) {
	switch r.PriceType {
	case OFFICIAL_PRICE_OPENING:
		return fmt.Sprintf("%s OfficialPriceMessage %s OFFICIAL_PRICE_OPENING price=%d",
			r.Timestamp.String(), r.Symbol(), r.OfficialPrice), nil
	case OFFICIAL_PRICE_CLOSING:
		return fmt.Sprintf("%s OfficialPriceMessage %s OFFICIAL_PRICE_CLOSING price=%d",
			r.Timestamp.String(), r.Symbol(), r.OfficialPrice), nil
	}
	return "", fmt.Errorf("malformed OfficialPriceMessage packet, %+v", r)
}
