package mdb

import (
	"database/sql"
	"fmt"
	_ "log"
	_ "os"
)

/*

	ALTER STATEMENT

*/

type AlterStatement struct {
	tableName, changeStr string
}

func NewAlterStatement(cd string, qs string) AlterStatement {
	return AlterStatement{tableName: cd, changeStr: qs}
}

func (a *AlterStatement) Serialize() string {
	return "ALTER TABLE " + a.tableName + " " + a.changeStr
}

func (a *AlterStatement) Apply(db *sql.DB) (result AlterResult) {

	res, err := db.Exec(a.Serialize())

	if err != nil {
		result = AlterResult{alter: *a, rowsAffected: -1, err: err}
		fmt.Printf("error occured with %s: %#v\n",
			a.tableName,
			result.err)
		return result
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		result = AlterResult{alter: *a, rowsAffected: -1, err: err}
		fmt.Printf("error occured getting affected rows with %s: %#v\n",
			a.changeStr,
			result.err)
		return result
	}

	return AlterResult{alter: *a, rowsAffected: int(rowsAffected), err: nil}
}
