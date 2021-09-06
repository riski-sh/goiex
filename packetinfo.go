package goiex

import (
	"time"

	. "github.com/riski-sh/goiex/messages"
)

// Below are the trading status that IEX will show for both IEX and non IEX
// stocks.
const (
	TRADE_STATUS_TRADING_HAULTED        = 0x48
	TRADE_STATUS_TRADING_HAULT_RELEASED = 0x4f
	TRADE_STATUS_TRADING_PAUSED         = 0x50
	TRADE_STATUS_TRADING                = 0x54
)

// The IEXTP (IEX Transport Protocol) operates under UDP multicast.
// The header of the UDP payload is defined by the IEXTPHeader struct below.
type IEXTPHeader struct {

	// Version of the IEX-TP protocol.
	Version IEX_BYTE

	// Reserved is a reserve bit due to structure packing on the wire level.
	Reserved IEX_BYTE

	// MessageProtocolID is a unique identifier of the higher-layer protocol
	MessageProtocolID IEX_SHORT

	// ChannelID identifies the stream of bytes/sequenced messages
	ChannelID IEX_INT

	// SessionID identifies the session
	SessionID IEX_INT

	// PayloadLength byte length of the payload
	PayloadLength IEX_SHORT

	// MessageCount number of messages in the payload
	MessageCount IEX_SHORT

	// StreamOffset byte offset of the data stream
	StreamOffset IEX_LONG

	// FirstMessageSequenceNumber is the sequence number of the first message
	// in the segment.
	FirstMessageSequenceNumber IEX_LONG

	// SendTime is the time that this segment left the sender.
	SendTime IEX_TIMESTAMP
}

// MessageBlock contains the length and the message payload.
type MessageBlock struct {
	// MessageLength is known and is first short or the message block.
	// MessageLength tells us how many bytes to read to collect the
	// message data associated with this message block.
	MessageLength IEX_SHORT

	//MessageData is the message payload with a length of MessageLength bytes
	MessageData []byte
}

// CallbackConfig holds all the callbacks that are possible when communicating
// with the IEX exchange.
type CallbackConfig struct {
	// OnSystemEventMessage callback
	OnSystemEventMessage func(event SystemEventMessage)

	// OnSecurityDirectoryMessage callback
	OnSecurityDirectoryMessage func(event SecurityDirectoryMessage)

	// OnTradingStatusMessage callback
	OnTradingStatusMessage func(event TradingStatusMessage)

	// OnOperationalHaultStatusMessage callback
	OnOperationalHaultStatusMessage func(event OperationalHaultStatusMessage)

  // OneShortSalePriceTestStatusMessage callback
  OnShortSalePriceTestStatusMessage func(event ShortSalePriceTestStatusMessage)

  // OnSecurityEventMessage callback
  OnSecurityEventMessage func(event SecurityEventMessage)
}

// Timestamp returns the SendTime of the IEXTPHeader as a nanosecond date time
// in UTC
func (r *IEXTPHeader) Timestamp() time.Time {
	return time.Unix(0, int64(r.SendTime))
}
