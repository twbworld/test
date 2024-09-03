package dao

import (
	"bytes"

	"github.com/twbworld/dating/model/db"
)

type dbUtils struct{}

func (u *dbUtils) getInsertSql(d db.Dbfunc, data map[string]interface{}) (string, []interface{}) {
	if len(data) < 1 {
		return ``, []interface{}{}
	}

	var (
		fields bytes.Buffer
		values bytes.Buffer
		sql    bytes.Buffer
		args   []interface{} = make([]interface{}, 0, len(data))
	)
	for k, v := range data {
		fields.WriteString("`")
		fields.WriteString(k)
		fields.WriteString("`,")
		values.WriteString(`?,`)
		args = append(args, v)
	}

	sql.WriteString("INSERT INTO `")
	sql.WriteString(d.TableName())
	sql.WriteString("`(")
	sql.Write(bytes.TrimRight(fields.Bytes(), `,`))
	sql.WriteString(`) VALUES(`)
	sql.Write(bytes.TrimRight(values.Bytes(), `,`))
	sql.WriteString(`)`)

	return sql.String(), args
}

func (u *dbUtils) getUpdateSql(d db.Dbfunc, id uint, data map[string]interface{}) (string, []interface{}) {
	if len(data) < 1 {
		return ``, []interface{}{}
	}

	var (
		fields bytes.Buffer
		sql    bytes.Buffer
		args   []interface{} = make([]interface{}, 0, len(data))
	)

	for k, v := range data {
		fields.WriteString(" `")
		fields.WriteString(k)
		fields.WriteString("` = ?,")
		args = append(args, v)
	}

	sql.WriteString("UPDATE `")
	sql.WriteString(d.TableName())
	sql.WriteString("` SET")
	sql.Write(bytes.TrimRight(fields.Bytes(), `,`))
	sql.WriteString(" WHERE `id` = ?")
	args = append(args, id)

	return sql.String(), args
}
