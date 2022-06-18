package joblet

import (
	"github.com/spf13/cobra"
	"job/internal/joblet/manager"
	"os"
	"os/signal"
	"syscall"
)

var (
	appName = "joblet"
	server  string
	ip      string
)

func NewJobletCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          appName,
		Short:        "joblet",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	cmd.Flags().StringVarP(&server, "server", "", "", "--server=192.168.0.1:8080")
	cmd.Flags().StringVarP(&ip, "ip", "", "", "--ip=192.168.0.1")
	return cmd
}

// 定时上报存活
// 定时List task, 运行任务
func run() error {
	done := make(chan os.Signal)
	m, err := manager.NewLetManager(server, ip)
	if err != nil {
		return err
	}
	go m.Live()
	go m.ListTasksAndRun()
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	return nil
}
