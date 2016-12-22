package sysmb

import (
	"os/exec"
	"testing"
	"time"
)

var (
	TEST_MSG_ID = uint16(231)
	TEST_MSG    = []byte{6, 7, 23, 1}
)

func Test_Basic(t *testing.T) {
	cmd := exec.Command("go", "run", "./sysmb-server/server-run.go")
	go cmd.Run()

	/*Publish message*/
	err := PublishMsg(TEST_MSG_ID, TEST_MSG)
	if nil != err {
		t.Error("Publish message error")
	}

	/*Subscribe message*/
	conn, err := SubscribeMsg(TEST_MSG_ID)
	if nil != err {
		t.Error("Subscribe message error")
		return
	}
	buff := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	n, err := conn.Read(buff)
	if nil != err {
		t.Error("Read message error")
		return
	}

	if n != len(TEST_MSG) {
		t.Error("Len is not equal")
		return
	}

	for i, v := range TEST_MSG {
		if v != buff[i] {
			t.Error("Value is not equal")
		}
	}
}
