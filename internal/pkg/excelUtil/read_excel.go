package excelUtil

import (
	"fmt"

	xlsx "github.com/tealeg/xlsx/v3"
)

func FetchAll(path, sheetName string) ([][]string, error) {
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	sh, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Print(err)
		return nil, err
	}
	data := make([][]string, 0, sh.MaxRow)
	sh.ForEachRow(func(r *xlsx.Row) error {
		cell := make([]string, 0, sh.MaxCol)
		r.ForEachCell(func(c *xlsx.Cell) error {
			cell = append(cell, c.Value)
			return nil
		})
		data = append(data, cell)
		return nil
	})

	// for _, row := range data {
	// 	fmt.Print(row)
	// 	fmt.Printf("n: %s, v: %s\n", strings.TrimSpace(row[0]), strings.TrimSpace(row[1]))
	// }
	return data, nil
}

func FetchAllFromBuf(buf []byte, sheetIdx int) ([][]string, error) {
	wb, err := xlsx.OpenBinary(buf)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	sh := wb.Sheets[sheetIdx]
	data := make([][]string, 0, sh.MaxRow)
	for i := 0; i < sh.MaxRow; i++ {
		row, err := sh.Row(i)
		if err != nil {
			continue
		}
		cell := make([]string, 0, sh.MaxCol)
		for j := 0; j < sh.MaxCol; j++ {
			col := row.GetCell(j)
			cell = append(cell, col.String())
		}
		data = append(data, cell)
	}

	return data, nil
}
