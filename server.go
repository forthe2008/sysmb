package sysmb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type messageId uint16

type Server struct {
	conn net.PacketConn
	subs map[messageId]*subscribe
	pubs map[messageId]*publish
}

type publish struct {
	tag    byte
	puber  net.Addr
	pubMsg []byte
}

type subscribe struct {
	tag   byte
	suber []net.Addr
}

func (s *Server) Startup() {
	err := s.createUdpServer()
	if nil != err {
		fmt.Println("Create server error")
		return
	}
}

func (s *Server) Shutdown() {
	s.conn.Close()
	s.pubs = nil
	s.subs = nil
	s.conn = nil
}

func (s *Server) handle(buff []byte, n int, addr net.Addr) {
	if n < 4 {
		fmt.Println("Receive message error. n=", n)
		return
	}

	switch buff[0] {
	case PUBLISH_MSG:
		s.handlePublish(buff, n, addr)
	case SUBLISH_MSG:
		s.handleSubscrube(buff, n, addr)
	case UNPUBLISH_MSG:
		s.handleUnPublish(buff, n, addr)
	case UNSUBLISH_MSG:
		s.handleUnSubscrube(buff, n, addr)
	case FLASH_MSG:
		s.handleFlash()
	default:
		fmt.Println("Receive unknown message")
	}
}

func (s *Server) createUdpServer() error {
	addr, err := net.ResolveUDPAddr("udp", "localhost:15001")
	if nil != err {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if nil != err {
		return err
	}

	s.conn = conn

	go s.receive(conn)

	return nil

}

func (s *Server) receive(conn net.PacketConn) {
	buff := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buff)
		if nil != err {
			fmt.Println(err)
			return
		}

		go s.handle(buff, n, addr)
	}

}

func (s *Server) handlePublish(buff []byte, n int, addr net.Addr) {
	/*To do ? handle the tag value*/
	tag := buff[1]
	b_buff := bytes.NewBuffer(buff[2:4])
	var msgId messageId
	binary.Read(b_buff, binary.BigEndian, &msgId)
	buff = buff[4:n]

	if pub, ok := s.pubs[msgId]; ok {
		fmt.Println("Already exit, update it")
		pub.puber = addr
		pub.tag = tag
		pub.pubMsg = buff
	} else {
		fmt.Println("Add a new publisher")
		s.pubs[msgId] = &publish{
			puber:  addr,
			tag:    tag,
			pubMsg: buff,
		}
	}

	/*go throug the subscriber*/
	if sub, ok := s.subs[msgId]; ok {
		fmt.Println("Exit subscriber, start to forward message")
		for _, subAddr := range sub.suber {
			fmt.Println("Send message to ", subAddr)
			s.conn.WriteTo(s.pubs[msgId].pubMsg, subAddr)
		}
	}

}

func (s *Server) handleSubscrube(buff []byte, n int, addr net.Addr) {
	if n > 4 {
		fmt.Println("Subscribe with payload")
	}
	/*To do ? handle the tag value*/
	tag := buff[1]
	isNew := true
	var msgId messageId
	b_buff := bytes.NewBuffer(buff[2:4])
	binary.Read(b_buff, binary.BigEndian, &msgId)

	if sub, ok := s.subs[msgId]; ok {
		fmt.Println("Already exit, add or update it")
		for _, subAddr := range sub.suber {
			if subAddr == addr {
				isNew = false
				break
			}
		}
		if isNew {
			sub.suber = append(sub.suber, addr)
		}

	} else {
		fmt.Println("Add a new subscriber")
		s.subs[msgId] = &subscribe{
			suber: []net.Addr{addr},
			tag:   tag,
		}
	}

	/*go throug the publisher*/
	if _, ok := s.pubs[msgId]; ok && isNew {
		fmt.Println("Exit publish, start to forward message")
		fmt.Println("Send message to ", addr)
		s.conn.WriteTo(s.pubs[msgId].pubMsg, addr)

	}
}

func (s *Server) handleUnPublish(buff []byte, n int, addr net.Addr) {
	/*To do ? handle the tag value*/
	b_buff := bytes.NewBuffer(buff[2:4])
	var msgId messageId
	binary.Read(b_buff, binary.BigEndian, &msgId)

	if _, ok := s.pubs[msgId]; ok {
		fmt.Println("Already exit, delete it")
		delete(s.pubs, msgId)
	}
}

func (s *Server) handleUnSubscrube(buff []byte, n int, addr net.Addr) {
	/*To do ? handle the tag value*/
	b_buff := bytes.NewBuffer(buff[2:4])
	var msgId messageId
	binary.Read(b_buff, binary.BigEndian, &msgId)

	if _, ok := s.subs[msgId]; ok {
		fmt.Println("Already exit, delete it")
		delete(s.subs, msgId)
	}
}

func (s *Server) handleFlash() {
	s.subs = make(map[messageId]*subscribe)
	s.pubs = make(map[messageId]*publish)
}
