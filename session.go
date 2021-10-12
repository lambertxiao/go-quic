package quic

import "go-quic/protocol"

type PacketHandler interface {
	// handlePacket(addr interface{}, hdr *PublicHeader, data []byte)
}

type Session struct {
	sessionId protocol.ConnectionID
}
