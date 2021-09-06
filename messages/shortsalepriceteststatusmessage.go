package messages

import (
	"fmt"
)

// Below are the constants used to determine the short sale price test status
type SHORT_SALE_PRICE_TEST_STATUS IEX_BYTE

const (
	SHORT_SALE_PRICE_TEST_STATUS_NOT_IN_EFFECT SHORT_SALE_PRICE_TEST_STATUS = 0x0
	SHORT_SALE_PRICE_TEST_STATUS_IN_EFFECT     SHORT_SALE_PRICE_TEST_STATUS = 0x1
)

// Below are the constants used to determine the details of the restriction
type SHORT_SALE_PRICE_TEST_STATUS_DETAIL IEX_BYTE

const (
	SHORT_SALE_PRICE_TEST_STATUS_DETAIL_NONE        SHORT_SALE_PRICE_TEST_STATUS_DETAIL = 0x20
	SHORT_SALE_PRICE_TEST_STATUS_DETAIL_ACTIVATED   SHORT_SALE_PRICE_TEST_STATUS_DETAIL = 0x41
	SHORT_SALE_PRICE_TEST_STATUS_DETAIL_CONTINUED   SHORT_SALE_PRICE_TEST_STATUS_DETAIL = 0x43
	SHORT_SALE_PRICE_TEST_STATUS_DETAIL_DEACTIVATED SHORT_SALE_PRICE_TEST_STATUS_DETAIL = 0x44
	SHORT_SALE_PRICE_TEST_STATUS_DETAIL_UNKNOWN     SHORT_SALE_PRICE_TEST_STATUS_DETAIL = 0x4e
)

// ShortSalePriceTestStatusMessage In association with Rule 201 of Regulation
// SHO, the Short Sale Price Test Message is used to indicate when a short sale
// price test restriction is in effect for a security.
type ShortSalePriceTestStatusMessage struct {
	// MessageType will always be MESSAGES_DEEP10_SHORTSALEPRICETESTMESSAGE
	MessageType IEX_BYTE

	// ShortSalePriceTestStatus will show the effect of the short sale price
	// status
	ShortSalePriceTestStatus SHORT_SALE_PRICE_TEST_STATUS

	// Timestamp is the time of the update event as set by the IEX Trading System
	// logic.
	Timestamp IEX_TIMESTAMP

	// RawSymbol is the security represented in Nasdaq Integrated Symbology
	RawSymbol [8]byte

	// Detail is the reason why the short sale price test was activatd or none
	// for no ongoing short sale price
	Detail SHORT_SALE_PRICE_TEST_STATUS_DETAIL
}

// Symbol returns the RawSymbol as a string type
func (r *ShortSalePriceTestStatusMessage) Symbol() string {
	return string(r.RawSymbol[:])
}

// String converts the ShortSalePriceTestStatusMessage into a readable text
// string containing useful information.
func (r *ShortSalePriceTestStatusMessage) String() (string, error) {
	switch r.ShortSalePriceTestStatus {
	case SHORT_SALE_PRICE_TEST_STATUS_NOT_IN_EFFECT:
		return fmt.Sprintf("%s ShortSalePriceTestMessage %s %s", r.Timestamp.String(),
			r.Symbol(), "SHORT_SALE_PRICE_TEST_STATUS_NOT_IN_EFFECT"), nil
	case SHORT_SALE_PRICE_TEST_STATUS_IN_EFFECT:
		return fmt.Sprintf("%s ShortSalePriceTestMessage %s %s", r.Timestamp.String(),
			r.Symbol(), "SHORT_SALE_PRICE_TEST_STATUS_IN_EFFECT"), nil
	}
	return "", fmt.Errorf("%s", "malformed ShortSalePriceTestStatusMessage package")
}
