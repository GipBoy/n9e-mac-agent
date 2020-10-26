// +build darwin
// +build !cgo

package disk

import "github.com/didi/nightingale/src/modules/collector/tools/internal/common"

func IOCounters(names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
