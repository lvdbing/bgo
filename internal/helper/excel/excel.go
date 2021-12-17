package excel

import (
	"github.com/lvdbing/bgo/pkg/convert"
)

// GetCell 获取第row行、第col列的单元格，row和col都从1开始。
func GetCell(row, col int) string {
	return GetCol(col) + convert.Itos(row)
}

// GetCol 获取第col列的字母标识，col从1开始。
func GetCol(col int) string {
	a := 'A'
	var buf []rune
	for col > 0 {
		col--
		b := a + rune(col%26)
		buf = append([]rune{b}, buf...)
		col = col / 26
	}

	return string(buf)
}
