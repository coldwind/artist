package isms

import "testing"

func TestQcloudSend(t *testing.T) {
	s := New(QCloud, &Config{})
	res, err := s.p.Send("+8613816207221", []string{"6666"})
	t.Log(res, err)
}
