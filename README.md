# GoIEX

GoIEX is a library to provide backtesting and replay capabilities the IEX
exchange. Leveraging packet replay, the event stream of the PCAP can be replayed
with very consistent time replication inbetween packets for "fake-live" testing.
Also a mass backtest is possible through just reading the pcap file.

GoIEX does **not** maintain orderbook, best bid/ask, or retain any information
when reading and dispatching events. It is _impossible_ for the GoIEX reader
to leak forward data on its own or remember previous information in any way.

It is up to the implementor to implement the callbacks and parse the data
correctly.

## iexdump
iexdump is a simple utility that dumps IEX pcap data in readable format. This
tool is simply used a verification method that the PCAP file can be read and
all callback events are accounted for.
