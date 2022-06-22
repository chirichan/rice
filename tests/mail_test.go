package tests

import (
	"testing"

	"github.com/woxingliu/rice"
)

func TestMail_SendMail(t *testing.T) {

	m := rice.NewMail(rice.NewGmailAuth("", ""), "", rice.GmailSmtp, rice.GmailTLSPort)

	err := m.SendMail("test", "testbody", []string{""})
	if err != nil {
		t.Error(err)
	}
}
