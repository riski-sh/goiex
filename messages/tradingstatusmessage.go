package messages

import (
	"fmt"
)

// Below are the trading status that IEX will show for both IEX and non IEX
// stocks.
type TRADING_STATUS IEX_BYTE

const (
	TRADING_STATUS_HAULTED     TRADING_STATUS = 0x48
	TRADING_STATUS_OPERATIONAL TRADING_STATUS = 0x4f
	TRADING_STATUS_PAUSED      TRADING_STATUS = 0x50
	TRADING_STATUS_TRADING     TRADING_STATUS = 0x54
)

// TradingStatusMessage is used to indicated the trading status of a security.
// TradingStatusMessage will trigger when a hault happens and when a resume
// happens.
type TradingStatusMessage struct {
	// MessageType will always be what is defined above for TradingStatusMessage
	MessageType IEX_BYTE

	// TradingStatus is the status and will always be a value of TRADE_STATUS_*
	TradingStatus TRADING_STATUS

	// Timestamp is the time of the update event set by the IEX Trading System
	// logic.
	Timestamp IEX_TIMESTAMP

	// RawSymbol is the symbol that this trading status message is reffering to
	// as represented in Nasaq Integrated Symbology
	RawSymbol [8]byte

	// Reason is the reason why the trading status has changed for this symbol.
	// There are a host of reasons and values. Here are them all along with the
	// string that corrisponds to them.
	//
	// Trading Hault Reasons
	//  T1: Hault News Pending
	//  IP1: IPO Not Yet Trading
	//  IPOD: IPO Deferred
	//  MCB3: Market-Wide Circuit Breaker Level 3 - Breached
	//  NA: Reason Not Available
	// Order Acceptance Period Resons
	//  T2: Hault News Dissemination
	//  IPO2: IPO Order Acceptance Period
	//  IPO3: IPO Pre-Launch Period
	//  MCB1: Market-Wide Circuit Breaker Level 1 - Breached
	//  MCB2: Market-Wide Circuit Breaker Level 2 - Breached
	RawReason [4]byte
}

// Symbol returns a string version of the RawSymbol
func (r *TradingStatusMessage) Symbol() string {
	return string(r.RawSymbol[:])
}

// Reaons returns a string version of the RawReason
func (r *TradingStatusMessage) Reason() string {
	return string(r.RawReason[:])
}

// String prints out this packet in human readable format
func (r *TradingStatusMessage) String() (string, error) {
	switch r.TradingStatus {
	case TRADING_STATUS_TRADING:
		return fmt.Sprintf("%s TradingStatusMessage %s %s", r.Timestamp.String(),
           r.Symbol(), "TRADING_STATUS_TRADING"), nil
	case TRADING_STATUS_PAUSED:
		return fmt.Sprintf("%s TradingStatusMessage %s %s", r.Timestamp.String(),
           r.Symbol(), "TRADING_STATUS_PAUSED"), nil
	case TRADING_STATUS_OPERATIONAL:
		return fmt.Sprintf("%s TradingStatusMessage %s %s", r.Timestamp.String(),
           r.Symbol(), "TRADING_STATUS_OPERATIONAL"), nil
	case TRADING_STATUS_HAULTED:
		return fmt.Sprintf("%s TradingStatusMessage %s %s", r.Timestamp.String(),
           r.Symbol(), "TRADING_STATUS_HAULTED"), nil
	}
	return "", fmt.Errorf("malformed TradingStatusMessage packet")
}
