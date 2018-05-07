package file

import "time"

type FileInfo struct {
	Name string
	Path string
	Size int64
	MTime time.Time
	CTime time.Duration
}

