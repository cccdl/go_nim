package tests

import (
	"github.com/cccdl/go_nim"
	"testing"
)

func TestSendCustomMessage(t *testing.T) {
	msg := &nim.Message{}
	c := nim.CreateImClient("", "")
	opt := &nim.ImSendMessageOption{
		MsgSenderNoSense:   0,
		MsgReceiverNoSense: 1,
	}

	err := c.SendCustomMessage("test1", "test2", msg, opt)
	if err != nil {
		t.Error(err)
	}
}
