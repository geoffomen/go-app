package excelUtil

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestExcelExport(t *testing.T) {
	tests := []struct {
		A1       int    `exnm:"字段1" en:"field1"`
		B字段名称不重要 bool   `exnm:"字段2" en:"field2"`
		C3       string `exnm:"字段3" en:"field3"`
	}{
		{2, true, "字符串001"},
		{3, false, "字符串002"},
		{5, true, "字符串003"},
		{7, true, "字符串004"},
		{11, false, "字符串005"},
		{13, true, "字符串006"},
	}
	bys, err := WriteToBuff(tests)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(bys)
	}

	err = WriteToFile(tests, "/tmp/a.xlsx")
	if err != nil {
		fmt.Println(err.Error())
	}

	bys, err = WriteToBuffWithSelectedColumns([]string{"A1", "C3"}, tests)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(bys)
	}
	err = WriteToFileWithSelectedColumns([]string{"A1", "C3"}, tests, "/tmp/b.xlsx")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestRead(t *testing.T) {
	f, _ := os.OpenFile("/tmp/导入模板.xlsm", os.O_RDONLY, os.ModePerm)
	buf, _ := io.ReadAll(f)
	data, _ := FetchAllFromBuf(buf, 1)
	fmt.Println(data)
}
