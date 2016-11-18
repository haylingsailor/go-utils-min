package datetime

import (
	"encoding/json"
	"time"
)

// ITime implements json.Marshaler interface, and is a replacement for time.Time
// which is to be used solely in structs which are intended to be passed to
// json.Marshal(). Ordinarily, Go by default uses the RFC3339 Nano format, which
// includes floating point second precision. Use this instead when you want to
// write integer-second UTC timestamps.
type ITime time.Time

// MarshalJSON implements interface json.Marshaler
func (t ITime) MarshalJSON() ([]byte, error) {
	secs := time.Time(t).Unix()
	return json.Marshal(time.Unix(secs, 0))
}
