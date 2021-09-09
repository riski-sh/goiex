package messages

import (
	"fmt"
)

// TradeBreakMessage are sent when an execution on IEX is broken on that same
// trading day.  Trade breaks are rare and only affect applications that rely
// upon IEX execution based data
type TradeBreakMessage struct {
	// MessageType is always MESSAGES_DEEP10_TRADE_BREAK_MESSAGE
	MessageType IEX_BYTE

	// SaleConditionFlags, see Appendex A in the IEX DEEP Specification to
	// understand how to interpret this field
	SaleConditionFlags IEX_BYTE

	// Timestamp the time an event triggered the trade as set by the IEX Trading
	// System logic.
	Timestamp IEX_TIMESTAMP

	// RawSymbol is the traded security represented in Nasdaq Intregrated Symbology
	RawSymbol [8]byte

	// Size is the size of the broken trade
	Size IEX_INT

	// Price is the execution price of the broken trade
	Price IEX_PRICE

	// TradeID is IEX Generated trade identifier, which is unique for the day
	// the trade it was traded, for the broken trade.
	TradeID IEX_LONG
}

// Symbol returns RawSymbol as a string
func (r *TradeBreakMessage) Symbol() string {
	return string(r.RawSymbol[:])
}

// String will convert the TradeReportMessage packet into a human readable
// string containing important information
func (r *TradeBreakMessage) String() (string, error) {
	return fmt.Sprintf("%s TradeBreakMessage %s price=%d size=%d tradeid=%d",
		r.Timestamp.String(), r.Symbol(), r.Price, r.Size, r.TradeID), nil
}
