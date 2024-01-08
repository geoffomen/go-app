package excelUtil

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"time"

	xlsx "github.com/tealeg/xlsx/v3"
)

// WriteToBuff 将一组结构体生成EXCEL文件，返回EXCEL文件的字节流
func WriteToBuff(dtoArray interface{}) ([]byte, error) {
	fVal := reflect.ValueOf(dtoArray)
	if fVal.Kind() != reflect.Slice || fVal.Len() == 0 {
		return nil, fmt.Errorf("参数错误，参数非切片类型或者切片长度为0")
	}
	f, err := prepareFile(fVal)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	err = f.Write(writer)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// WriteToFile 将一组结构体生成EXCEL文件
func WriteToFile(dtoArray interface{}, filePath string) error {
	fVal := reflect.ValueOf(dtoArray)
	if fVal.Kind() != reflect.Slice || fVal.Len() == 0 {
		return fmt.Errorf("参数错误，参数非切片类型或者切片长度为0")
	}
	f, err := prepareFile(fVal)
	if err != nil {
		return err
	}
	err = f.Save(filePath)
	if err != nil {
		return err
	}
	return nil
}

// WriteToBuffWithSelectedColumns 将一组结构体生成EXCEL文件，返回EXCEL文件的字节流
func WriteToBuffWithSelectedColumns(fieldNames []string, dtoArray interface{}) ([]byte, error) {
	fVal := reflect.ValueOf(dtoArray)
	if fVal.Kind() != reflect.Slice || fVal.Len() == 0 || fVal.Len() < len(fieldNames) {
		return nil, fmt.Errorf("参数错误，参数非切片类型或者切片长度为0, 或者选定导出的字段多于实际存在的字段")
	}
	f, err := prepareFileWithSelectedColumns(fieldNames, fVal)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	err = f.Write(writer)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// WriteToFileWithSelectedColumns 将一组结构体生成EXCEL文件
func WriteToFileWithSelectedColumns(fieldNames []string, dtoArray interface{}, filePath string) error {
	fVal := reflect.ValueOf(dtoArray)
	if fVal.Kind() != reflect.Slice || fVal.Len() == 0 || fVal.Len() < len(fieldNames) {
		return fmt.Errorf("参数错误，参数非切片类型或者切片长度为0, 或者选定导出的字段多于实际存在的字段")
	}
	f, err := prepareFileWithSelectedColumns(fieldNames, fVal)
	if err != nil {
		return err
	}
	err = f.Save(filePath)
	if err != nil {
		return err
	}
	return nil
}

func prepareFile(fVal reflect.Value) (*xlsx.File, error) {
	rows := make([]reflect.Value, fVal.Len())
	for i := 0; i < fVal.Len(); i++ {
		rows[i] = fVal.Index(i)
	}
	wb := xlsx.NewFile()
	sh, err := wb.AddSheet("Sheet")
	if err != nil {
		return nil, err
	}

	fns, err := writeExcelHeaders(rows[0], sh)
	if err != nil {
		return nil, err
	}
	for _, fv := range rows {
		row := sh.AddRow()
		for _, header := range fns {
			field := fv.FieldByName(header)
			cell := row.AddCell()
			if field.Type().String() == "time.Time" {
				t := field.Interface().(time.Time)
				if t.Year() < 2 {
					cell.SetValue("")
				}
				str := t.Format("2006-01-02 15:04:05")
				cell.SetValue(str)
			} else {
				cell.SetValue(field.Interface())
			}
		}
	}

	return wb, nil
}

func writeExcelHeaders(aDtoFVal reflect.Value, sh *xlsx.Sheet) ([]string, error) {
	fType := aDtoFVal.Type()
	numFields := fType.NumField()

	fns := make([]string, 0, numFields)

	myStyle := xlsx.NewStyle()
	myStyle.Alignment.Horizontal = "right"
	myStyle.Fill.FgColor = "FFFFFFFF"
	myStyle.Fill.BgColor = "00000000"
	myStyle.Font.Bold = true
	myStyle.ApplyAlignment = true
	myStyle.ApplyFill = true
	myStyle.ApplyFont = true

	row := sh.AddRow()
	for i := 0; i < numFields; i++ {
		field := fType.Field(i)
		exnm := field.Tag.Get("exnm")
		if exnm == "" {
			continue
		}
		cell := row.AddCell()
		fns = append(fns, field.Name)
		cell.SetValue(exnm)
		cell.SetStyle(myStyle)
		cell.Row.SetHeight(30)
		cell.Row.Sheet.SetColWidth(10, 100, 30)
	}

	return fns, nil
}

func prepareFileWithSelectedColumns(fieldNames []string, fVal reflect.Value) (*xlsx.File, error) {
	rows := make([]reflect.Value, fVal.Len())
	for i := 0; i < fVal.Len(); i++ {
		rows[i] = fVal.Index(i)
	}
	wb := xlsx.NewFile()
	sh, err := wb.AddSheet("Sheet")
	if err != nil {
		return nil, err
	}

	fns, err := writeExcelHeadersWithSelectedColumns(fieldNames, rows[0], sh)
	if err != nil {
		return nil, err
	}
	for _, fv := range rows {
		row := sh.AddRow()
		for _, header := range fns {
			field := fv.FieldByName(header)
			cell := row.AddCell()
			if field.Type().String() == "time.Time" {
				t := field.Interface().(time.Time)
				if t.Year() < 2 {
					cell.SetValue("")
					continue
				}
				str := t.Format("2006-01-02 15:04:05")
				cell.SetValue(str)
			} else {
				cell.SetValue(field.Interface())
			}
		}
	}

	return wb, nil
}

func writeExcelHeadersWithSelectedColumns(fieldNames []string, aDtoFVal reflect.Value, sh *xlsx.Sheet) ([]string, error) {
	fns := make([]string, 0)

	fType := aDtoFVal.Type()
	if fType.NumField() < len(fieldNames) {
		return nil, fmt.Errorf("参数错误，选定导出的字段多于实际存在的字段")
	}

	myStyle := xlsx.NewStyle()
	myStyle.Alignment.Horizontal = "right"
	myStyle.Fill.FgColor = "FFFFFFFF"
	myStyle.Fill.BgColor = "00000000"
	myStyle.Font.Bold = true
	myStyle.ApplyAlignment = true
	myStyle.ApplyFill = true
	myStyle.ApplyFont = true

	row := sh.AddRow()
	for _, colNm := range fieldNames {
		cell := row.AddCell()
		field, ok := fType.FieldByName(colNm)
		if !ok {
			return nil, fmt.Errorf("无法找到字段")
		}
		fns = append(fns, field.Name)
		cell.SetValue(field.Tag.Get("exnm"))
		cell.SetStyle(myStyle)
		cell.Row.SetHeight(30)
		cell.Row.Sheet.SetColWidth(10, 100, 30)
	}

	return fns, nil
}
