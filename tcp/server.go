package tcp

import (
	"github.com/panjf2000/gnet"
)

type IListener interface {
	OnAccept(conn IConnection) IConnection
	OnTcpClose(conn IConnection) error
}

type Server struct {
	*gnet.EventServer
	addr     string
	listener IListener
}

func NewServer(addr string, listener IListener, opt gnet.Option) (*Server, error) {
	s := &Server{
		addr:     addr,
		listener: listener,
	}

	err := gnet.Serve(s, addr, opt)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// OnShutdown fires when the engine is being shut down, it is called right after
// all event-loops and connections are closed.
func (s *Server) OnShutdown(gs gnet.Server) {
}

// OnOpen fires when a new connection has been opened.
// The parameter out is the return value which is going to be sent back to the peer.
func (s *Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	baseConn := NewConnection(c)
	conn := s.listener.OnAccept(baseConn)
	c.SetContext(conn)

	return
}

// OnClose fires when a connection has been closed.
// The parameter err is the last known connection error.
func (s *Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	conn := getConnection(c)
	conn.OnTcpClose()
	s.listener.OnTcpClose(conn)

	return
}

// OnTraffic fires when a local socket receives data from the peer.
func (s *Server) React(packet []byte, c gnet.Conn) ([]byte, gnet.Action) {

	conn := getConnection(c)

	return conn.OnTcpRread(packet)
}

func getConnection(c gnet.Conn) IConnection {
	ctxPtr := c.Context()
	if ctxPtr == nil {
		return nil
	}

	return ctxPtr.(IConnection)
}
