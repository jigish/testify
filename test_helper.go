package common

import (
  "testing"
  "fmt"
)

type TestHelper struct {
  t *testing.T
}

func (self *TestHelper) SetT(newT *testing.T) {
  self.t = newT
}

func (self *TestHelper) Refute(b bool, msg string) {
  if b {
    self.t.Error(msg)
  }
}

func (self *TestHelper) Assert(b bool, msg string) {
  self.Refute(!b, msg)
}

func (self *TestHelper) AssertNotNil(obj interface{}, msg string) {
  self.Refute(obj == nil, msg)
}

func (self *TestHelper) AssertNil(obj interface{}, msg string) {
  self.Refute(obj != nil, msg)
}

func (self *TestHelper) RefuteError(err error, msg string) {
  if err != nil {
    self.t.Error(msg+"\n"+err.Error())
  }
}

func (self *TestHelper) AssertEqual(s1, s2 interface{}, msg string) {
  self.Refute(s1 != s2, fmt.Sprintf(msg+"\n%v != %v", s1, s2))
}
