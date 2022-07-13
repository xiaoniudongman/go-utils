package hbases_test

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"testing"

	"github.com/xiaoniudongman/go-utils/dbs/hbases/gen-go/hbase"

	"github.com/xiaoniudongman/go-utils/dbs/hbases"

	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/tools/errs"
)

const (
	Config_path = "../../config/test.yaml"
)

func TestGet(t *testing.T) {
	db := &hbases.HBaseThriftAgent{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db2")
	t.Log("db", db)
	t.Log("db", db.TableName)
	getRes, _ := db.Get(context.Background(), "recommend_samh:alg_nrt_ar_fpgrowth", &hbase.TGet{Row: []byte("5c341a7dda910511:comic")})
	fmt.Println(getRes)
}

func TestGetRow(t *testing.T) {
	db := &hbases.HBaseThriftAgent{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db2")
	t.Log("db", db)
	t.Log("db", db.TableName)
	getRes, _ := db.GetRow(context.Background(), "recommend_samh:alg_nrt_ar_fpgrowth", "5c341a7dda910511:comic", nil, true)
	fmt.Println(getRes)
}

func TestGetMultipleRows(t *testing.T) {
	db := &hbases.HBaseThriftAgent{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db2")
	t.Log("db", db)
	t.Log("db", db.TableName)
	var tColumns = []*hbase.TColumn{
		{
			Family:    []byte("info"),
			Qualifier: []byte("recs"),
		},
	}
	getRes, _ := db.GetMultipleRows(context.Background(), "recommend_samh:alg_nrt_ar_fpgrowth",
		[]string{"5c341a7dda910511:comic"},
		tColumns, true)
	fmt.Println(getRes)
}

func TestGetScannerResultsAll(t *testing.T) {
	db := &hbases.HBaseThriftAgent{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db2")
	t.Log("db", db)
	t.Log("db", db.TableName)
	var tColumns = []*hbase.TColumn{
		{
			Family:    []byte("info"),
			Qualifier: []byte("weight"),
		},
	}
	getRes, _ := db.GetScannerResultsAll(context.Background(), "recommend_samh:item_base", tColumns, 1000000, true)
	fmt.Println(getRes)
}
