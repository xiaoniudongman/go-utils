package hbases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/xiaoniudongman/go-utils/common/consts"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/xiaoniudongman/go-utils/config"
	"github.com/xiaoniudongman/go-utils/dbs/hbases/gen-go/hbase"
)

var (
	connTimeout     = time.Second * 2
	idleTimeout     = time.Second * 120
	timeout         = time.Second * 10
	maxConn         = int32(100)
	thriftPoolAgent *ThriftPoolAgent
)

func thriftDial(addr, user, passwd string, connTimeout time.Duration) (*IdleClient, error) {
	transport, err := thrift.NewTHttpClient(addr)
	if err != nil {
		return nil, err
	}
	if err := transport.Open(); err != nil {
		return nil, err
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	// 设置用户名密码
	httClient := transport.(*thrift.THttpClient)
	httClient.SetHeader("ACCESSKEYID", user)
	httClient.SetHeader("ACCESSSIGNATURE", passwd)
	return &IdleClient{
		Transport: transport,
		RawClient: hbase.NewTHBaseServiceClientFactory(transport, protocolFactory),
	}, nil
}

func closeThriftClient(c *IdleClient) error {
	if c == nil {
		return nil
	}
	return c.Transport.Close()
}

type HBaseThriftAgent struct {
	_client   *ThriftPoolAgent
	Namespace string            `yaml:"Namespace"`
	TableName map[string]string `yaml:"Table_name"`
}

func NewHBaseThriftAgent() *HBaseThriftAgent {
	return &HBaseThriftAgent{}
}

func (hb *HBaseThriftAgent) NewTHttpClient(host, user, passwd string) {
	pool := &ThriftPoolConfig{
		Addr:        host,
		User:        user,
		Passwd:      passwd,
		MaxConn:     maxConn,
		ConnTimeout: connTimeout,
		IdleTimeout: idleTimeout,
		Timeout:     timeout,
	}
	thriftPool := NewThriftPool(pool, thriftDial, closeThriftClient)
	thriftPoolAgent = new(ThriftPoolAgent)
	thriftPoolAgent.Init(thriftPool)
	hb._client = thriftPoolAgent
}

func (this *HBaseThriftAgent) GetDbConnFromConf(c *config.ConfigEngine, name string) {
	conf := c.GetHBaseV2FromConf(name)
	this.NewTHttpClient(conf.Thrift, conf.User, conf.Passwd)
	this.TableName = conf.TableName
	this.Namespace = conf.Namespace
}

// Method for getting data from a row.
//
// If the row cannot be found an empty Result is returned.
// This can be checked by the empty field of the TResult
//
// @return the result
//
// Parameters:
//  - Table: the table to get from
//  - Tget: the TGet to fetch
func (this *HBaseThriftAgent) Get(ctx context.Context, table string, tget *hbase.TGet) (*hbase.TResult_, error) {
	var tResult_ *hbase.TResult_
	err := thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		var err2 error
		tResult_, err2 = client.Get(ctx, []byte(table), tget)
		return err2
	})
	if nil == err {
		return tResult_, err
	}
	return nil, err
}

// Method for getting data from a row.
//
// If the row cannot be found an empty Result is returned.
// This can be checked by the empty field of the TResult
//
// @return the result
//
// Parameters:
//  - Table: the table to get from
//  - Tget: the TGet to fetch
func (this *HBaseThriftAgent) GetRow(ctx context.Context, table, rowKey string, cols []*hbase.TColumn, hasFamily bool) (map[string]string, error) {
	var tResult_ *hbase.TResult_
	err := thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		var err2 error
		tResult_, err2 = client.Get(ctx, []byte(table), &hbase.TGet{Row: []byte(rowKey), Columns: cols})
		return err2
	})
	if nil == err {
		var result = make(map[string]string, len(tResult_.ColumnValues))
		if hasFamily {
			for _, tColumnValue := range tResult_.ColumnValues {
				q, f, v := tColumnValue.GetQualifier(), tColumnValue.GetFamily(), tColumnValue.GetValue()
				result[fmt.Sprintf("%s:%s", string(f), string(q))] = string(v)
			}
		} else {
			for _, tColumnValue := range tResult_.ColumnValues {
				q, v := tColumnValue.GetQualifier(), tColumnValue.GetValue()
				result[string(q)] = string(v)
			}
		}
		return result, err
	}
	return nil, err
}

// Method for getting multiple rows.
//
// If a row cannot be found there will be a null
// value in the result list for that TGet at the
// same position.
//
// So the Results are in the same order as the TGets.
//
// Parameters:
//  - Table: the table to get from
//  - Tgets: a list of TGets to fetch, the Result list
// will have the Results at corresponding positions
// or null if there was an error
func (this *HBaseThriftAgent) GetMultiple(ctx context.Context, table string, tgets []*hbase.TGet) ([]*hbase.TResult_, error) {
	var tResult_ []*hbase.TResult_
	err := thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		var err2 error
		tResult_, err2 = client.GetMultiple(ctx, []byte(table), tgets)
		return err2
	})
	if nil == err {
		return tResult_, err
	}
	return nil, err
}

