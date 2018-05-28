package sample

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

/*AsTimestamp returns a timestamp.Timestamp from a go time.Time structure
* This is mostly a helper for the lazy*/
func AsTimestamp(whence time.Time) *timestamp.Timestamp {
	s, ns := whence.Unix(), whence.UnixNano()
	ns = ns - s*int64(time.Second)
	return &timestamp.Timestamp{Seconds: s, Nanos: int32(ns)}
}

/*AsTime 'converts' a timestamp.Timestamp to go's time.Time structure.
*
* If whence is nil, it returns an empty time.Time
 */
func AsTime(whence *timestamp.Timestamp) time.Time {
	if whence == nil {
		return time.Time{}
	}
	return time.Unix(whence.Seconds, int64(whence.Nanos))
}
