// GoIEX is a custom golang implementation of the IEX exchange protocol.
// Although GoIEX could be used for a live feed (if cross connected) this
// package mainly focuses on creating a historical event based stream from
// the PCAP archive provided by IEX. GoIEX also does not manipulate data in
// any way, instead a callback for the event must be defined and the parent
// program calling GoIEX must handle the incoming data.
package goiex
