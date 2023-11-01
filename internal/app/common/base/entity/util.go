package entity

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"example.com/internal/app/common/base/vo"
	"example.com/internal/pkg/myerr/myerrimp"
	"example.com/internal/pkg/stringutil"
)

type LoggerIface interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
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
		if rfv.Field(i).Kind() == reflect.Struct {
			for j := 0; j < rfv.Field(i).NumField(); j++ {
				fieldNames = append(fieldNames, stringutil.CamelToUnderscore(rfv.Field(i).Type().Field(j).Name))
				fieldAddrs = append(fieldAddrs, rfv.Field(i).Field(j).Addr().Interface())
				fieldValuePlaceHolders = append(fieldValuePlaceHolders, "?")
				fieldValues = append(fieldValues, rfv.Field(i).Field(j).Interface())
			}
		} else {
			fieldNames = append(fieldNames, stringutil.CamelToUnderscore(rfv.Type().Field(i).Name))
			fieldAddrs = append(fieldAddrs, rfv.Field(i).Addr().Interface())
			fieldValuePlaceHolders = append(fieldValuePlaceHolders, "?")
			fieldValues = append(fieldValues, rfv.Field(i).Interface())
		}
	}

	return fieldNames, fieldAddrs, fieldValues, fieldValuePlaceHolders, nil
}

func Create[T any](ctx vo.SessionInfo, db *sql.DB, logger LoggerIface, e T) (insertedRecordId int64, err error) {
	tableName, fieldNames, _, fieldValues, fieldValuePlaceHolders, err := reflectEntity(&e)
	if err != nil {
		return 0, myerrimp.New(err)
	}
	nfns := make([]string, 0)
	nfvs := make([]any, 0)
	nfvphs := make([]string, 0)
	for idx := range fieldNames {
		if strings.EqualFold(fieldNames[idx], "id") {
			continue
		} else if strings.EqualFold(fieldNames[idx], "created_by") && fieldValues[idx].(int) == 0 {
			nfns = append(nfns, fieldNames[idx])
			nfvs = append(nfvs, ctx.Uid)
			nfvphs = append(nfvphs, fieldValuePlaceHolders[idx])
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
	logger.Infof("[%s], %v", stmt, nfvs)
	result, err := db.ExecContext(ctx.Ctx, stmt, nfvs...)
	if err != nil {
		return 0, myerrimp.New(err)
	}
	return result.LastInsertId()
}

func PhysicalDeleteById[T any](ctx vo.SessionInfo, db *sql.DB, logger LoggerIface, id int64) error {
	obj := new(T)
	tableName, _, _, _, _, err := reflectEntity(obj)
	if err != nil {
		return myerrimp.New(err)
	}
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableName)
	logger.Infof("[%s], %v", stmt, id)
	_, err = db.ExecContext(ctx.Ctx, stmt, id)
	if err != nil {
		return myerrimp.New(err)
	}
	return nil
}

func LogicalDeleteById[T any](ctx vo.SessionInfo, db *sql.DB, logger LoggerIface, id int64) error {
	obj := new(T)
	tableName, _, _, _, _, err := reflectEntity(obj)
	if err != nil {
		return myerrimp.New(err)
	}
	stmt := fmt.Sprintf("UPDATE %s SET deleted_by = ?, deleted_time = ? WHERE id = %d", tableName, id)
	now := time.Now()
	logger.Infof("[%s], %d, %s", stmt, ctx.Uid, now)
	_, err = db.ExecContext(ctx.Ctx, stmt, ctx.Uid, now)
	if err != nil {
		return myerrimp.New(err)
	}
	return nil
}

func UpdateById[T any](ctx vo.SessionInfo, db *sql.DB, logger LoggerIface, e T) error {
	tableName, fieldNames, _, fieldValues, fieldValuePlaceHolders, err := reflectEntity(&e)
	if err != nil {
		return myerrimp.New(err)
	}
	var id int
	nfns := make([]string, 0)
	nfvs := make([]any, 0)
	nfvphs := make([]string, 0)
	for idx := range fieldNames {
		if strings.EqualFold(fieldNames[idx], "id") {
			id = fieldValues[idx].(int)
			continue
		} else if strings.EqualFold(fieldNames[idx], "updated_by") {
			nfns = append(nfns, fieldNames[idx])
			nfvs = append(nfvs, ctx.Uid)
			nfvphs = append(nfvphs, fieldValuePlaceHolders[idx])
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
	logger.Infof("[%s], %v", stmt, nfvs)
	_, err = db.ExecContext(ctx.Ctx, stmt, nfvs...)
	if err != nil {
		return myerrimp.New(err)
	}
	return nil
}

func SelectById[T any](ctx vo.SessionInfo, db *sql.DB, logger LoggerIface, id int64) (dst *T, err error) {
	dst = new(T)
	tableName, fieldNames, fieldAddrs, _, _, err := reflectEntity(dst)
	if err != nil {
		return nil, myerrimp.New(err)
	}
	stmt := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", strings.Join(fieldNames, ","), tableName)
	logger.Infof("[%s], %v", stmt, id)
	row := db.QueryRowContext(ctx.Ctx, stmt, id)
	if row.Err() != nil {
		return nil, myerrimp.New(row.Err())
	}

	if err := row.Scan(fieldAddrs...); err != nil {
		return nil, myerrimp.New(err)
	}
	return dst, nil
}

func SelectPage[T any](ctx vo.SessionInfo, db *sql.DB, logger LoggerIface, conditions []string, orderBy string, offset int64, limit int64) ([]T, int64, error) {
	tableName, fieldNames, _, _, _, err := reflectEntity(new(T))
	if err != nil {
		return nil, 0, myerrimp.New(err)
	}
	if orderBy == "" {
		orderBy = "id ASC"
	}

	conditions = append(conditions, "1=1")
	conds := strings.Join(conditions, " AND ")
	stmt := fmt.Sprintf("SELECT count(*) FROM %s WHERE %s", tableName, conds)
	logger.Infof("[%s]", stmt)
	row := db.QueryRowContext(ctx.Ctx, stmt)
	var total int64
	row.Scan(&total)

	stmt = fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY %s LIMIT %d, %d", strings.Join(fieldNames, ","),
		tableName, conds, orderBy, offset, limit)
	logger.Infof("[%s]", stmt)
	rows, err := db.QueryContext(ctx.Ctx, stmt)
	if err != nil {
		return nil, 0, myerrimp.New(err)
	}

	data := make([]T, 0)
	for rows.Next() {
		o := new(T)
		_, fieldAddrs, _, _, _ := reflectObject(o)
		if err := rows.Scan(fieldAddrs...); err != nil {
			return nil, 0, myerrimp.New(err)
		}
		data = append(data, *o)
	}

	return data, total, nil
}
