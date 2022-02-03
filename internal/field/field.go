package field

import (
	"fmt"
	"go-tic-tac-toe/internal/player"
)

type Field struct {
	Cols    int
	Rows    int
	Players []player.Player
}

func New(cols, rows int) (Field, error) {
	if cols > 1 && rows > 1 {
		return Field{
			Cols:    cols,
			Rows:    rows,
			Players: make([]player.Player, 0, 2),
		}, nil
	}
	return Field{}, fmt.Errorf("Incorrect cols=%d or rows=%d value", cols, rows)
}

func (f Field) String() string {
	return fmt.Sprintf("Field content %d columns, %d rows, Players: %v", f.Cols, f.Rows, f.Players)
}

func (f Field) Check(col, row int) error {
	// fmt.Printf("check col:%d, row:%d\n", col, row)
	cols_ok := false
	rows_ok := false
	if col >= 0 && col < f.Cols {
		cols_ok = true
	}
	if row >= 0 && row < f.Rows {
		rows_ok = true
	}
	if rows_ok && cols_ok {
		return nil
	}
	return fmt.Errorf("Incorrect cell: column check is %t, row check is %t\n", cols_ok, rows_ok)
}

func (f *Field) AddPlayer(player player.Player) error {
	// TODO validate if player already exist
	f.Players = append(f.Players, player)
	// TODO add sort by Priority
	return nil
}
