package goiex

import (
	"bytes"
	"encoding/binary"
  "fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	. "github.com/riski-sh/goiex/messages"
)

// Playback reads an IEX deep pcap file and calls the appropriate event
// on event functions. Playback takes a file variable which is the relative
// path to the pcap file to read. PlaybackDeep also requires the
// callbackconfig structure in order to dispatch events to be processed
// somewhere else.
func PlaybackDeep(file string, callbacks CallbackConfig) error {
	if handle, err := pcap.OpenOffline(file); err != nil {
	  return err
  } else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			// Verify that the network packet was decoded successfully.
			if packet.ErrorLayer() != nil {
				return packet.ErrorLayer().Error()
			}

			payload := packet.Layer(layers.LayerTypeUDP).LayerPayload()

			header := IEXTPHeader{}
			err := binary.Read(bytes.NewBuffer(payload[:40]), binary.LittleEndian, &header)

			if err != nil {
				return err
			}

			// Check if the message count and payload length are zero if they are
			// then this packet contains no state changes, IEX-TP defines this
			// message as a heartbeat
			if header.MessageCount == 0 && header.PayloadLength == 0 {
				continue
			}

			// Loop through all the messages in the payload
			messagesRead := uint16(0)
			packetStride := uint16(40)
			for messagesRead < header.MessageCount {
				var msgBlock MessageBlock
				msgBlock.MessageLength = binary.LittleEndian.Uint16(payload[packetStride : packetStride+2])
				packetStride += 2
				msgBlock.MessageData = payload[packetStride : packetStride+msgBlock.MessageLength]
				packetStride += msgBlock.MessageLength

				// Now that the MessageBlock is fully read we must cast the data portion
				// of the message block. The first byte of the data message describes
				// what message we have recieved. Only then we can cast to the
				// appropirate data structure to be passed along to the callback
				messageDataBuff := bytes.NewBuffer(msgBlock.MessageData)
				switch msgBlock.MessageData[0] {
				case MESSAGES_DEEP10_SYSTEM_EVENT_MESSAGE:
					event := SystemEventMessage{}
          err := binary.Read(messageDataBuff, binary.LittleEndian, &event)
					if err != nil {
            return err
          }
          callbacks.OnSystemEventMessage(event)
					break
				case MESSAGES_DEEP10_SECURITY_DIRECTORY_MESSAGE:
					event := SecurityDirectoryMessage{}
          err := binary.Read(messageDataBuff, binary.LittleEndian, &event)
					if err != nil {
            return err
          }
          callbacks.OnSecurityDirectoryMessage(event)
					break
				case MESSAGES_DEEP10_TRADING_STATUS_MESSAGE:
					event := TradingStatusMessage{}
					err := binary.Read(messageDataBuff, binary.LittleEndian, &event)
					if err != nil {
            return err
          }
          callbacks.OnTradingStatusMessage(event)
					break
				case MESSAGES_DEEP10_OPERATIONAL_HAULT_STATUS_MESSAGE:
					event := OperationalHaultStatusMessage{}
          err := binary.Read(messageDataBuff, binary.LittleEndian, &event)
					if err != nil {
            return err
          }
          callbacks.OnOperationalHaultStatusMessage(event)
					break
				}
				messagesRead += 1
			}

			if packetStride-40 != header.PayloadLength {
				return fmt.Errorf("expected to read %d bytes but only read %d", packetStride-40, header.PayloadLength)
			}
		}
	}
  return nil
}
