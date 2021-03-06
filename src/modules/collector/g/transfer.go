package g

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/didi/nightingale/src/dataobj"
)

var (
	TransferClientsLock *sync.RWMutex                   = new(sync.RWMutex)
	TransferClients     map[string]*SingleConnRpcClient = map[string]*SingleConnRpcClient{}
)

func SendMetrics(metrics []*dataobj.MetricValue, resp *dataobj.TransferResponse) {
	rand.Seed(time.Now().UnixNano())
	//log.Print("SendMetrics:")
	//log.Print(len(Config().Transfer.Addrs))
	for _, i := range rand.Perm(len(Config().Transfer.Addrs)) {
		addr := Config().Transfer.Addrs[i]
		//log.Print(addr)
		if _, ok := TransferClients[addr]; !ok {
			initTransferClient(addr)
		}
		if updateMetrics(addr, metrics, resp) {
			break
		}
	}
}

func initTransferClient(addr string) {
	TransferClientsLock.Lock()
	defer TransferClientsLock.Unlock()
	TransferClients[addr] = &SingleConnRpcClient{
		RpcServer: addr,
		Timeout:   time.Duration(Config().Transfer.Timeout) * time.Millisecond,
	}
}

func updateMetrics(addr string, metrics []*dataobj.MetricValue, resp *dataobj.TransferResponse) bool {
	TransferClientsLock.RLock()
	defer TransferClientsLock.RUnlock()
	err := TransferClients[addr].Call("Transfer.Update", metrics, resp)
	if err != nil {
		log.Println("call Transfer.Update fail", addr, err)
		return false
	}
	return true
}
