package funcs

import (
	"log"
	//"github.com/toolkits/pkg/logger"
	//"github.com/toolkits/pkg/nux"

	"github.com/didi/nightingale/src/dataobj"
	"github.com/didi/nightingale/src/modules/collector/core"
	"github.com/didi/nightingale/src/modules/collector/tools/mem"
)

func MemMetrics() []*dataobj.MetricValue {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Println(err)
		return nil
	}

	sm, err := mem.SwapMemory()
	if err != nil {
		log.Println(err)
		return nil
	}

	memFree := m.Free + m.Buffers + m.Cached
	memUsed := m.Total - memFree

	pmemFree := 0.0
	pmemUsed := 0.0
	if m.Total != 0 {
		pmemFree = float64(memFree) * 100.0 / float64(m.Total)
		pmemUsed = float64(memUsed) * 100.0 / float64(m.Total)
	}

	pswapFree := 0.0
	pswapUsed := 0.0
	if sm.Total != 0 {
		pswapFree = float64(sm.Free) * 100.0 / float64(sm.Total)
		pswapUsed = float64(sm.Used) * 100.0 / float64(sm.Total)
	}

	return []*dataobj.MetricValue{
		core.GaugeValue("mem.memtotal", m.Total),
		core.GaugeValue("mem.memused", memUsed),
		core.GaugeValue("mem.memfree", memFree),
		core.GaugeValue("mem.swaptotal", sm.Total),
		core.GaugeValue("mem.swapused", sm.Used),
		core.GaugeValue("mem.swapfree", sm.Free),
		core.GaugeValue("mem.memfree.percent", pmemFree),
		core.GaugeValue("mem.memused.percent", pmemUsed),
		core.GaugeValue("mem.swapfree.percent", pswapFree),
		core.GaugeValue("mem.swapused.percent", pswapUsed),
	}
}