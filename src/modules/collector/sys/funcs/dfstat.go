package funcs

import (
	"fmt"
	//"strings"
	"log"

	"github.com/didi/nightingale/src/dataobj"
	"github.com/didi/nightingale/src/modules/collector/core"
	//"github.com/didi/nightingale/src/modules/collector/sys"
	"github.com/didi/nightingale/src/modules/collector/tools/disk"

	//"github.com/toolkits/pkg/logger"
	//"github.com/toolkits/pkg/nux"
	//"github.com/toolkits/pkg/slice"
)

func DeviceMetrics() (L []*dataobj.MetricValue) {

	var diskTotal uint64 = 0
	var diskUsed uint64 = 0

	var du, err = disk.Partitions(false)
	if err != nil {
		log.Println(err)
	}
	for _, ds := range du {

		var dusage, err = disk.Usage(ds.Mountpoint)
		if err != nil {
			log.Println(err)
		}

		diskTotal += dusage.Total
		diskUsed += dusage.Used

		tags := fmt.Sprintf("mount=%s,fstype=%s", dusage.Path, dusage.Fstype)
		L = append(L, core.GaugeValue("df.bytes.total", dusage.Total, tags))
		L = append(L, core.GaugeValue("df.bytes.used", dusage.Used, tags))
		L = append(L, core.GaugeValue("df.bytes.free", dusage.Free, tags))
		L = append(L, core.GaugeValue("df.bytes.used.percent", dusage.UsedPercent, tags))
		L = append(L, core.GaugeValue("df.bytes.free.percent", 100-dusage.UsedPercent, tags))
		L = append(L, core.GaugeValue("df.inodes.total", dusage.InodesTotal, tags))
		L = append(L, core.GaugeValue("df.inodes.used", dusage.InodesUsed, tags))
		L = append(L, core.GaugeValue("df.inodes.free", dusage.InodesFree, tags))
		L = append(L, core.GaugeValue("df.inodes.used.percent", dusage.InodesUsedPercent, tags))
		L = append(L, core.GaugeValue("df.inodes.free.percent", 100-dusage.InodesUsedPercent, tags))
	}

	if len(L) > 0 && diskTotal > 0 {
		L = append(L, core.GaugeValue("df.statistics.total", float64(diskTotal)))
		L = append(L, core.GaugeValue("df.statistics.used", float64(diskUsed)))
		L = append(L, core.GaugeValue("df.statistics.used.percent", float64(diskUsed)*100.0/float64(diskTotal)))
	}

	return
}