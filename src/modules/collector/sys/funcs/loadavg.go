package funcs

import (
	"github.com/toolkits/pkg/logger"
	//"github.com/toolkits/pkg/nux"
	"github.com/didi/nightingale/src/modules/collector/tools/load"

	"github.com/didi/nightingale/src/dataobj"
	"github.com/didi/nightingale/src/modules/collector/core"
)

func LoadAvgMetrics() []*dataobj.MetricValue {
	load, err := load.Avg()
	if err != nil {
		logger.Error(err)
		return nil
	}

	return []*dataobj.MetricValue{
		core.GaugeValue("cpu.loadavg.1", load.Load1),
		core.GaugeValue("cpu.loadavg.5", load.Load5),
		core.GaugeValue("cpu.loadavg.15", load.Load15),
	}
}
