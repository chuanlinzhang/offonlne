package main

import (
	"net"
	"bufio"
)

type Packet struct {

}
type Channel struct {
	conn net.Conn //websocket连接
	send chan Packet
}

func NewChannel(conn net.Conn) *Channel  {
	c:=&Channel{
		conn:conn,
		send:make(chan Packet,N)
	}

}
func (c *Channel) reader()  {
	buf:=bufio.NewReader(c.conn)

}