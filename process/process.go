package process

import (
	"log"
	"sync"
	"time"

	"github.com/mdh67899/openfalcon-flume-monitor/config"
	"github.com/mdh67899/openfalcon-flume-monitor/funcs"
	"github.com/mdh67899/openfalcon-flume-monitor/model"
)

type Program struct {
	*sync.RWMutex
	*sync.WaitGroup

	Cfg    *model.Cfg
	closed chan struct{}

	conn *funcs.SingleConnRpcClient
}

func NewProgram(filepath string) *Program {
	cfg := config.ParseConfig(filepath)

	return &Program{
		new(sync.RWMutex),
		new(sync.WaitGroup),

		cfg,
		make(chan struct{}),
		funcs.NewSingleConnRpcClient(cfg.Transfer.Addrs, time.Second*time.Duration(cfg.Transfer.Timeout)),
	}
}

func (this *Program) Process() {
	this.Lock()
	defer this.Unlock()

	for i := 0; i < len(this.Cfg.Instance); i++ {
		this.Add(1)

		go func(i int, needSend2Transfer bool, fn func([]*model.MetricValue)) {
			defer this.Done()
			process(this.Cfg.Instance[i], needSend2Transfer, fn, this.closed)
		}(i, this.Cfg.Transfer.Enabled, this.conn.SendMetrics)
	}
}

func process(flume model.FlumeConfig, needSend2Transfer bool, fn func([]*model.MetricValue), closed chan struct{}) {
	if !flume.Enabled {
		return
	}

	ticker := time.NewTicker(time.Second * time.Duration(flume.Step))

	for {
		select {
		case <-ticker.C:
			body, err := funcs.HttpGet(flume.MetricUrl)
			if err != nil {
				log.Println(err)
				continue
			}

			metrics, err := funcs.ToMetric(body, flume.Hostname, flume.Step, flume.Tags, time.Now().Unix())
			if err != nil {
				log.Println(err)
				continue
			}

			if needSend2Transfer {
				fn(metrics)
			}

		case <-closed:
			return
		}
	}
}

func (this *Program) Stop() {
	this.Lock()
	defer this.Unlock()
	close(this.closed)

	this.Wait()

	this.conn.Close()
}
