package main

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"math/rand"
)

func main() {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		println(err.Error())
	}
	styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	if err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("A1", []interface{}{excelize.Cell{StyleID: styleID, Value: "Data"}}); err != nil {
		println(err.Error())
	}
	for rowID := 2; rowID <= 102400; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			println(err.Error())
		}
	}
	if err := streamWriter.Flush(); err != nil {
		println(err.Error())
	}
	if err := file.SaveAs("Book1.xlsx"); err != nil {
		println(err.Error())
	}
}
