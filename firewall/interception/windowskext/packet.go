package windowskext

import (
	"github.com/tevino/abool"

	"github.com/Safing/portmaster/network/packet"
)

type Packet struct {
	packet.PacketBase

	kextID     uint32
	packetData []byte

	verdictSet *abool.AtomicBool
}

func (pkt *Packet) Accept() error {
	if pkt.verdictSet.SetToIf(false, true) {
		return pkt.windivert.Send(pkt.packetData, pkt.packetAddress)
	}
	return nil
}

func (pkt *Packet) Block() error {
	if pkt.verdictSet.SetToIf(false, true) {
		// TODO: implement blocking mechanism
		return nil
	}
	return nil
}

func (pkt *Packet) Drop() error {
	return nil
}

func (pkt *Packet) PermanentAccept() error {
	return pkt.Accept()
}

func (pkt *Packet) PermanentBlock() error {
	return pkt.Block()
}

func (pkt *Packet) PermanentDrop() error {
	return pkt.Drop()
}

func (pkt *Packet) RerouteToNameserver() error {
	return nil
}

func (pkt *Packet) RerouteToTunnel() error {
	return nil
}
