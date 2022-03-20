package field

import (
	"fmt"
)

type Field struct {
	size  int
	cells [][]rune
}

var (
	ErrFieldSize  = fmt.Errorf("Incorrect 'size' value. 'size' should be greater than '2'")
	ErrCellColumn = fmt.Errorf("Incorrect cell: column check failed")
	ErrCellRow    = fmt.Errorf("Incorrect cell: row check failed")
	ErrCellBusy   = fmt.Errorf("Cell is not free")
)

func New(size int) (Field, error) {
	f := Field{}
	if size < 3 {
		return f, ErrFieldSize
	}
	f.size = size
	f.cells = make([][]rune, size)
	for i := range f.cells {
		f.cells[i] = make([]rune, size)
	}
	return f, nil
}

func (f Field) isCellValid(col, row int) error {
	if col < 0 || col >= f.size {
		return ErrCellColumn
	}
	if row < 0 || row >= f.size {
		return ErrCellRow
	}
	return nil
}

func (f Field) isCellFree(col, row int) error {
	if f.cells[col][row] != rune(0) {
		return ErrCellBusy
	}
	return nil
}

func (f Field) AssignCell(col, row int, mark rune) error {
	var err error
	err = f.isCellValid(col, row)
	if err != nil {
		return err
	}
	err = f.isCellFree(col, row)
	if err != nil {
		return err
	}
	f.cells[col][row] = mark
	return nil
}

func (f Field) GetCellValue(col, row int) (rune, error) {
	var err error
	err = f.isCellValid(col, row)
	if err != nil {
		return 0, err
	}
	return f.cells[col][row], nil
}

func (f Field) IsCellWinner(col, row int) rune {
	return f.isCellWinner(col, row)
}

func (f Field) isCellWinner(col, row int) rune {
	var win rune
	win = f.isColumnWinner(col, row)
	if win != NoWinner {
		return win
	}
	win = f.isRowWinner(col, row)
	if win != NoWinner {
		return win
	}
	win = f.isDiagStreightWinner(col, row)
	if win != NoWinner {
		return win
	}
	win = f.isDiagReverseWinner(col, row)
	if win != NoWinner {
		return win
	}
	return NoWinner
}

func (f Field) IsFieldFull() bool {
	for col := 0; col < f.size; col++ {
		for row := 0; row < f.size; row++ {
			if winner, _ := f.GetCellValue(col, row); winner == NoWinner {
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
			mark, _ := f.GetCellValue(col, row)
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

var htmlTemplate string = `<html>
<head>
	<title>Tic-tac-toe</title>
	<style>
.field { display:table; outline:2px solid black; border-collapse:collapse; }
.row { display:table-row; }
.cell { display:table-cell; outline:1px solid black; border-collapse:collapse; margin:0; padding:0; width:%dpx; height:%dpx; font-size:50px; text-align: center; vertical-align:middle; }
	</style>
</head>
<body>
	<div class="field">%s</div>
	<div class="extra">%s</div>
</body>
</html>`

func (f Field) ToHtml(extra ...string) string {
	extraByte := make([]byte, 0, 512)
	for _, v := range extra {
		extraByte = append(extraByte, []byte(v)...)
		extraByte = append(extraByte, []byte("<br/>")...)
	}
	table := make([]byte, 0, 512)
	for row := 0; row < f.size; row++ {
		table = append(table, []byte("<div class=\"row\">")...)
		for col := 0; col < f.size; col++ {
			mark, _ := f.GetCellValue(col, row)
			if mark == NoWinner {
				mark = ' '
			}
			table = append(table, []byte(fmt.Sprintf("<div class=\"cell\" onclick=\"location.href='/%d/%d';\">%c</div>", col, row, mark))...)
		}
		table = append(table, []byte("</div>")...)
	}
	// TODO width = 100, height = 100. Make it parametrized
	return fmt.Sprintf(htmlTemplate, 100, 100, string(table), string(extraByte))
}

func (f Field) Print() {
	print(f.ToString())
}
