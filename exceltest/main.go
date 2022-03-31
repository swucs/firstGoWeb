package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	file := excelize.NewFile()
	sheet := file.NewSheet("Sheet1")
	sumSheet := file.NewSheet("sum-sheet")

	file.SetActiveSheet(sheet)
	file.SetActiveSheet(sumSheet)

	file.SetCellValue("Sheet1", "A1", "VALUE")

	file.SetCellValue("sum-sheet", "A1", "SUM")
	file.SetCellValue("sum-sheet", "A2", 100)
	file.SetCellValue("sum-sheet", "B2", 250)
	file.SetCellFormula("sum-sheet", "B1", "SUM('sum-sheet'!A2,'sum-sheet'!B2)")

	err := file.SaveAs("TEST_1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

}
