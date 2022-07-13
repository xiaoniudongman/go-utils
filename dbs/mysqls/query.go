package mysqls

import (
	sql_ "database/sql"

	"github.com/xiaoniudongman/go-utils/tools/errs"
)

//http://jmoiron.github.io/sqlx/
func (this *MysqlDbInfo) SelectStrList(sql string, para ...interface{}) (dest []string, err error) {
	errs.CheckCommonErr(this.sqlDataDb.Select(&dest, sql, para...))
	return
}

func (this *MysqlDbInfo) SelectIntList(sql string, para ...interface{}) (dest []int, err error) {
	errs.CheckCommonErr(this.sqlDataDb.Select(&dest, sql, para...))
	return
}

func (this *MysqlDbInfo) QueryStruct(sql string, dest ...interface{}) (err error) {
	rows, err := this.sqlDataDb.Query(sql)
	errs.CheckCommonErr(err)
	defer rows.Close()
	rows.Next()
	err = rows.Scan(dest...)
	errs.CheckCommonErr(err)
	return
}

func (this *MysqlDbInfo) QueryIdMap(sql string, para ...interface{}) (dest map[string]string, err error) {
	dest = make(map[string]string, 100)
	// 查询数据
	var key, val sql_.NullString
	row, err := this.sqlDataDb.Query(sql, para...)
	defer row.Close()
	errs.CheckCommonErr(err)
	for row.Next() {
		err = row.Scan(&key, &val)
		errs.CheckCommonErr(err)
		dest[key.String] = val.String
	}
	return
}
