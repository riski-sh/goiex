package messages

import (
	"fmt"
)

// Below are constants describing the possible System Event values
type SYSTEM_EVENT_MESSAGE IEX_BYTE

const (
	SYSTEM_EVENT_MESSAGE_START_OF_MESSAGES             SYSTEM_EVENT_MESSAGE = 0x4f
	SYSTEM_EVENT_MESSAGE_START_OF_SYSTEM_HOURS         SYSTEM_EVENT_MESSAGE = 0x53
	SYSTEM_EVENT_MESSAGE_START_OF_REGULAR_MARKET_HOURS SYSTEM_EVENT_MESSAGE = 0x52
	SYSTEM_EVENT_MESSAGE_END_OF_REGULAR_MARKET_HOURS   SYSTEM_EVENT_MESSAGE = 0x4d
	SYSTEM_EVENT_MESSAGE_END_OF_SYSTEM_HOURS           SYSTEM_EVENT_MESSAGE = 0x45
	SYSTEM_EVENT_MESSAGE_END_OF_MESSAGES               SYSTEM_EVENT_MESSAGE = 0x43
)

// SystemEventMessage is used to indicate events that apply to the market or
// the data feed. There will be a single message disseminated per channel for
// each System Event type within a given trading session.
type SystemEventMessage struct {
	// MessageType for SystemEventMessage will always be what is defined above
	MessageType IEX_BYTE

	// SystemEvent has a value of one of the SYSTEM_EVENT_MESSAGE_* constants
	// witch describes the event.
	SystemEvent SYSTEM_EVENT_MESSAGE

	// Timestamp is the time of the event when it happened on the IEX Trading
	// system
	Timestamp IEX_TIMESTAMP
}

// String formats the SystemEventMessage into a human readable string containing
// important information.
func (r *SystemEventMessage) String() (string, error) {
	switch r.SystemEvent {
	case SYSTEM_EVENT_MESSAGE_START_OF_MESSAGES:
		return fmt.Sprintf("%s SystemEventMessage %s", r.Timestamp.String(),
			"SYSTEM_EVENT_MESSAGE_START_OF_MESSAGES"), nil
	case SYSTEM_EVENT_MESSAGE_START_OF_SYSTEM_HOURS:
		return fmt.Sprintf("%s SystemEventMessage %s", r.Timestamp.String(),
			"SYSTEM_EVENT_MESSAGE_START_OF_SYSTEM_HOURS"), nil
	case SYSTEM_EVENT_MESSAGE_START_OF_REGULAR_MARKET_HOURS:
		return fmt.Sprintf("%s SystemEventMessage %s", r.Timestamp.String(),
			"SYSTEM_EVENT_MESSAGE_START_OF_REGULAR_MARKET_HOURS"), nil
	case SYSTEM_EVENT_MESSAGE_END_OF_REGULAR_MARKET_HOURS:
		return fmt.Sprintf("%s SystemEventMessage %s", r.Timestamp.String(),
			"SYSTEM_EVENT_MESSAGE_END_OF_REGULAR_MARKET_HOURS"), nil
	case SYSTEM_EVENT_MESSAGE_END_OF_SYSTEM_HOURS:
		return fmt.Sprintf("%s SystemEventMessage %s", r.Timestamp.String(),
			"SYSTEM_EVENT_MESSAGE_END_OF_SYSTEM_HOURS"), nil
	case SYSTEM_EVENT_MESSAGE_END_OF_MESSAGES:
		return fmt.Sprintf("%s SystemEventMessage %s", r.Timestamp.String(),
			"SYSTEM_EVENT_MESSAGEE_END_OF_MESSAGES"), nil
	}
	return "", fmt.Errorf("malformed SystemEVentMessage packet, %+v", r)
}
