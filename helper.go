package testify

import (
  "testing"
  "fmt"
)

type Helper struct {
  t *testing.T
}

func (self *Helper) SetT(newT *testing.T) {
  self.t = newT
}

func (self *Helper) Refute(b bool, msg string) {
  if b {
    self.t.Error(msg)
  }
}

func (self *Helper) Assert(b bool, msg string) {
  self.Refute(!b, msg)
}

func (self *Helper) AssertNotNil(obj interface{}, msg string) {
  self.Refute(obj == nil, msg)
}

func (self *Helper) AssertNil(obj interface{}, msg string) {
  self.Refute(obj != nil, msg)
}

func (self *Helper) RefuteError(err error, msg string) {
  if err != nil {
    self.t.Error(msg+"\n"+err.Error())
  }
}

func (self *Helper) AssertEqual(s1, s2 interface{}, msg string) {
  self.Refute(s1 != s2, fmt.Sprintf(msg+"\n%v != %v", s1, s2))
}
