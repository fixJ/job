package manager

import (
	"job/internal/jobserver/pkg/core"
	"sync"
	"time"
)

type Manager struct {
	liveNodes []core.NodeInfo
	deadTTL   time.Duration
	l         sync.Mutex
}

// 更新存活节点列表
func (m *Manager) UpdateLiveNode(ip, hostname string) {
	for _, node := range m.liveNodes {
		// 已经存在在列表中，则更新时间戳
		if node.IP == ip && node.Hostname == hostname {
			m.l.Lock()
			node.LastProbe = time.Now().Unix()
			m.l.Unlock()
		}
	}
	// 不存在列表中，可能是第一次上报，也可能之前被移除了再次上报
	m.l.Lock()
	m.liveNodes = append(m.liveNodes, core.NodeInfo{
		IP:        ip,
		Hostname:  hostname,
		CreatedAt: time.Now().Unix(),
		LastProbe: time.Now().Unix(),
	})
	m.l.Unlock()
}

// 定期移除长时间没报存活的节点
func (m *Manager) RemoveDeadNode() {
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
