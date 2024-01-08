package myDatabase

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type loggerIface interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// camelToUnderscore case to CamelToUnderscore style.
func camelToUnderscore(camelStr string) string {
	l := utf8.RuneCountInString(camelStr)
	ss := strings.Split(camelStr, "")

	// we just care about the key of idx map,
	// the value of map is meaningless
	idx := make(map[int]int, 1)

	var rs []rune
	for _, s := range camelStr {
		rs = append(rs, []rune(string(s))...)
	}

	for i := l - 1; i >= 0; {
		if unicode.IsUpper(rs[i]) {
			var start, end int
			end = i
			j := i - 1
			for ; j >= 0; j-- {
				if unicode.IsLower(rs[j]) {
					start = j + 1
					break
				}
			}
			// handle the case: "BBC" or "AaBBB" case
			if end == l-1 {
				idx[start] = 1
			} else {
				if start == end {
					// value=1 is meaningless
					idx[start] = 1
				} else {
					idx[start] = 1
					idx[end] = 1
				}
			}
			i = j - 1
		} else {
			i--
		}
	}

	for i := l - 1; i >= 0; i-- {
		ss[i] = strings.ToLower(ss[i])
		if _, ok := idx[i]; ok && i != 0 {
			ss = append(ss[0:i], append([]string{"_"}, ss[i:]...)...)
		}
	}

	return strings.Join(ss, "")
}

func reflectEntity(obj any) (tableName string, fieldNames []string, fieldAddrs []any, fieldValues []any, fieldValuePlaceHolders []string, err error) {
	rfv := reflect.ValueOf(obj)
	if rfv.Kind() != reflect.Pointer {
		return "", nil, nil, nil, nil, fmt.Errorf("obj must be a pointer")
	}
	rfv = rfv.Elem()

	m := rfv.MethodByName("TableName")
	if !m.IsValid() {
		return "", nil, nil, nil, nil, fmt.Errorf("entity not implemented method[TableName]")
	}
	tableName = m.Call([]reflect.Value{})[0].Interface().(string)

	fieldNames, fieldAddrs, fieldValues, fieldValuePlaceHolders, err = reflectObject(obj)

	return tableName, fieldNames, fieldAddrs, fieldValues, fieldValuePlaceHolders, err
}

func reflectObject(obj any) (fieldNames []string, fieldAddrs []any, fieldValues []any, fieldValuePlaceHolders []string, err error) {
	rfv := reflect.ValueOf(obj)
	if rfv.Kind() != reflect.Pointer {
		return nil, nil, nil, nil, fmt.Errorf("obj must be a pointer")
	}
	rfv = rfv.Elem()
	fieldNames = make([]string, 0)
	fieldValues = make([]any, 0)
	fieldValuePlaceHolders = make([]string, 0)
	for i := 0; i < rfv.NumField(); i++ {
		if rfv.Field(i).Kind() == reflect.Struct && rfv.Field(i).Type() == reflect.TypeOf((*BaseEntity)(nil)).Elem() {
			for j := 0; j < rfv.Field(i).NumField(); j++ {
				fieldNames = append(fieldNames, camelToUnderscore(rfv.Field(i).Type().Field(j).Name))
				fieldAddrs = append(fieldAddrs, rfv.Field(i).Field(j).Addr().Interface())
				fieldValuePlaceHolders = append(fieldValuePlaceHolders, "?")
				if rfv.Field(i).Field(j).Type() == reflect.TypeOf((*time.Time)(nil)).Elem() {
					if rfv.Field(i).Field(j).Interface().(time.Time).IsZero() {
						t, _ := time.Parse(time.DateTime, "0001-01-01 08:00:00")
						fieldValues = append(fieldValues, t)
					} else {
						fieldValues = append(fieldValues, rfv.Field(i).Field(j).Interface())
					}
				} else {
					fieldValues = append(fieldValues, rfv.Field(i).Field(j).Interface())
				}
			}
		} else {
			fieldNames = append(fieldNames, camelToUnderscore(rfv.Type().Field(i).Name))
			fieldAddrs = append(fieldAddrs, rfv.Field(i).Addr().Interface())
			fieldValuePlaceHolders = append(fieldValuePlaceHolders, "?")
			fieldValues = append(fieldValues, rfv.Field(i).Interface())
		}
	}

	return fieldNames, fieldAddrs, fieldValues, fieldValuePlaceHolders, nil
}

