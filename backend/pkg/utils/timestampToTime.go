package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func TimestampToTime(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	t := ts.AsTime()
	return &t
}
