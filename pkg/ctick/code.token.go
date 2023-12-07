package ctick

import "time"

type CodeToken struct {
	Code        string
	ExpiredTime time.Time
}
