package config

import "testing"

func TestNewGetContract(t *testing.T) {
	m := map[string]string{"foo": "bar"}
	c := New(nil, m)

	ret, exist := c.Get("foo")

	if !(exist && ret == "bar") {
		t.Fail()
	}
}

func TestBackup(t *testing.T) {
	m1 := map[string]string{"foo": "bar"}
	c1 := New(nil, m1)

	m2 := map[string]string{}
	c2 := New(c1, m2)

	ret, exist := c2.Get("foo")

	if !(exist && ret == "bar") {
		t.Fail()
	}
}

func TestShadowBackup(t *testing.T) {
	m1 := map[string]string{"foo": "bar"}
	c1 := New(nil, m1)

	m2 := map[string]string{"foo": "baz"}
	c2 := New(c1, m2)

	ret, exist := c2.Get("foo")

	if !(exist && ret == "baz") {
		t.Fail()
	}
}
