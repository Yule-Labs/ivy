package server

import (
	"bufio"
	"log"
	"net"
)

type Server interface {
	ListenAndServe() error
}

type defaultServer struct {
	port string
	network string
}

func (s *defaultServer) ListenAndServe() error {
	listener, err := net.Listen(s.network, s.port)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go func() {
			err := s.handleConn(conn)
			if err != nil {
				log.Println(err)
			}
		}()
	}
}

func (s *defaultServer) handleConn(conn net.Conn) (err error) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	message, _ := bufio.NewReader(conn).ReadString('\n')
	response := s.handleMessage(message)

	_, err = conn.Write([]byte(response))
	return
}

func (s *defaultServer) handleMessage(msg string) string {
	return "+OK\r\n"
}

func newServer() Server {
	return &defaultServer{
		port:    ":9000",
		network: "tcp",
	}
}

func InitServer() {
	err := newServer().ListenAndServe()
	if err != nil {
		panic(err)
	}
}
