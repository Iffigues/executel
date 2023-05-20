package tcp

import (
	"executel-server/src/conf"
	"net"
)

type Tcp struct {
	s net.Conn
}

func NewTcp(c *conf.Conf) (t *Tcp, err error) {
	println("localhost:8889")
	con, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		return nil, err
	}
	return &Tcp{con}, nil
}

func (t *Tcp) Close() {
	t.s.Close()
}
