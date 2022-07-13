package hbases_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	_ "net/http/pprof"
	"sync/atomic"
	"testing"
	"time"

	"github.com/xiaoniudongman/go-utils/dbs/hbases"
	"github.com/xiaoniudongman/go-utils/dbs/hbases/gen-go/hbase"

	"github.com/apache/thrift/lib/go/thrift"
)

var (
	addr            = "http://ld-bp17y8n3j6f45p944-proxy-hbaseue.hbaseue.rds.aliyuncs.com:9190"
	connTimeout     = time.Second * 2
	idleTimeout     = time.Second * 120
	timeout         = time.Second * 10
	maxConn         = int32(100)
	service         = "com.czl.api.ApiService1$Client"
	bct             = context.Background()
	delay           int64 // 单位微妙
	successCount    = int64(0)
	failCount       = int64(0)
	thriftPoolAgent *hbases.ThriftPoolAgent
)

// 初始化Thrift连接池代理
func init() {
	config := &hbases.ThriftPoolConfig{
		Addr:        addr,
		MaxConn:     maxConn,
		ConnTimeout: connTimeout,
		IdleTimeout: idleTimeout,
		Timeout:     timeout,
	}
	thriftPool := hbases.NewThriftPool(config, thriftDial, closeThriftClient)
	thriftPoolAgent = new(hbases.ThriftPoolAgent)
	thriftPoolAgent.Init(thriftPool)
}

func thriftDial(addr string, connTimeout time.Duration) (*hbases.IdleClient, error) {
	//socket, err := thrift.NewTSocketTimeout(addr, connTimeout)
	//if err != nil {
	//	return nil, err
	//}
	//transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	//transport, err := transportFactory.GetTransport(socket)
	transport, err := thrift.NewTHttpClient(addr)
	if err != nil {
		return nil, err
	}
	if err := transport.Open(); err != nil {
		return nil, err
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//iprot := thrift.NewTMultiplexedProtocol(protocolFactory.GetProtocol(transport), service)
	//oprot := thrift.NewTMultiplexedProtocol(protocolFactory.GetProtocol(transport), service)

	// 设置用户名密码
	httClient := transport.(*thrift.THttpClient)
	httClient.SetHeader("ACCESSKEYID", "root")
	httClient.SetHeader("ACCESSSIGNATURE", "root")
	//hb._client = hbase.NewTHBaseServiceClientFactory(trans, protocolFactory)

	//client := hbase.new(transport, iprot, oprot)
	return &hbases.IdleClient{
		Transport: transport,
		RawClient: hbase.NewTHBaseServiceClientFactory(transport, protocolFactory),
	}, nil
}

func closeThriftClient(c *hbases.IdleClient) error {
	if c == nil {
		return nil
	}
	return c.Transport.Close()
}

func TestCCCCC(t *testing.T) {
	//go startPprof()
	for i := 0; i < 1; i++ {
		go start()
	}
	time.Sleep(time.Second * 600)
	avgQps := float64(successCount) / float64(600)
	avgDelay := float64(delay) / float64(successCount) / 1000
	log.Println(fmt.Sprintf("总运行时间：600s, 并发协程数：100，平均吞吐量：%v，平均延迟（ms）：%v，总成功数：%d，总失败数：%d",
		avgQps, avgDelay, successCount, failCount))
}

//func startPprof() {
//	http.ListenAndServe("0.0.0.0:9999", nil)
//}

func start() {
	//for {
	if err := dialApi(); err != nil {
		log.Println(err)
	} else {
		log.Println("ok")
	}
	//}
}

func dialApi() error {
	st := time.Now()
	err := thriftPoolAgent.Do(func(rawClient interface{}) error {
		client, ok := rawClient.(*hbase.THBaseServiceClient)
		if !ok {
			return errors.New("unknown client type")
		}
		var err2 error
		resp, err2 := client.Get(bct, []byte("recommend_samh:alg_nrt_ar_fpgrowth"), &hbase.TGet{Row: []byte("5c341a7dda910511:comic")})
		fmt.Println(resp, err2)
		return err2
	})
	if err != nil {
		fail()
	} else {
		success(st)
	}
	return nil
}

func success(st time.Time) {
	atomic.AddInt64(&delay, time.Now().Sub(st).Microseconds())
	atomic.AddInt64(&successCount, 1)
}

func fail() {
	atomic.AddInt64(&failCount, 1)
}
