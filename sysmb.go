// sysmb project sysmb.go
package sysmb

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

func NewServer() *Server {
	return &Server{
		subs: make(map[messageId]*subscribe),
		pubs: make(map[messageId]*publish),
	}
}

func PublishMsg(msgId uint16, pck []byte) error {
	if pck == nil {
		fmt.Println("Can't publish a empty message")
		return errors.New("Publish empty message")
	}

	msgheader := []byte{PUBLISH_MSG, 0}
	b_buff := bytes.NewBuffer([]byte{})
	binary.Write(b_buff, binary.BigEndian, msgId)
	msgheader = append(msgheader, b_buff.Bytes()...)
	fmt.Println("Publish messsage:", msgheader)

	msg := append(msgheader, pck...)
	addr, err := net.ResolveUDPAddr("udp", "localhost:15001")
	if nil != err {
		fmt.Println("PublishMsg(1)", err)
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if nil != err {
		fmt.Println("PublishMsg(2)", err)
		return err
	}
	defer conn.Close()

	_, err = conn.Write(msg)
	if nil != err {
		fmt.Println("PublishMsg(3)", err)
		return err
	}

	return nil
}

func SubscribeMsg(msgId uint16) (net.Conn, error) {
	msg := []byte{SUBLISH_MSG, 0}
	b_buff := bytes.NewBuffer([]byte{})
	binary.Write(b_buff, binary.BigEndian, msgId)
	msg = append(msg, b_buff.Bytes()...)
	fmt.Println("Subscribe messsage:", msg)

	addr, err := net.ResolveUDPAddr("udp", "localhost:15001")
	if nil != err {
		fmt.Println("SubscribeMsg(1)", err)
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if nil != err {
		fmt.Println("SubscribeMsg(2)", err)
		return nil, err
	}

	_, err = conn.Write(msg)
	if nil != err {
		fmt.Println("SubscribeMsg(3)", err)
		conn.Close()
		return nil, err
	}
	return conn, nil
}

func UnPublishMsg(msgId uint16) error {
	msg := []byte{UNPUBLISH_MSG, 0}
	b_buff := bytes.NewBuffer([]byte{})
	binary.Write(b_buff, binary.BigEndian, msgId)
	msg = append(msg, b_buff.Bytes()...)
	fmt.Println("UnPublish messsage:", msg)

	addr, err := net.ResolveUDPAddr("udp", "localhost:15001")
	if nil != err {
		fmt.Println("UnPublishMsg(1)", err)
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if nil != err {
		fmt.Println("UnPublishMsg(2)", err)
		return err
	}
	defer conn.Close()

	_, err = conn.Write(msg)
	if nil != err {
		fmt.Println("UnPublishMsg(3)", err)
		return err
	}

	return nil
}

func UnSubscribeMsg(msgId uint16) error {
	msg := []byte{UNSUBLISH_MSG, 0}
	b_buff := bytes.NewBuffer([]byte{})
	binary.Write(b_buff, binary.BigEndian, msgId)
	msg = append(msg, b_buff.Bytes()...)
	fmt.Println("UnSubscribe messsage:", msg)

	addr, err := net.ResolveUDPAddr("udp", "localhost:15001")
	if nil != err {
		fmt.Println("UnSubscribeMsg(1)", err)
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if nil != err {
		fmt.Println("UnSubscribeMsg(2)", err)
		return err
	}
	defer conn.Close()

	_, err = conn.Write(msg)
	if nil != err {
		fmt.Println("UnSubscribeMsg(3)", err)
		return err
	}

	return nil
}

func FlishMsg() error {
	msg := []byte{FLASH_MSG, 0, 0, 0}
	fmt.Println("FlishMsg messsage:", msg)

	addr, err := net.ResolveUDPAddr("udp", "localhost:15001")
	if nil != err {
		fmt.Println("FlishMsg(1)", err)
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if nil != err {
		fmt.Println("FlishMsg(2)", err)
		return err
	}
	defer conn.Close()

	_, err = conn.Write(msg)
	if nil != err {
		fmt.Println("FlishMsg(3)", err)
		return err
	}

	return nil
}
