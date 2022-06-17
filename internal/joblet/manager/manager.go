package manager

import (
	"encoding/json"
	"fmt"
	"job/internal/joblet/pkg/types"
	"job/pkg/client"
	"sync"
)

type LetManager struct {
	IP          string
	RunningTask []types.TaskInfo
	l           sync.Mutex
}

func (m *LetManager) ListTasksAndRun() {

}

func (m *LetManager) runCommand(cmd string) error {
	return nil
}

func (m *LetManager) Live(server string) {
	d, _ := json.Marshal(types.LiveReq{Target: m.IP})
	_, err := client.DoPost(server, d)
	if err != nil {
		fmt.Println("live failed")
	}
}
