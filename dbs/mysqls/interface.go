package mysqls

/*
有关mysql数据库连接的封装
*/
import (
	"fmt"
	"time"

	"github.com/xiaoniudongman/go-utils/tools/errs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xiaoniudongman/go-utils/common/consts"
	"github.com/xiaoniudongman/go-utils/config"
)

type MysqlMethod interface {
	GetDbConnFromConf(c *config.ConfigEngine, name string)
	SelectStrList(sql string, para ...interface{}) (dest []string, err error)
	SelectIntList(sql string, para ...interface{}) (dest []int, err error)
	QueryStruct(sql string, dest ...interface{}) (err error)
	QueryIdMap(sql string, para ...interface{}) (dest map[string]string, err error)
	GetDb() *sqlx.DB
	GetTableName(key string) string
}

type MysqlDbInfo struct {
	sqlDataDb *sqlx.DB
	tableName map[string]string
	maxConns  int
	dbTimeOut int
}

func getMySqlLoginStr(data *config.MysqlDbData) string {
	return data.User + ":" +
		data.Password + "@tcp(" + data.Host + ":" +
		data.Port + ")/" + data.Db_name + "?charset=utf8"
}

func (this *MysqlDbInfo) createDatabaseConns(login *config.MysqlDbData) {
	db, err := sqlx.Open("mysql", getMySqlLoginStr(login))
	errs.CheckFatalErr(err)
	db.SetConnMaxLifetime(time.Duration(login.Time_out) * time.Second)
	db.SetMaxOpenConns(login.Max_conns)
	db.SetMaxIdleConns(login.Max_conns)
	errs.CheckFatalErr(db.Ping())
	this.sqlDataDb = db
	this.tableName = login.Table_name
	this.maxConns = login.Max_conns
	this.dbTimeOut = login.Time_out
}

func (this *MysqlDbInfo) GetDbConnFromConf(c *config.ConfigEngine, name string) {
	this.createDatabaseConns(c.GetMySqlFromConf(name))
}

func (this *MysqlDbInfo) GetDb() *sqlx.DB {
	return this.sqlDataDb
}

func (this *MysqlDbInfo) GetTableName(key string) string {
	if val, ok := this.tableName[key]; ok {
		return val
	} else {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("key %s not in tablenames.", key)))
		return consts.BLANK
	}
}
