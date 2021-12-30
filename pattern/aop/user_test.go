package aop

import "testing"

func TestNewUser(t *testing.T) {
	user := &User{Account: "demo@gmail.com", Name: "demo", Pwd: "demo-mail"}
	proxy := NewUserProxy(user)
	proxy.Auth()
}
