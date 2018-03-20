package model

type Source struct {
	OpenConnectionCount      string `json:"OpenConnectionCount"`
	Type                     string `json:"Type"`
	AppendBatchAcceptedCount string `json:"AppendBatchAcceptedCount"`
	AppendBatchReceivedCount string `json:"AppendBatchReceivedCount"`
	EventAcceptedCount       string `json:"EventAcceptedCount"`
	StopTime                 string `json:"StopTime"`
	AppendReceivedCount      string `json:"AppendReceivedCount"`
	StartTime                string `json:"StartTime"`
	EventReceivedCount       string `json:"EventReceivedCount"`
	AppendAcceptedCount      string `json:"AppendAcceptedCount"`
}

type Channel struct {
	EventPutSuccessCount  string `json:"EventPutSuccessCount"`
	ChannelFillPercentage string `json:"ChannelFillPercentage"`
	Type                  string `json:"Type"`
	EventPutAttemptCount  string `json:"EventPutAttemptCount"`
	ChannelSize           string `json:"ChannelSize"`
	StopTime              string `json:"StopTime"`
	StartTime             string `json:"StartTime"`
	EventTakeSuccessCount string `json:"EventTakeSuccessCount"`
	ChannelCapacity       string `json:"ChannelCapacity"`
	EventTakeAttemptCount string `json:"EventTakeAttemptCount"`
}

type Sink struct {
	BatchCompleteCount     string `json:"BatchCompleteCount"`
	ConnectionFailedCount  string `json:"ConnectionFailedCount"`
	EventDrainAttemptCount string `json:"EventDrainAttemptCount"`
	ConnectionCreatedCount string `json:"ConnectionCreatedCount"`
	BatchEmptyCount        string `json:"BatchEmptyCount"`
	ConnectionClosedCount  string `json:"ConnectionClosedCount"`
	EventDrainSuccessCount string `json:"EventDrainSuccessCount"`
	BatchUnderflowCount    string `json:"BatchUnderflowCount"`
}
