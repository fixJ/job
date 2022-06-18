package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"job/internal/joblet/pkg/types"
	"job/pkg/client"
	"job/pkg/constant"
	"job/pkg/utils"
	"os/exec"
	"sync"
	"time"
)

var (
	m    *LetManager
	once sync.Once
)

type LetManager struct {
	Server      string
	IP          string
	RunningTask []types.TaskInfo
	l           sync.Mutex
}

func NewLetManager(server, ip string) (*LetManager, error) {
	once.Do(func() {
		m = &LetManager{
			IP:     ip,
			Server: server,
		}
	})
	if m == nil {
		return nil, fmt.Errorf("get let manager failed")
	}
	return m, nil
}

func (m *LetManager) ListTasksAndRun() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		d, _ := json.Marshal(types.ListReq{Target: m.IP})
		resp, err := client.DoPost(utils.URL(m.Server, constant.TASKLISTURI), d)
		//fmt.Println(string(resp))
		if err != nil {
			continue
		}
		var listObj types.ListResp
		_ = json.Unmarshal(resp, &listObj)
		tasks := listObj.Data
		for _, task := range tasks {
			if !m.IsRunning(task) {
				m.l.Lock()
				m.RunningTask = append(m.RunningTask, task)
				m.l.Unlock()
				go func(task types.TaskInfo) {
					err := m.runCommand(task.Command)
					if err != nil {
						m.removeFromRunning(task)
						return
					}
					d, _ := json.Marshal(types.UpdateReq{
						ID:     task.ID,
						Status: 1,
					})
					_, err = client.DoPost(utils.URL(m.Server, constant.TASKUPDATEURI), d)
					m.removeFromRunning(task)
				}(task)
			}
		}
	}
}

func (m *LetManager) runCommand(cmd string) error {
	in := bytes.NewBuffer(nil)
	c := exec.Command("sh")
	c.Stdin = in
	in.WriteString(cmd)
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (m *LetManager) Live() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		d, _ := json.Marshal(types.LiveReq{Target: m.IP})
		_, err := client.DoPost(utils.URL(m.Server, constant.LIVEURI), d)
		if err != nil {
			fmt.Println("live failed", err)
			continue
		}
		fmt.Println("live success")
	}
}

func (m *LetManager) IsRunning(task types.TaskInfo) bool {
	for _, rt := range m.RunningTask {
		if rt.ID == task.ID {
			return true
		}
	}
	return false
}

func (m *LetManager) removeFromRunning(task types.TaskInfo) {
	for i, t := range m.RunningTask {
		if t.ID == task.ID {
			m.l.Lock()
			m.RunningTask = append(m.RunningTask[:i], m.RunningTask[i+1:]...)
			m.l.Unlock()
		}
	}
}