func buildSql(stmt string, arg ...any) string {
	cnt := 0
	for _, c := range stmt {
		if c == '?' {
			v := arg[cnt]
			rv := reflect.ValueOf(v)
			switch rv.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				stmt = strings.Replace(stmt, "?", fmt.Sprintf("%d", arg[cnt]), 1)
			case reflect.Float32, reflect.Float64:
				stmt = strings.Replace(stmt, "?", fmt.Sprintf("%f", arg[cnt]), 1)
			case reflect.Struct:
				if rv.Type() == reflect.TypeOf((*time.Time)(nil)).Elem() {
					stmt = strings.Replace(stmt, "?", fmt.Sprintf("'%s'", arg[cnt].(time.Time).Format("2006-01-02 15:04:05.000")), 1)
				} else {
					stmt = strings.Replace(stmt, "?", fmt.Sprintf("'%s'", arg[cnt]), 1)
				}
			default:
				stmt = strings.Replace(stmt, "?", fmt.Sprintf("'%s'", arg[cnt]), 1)
			}
			cnt++
		}
	}
	return stmt
}

func Create[T any](db Iface, logger loggerIface, e *T) (insertedRecordId int64, err error) {
	tableName, fieldNames, _, fieldValues, fieldValuePlaceHolders, err := reflectEntity(e)
	if err != nil {
		return 0, err
	}
	nfns := make([]string, 0)
	nfvs := make([]any, 0)
	nfvphs := make([]string, 0)
	for idx := range fieldNames {
		if strings.EqualFold(fieldNames[idx], "id") {
			continue
		} else if strings.EqualFold(fieldNames[idx], "created_time") && fieldValues[idx].(time.Time).Before(time.UnixMilli(0)) {
			nfns = append(nfns, fieldNames[idx])
			nfvs = append(nfvs, time.Now())
			nfvphs = append(nfvphs, fieldValuePlaceHolders[idx])
		} else {
			nfns = append(nfns, fieldNames[idx])
			nfvs = append(nfvs, fieldValues[idx])
			nfvphs = append(nfvphs, fieldValuePlaceHolders[idx])
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(nfns, ","), strings.Join(nfvphs, ","))
	result, err := db.Exec(stmt, nfvs...)
	rs, _ := result.RowsAffected()
	logger.Infof("[%s], row affected: %d", buildSql(stmt, nfvs...), rs)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func PhysicalDeleteById[T any](db Iface, logger loggerIface, id int64) error {
	obj := new(T)
	tableName, _, _, _, _, err := reflectEntity(obj)
	if err != nil {
		return err
	}
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableName)
	result, err := db.Exec(stmt, id)
	rs, _ := result.RowsAffected()
	logger.Infof("[%s], row affected: %d", buildSql(stmt, id), rs)
	if err != nil {
		return err
	}
	return nil
}

func LogicalDeleteById[T any](deleteByUid int64, db Iface, logger loggerIface, id int64) error {
	obj := new(T)
	tableName, _, _, _, _, err := reflectEntity(obj)
	if err != nil {
		return err
	}
	stmt := fmt.Sprintf("UPDATE %s SET deleted_by = ?, deleted_time = ? WHERE id = %d", tableName, id)
	now := time.Now()
	result, err := db.Exec(stmt, deleteByUid, now)
	rs, _ := result.RowsAffected()
	logger.Infof("[%s], row affected: %d", buildSql(stmt, deleteByUid, now), rs)
	if err != nil {
		return err
	}
	return nil
}

func UpdateById[T any](db Iface, logger loggerIface, e *T) error {

	tableName, fieldNames, _, fieldValues, fieldValuePlaceHolders, err := reflectEntity(e)
	if err != nil {
		return err
	}
	var id int64
	nfns := make([]string, 0)
	nfvs := make([]any, 0)
	nfvphs := make([]string, 0)
	for idx := range fieldNames {
		if strings.EqualFold(fieldNames[idx], "id") {
			id = fieldValues[idx].(int64)
			continue
		} else if strings.EqualFold(fieldNames[idx], "updated_time") {
			nfns = append(nfns, fieldNames[idx])
			nfvs = append(nfvs, time.Now())
			nfvphs = append(nfvphs, fieldValuePlaceHolders[idx])
		} else {
			nfns = append(nfns, fieldNames[idx])
			nfvs = append(nfvs, fieldValues[idx])
			nfvphs = append(nfvphs, fieldValuePlaceHolders[idx])
		}
	}
	updates := make([]string, 0, len(nfns)+2)
	for i := 0; i < len(nfns); i++ {
		updates = append(updates, fmt.Sprintf("%s=%s", nfns[i], nfvphs[i]))
	}
	stmt := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", tableName, strings.Join(updates, ", "), id)
	result, err := db.Exec(stmt, nfvs...)
	rs, _ := result.RowsAffected()
	logger.Infof("[%s], row affected: %d", buildSql(stmt, nfvs...), rs)
	if err != nil {
		return err
	}
	return nil
}

func SelectById[T any](db Iface, logger loggerIface, id int64) (dst *T, err error) {
	dst = new(T)
	tableName, fieldNames, fieldAddrs, _, _, err := reflectEntity(dst)
	if err != nil {
		return nil, err
	}
	stmt := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", strings.Join(fieldNames, ","), tableName)
	row := db.QueryRow(stmt, id)
	logger.Infof("[%s]", buildSql(stmt, id))
	if row.Err() != nil {
		return nil, row.Err()
	}

	if err := row.Scan(fieldAddrs...); err != nil {
		return nil, err
	}
	return dst, nil
}

func SelectPage[T any](db Iface, logger loggerIface, filter *Filter) ([]T, int64, error) {
	tableName, fieldNames, _, _, _, err := reflectEntity(new(T))
	if err != nil {
		return nil, 0, err
	}

	var total int64
	if filter.GetIsTotal() {
		stmt := fmt.Sprintf("SELECT count(*) FROM %s WHERE %s", tableName, filter.BuildQueryString())
		logger.Infof("[%s]", buildSql(stmt, filter.GetArgs()...))
		row := db.QueryRow(stmt, filter.GetArgs()...)
		row.Scan(&total)
	}

	sc := ""
	if s := filter.BuildSelectString(); s == "*" {
		sc = strings.Join(fieldNames, ", ")
	} else {
		sc = s
	}
	stmt := fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY %s LIMIT %d, %d", sc,
		tableName, filter.BuildQueryString(), filter.GetOrder(), filter.GetOffset(), filter.GetLimit())
	logger.Infof("[%s]", buildSql(stmt, filter.GetArgs()...))
	rows, err := db.Query(stmt, filter.GetArgs()...)
	if err != nil {
		return nil, 0, err
	}

	data := make([]T, 0)
	for rows.Next() {
		o := new(T)
		fieldNames, fieldAddrs, _, _, _ := reflectObject(o)
		fa := make([]any, 0)
		if cs := filter.GetSelect(); len(cs) > 0 && len(cs) < len(fieldNames) {
			for _, c := range cs {
				for i, n := range fieldNames {
					if c == n {
						fa = append(fa, fieldAddrs[i])
						break
					}
				}
			}
		} else {
			fa = fieldAddrs
		}
		if err := rows.Scan(fa...); err != nil {
			return nil, 0, err
		}
		data = append(data, *o)
	}

	return data, total, nil
}

func DoTransaction(db Iface, f func(tx Iface) error) error {
	tx, err := db.(*sql.DB).Begin()
	if err != nil {
		return fmt.Errorf("occur error while begin transaction: %s", err)
	}
	err = f(tx)
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}
