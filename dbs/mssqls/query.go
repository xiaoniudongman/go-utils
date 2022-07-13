package mssqls

import (
	"github.com/xiaoniudongman/go-utils/tools/errs"
)

//http://jmoiron.github.io/sqlx/
func (this *MssqlDbInfo) QueryIdList(sql string, para ...interface{}) (ids []string, err error) {
	err = this.SqlDataDb.Select(&ids, sql, para...)
	if err != nil {
		err = this.SqlDataDb.Select(&ids, sql, para...)
		if err != nil {
			errs.CheckCommonErr(err)
			return
		}
	}
	return
}

func (this *MssqlDbInfo) QueryIdIntList(sql string, para ...interface{}) (ids []int, err error) {
	err = this.SqlDataDb.Select(&ids, sql, para...)
	if err != nil {
		err = this.SqlDataDb.Select(&ids, sql, para...)
		if err != nil {
			errs.CheckCommonErr(err)
			return
		}
	}
	return
}

func (this *MssqlDbInfo) QueryStruct(sql string, dest ...interface{}) (err error) {
	rows, err := this.SqlDataDb.Query(sql)
	if err != nil {
		rows, err = this.SqlDataDb.Query(sql)
		if err != nil {
			errs.CheckCommonErr(err)
			return
		}
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(dest...)
	errs.CheckCommonErr(err)
	return
}

func (this *MssqlDbInfo) QueryIdMap(sql string) (result map[string]string, err error) {
	result = make(map[string]string, 0)
	// 查询数据
	var key, val string
	row, err := this.SqlDataDb.Query(sql)
	if err != nil {
		row, err = this.SqlDataDb.Query(sql)
		if err != nil {
			errs.CheckCommonErr(err)
			return
		}
	}
	defer row.Close()
	errs.CheckCommonErr(err)
	for row.Next() {
		err = row.Scan(&key, &val)
		errs.CheckCommonErr(err)
		result[key] = val
	}
	return
}
