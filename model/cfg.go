package model

type FlumeConfig struct {
	Enabled   bool   `json:"enabled"`
	Hostname  string `json:"hostname"`
	Step      int64  `json:"step"`
	MetricUrl string `json:"metric_url"`
	Tags      string `json:"tags"`
}

type TransferConfig struct {
	Enabled bool   `json:"enabled"`
	Addrs   string `json:"addrs"`
	Timeout int64  `json:"timeout"`
}

type Cfg struct {
	Instance []FlumeConfig  `json:"instance"`
	Transfer TransferConfig `json:"transfer"`
}
