package testify

import (
  "testing"
  "fmt"
)

type Testify struct {
  t *testing.T
}

func Create(newT *testing.T) *Testify {
  return &Testify{newT}
}

func (self *Testify) SetT(newT *testing.T) {
  self.t = newT
}

func (self *Testify) Refute(b bool, msg string) {
  if b {
    self.t.Error(msg)
  }
}

func (self *Testify) Assert(b bool, msg string) {
  self.Refute(!b, msg)
}

func (self *Testify) AssertNotNil(obj interface{}, msg string) {
  self.Refute(obj == nil, msg)
}

func (self *Testify) AssertNil(obj interface{}, msg string) {
  self.Refute(obj != nil, msg)
}

func (self *Testify) RefuteError(err error, msg string) {
  if err != nil {
    self.t.Error(msg+"\n"+err.Error())
  }
}

func (self *Testify) AssertEqual(s1, s2 interface{}, msg string) {
  self.Refute(s1 != s2, fmt.Sprintf(msg+"\n%v != %v", s1, s2))
}
