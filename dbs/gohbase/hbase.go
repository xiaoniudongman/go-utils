package gohbase

import (
	"context"
	"errors"
	"io"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/tools/errs"
)

type HBaseDbInfo struct {
	Zkquorum  string
	Option    string
	Namespace string
	QueueSize int
	TableName map[string]string
	_client   gohbase.Client
}

func (hb *HBaseDbInfo) ConnectHBase(account string, zkquorum string, options ...gohbase.Option) {
	hb.Zkquorum = zkquorum
	hb._client = gohbase.NewClient(zkquorum, options...)
}

func (hb *HBaseDbInfo) connectHBase(db *config.HBaseDbData) {
	hb.Zkquorum = db.ZK
	hb.Namespace = db.Namespace
	hb.TableName = db.TableName
	hb._client = gohbase.NewClient(db.ZK, gohbase.RpcQueueSize(db.QueueSize), gohbase.CompressionCodec("snappy"))
}

//通过hb.PutsByRowkeyVersion(table, rowkey, values, hrpc.Timestamp(timestamp))调用，其中timestamp是time.Time类型，options也可以是其他 func(hrpc.Call)的函数
func (hb *HBaseDbInfo) PutsByRowkeyVersion(table, rowKey string, values map[string]map[string][]byte, options func(hrpc.Call) error) (err error) {
	putRequest, err := hrpc.NewPutStr(context.Background(), table, rowKey, values, options)
	errs.CheckCommonErr(err)
	_, err = hb._client.Put(putRequest)
	errs.CheckCommonErr(err)

	return
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBaseDbInfo) GetsByOption(table, rowkey string, options ...func(hrpc.Call) error) (*hrpc.Result, error) {
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowkey, options...)
	if nil != err {
		errs.CheckCommonErr(err)
		return nil, err
	}
	res, err := hb._client.Get(getRequest)
	if nil != err {
		errs.CheckCommonErr(err)
		return nil, err
	}
	return res, err
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBaseDbInfo) GetsByScanOption(table string, options ...func(hrpc.Call) error) (rsp []*hrpc.Result, err error) {
	var (
		scanRequest *hrpc.Scan
		res         *hrpc.Result
		err1        error
	)
	scanRequest, err = hrpc.NewScanStr(context.Background(), table, options...)
	if nil != err {
		errs.CheckCommonErr(err)
		return nil, err
	}
	scanner := hb._client.Scan(scanRequest)
	for {
		res, err1 = scanner.Next()
		if err1 == io.EOF || res == nil {
			break
		}
		if nil != err1 {
			errs.CheckCommonErr(err1)
			return nil, err1
		}
		rsp = append(rsp, res)
	}
	return
}

func (this *HBaseDbInfo) Ping() error {
	if this._client != nil {
		return nil
	}
	return errors.New("连接为空")
}

func (this *HBaseDbInfo) GetClient() gohbase.Client {
	return this._client
}

func (this *HBaseDbInfo) GetDbConnFromConf(c *config.ConfigEngine, name string) {
	this.connectHBase(c.GetHBaseFromConf(name))
}
