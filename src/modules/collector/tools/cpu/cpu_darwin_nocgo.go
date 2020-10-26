// +build darwin
// +build !cgo

package cpu

import "github.com/didi/nightingale/src/modules/collector/tools/internal/common"

func perCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}

func allCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}
