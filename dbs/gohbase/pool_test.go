package gohbase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/xiaoniudongman/go-utils/dbs/gohbase"
	"github.com/xiaoniudongman/go-utils/dbs/hbases/gen-go/hbase"
)

//func TestHbaseGets(t *testing.T) {
//	defaultCtx := context.Background()
//	cli, _ := hbases.CreateClient()
//	fmt.Println(hbases.ConnPool)
//	t.Log(cli)
//	t.Log(hbases.Ping(cli))
//	kkk := cli.(hbases.MyHbaseClient)
//	jj, err := kkk.Client.Get(defaultCtx, []byte("recommend_samh:alg_nrt_ar_fpgrowth"),
//		&hbase.TGet{Row: []byte("5c341a7dda910511:comic")})
//	t.Log(err)
//	t.Log(jj.ColumnValues)
//	for _, xx := range jj.ColumnValues {
//		fmt.Println("string(xx.Family)", string(xx.Family))
//		fmt.Println("string(xx.Qualifier)", string(xx.Qualifier))
//		fmt.Println("string(xx.Value)", string(xx.Value))
//	}
//}

func TestHbaseGets2(t *testing.T) {
	//ConnPool
	defaultCtx := context.Background()
	//Pool:= hbases.Init()
	fmt.Println(gohbase.ConnPool)
	//t.Log(cli)
	//t.Log(hbases.Ping(cli))
	aaa, _ := gohbase.ConnPool.Get()
	kkk := aaa.(gohbase.MyHbaseClient)
	jj, err := kkk.Client.Get(defaultCtx,
		[]byte("recommend_samh:alg_nrt_ar_fpgrowth"),
		&hbase.TGet{Row: []byte("5c341a7dda910511:comic")})
	t.Log(err)
	t.Log(jj.ColumnValues)
	for _, xx := range jj.ColumnValues {
		fmt.Println("string(xx.Family)", string(xx.Family))
		fmt.Println("string(xx.Qualifier)", string(xx.Qualifier))
		fmt.Println("string(xx.Value)", string(xx.Value))
	}
}
