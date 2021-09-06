package messages

import (
  "fmt"
)

// SecurityDirectoryMessage is IEX's way of disseminating a full pre-market
// spin of all IEX-listed securities.  After the pre-market spin, IEX will use
// the SecurityDirectoryMessage to relay changes for an individual security.
type SecurityDirectoryMessage struct {
	// MessageType for SecurityDirectoryMessage will always be what is defined
	// above.
	MessageType IEX_BYTE

	// Flags are the unique options identifiable for the SecurityDirectoryMessage
	// For more information on the flags please see Appendix A of the IEX
	// DEEP1.0 specification.
	Flags IEX_BYTE

	// Timestamp is the time of the event when it happened on the IEX Trading
	// system.
	Timestamp IEX_TIMESTAMP

	// Symbol of the IEX-listed security represented in Nasdaq Integrated
	// symbology
	RawSymbol [8]byte

	// RoundLotSize is the number of shares that represent a round lot for the
	// security.
	RoundLotSize IEX_INT

	// AdjustedPOCPrice is the adjusted previous close price given the event of
	// a stock split / dividend etc.., under no event it is simply the previous
	// trading sessions close and if this security has just IPO this will be the
	// issue price.
	AdjustedPOCPrice IEX_PRICE

	// LULDTier, if you need to read this comment to figure out what an LULDTier
	// is you have no need to know what the LULDTier of an asset is.
	LULDTier IEX_BYTE
}

// Symbol converst the 8 character byte array to a string
func (r *SecurityDirectoryMessage) Symbol() string {
	return string(r.RawSymbol[:])
}

// String converts the SecurityDirectoryMessage into a human readable packet
// containing usefull information
func (r *SecurityDirectoryMessage) String() string {
  return fmt.Sprintf("%s SecurityDirectoryMessage %s", r.Timestamp.String(), r.Symbol())
}
