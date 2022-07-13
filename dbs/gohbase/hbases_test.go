package gohbase

import (
	"fmt"
	"testing"

	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/tools/errs"
)

const (
	Config_path = "../../config/test.yaml"
)

func TestHbaseConnectionFromConfig(t *testing.T) {
	db := &HBaseDbInfo{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db")
	t.Log(db.Ping())
	t.Log("db", db)
	t.Log("db", db.TableName)
	f := map[string][]string{"alg_nrt_comicsim_cb": []string{"recs"}}
	getRes, _ := db.GetsByOption("recommender:recall_samh_alg_nrt_comicsim_cb", "100012", hrpc.Families(f))
	fmt.Println(getRes)
}

func TestHbaseConnectionFromConfig1(t *testing.T) {
	db := &HBaseDbInfo{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db")
	t.Log(db.Ping())
	t.Log("db", db)
	t.Log("db", db.TableName)
	f := map[string][]string{"alg_nrt_comicsim_cb": []string{"recs"}}
	f1 := filter.NewPrefixFilter([]byte("100193"))
	getRes, _ := db.GetsByScanOption("recommender:recall_samh_alg_nrt_comicsim_cb", hrpc.Families(f), hrpc.Filters(f1))
	fmt.Println(getRes)
}
