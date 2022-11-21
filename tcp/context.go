package tcp

import (
	"github.com/panjf2000/gnet"
)

type IContext interface {
	SetConn(c interface{})
	OnTcpClose() error
	OnTcpRread(buf []byte) ([]byte, gnet.Action)
	RemoteAddr() string
	LocalAddr() string
	Network() string
	Write(data []byte) error
}

type Context struct {
	Conn gnet.Conn
}

func NewContext() IContext {
	return &Context{}
}

func (ctx *Context) Write(data []byte) error {
	return ctx.Conn.AsyncWrite(data)
}

func (ctx *Context) SetConn(c interface{}) {
	if c == nil {
		ctx.Conn = nil
	} else {
		ctx.Conn = c.(gnet.Conn)
	}
}

func (ctx *Context) OnTcpClose() error {
	return nil
}

func (ctx *Context) OnTcpRread(buf []byte) ([]byte, gnet.Action) {
	return nil, gnet.None
}

func (ctx *Context) RemoteAddr() string {
	if (ctx.Conn == nil) || (ctx.Conn.RemoteAddr() == nil) {
		return ""
	}

	return ctx.Conn.RemoteAddr().String()
}

func (ctx *Context) LocalAddr() string {
	if (ctx.Conn == nil) || (ctx.Conn.RemoteAddr() == nil) {
		return ""
	}

	return ctx.Conn.LocalAddr().String()
}

func (ctx *Context) Network() string {
	if ctx.Conn == nil {
		return ""
	}

	return ctx.Conn.RemoteAddr().Network()
}
