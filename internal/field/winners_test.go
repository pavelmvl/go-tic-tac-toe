package field

import (
	"testing"
)

func TestIsColumnWinner(t *testing.T) {
	f, err := New(4)
	if err != nil {
		t.Fatal(err)
	}
	f.AssignCell(0, 0, 'X')
	f.AssignCell(0, 1, 'X')
	f.AssignCell(0, 2, 'X')
	f.AssignCell(0, 3, 'X')
	if f.isColumnWinner(0, 1) == NoWinner {
		t.Fatal(f.ToString(), "IsColumnWinner(0,1) should return winner")
	}
	f.AssignCell(1, 0, 'X')
	f.AssignCell(1, 1, 'X')
	if f.isColumnWinner(1, 0) != NoWinner {
		t.Fatal("\n", f.ToString(), "IsColumnWinner(1,0) should return no winner")
	}
	if f.isColumnWinner(1, 3) != NoWinner {
		t.Fatal("\n", f.ToString(), "IsColumnWinner(1,3) should return no winner")
	}
}

func TestIsRowWinner(t *testing.T) {
	f, err := New(4)
	if err != nil {
		t.Fatal(err)
	}
	f.AssignCell(0, 1, 'X')
	f.AssignCell(1, 1, 'X')
	f.AssignCell(2, 1, 'X')
	f.AssignCell(3, 1, 'X')
	if f.isRowWinner(1, 1) == NoWinner {
		t.Fatal(f.ToString(), "IsColumnWinner(1,1) should return winner")
	}
	f.AssignCell(0, 2, 'X')
	f.AssignCell(1, 2, 'X')
	if f.isRowWinner(1, 2) != NoWinner {
		t.Fatal("\n", f.ToString(), "IsColumnWinner(1,2) should return no winner")
	}
	if f.isRowWinner(2, 2) != NoWinner {
		t.Fatal("\n", f.ToString(), "IsColumnWinner(2,2) should return no winner")
	}
}

func TestIsDiagStreightWinner(t *testing.T) {
	f, err := New(5)
	if err != nil {
		t.Fatal(err)
	}
	f.AssignCell(0, 0, 'X')
	f.AssignCell(1, 1, 'X')
	f.AssignCell(2, 2, 'X')
	f.AssignCell(3, 3, 'X')
	f.AssignCell(4, 4, 'X')
	if f.isDiagStreightWinner(2, 2) == NoWinner {
		t.Fatal("\n", f.ToString(), "f.IsDiagStreightWinner(2,2) should return winner")
	}
	if f.isDiagReverseWinner(2, 2) != NoWinner {
		t.Fatal("\n", f.ToString(), "f.IsDiagReverseWinner(2,2) should return no winner")
	}
	if f.isDiagReverseWinner(3, 3) != NoWinner {
		t.Fatal("\n", f.ToString(), "f.IsDiagReverseWinner(3,3) should return no winner")
	}
}

func TestIsDiagReverseWinner(t *testing.T) {
	f, err := New(5)
	if err != nil {
		t.Fatal(err)
	}
	f.AssignCell(4, 0, 'X')
	f.AssignCell(3, 1, 'X')
	f.AssignCell(2, 2, 'X')
	f.AssignCell(1, 3, 'X')
	f.AssignCell(0, 4, 'X')
	if f.isDiagStreightWinner(2, 2) != NoWinner {
		t.Fatal("\n", f.ToString(), "f.IsDiagStreightWinner(2,2) should return no winner")
	}
	if f.isDiagStreightWinner(1, 3) != NoWinner {
		t.Fatal("\n", f.ToString(), "f.IsDiagStreightWinner(1,3) should return no winner")
	}
	if f.isDiagReverseWinner(2, 2) == NoWinner {
		t.Fatal("\n", f.ToString(), "f.IsDiagReverseWinner(2,2) should return winner")
	}
}
