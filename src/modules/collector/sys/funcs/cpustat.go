package funcs

import (
	//"fmt"
	"sync"
	"time"

	"github.com/toolkits/pkg/logger"
	//"github.com/toolkits/pkg/nux"

	"github.com/didi/nightingale/src/dataobj"
	"github.com/didi/nightingale/src/modules/collector/core"
	"github.com/didi/nightingale/src/modules/collector/tools/cpu"
)

const (
	historyCount int = 2
)

var (
	procStatHistory [historyCount]*cpu.TimesStat
	psLock    = new(sync.RWMutex)
)

func PrepareCpuStat() {
	d := time.Duration(3) * time.Second
	for {
		err := UpdateCpuStat()
		if err != nil {
			logger.Error("update cpu stat fail", err)
		}
		time.Sleep(d)
	}
}

type CpuStats struct {
	User    float64
	Nice    float64
	System  float64
	Idle    float64
	Iowait  float64
	Irq     float64
	SoftIrq float64
	Steal   float64
	Guest   float64
	Total   float64
}

func UpdateCpuStat() error {
	ps, err := cpu.Times(false)
	if err != nil {
		return err
	}
	//log.Print(ps)
	psLock.Lock()
	defer psLock.Unlock()
	for i := historyCount - 1; i > 0; i-- {
		procStatHistory[i] = procStatHistory[i-1]
	}

	procStatHistory[0] = &ps[0]
	return nil
}

func deltaTotal() uint64 {
	if procStatHistory[1] == nil {
		return 0
	}
	return (uint64)(procStatHistory[0].Total() - procStatHistory[1].Total())
}

func CpuIdle() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].Idle-procStatHistory[1].Idle) * invQuotient
}

func CpuUser() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].User-procStatHistory[1].User) * invQuotient
}

func CpuNice() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].Nice-procStatHistory[1].Nice) * invQuotient
}

func CpuSystem() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].System-procStatHistory[1].System) * invQuotient
}



func CpuIowait() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].Iowait-procStatHistory[1].Iowait) * invQuotient
}

func CpuIrq() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].Irq-procStatHistory[1].Irq) * invQuotient
}

func CpuSoftIrq() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].Softirq-procStatHistory[1].Softirq) * invQuotient
}

func CpuSteal() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].Steal-procStatHistory[1].Steal) * invQuotient
}

func CpuGuest() float64 {
	psLock.RLock()
	defer psLock.RUnlock()
	dt := deltaTotal()
	if dt == 0 {
		return 0.0
	}
	invQuotient := 100.00 / float64(dt)
	return float64(procStatHistory[0].Guest-procStatHistory[1].Guest) * invQuotient
}


func CpuPrepared() bool {
	psLock.RLock()
	defer psLock.RUnlock()
	return procStatHistory[1] != nil
}

func CpuMetrics() []*dataobj.MetricValue {
	if !CpuPrepared() {
		return []*dataobj.MetricValue{}
	}

	var ret []*dataobj.MetricValue

	//cpuIdleVal := CpuIdle()
	//idle := core.GaugeValue("cpu.idle", cpuIdleVal)
	//util := core.GaugeValue("cpu.util", 100.0-cpuIdleVal)
	user := core.GaugeValue("cpu.user", CpuUser())
	//system := core.GaugeValue("cpu.sys", CpuSystem())
	//nice := core.GaugeValue("cpu.nice", CpuNice())
	//iowait := core.GaugeValue("cpu.iowait", CpuIowait())
	//irq := core.GaugeValue("cpu.irq", CpuIrq())
	//softirq := core.GaugeValue("cpu.softirq", CpuSoftIrq())
	//steal := core.GaugeValue("cpu.steal", CpuSteal())
	//guest := core.GaugeValue("cpu.guest", CpuGuest())
	//switches := core.GaugeValue("cpu.switches", CpuContentSwitches())
	ret = []*dataobj.MetricValue{user}

	return ret
}
