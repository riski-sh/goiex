package messages

import (
	"fmt"
)

// Below define the two types of security event messages
type SECURITY_EVENT IEX_BYTE

const (
	SECURITY_EVENT_OPENING_PROCESS_COMPLETE SECURITY_EVENT = 0x4f
	SECURITY_EVENT_CLOSING_PROCESS_COMPLETE SECURITY_EVENT = 0x43
)

// SecurityEventMessage is used to indicate events applied to a security
type SecurityEventMessage struct {
	// MessageType will always be MESSAGES_DEEP10_SECURITY_EVENT_MESSAGE
	MessageType IEX_BYTE

	// SecurityEvent either tells us that opening or closing process is complete
	SecurityEvent SECURITY_EVENT

	// Timestamp is the time of the update event as set by the IEX Trading System
	// logic
	Timestamp IEX_TIMESTAMP

	// RawSymbol the raw symbol this security event message corrisponds to
	RawSymbol [8]byte
}

// Symbol returns the RawSymbol as a string type
func (r *SecurityEventMessage) Symbol() string {
	return string(r.RawSymbol[:])
}

// String converts a SecurityEventMessage packet into readable text containing
// usefull information.
func (r *SecurityEventMessage) String() (string, error) {
	switch r.SecurityEvent {
	case SECURITY_EVENT_OPENING_PROCESS_COMPLETE:
		return fmt.Sprintf("%s SecurityEventMessage %s %s", r.Timestamp.String(),
			r.Symbol(), "SECURITY_EVENT_OPENING_PROCESS_COMPLETE"), nil
	case SECURITY_EVENT_CLOSING_PROCESS_COMPLETE:
		return fmt.Sprintf("%s SecurityEventMessage %s %s", r.Timestamp.String(),
			r.Symbol(), "SECURITY_EVENT_CLOSING_PROCESS_COMPLETE"), nil
	}
	return "", fmt.Errorf("malformed SecurityEventMessge packet, %+v", r)
}
