package field

import (
	"fmt"
)

type Field struct {
	size  int
	cells [][]rune
}

func New(size int) (Field, error) {
	f := Field{}
	if size < 3 {
		return f, fmt.Errorf("Incorrect 'size' value (%d). 'size' should be greater than '2'", size)
	}
	f.size = size
	f.cells = make([][]rune, size)
	for i := range f.cells {
		f.cells[i] = make([]rune, size)
	}
	return f, nil
}

func (f Field) IsCellValid(col, row int) error {
	if col < 0 || col >= f.size {
		return fmt.Errorf("Incorrect cell: column check failed")
	}
	if row < 0 || row >= f.size {
		return fmt.Errorf("Incorrect cell: row check failed")
	}
	return nil
}

func (f Field) IsCellFree(col, row int) error {
	if f.cells[col][row] != rune(0) {
		return fmt.Errorf("Cell is not free (%c)", f.cells[col][row])
	}
	return nil
}

func (f *Field) AssignCell(col, row int, mark rune) error {
	f.cells[col][row] = mark
	return nil
}

func (f Field) GetCellValue(col, row int) rune {
	return f.cells[col][row]
}

func (f Field) IsCellWinner(col, row int) rune {
	var win rune
	win = f.IsColumnWinner(col, row)
	if win != NoWinner {
		return win
	}
	win = f.IsRowWinner(col, row)
	if win != NoWinner {
		return win
	}
	win = f.IsDiagStreightWinner(col, row)
	if win != NoWinner {
		return win
	}
	win = f.IsDiagReverseWinner(col, row)
	if win != NoWinner {
		return win
	}
	return NoWinner
}

func (f Field) IsFieldFull() bool {
	for col := 0; col < f.size; col++ {
		for row := 0; row < f.size; row++ {
			if f.GetCellValue(col, row) == NoWinner {
				return false
			}
		}
	}
	return true
}

func (f Field) ToString() string {
	buf := make([]byte, 0, f.size*(2*f.size+1)+1)
	for row := 0; row < f.size; row++ {
		buf = append(buf, []byte(" ")...)
		for col := 0; col < f.size; col++ {
			mark := f.GetCellValue(col, row)
			if mark == NoWinner {
				mark = ' '
			}
			buf = append(buf, []byte(string(mark))...)
			buf = append(buf, []byte(" ")...)
		}
		buf = append(buf, []byte("\n")...)
	}
	return string(buf)
}

func (f Field) ToHtml() string {
	html := make([]byte, 2048)
	html = append(html, []byte(`
<html>
<head>
	<title>Tic-tac-toe</title>
	<style>
.field { display:table; outline:2px solid black; border-collapse:collapse; }
.row { display:table-row; }
.cell { display:table-cell; outline:1px solid black; border-collapse:collapse; margin:0; padding:0; width:100px; height:100px; font-size:50px; text-align: center; vertical-align:middle; }
	</style>
</head>
<body>
	<div class="field">
`)...)
	for row := 0; row < f.size; row++ {
		html = append(html, []byte("<div class=\"row\">")...)
		for col := 0; col < f.size; col++ {
			mark := f.GetCellValue(col, row)
			if mark == NoWinner {
				mark = ' '
			}
			html = append(html, []byte(fmt.Sprintf("<div class=\"cell\">%c</div>", mark))...)
		}
		html = append(html, []byte("</div>")...)
	}
	html = append(html, []byte(`
	</div>
</body>
</html>`)...)
	return string(html)
}

func (f Field) Print() {
	print(f.ToString())
}
