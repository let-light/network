package tcp

import (
	"github.com/panjf2000/gnet"
)

type IConnection interface {
	SetConn(c interface{})
	OnTcpClose() error
	OnTcpRread(buf []byte) ([]byte, gnet.Action)
	RemoteAddr() string
	LocalAddr() string
	Network() string
	Write(data []byte) error
}

type Connection struct {
	Conn gnet.Conn
}

func NewConnection(c gnet.Conn) IConnection {
	return &Connection{
		Conn: c,
	}
}

func (conn *Connection) Write(data []byte) error {
	return conn.Conn.AsyncWrite(data)
}

func (conn *Connection) SetConn(c interface{}) {
	if c == nil {
		conn.Conn = nil
	} else {
		conn.Conn = c.(gnet.Conn)
	}
}

func (conn *Connection) OnTcpClose() error {
	return nil
}

func (conn *Connection) OnTcpRread(buf []byte) ([]byte, gnet.Action) {
	return nil, gnet.None
}

func (conn *Connection) RemoteAddr() string {
	if (conn.Conn == nil) || (conn.Conn.RemoteAddr() == nil) {
		return ""
	}

	return conn.Conn.RemoteAddr().String()
}

func (conn *Connection) LocalAddr() string {
	if (conn.Conn == nil) || (conn.Conn.RemoteAddr() == nil) {
		return ""
	}

	return conn.Conn.LocalAddr().String()
}

func (conn *Connection) Network() string {
	if conn.Conn == nil {
		return ""
	}

	return conn.Conn.RemoteAddr().Network()
}
