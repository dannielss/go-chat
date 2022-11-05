package chat

import "net"

type client struct {
	conn     net.Addr
	nick     string
	room     *room
	commands chan<- command
}
