package mssqls

/*
有关sqlServer数据库连接的封装
*/
import (
	"fmt"
	"time"

	"github.com/xiaoniudongman/go-utils/tools/errs"

	"github.com/jmoiron/sqlx"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/xiaoniudongman/go-utils/config"
)

type MssqlMethod interface {
	GetMssqlConnFromConf(c *config.ConfigEngine, name string)
	QueryIdList(sql string, para ...string) (ids []string, err error)
	QueryIdIntList(sql string, para ...string) (ids []int, err error)
	QueryStruct(sql string, dest ...interface{}) (err error)
	QueryIdMap(sql string) (result map[string]string, err error)
}

type MssqlDbInfo struct {
	SqlDataDb *sqlx.DB
	TableName map[string]string
	MaxConns  int
	DbTimeOut int
}

func getMssqlLoginStr(section *config.MssqlDbData) string {
	//连接字符串
	return fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s",
		section.Host, section.Port, section.Db_name, section.User, section.Password)
}

func (this *MssqlDbInfo) createMssqlConns(login *config.MssqlDbData) {
	db, err := sqlx.Open("mssql", getMssqlLoginStr(login))
	errs.CheckFatalErr(err)
	db.SetConnMaxLifetime(time.Duration(login.Time_out) * time.Second)
	db.SetMaxOpenConns(login.Max_conns)
	db.SetMaxIdleConns(login.Max_conns)
	err = db.Ping()
	errs.CheckFatalErr(err)
	this.SqlDataDb = db
	this.TableName = login.Table_name
	this.MaxConns = login.Max_conns
	this.DbTimeOut = login.Time_out
}

func (this *MssqlDbInfo) GetMssqlConnFromConf(c *config.ConfigEngine, name string) {
	this.createMssqlConns(c.GetMssqlDataFromConf(name))
}
