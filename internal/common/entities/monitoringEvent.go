package entities

type MonitoringEvent struct {
	OrderBy   string
	PageToken int64
	PageSize  int64
	SortDesc  bool
}
