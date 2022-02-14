package field

import (
	"testing"
)

func TestNew(t *testing.T) {
	var e error
	_, e = New(1)
	if e == nil {
		t.Fatal("field.New(1) should be fail")
	}
	_, e = New(-1)
	if e == nil {
		t.Fatal("field.New(-1) should be fail")
	}
	_, e = New(0)
	if e == nil {
		t.Fatal("field.New(0) should be fail")
	}
	_, e = New(3)
	if e != nil {
		t.Fatal("field.New(3) should be OK, error:", e)
	}
	_, e = New(500)
	if e != nil {
		t.Fatal("field.New(500) should be OK, error:", e)
	}
}

func TestIsCellValidDefault(t *testing.T) {
	f, e := New(3)
	if e != nil {
		t.Fatal("field.New(3) return error: ", e)
	}
	e = f.IsCellValid(-1, -1)
	if e == nil {
		t.Fatal("field.IsCellValid(-1, -1) should be fail")
	}
	e = f.IsCellValid(-1, 0)
	if e == nil {
		t.Fatal("field.IsCellValid(-1, 0) should be fail")
	}
	e = f.IsCellValid(0, -1)
	if e == nil {
		t.Fatal("field.IsCellValid(0, -1) should be fail")
	}
	e = f.IsCellValid(0, 0)
	if e != nil {
		t.Fatal("field.IsCellValid(0, 0) return error: ", e)
	}
	e = f.IsCellValid(1, 1)
	if e != nil {
		t.Fatal("field.IsCellValid(1, 1) return error: ", e)
	}
	e = f.IsCellValid(2, 2)
	if e != nil {
		t.Fatal("field.IsCellValid(2, 2) return error: ", e)
	}
	e = f.IsCellValid(3, 2)
	if e == nil {
		t.Fatal("field.IsCellValid(3, 2) should be fail")
	}
	e = f.IsCellValid(2, 3)
	if e == nil {
		t.Fatal("field.IsCellValid(2, 3) should be fail")
	}
	e = f.IsCellValid(3, 3)
	if e == nil {
		t.Fatal("field.IsCellValid(3, 3) should be fail")
	}
}

func TestIsCellFree(t *testing.T) {
	f, _ := New(3)
	f.AssignCell(0, 0, 'X')
	if f.IsCellFree(0, 0) == nil {
		t.Fatal("\n", f.ToString(), "f.IsCellFree(0,0) should be full")
	}
	if f.IsCellFree(1, 1) != nil {
		t.Fatal("\n", f.ToString(), "f.IsCellFree(1,1) should be empty")
	}
}