// Method for getting multiple rows.
//
// If a row cannot be found there will be a null
// value in the result list for that TGet at the
// same position.
//
// So the Results are in the same order as the TGets.
//
// Parameters:
//  - Table: the table to get from
//  - Tgets: a list of TGets to fetch, the Result list
// will have the Results at corresponding positions
// or null if there was an error
func (this *HBaseThriftAgent) GetMultipleRows(ctx context.Context, table string, rowKeys []string, cols []*hbase.TColumn, hasFamily bool) ([]map[string]string, error) {
	var tResults []*hbase.TResult_
	err := thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		var err2 error
		var tGets []*hbase.TGet
		for _, r := range rowKeys {
			tGets = append(tGets, &hbase.TGet{
				Row:     []byte(r),
				Columns: cols,
			})
		}
		tResults, err2 = client.GetMultiple(ctx, []byte(table), tGets)
		return err2
	})
	if nil == err {
		var result = make([]map[string]string, consts.ZERO, len(rowKeys))
		if hasFamily {
			for _, tResult := range tResults {
				var tmp = make(map[string]string)
				for _, tColumnValue := range tResult.ColumnValues {
					q, f, v := tColumnValue.GetQualifier(), tColumnValue.GetFamily(), tColumnValue.GetValue()
					tmp[fmt.Sprintf("%s:%s", string(f), string(q))] = string(v)
				}
				result = append(result, tmp)
			}
		} else {
			for _, tResult := range tResults {
				var tmp = make(map[string]string)
				for _, tColumnValue := range tResult.ColumnValues {
					q, v := tColumnValue.GetQualifier(), tColumnValue.GetValue()
					tmp[string(q)] = string(v)
				}
				result = append(result, tmp)
			}
		}
		return result, err
	}
	return nil, err
}

// Commit a TPut to a table.
//
// Parameters:
//  - Table: the table to put data in
//  - Tput: the TPut to put
func (this *HBaseThriftAgent) Put(ctx context.Context, table string, tput *hbase.TPut) error {
	err := thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		return client.Put(ctx, []byte(table), tput)
	})
	return err
}

// Commit a TPut to a table.
//
// Parameters:
//  - Table: the table to put data in
//  - Tput: the TPut to put
func (this *HBaseThriftAgent) PutMultiple(ctx context.Context, table string, tput []*hbase.TPut) error {
	err := thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		return client.PutMultiple(ctx, []byte(table), tput)
	})
	return err
}

// Atomically checks if a row/family/qualifier value matches the expected
// value. If it does, it adds the TPut.
//
// @return true if the new put was executed, false otherwise
//
// Parameters:
//  - Table: to check in and put to
//  - Row: row to check
//  - Family: column family to check
//  - Qualifier: column qualifier to check
//  - Value: the expected value, if not provided the
// check is for the non-existence of the
// column in question
//  - Tput: the TPut to put if the check succeeds
func (this *HBaseThriftAgent) CheckAndPut(ctx context.Context, table, row, family, qualifier, value string, tput *hbase.TPut) error {
	err := thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		var err2 error
		_, err2 = client.CheckAndPut(ctx, []byte(table), []byte(row), []byte(family), []byte(qualifier), []byte(value), tput)
		return err2
	})
	return err
}

// Deletes as specified by the TDelete.
//
// Note: "delete" is a reserved keyword and cannot be used in Thrift
// thus the inconsistent naming scheme from the other functions.
//
// Parameters:
//  - Table: the table to delete from
//  - Tdelete: the TDelete to delete
func (this *HBaseThriftAgent) DeleteSingle(ctx context.Context, table string, tdelete *hbase.TDelete) error {
	err := thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		return client.DeleteSingle(ctx, []byte(table), tdelete)
	})
	return err
}

// Get results for the provided TScan object.
// This helper function opens a scanner, get the results and close the scanner.
//
// @return between zero and numRows TResults
//
// Parameters:
//  - Table: the table to get the Scanner for
//  - Tscan: the scan object to get a Scanner for
//  - NumRows: number of rows to return
func (this *HBaseThriftAgent) GetScannerResults(ctx context.Context, table string, tscan *hbase.TScan, numRows int32) (r []*hbase.TResult_, err error) {
	var tResults []*hbase.TResult_
	err = thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		var err2 error
		tResults, err2 = client.GetScannerResults(ctx, []byte(table), tscan, numRows)
		return err2
	})
	if nil == err {
		return tResults, err
	}
	return nil, err
}

// Get results for the provided TScan object.
// This helper function opens a scanner, get the results and close the scanner.
//
// @return between zero and numRows TResults
//
// Parameters:
//  - Table: the table to get the Scanner for
//  - Tscan: the scan object to get a Scanner for
//  - NumRows: number of rows to return
func (this *HBaseThriftAgent) GetScannerResultsAll(ctx context.Context, table string, cols []*hbase.TColumn, numRows int32, hasFamily bool) (r []map[string]string, err error) {
	var tResults []*hbase.TResult_
	err = thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		var err2 error
		tResults, err2 = client.GetScannerResults(ctx, []byte(table), &hbase.TScan{
			Columns: cols,
		}, numRows)
		return err2
	})
	if nil == err {
		var result = make([]map[string]string, consts.ZERO)
		if hasFamily {
			for _, tResult := range tResults {
				var tmp = make(map[string]string)
				for _, tColumnValue := range tResult.ColumnValues {
					q, f, v := tColumnValue.GetQualifier(), tColumnValue.GetFamily(), tColumnValue.GetValue()
					tmp[fmt.Sprintf("%s:%s", string(f), string(q))] = string(v)
				}
				result = append(result, tmp)
			}
		} else {
			for _, tResult := range tResults {
				var tmp = make(map[string]string)
				for _, tColumnValue := range tResult.ColumnValues {
					q, v := tColumnValue.GetQualifier(), tColumnValue.GetValue()
					tmp[string(q)] = string(v)
				}
				result = append(result, tmp)
			}
		}
		return result, err
	}
	return nil, err
}
