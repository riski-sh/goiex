package messages

import (
	"fmt"
)

// Below are the two types of update events, transition where the order book
// is performing an action, and complete where the action has been completed.
type PRICE_LEVEL_UPDATE_EVENT IEX_BYTE

const (
	PRICE_LEVEL_UPDATE_EVENT_TRANSITION PRICE_LEVEL_UPDATE_EVENT = 0x0
	PRICE_LEVEL_UPDATE_EVENT_COMPLETE   PRICE_LEVEL_UPDATE_EVENT = 0x1
)

// PriceLevelUpdateMessage message tells us that a level on the bid or ask
// side is going to be changed or has changed.
// See "Consuming Price Level Update Messages and Updating the IEX BBO" in the
// IEX DEEP Specification to understand how to maintain an orderbook using this
// data.
type PriceLevelUpdateMessage struct {
	// MessageType will either be MESSAGES_DEEP10_PRICE_LEVEL_UPDATE_MESSAGE_BUY
	// for level updates on the buy side of the book or
	// MESSAGES_DEEP10_PRICE_LEVEL_UPDATE_MESSAGE_SELL for level updates on the
	// sell side of the book
	MessageType IEX_BYTE

	// EventFlags describes the event of this orderbook change
	EventFlags PRICE_LEVEL_UPDATE_EVENT

	// Timestamp is the time an event triggered the price level update as set
	// by the IEX Trading System logic.
	Timestamp IEX_TIMESTAMP

	// RawSymbol is the symbol this event is happening to represented in
	// Nasdaq Integrated Symbology
	RawSymbol [8]byte

	// Size is the aggregated quote size of the Price level defined by the Price
	// below
	Size IEX_INT

	// Price is the price level to add/update in the IEX Order Book
	Price IEX_PRICE
}

// Symbol returns the RawSymbol as a string type
func (r *PriceLevelUpdateMessage) Symbol() string {
	return string(r.RawSymbol[:])
}

// String prints out this packet in human readable format
func (r *PriceLevelUpdateMessage) String() (string, error) {
	switch r.MessageType {
	case MESSAGES_DEEP10_PRICE_LEVEL_UPDATE_MESSAGE_BUY:
		switch r.EventFlags {
		case PRICE_LEVEL_UPDATE_EVENT_TRANSITION:
			return fmt.Sprintf("%s PriceLevelUpdateMessage %s %s BUY_SIDE level=%d size=%d",
				r.Timestamp.String(), r.Symbol(), "PRICE_LEVEL_UPDATE_EVENT_TRANSITION", r.Price, r.Size), nil
		case PRICE_LEVEL_UPDATE_EVENT_COMPLETE:
			return fmt.Sprintf("%s PriceLevelUpdateMessage %s %s BUY_SIDE level=%d size=%d",
				r.Timestamp.String(), r.Symbol(), "PRICE_LEVEL_UPDATE_EVENT_COMPLETE", r.Price, r.Size), nil
		}
	case MESSAGES_DEEP10_PRICE_LEVEL_UPDATE_MESSAGE_SELL:
		switch r.EventFlags {
		case PRICE_LEVEL_UPDATE_EVENT_TRANSITION:
			return fmt.Sprintf("%s PriceLevelUpdateMessage %s %s SELL_SIDE level=%d size=%d",
				r.Timestamp.String(), r.Symbol(), "PRICE_LEVEL_UPDATE_EVENT_TRANSITION", r.Price, r.Size), nil
		case PRICE_LEVEL_UPDATE_EVENT_COMPLETE:
			return fmt.Sprintf("%s PriceLevelUpdateMessage %s %s SELL_SIDE level=%d size=%d",
				r.Timestamp.String(), r.Symbol(), "PRICE_LEVEL_UPDATE_EVENT_COMPLETE", r.Price, r.Size), nil
		}
	}
	return "", fmt.Errorf("malformed PriceLevelUpdateMessage packet, %+v", r)
}
