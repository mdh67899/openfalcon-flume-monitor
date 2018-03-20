package funcs

import (
	"encoding/json"
	"strings"

	"github.com/mdh67899/openfalcon-flume-monitor/model"
)

const (
	Guage   string = "GAUGE"
	Counter string = "COUNTER"
)

var (
	Item_Ignore map[string]struct{} = map[string]struct{}{
		"Type":            struct{}{},
		"StartTime":       struct{}{},
		"StopTime":        struct{}{},
		"ChannelCapacity": struct{}{},
	}

	Counter_Type map[string]struct{} = map[string]struct{}{
		"AppendBatchReceivedCount": struct{}{},
		"AppendBatchAcceptedCount": struct{}{},
		"AppendReceivedCount":      struct{}{},
		"AppendAcceptedCount":      struct{}{},
		"EventReceivedCount":       struct{}{},
		"EventAcceptedCount":       struct{}{},
		"EventPutSuccessCount":     struct{}{},
		"EventPutAttemptCount":     struct{}{},
		"EventTakeSuccessCount":    struct{}{},
		"EventTakeAttemptCount":    struct{}{},
		"BatchCompleteCount":       struct{}{},
		"ConnectionFailedCount":    struct{}{},
		"EventDrainAttemptCount":   struct{}{},
		"ConnectionCreatedCount":   struct{}{},
		"BatchEmptyCount":          struct{}{},
		"ConnectionClosedCount":    struct{}{},
		"EventDrainSuccessCount":   struct{}{},
		"BatchUnderflowCount":      struct{}{},
	}

	Prefixs map[string]struct{} = map[string]struct{}{
		"SOURCE.":  struct{}{},
		"CHANNEL.": struct{}{},
		"SINK.":    struct{}{},
	}
)

func ToMetric(body []byte, hostname string, step int64, tags string, ts int64) ([]*model.MetricValue, error) {
	metrics := []*model.MetricValue{}

	var data map[string]map[string]string

	err := json.Unmarshal(body, &data)
	if err != nil {
		return []*model.MetricValue{}, err
	}

	for k, _ := range data {
		prefix := k
		for p, _ := range Prefixs {
			prefix = strings.Trim(prefix, p)
		}

		for k1, _ := range data[k] {
			if _, e := Item_Ignore[k1]; e {
				continue
			}

			counter_type := Guage
			if _, exist := Counter_Type[k1]; exist {
				counter_type = Counter
			}

			metrics = append(metrics,
				&model.MetricValue{
					Endpoint:  hostname,
					Metric:    prefix + "_" + k1,
					Value:     data[k][k1],
					Step:      step,
					Type:      counter_type,
					Tags:      tags,
					Timestamp: ts,
				},
			)
		}
	}

	return metrics, nil
}
