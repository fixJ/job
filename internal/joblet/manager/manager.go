package manager

import (
	"encoding/json"
	"fmt"
	"job/internal/joblet/pkg/types"
	"job/pkg/client"
	"sync"
	"time"
)

var (
	m    *LetManager
	once sync.Once
)

type LetManager struct {
	IP          string
	RunningTask []types.TaskInfo
	l           sync.Mutex
}

func NewLetManager() (*LetManager, error) {
	once.Do(func() {
		m = &LetManager{}
	})
	if m == nil {
		return nil, fmt.Errorf("get let manager failed")
	}
	return m, nil
}

func (m *LetManager) ListTasksAndRun() {

}

func (m *LetManager) runCommand(cmd string) error {
	return nil
}

func (m *LetManager) Live(server string) {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		d, _ := json.Marshal(types.LiveReq{Target: m.IP})
		_, err := client.DoPost(server, d)
		if err != nil {
			fmt.Println("live failed")
		}
	}
}
