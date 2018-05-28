/*sample is a package that contains structures that are time sequenced
* samples
*
* This is mostly a simple test of utilizing protocol buffers to create useful go
* structures*/
package sample

//go:generate protoc -I=. --go_out=. ./sample.proto

import (
	"fmt"
	"time"
)

/*NewSample returns a sample with timestamps taken as sampled now.*/
func New(raw []byte) *Sample {
	now := time.Now().UTC()
	s := &Sample{
		RecordTime: AsTimestamp(now),
		SampleTime: AsTimestamp(now),
		Raw:        raw,
		Tags:       []string{},
		Values:     map[string]float64{},
	}
	return s
}

/*Dump dumps the sample with values in a human readable format*/
func (s *Sample) Contents() string {
	r := fmt.Sprintf("%s:", AsTime(s.SampleTime).Format(time.RFC3339Nano))
	if len(s.Values) > 0 {
		for name, val := range s.Values {
			r += fmt.Sprintf(" %s=%4.4f", name, val)
		}
	}
	if len(s.Tags) > 0 {
		for _, tag := range s.Tags {
			r += fmt.Sprintf(" @%s", tag)
		}
		r += "}"
	}
	return r + "\n"
}

/*AddTags adds some set of tags to the taglist.*/
func (s *Sample) AddTags(tags ...string) {
	if s.Tags == nil {
		s.Tags = make([]string, 0)
	}
	s.Tags = append(s.Tags, tags...)
}
