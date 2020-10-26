// +build darwin
// +build !cgo

package host

import "github.com/didi/nightingale/src/modules/collector/tools/internal/common"

func SensorsTemperatures() ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
