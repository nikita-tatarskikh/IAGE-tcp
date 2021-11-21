package server

import (
	"bufio"
	"log"
	"net"
)

type TCPServer struct {
	Addr string
}

func (srv *TCPServer) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":8081"
	}
	log.Printf("starting server on %v\n", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection %v", err)
			continue
		}
		log.Printf("accepted connection from %v", conn.RemoteAddr())
		err = handle(conn)
		if err != nil {
			log.Println()
		}
	}
}

func handle(conn net.Conn) error {
	defer func() {
		log.Printf("closing connection from %v", conn.RemoteAddr())
		err := conn.Close()
		if err != nil {
			log.Println("Error while closing connection", err)
		}
	}()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	scanr := bufio.NewScanner(r)

	for {
		scanned := scanr.Scan()
		if !scanned {
			if err := scanr.Err(); err != nil {
				log.Printf("%v(%v)", err, conn.RemoteAddr())
				return err
			}
			break
		}
		w.WriteString(scanr.Text() + "\r\n")
		w.Flush()
	}
	return nil
}


