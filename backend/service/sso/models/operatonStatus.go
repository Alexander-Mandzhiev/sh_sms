package models

import "time"

type OperationStatus struct {
	Success       bool
	Message       string
	OperationTime time.Time
}
