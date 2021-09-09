package messages

import (
	"fmt"
)

// TradeReportMessage is a message sent when and order is executed in whole
// or in part
type TradeReportMessage struct {
	// MessageType is always MESSAGES_DEEP10_TRADE_REPORT_MESSAGE
	MessageType IEX_BYTE

	// SaleConditionFlags, see Appendex A in the IEX DEEP Specification to
	// understand how to interpret this field
	SaleConditionFlags IEX_BYTE

	// Timestamp the time an event triggered the trade as set by the IEX Trading
	// System logic.
	Timestamp IEX_TIMESTAMP

	// RawSymbol is the traded security represented in Nasdaq Intregrated Symbology
	RawSymbol [8]byte

	// Size is the size of th trade represented in the number of shares
	Size IEX_INT

	// Price is the execution price of the trade
	Price IEX_PRICE

	// TradeID is IEX Generated trade identifier, which is unique for the day
	// the trade it was traded.
	TradeID IEX_LONG
}

// Symbol returns RawSymbol as a string
func (r *TradeReportMessage) Symbol() string {
	return string(r.RawSymbol[:])
}

// String will convert the TradeReportMessage packet into a human readable
// string containing important information
func (r *TradeReportMessage) String() (string, error) {
	return fmt.Sprintf("%s TradeReportStatusMessage %s price=%d size=%d tradeid=%d",
		r.Timestamp.String(), r.Symbol(), r.Price, r.Size, r.TradeID), nil
}
