package quic

import (
	"go-quic/protocol"
	"net"
)

type Server struct {
	addr     *net.UDPAddr
	conn     *net.UDPConn
	sessions map[protocol.ConnectionID]PacketHandler
}

func NewServer() (*Server, error) {
	return nil, nil
}

func (s *Server) ListenAndServe() error {
	conn, err := net.ListenUDP("udp", s.addr)
	if err != nil {
		return err
	}

	return s.serve(conn)
}

func (s *Server) serve(conn *net.UDPConn) error {
	s.conn = conn

	for {
		buf := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			return err
		}

		packet := buf[:n]
		s.handlePacket(packet, remoteAddr)
	}
}

func (s *Server) handlePacket(packet []byte, remoteAddr *net.UDPAddr) error {
	// 1. 处理header
	// 2. 拿到connectionId
	// 3. 拿到connection对应的session
	// 4. 通过session处理对应的packet
	return nil
}
