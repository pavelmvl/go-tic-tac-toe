package field

import (
	"go-tic-tac-toe/internal/common"
	"testing"
)

func TestNew(t *testing.T) {
	var e error
	_, e = New(1, 0)
	if e == nil {
		t.Fatal("field.New(1) should be fail")
	}
	_, e = New(-1, 0)
	if e == nil {
		t.Fatal("field.New(-1) should be fail")
	}
	_, e = New(0, 0)
	if e == nil {
		t.Fatal("field.New(0) should be fail")
	}
	_, e = New(3, 3)
	if e != nil {
		t.Fatal("field.New(3) should be OK, error:", e)
	}
	_, e = New(500, 250)
	if e != nil {
		t.Fatal("field.New(500) should be OK, error:", e)
	}
}

type cellData struct {
	col  int
	row  int
	mark rune
	err  error
}

var cellDataTest = [...]cellData{
	cellData{-1, -1, 'X', common.ErrCellColumn},
	cellData{-1, 0, 'X', common.ErrCellColumn},
	cellData{0, -1, 'X', common.ErrCellRow},
	cellData{0, 0, 'X', nil},
	cellData{0, 0, 'O', common.ErrCellBusy},
	cellData{1, 1, 0, nil},
	cellData{99, 99, 'X', common.ErrCellColumn},
}

func TestAssignCell(t *testing.T) {
	f, e := New(3, 3)
	if e != nil {
		t.Fatal("field.New(3) return error: ", e)
	}
	for _, v := range cellDataTest {
		if err := f.AssignCell(v.col, v.row, v.mark); err != v.err {
			t.Fatal("Unexpected result: ", err, "\n",
				"Should be: ", v.err, "\n",
				"Test data: ", v)
		}
	}
}
