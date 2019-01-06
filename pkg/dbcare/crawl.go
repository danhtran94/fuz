package dbcare

import (
	"database/sql"
	"fmt"

	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

var DBClient sqlbuilder.Database

func SetupDBClient(connectString string, log bool) (err error) {
	var sess *sql.DB
	sess, err = sql.Open("postgres", connectString)

	DBClient, err = postgresql.New(sess)
	if err != nil {
		return err
	}
	DBClient.SetLogging(log)

	return err
}

type ColTypes map[string]string

func GetColTypes(tableName string) (ColTypes, error) {
	colTypes := make(ColTypes)

	queryItr := DBClient.SelectFrom("pg_attribute").
		Columns("attname AS col_name", db.Raw(fmt.Sprintf("%s::regtype AS data_type", "atttypid"))).
		Where(db.Cond{
			"attrelid =":   db.Raw(fmt.Sprintf("'%s'::regclass", tableName)),
			"attnum >":     0,
			"attisdropped": false,
		}).OrderBy("attnum").Iterator()

	for queryItr.Next() {
		var colName, dataType string

		err := queryItr.Scan(&colName, &dataType)
		if err != nil {
			break
		}

		colTypes[colName] = dataType
	}

	if err := queryItr.Err(); err != nil {
		return nil, err
	}

	return colTypes, nil
}
