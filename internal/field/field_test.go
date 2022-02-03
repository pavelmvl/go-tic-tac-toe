package field

import (
	"go-tic-tac-toe/internal/player"
	"testing"
)

func TestNew(t *testing.T) {
	var e error
	_, e = New(1, 1)
	if e == nil {
		t.Fatal("field.New(1,1) should be fail")
	}
	_, e = New(-1, -1)
	if e == nil {
		t.Fatal("field.New(0,0) should be fail")
	}
	_, e = New(-500, 1)
	if e == nil {
		t.Fatal("field.New(-500,1) should be fail")
	}
	_, e = New(0, 0)
	if e == nil {
		t.Fatal("field.New(0,0) should be fail")
	}
	_, e = New(2, 2)
	if e != nil {
		t.Fatal("field.New(2,2) should be OK, error:", e)
	}
}

func TestCheckDefault(t *testing.T) {
	f, e := New(3, 3)
	if e != nil {
		t.Fatal("field.New(3,3) return error: ", e)
	}
	e = f.Check(-1, -1)
	if e == nil {
		t.Fatal("field.Check(-1, -1) should be fail")
	}
	e = f.Check(-1, 0)
	if e == nil {
		t.Fatal("field.Check(-1, 0) should be fail")
	}
	e = f.Check(0, -1)
	if e == nil {
		t.Fatal("field.Check(0, -1) should be fail")
	}
	e = f.Check(0, 0)
	if e != nil {
		t.Fatal("field.Check(0, 0) return error: ", e)
	}
	e = f.Check(1, 1)
	if e != nil {
		t.Fatal("field.Check(1, 1) return error: ", e)
	}
	e = f.Check(2, 2)
	if e != nil {
		t.Fatal("field.Check(2, 2) return error: ", e)
	}
	e = f.Check(3, 2)
	if e == nil {
		t.Fatal("field.Check(3, 2) should be fail")
	}
	e = f.Check(2, 3)
	if e == nil {
		t.Fatal("field.Check(2, 3) should be fail")
	}
	e = f.Check(3, 3)
	if e == nil {
		t.Fatal("field.Check(3, 3) should be fail")
	}
}

func TestAddPlayer(t *testing.T) {
	f, e := New(3, 3)
	if e != nil {
		t.Fatal("field.New(3,3) return error: ", e)
	}
	p := player.New("x")
	e = f.AddPlayer(p)
	if e != nil {
		t.Fatal("field.AddPlayer(p) return error: ", e)
	}
	if p != f.Players[0] {
		t.Fatal("player struct missmatch")
	}
}
