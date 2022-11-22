package main

import (
	"github.com/let-light/network/tcp"
	"github.com/panjf2000/gnet"
	goPool "github.com/panjf2000/gnet/pkg/pool/goroutine"
)

type Manager struct {
	session *Session
}

type Session struct {
	tcp.IConnection
	pool *goPool.Pool
}

func (s *Session) OnTcpRread(buf []byte) ([]byte, gnet.Action) {
	s.pool.Submit(func() {
		// do something

		// write
		s.Write(buf)
	})

	return nil, gnet.None
}

func (m *Manager) OnAccept(conn tcp.IConnection) tcp.IConnection {
	s := &Session{
		IConnection: conn,
		pool:        goPool.Default(),
	}

	m.session = s

	return s
}

func (m *Manager) OnTcpClose(conn tcp.IConnection) error {
	return nil
}

// func (m *Manager) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
// 	return buf, nil
// }

// func (m *Manager) Decode(c gnet.Conn) ([]byte, error) {
// 	return nil, nil
// }

func main() {
	_, err := tcp.NewServer(":8082", &Manager{}, gnet.WithMulticore(true))
	if err != nil {
		panic(err)
	}
}
