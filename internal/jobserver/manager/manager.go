package manager

import (
	"fmt"
	"job/internal/jobserver/pkg/types"
	"sync"
	"time"
)

var (
	m    *ServerManager
	once sync.Once
)

type ServerManager struct {
	liveNodes []types.NodeInfo
	deadTTL   time.Duration
	l         sync.Mutex
}

func GetManagerOr() (*ServerManager, error) {
	once.Do(func() {
		m = &ServerManager{
			deadTTL: 5 * time.Second,
		}
	})
	if m == nil {
		return nil, fmt.Errorf("manager is nil")
	}
	return m, nil
}

// 更新存活节点列表
func (m *ServerManager) UpdateLiveNode(ip string) {
	for _, node := range m.liveNodes {
		// 已经存在在列表中，则更新时间戳
		if node.IP == ip {
			m.l.Lock()
			node.LastProbe = time.Now().Unix()
			m.l.Unlock()
		}
	}
	// 不存在列表中，可能是第一次上报，也可能之前被移除了再次上报
	m.l.Lock()
	m.liveNodes = append(m.liveNodes, types.NodeInfo{
		IP:        ip,
		CreatedAt: time.Now().Unix(),
		LastProbe: time.Now().Unix(),
	})
	m.l.Unlock()
}

// 定期移除长时间没报存活的节点
func (m *ServerManager) RemoveDeadNode() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		for i, node := range m.liveNodes {
			if time.Now().Unix()-node.LastProbe > int64(m.deadTTL.Seconds()) {
				m.l.Lock()
				m.liveNodes = append(m.liveNodes[:i], m.liveNodes[i+1:]...)
				m.l.Unlock()
			}
		}
	}
}

func (m *ServerManager) IsLive(target string) bool {
	for _, n := range m.liveNodes {
		if target == n.IP {
			return true
		}
	}
	return false
}
