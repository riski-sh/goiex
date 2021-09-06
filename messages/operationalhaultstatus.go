package messages

import (
	"fmt"
)

// Below are the operation hault statues if trades have been haulted due
// to an IEX reason and not a global hault
type OPERATIONAL_HAULT_STATUS IEX_BYTE

const (
	OPERATIONAL_HAULT_STATUS_HAULTED    OPERATIONAL_HAULT_STATUS = 0x4f
	OPERATIONAL_HAULT_STATUS_NOT_HALTED OPERATIONAL_HAULT_STATUS = 0x4e
)

// OperationalHaultStatusMessage are haults specificly happening on the IEX
// exchange due to an IEX reason.
type OperationalHaultStatusMessage struct {
	// MessageType is always what is defined above
	MessageType IEX_BYTE

	// OperationHaultStatus is always a value under OPERATIONAL_HAULT_STATUS_*
	OperationalHaultStatus OPERATIONAL_HAULT_STATUS

	// Timestamp the time of the update event as set by teh IEX Trading System
	// logic.
	Timestamp IEX_TIMESTAMP

	// RawSymbol is the symbol that this operation hault message refers to
	RawSymbol [8]byte
}

// Symbol returns RawSymbol as a string
func (r *OperationalHaultStatusMessage) Symbol() string {
	return string(r.RawSymbol[:])
}

// String converts this packet into a human readable string containing
// the important information.
func (r *OperationalHaultStatusMessage) String() (string, error) {
	switch r.OperationalHaultStatus {
	case OPERATIONAL_HAULT_STATUS_NOT_HALTED:
		return fmt.Sprintf("%s OperationalHaultStatus %s %s", r.Timestamp.String(),
           r.Symbol(), "OPERATIONAL_HAULT_STATUS_HAULTED"), nil
	case OPERATIONAL_HAULT_STATUS_HAULTED:
		return fmt.Sprintf("%s OperationalHaultStatus %s %s", r.Timestamp.String(),
           r.Symbol(), "OPERATIONAL_HAULT_HAULTED"), nil
	}
	return "", fmt.Errorf("malformed OperationalHaultStatusMessage packet, %+v", r)
}
