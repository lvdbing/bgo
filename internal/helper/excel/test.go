package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func CreateCalendar() {
	f := excelize.NewFile()

	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)

	f.SetColWidth("Sheet2", "A", "DZ", 2.7)

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("日历.xlsx"); err != nil {
		fmt.Println(err)
	}
}
