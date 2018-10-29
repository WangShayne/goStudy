package main

import "testing"

func TestAdd(t *testing.T) {
	a := 10
	b := 20
	c := Add(a, b)
	t.Logf("a=%d b=%d\n ", a, b)
	if c != 30 {
		t.Fatalf("invalied a+b ,c=%d\n", c)
	}
	t.Logf("a=%d,b=%d,c=%d\n", a, b, c)
}

func TestSub(t *testing.T) {
	a := 10
	b := 20
	c := Sub(a, b)
	t.Logf("a=%d,b=%d", a, b)
	t.Logf("a=%d,b=%d,c=%d", a, b, c)

}
