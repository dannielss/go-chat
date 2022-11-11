package main

import (
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick()
		case CMD_JOIN:
			s.join()
		case CMD_ROOMS:
			s.rooms()
		case CMD_MSG:
			s.msg()
		case CMD_QUIT:
			s.quit()
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}
