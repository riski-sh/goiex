package messages

import (
	"time"
)

// Type aliases
type (
	IEX_LONG  = int64
	IEX_INT   = uint32
	IEX_SHORT = uint16
	IEX_BYTE  = uint8
)

// Below are constants of the DEEP10 message types
const (
	MESSAGES_DEEP10_SYSTEM_EVENT_MESSAGE             = 0x53
	MESSAGES_DEEP10_SECURITY_DIRECTORY_MESSAGE       = 0x44
	MESSAGES_DEEP10_TRADING_STATUS_MESSAGE           = 0x48
	MESSAGES_DEEP10_OPERATIONAL_HAULT_STATUS_MESSAGE = 0x4f
	MESSAGES_DEEP10_SHORT_SALE_TEST_STATUS_MESSAGE   = 0x50
	MESSAGES_DEEP10_SECURITY_EVENT_MESSAGE           = 0x45
	MESSAGES_DEEP10_PRICE_LEVEL_UPDATE_MESSAGE_BUY   = 0x38
	MESSAGES_DEEP10_PRICE_LEVEL_UPDATE_MESSAGE_SELL  = 0x35
	MESSAGES_DEEP10_TRADE_REPORT_MESSAGE             = 0x54
	MESSAGES_DEEP10_OFFICIAL_PRICE_MESSAGE           = 0x58
	MESSAGES_DEEP10_TRADE_BREAK_MESSAGE              = 0x42
	MESSAGES_DEEP10_AUCTION_INFORMATION_MESSAGE      = 0x41
)

// IEX_TIMESTAMP is the number of nanoseconds since EPOCH
type IEX_TIMESTAMP int64

var easternTime, _ = time.LoadLocation("America/New_York")

// String converts the IEX_TIMESTAMP type to a time date string
// represented in the US/Eastern time
func (r IEX_TIMESTAMP) String() string {
	return time.Unix(0, int64(r)).In(easternTime).Format(time.RFC3339Nano)
}

// Time converts the IEX_TIMESTAMP type to a unix timestamp.
// This assumes the IEX_TIMESTAMP is in UTC time
func (r IEX_TIMESTAMP) Time() time.Time {
	return time.Unix(0, int64(r))
}

// IEX_PRICE is the price represented as a fixed-point number with 4 digits to
// the right of an implied decimal point
type IEX_PRICE int64
