package funcs

import (
	"log"
	"math"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
	"time"

	"github.com/mdh67899/openfalcon-flume-monitor/model"
)

type SingleConnRpcClient struct {
	sync.Mutex
	rpcClient *rpc.Client
	RpcServer string
	Timeout   time.Duration
}

func NewSingleConnRpcClient(addr string, timeout time.Duration) *SingleConnRpcClient {
	conn := SingleConnRpcClient{
		RpcServer: addr,
		Timeout:   timeout,
	}
	return &conn
}

func (this *SingleConnRpcClient) Close() {
	this.Lock()
	defer this.Unlock()

	if this.rpcClient != nil {
		this.rpcClient.Close()
		this.rpcClient = nil
	}
}

func JsonRpcClient(network, address string, timeout time.Duration) (*rpc.Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), err
}

func (this *SingleConnRpcClient) serverConn() error {
	if this.rpcClient != nil {
		return nil
	}

	var err error
	var retry int = 1

	for {
		if this.rpcClient != nil {
			return nil
		}

		this.rpcClient, err = JsonRpcClient("tcp", this.RpcServer, this.Timeout)
		if err != nil {
			log.Printf("dial %s fail: %v", this.RpcServer, err)
			if retry > 3 {
				return err
			}
			time.Sleep(time.Duration(math.Pow(2.0, float64(retry))) * time.Second)
			retry++
			continue
		}
		return err
	}
}

func (this *SingleConnRpcClient) call(method string, args interface{}, reply interface{}) error {
	err := this.serverConn()
	if err != nil {
		return err
	}

	timeout := time.Duration(3 * time.Second)
	done := make(chan error, 1)

	go func() {
		err := this.rpcClient.Call(method, args, reply)
		done <- err
	}()

	select {
	case <-time.After(timeout):
		log.Printf("[WARN] rpc call timeout %v => %v", this.rpcClient, this.RpcServer)
		this.Close()
	case err := <-done:
		if err != nil {
			this.Close()
			return err
		}
	}

	return nil
}

func (this *SingleConnRpcClient) SendMetrics(metrics []*model.MetricValue) {
	this.Lock()
	defer this.Unlock()

	log.Println("=> <Total =", len(metrics), ", MetricValue:", metrics[0].String())

	resp := &model.TransferResponse{}
	err := this.call("Transfer.Update", metrics, resp)
	if err != nil {
		log.Println("call Transfer.Update fail", this.RpcServer, err)
		return
	}

	log.Println(resp.String())
}
