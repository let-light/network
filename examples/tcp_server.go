package main

import (
	"github.com/let-light/network/tcp"
	"github.com/panjf2000/gnet"
	goPool "github.com/panjf2000/gnet/pkg/pool/goroutine"
)

type Manager struct {
	conn *Connection
}

type Connection struct {
	tcp.IContext
	pool *goPool.Pool
}

func (c *Connection) OnTcpRread(buf []byte) ([]byte, gnet.Action) {
	c.pool.Submit(func() {
		// do something

		// write
		c.Write(buf)
	})

	return nil, gnet.None
}

func (m *Manager) NewOrGet() tcp.IContext {
	conn := &Connection{
		IContext: tcp.NewContext(),
		pool:     goPool.Default(),
	}

	m.conn = conn

	return conn
}

func (m *Manager) OnTcpClose(ctx tcp.IContext) error {
	return nil
}

// func (m *Manager) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
// 	return buf, nil
// }

// func (m *Manager) Decode(c gnet.Conn) ([]byte, error) {
// 	return nil, nil
// }

func main() {
	_, err := tcp.NewServer(":8081", &Manager{}, gnet.WithMulticore(true))
	if err != nil {
		panic(err)
	}
}
