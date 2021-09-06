package main

import (
	"flag"
	"fmt"

	. "github.com/riski-sh/goiex"
	. "github.com/riski-sh/goiex/messages"
)

// SystemEventMessageHandler is the callback function for System Event Messages
// Every SYSTEM_EVENT_MESSAGE_* will be seen throughout the trading day.
func SystemEventMessageHandler(event SystemEventMessage) {
  str, err := event.String()
  if err != nil {
    panic(err)
  }
  fmt.Printf("%s\n", str)
}

// SecurityDirectoryMessageHandler is the callback function for the
// SecurityDirectoryMessage messages. This callback is pretty much useless now
// because IEX doesn't list any symbols anymore and the symbols they do
// propogate through their network are test symbols that can not be traded.
// It is safe to ignore these messages but the callback is still present due
// to historical reasons.
func SecurityDirectoryMessageHandler(event SecurityDirectoryMessage) {
  str, err := event.String()
  if err != nil {
    panic(err)
  }
  fmt.Printf("%s\n", str)
}

// TradingStatusMessageHandler is the callback function for the
// TradingStatusMessage messages. This callback is triggered when a security
// is haulted or resumed.
func TradingStatusMessageHandler(event TradingStatusMessage) {
  str, err := event.String()
  if err != nil {
    panic(err)
  }
  fmt.Printf("%s\n", str)
}

// OperationalHaultStatusMessageHandler is the callback function for the
// OperaationalHaultStatusMessage which is only triggered when IEX imposes a
// hault on one of its securities. This does not mean a hault is happening
// market wide.
func OperationalHaultStatusMessageHandler(event OperationalHaultStatusMessage) {
  str, err := event.String()
  if err != nil {
    panic(err)
  }
  fmt.Printf("%s\n", str)
}

func ShortSalePriceTestStatusMessageHandler(event ShortSalePriceTestStatusMessage) {
  str, err := event.String()
  if err != nil {
    panic(err)
  }
  fmt.Printf("%s\n", str)
}

func main() {
	pcapstring := flag.String("pcapdeep", "", "specify the pcap representing a DEEP IEX pcap dump")
	flag.Parse()

	if *pcapstring == "" {
		flag.PrintDefaults()
		return
	}

	err := PlaybackDeep(*pcapstring, CallbackConfig{
		OnSystemEventMessage:            SystemEventMessageHandler,
		OnSecurityDirectoryMessage:      SecurityDirectoryMessageHandler,
		OnTradingStatusMessage:          TradingStatusMessageHandler,
		OnOperationalHaultStatusMessage: OperationalHaultStatusMessageHandler,
    OnShortSalePriceTestStatusMessage: ShortSalePriceTestStatusMessageHandler,
  })

	if err != nil {
		panic(err)
	}
}
