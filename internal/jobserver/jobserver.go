package jobserver

import (
	"github.com/spf13/cobra"
	"job/internal/jobserver/controller"
	"job/internal/jobserver/manager"
	"job/internal/jobserver/server"
	"job/internal/jobserver/service"
	"job/internal/jobserver/store/mysql"
	"job/pkg/constant"
)

var (
	appName    = "jobserver"
	serverHost string
	serverPort string
	dbHost     string
	dbPort     string
	dbUsername string
	dbPassword string
	dbName     string
)

func NewJobServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          appName,
		Short:        "jobserver",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	cmd.Flags().StringVarP(&serverHost, "serverHost", "", "127.0.0.1", "--serverHost=127.0.0.1")
	cmd.Flags().StringVarP(&serverPort, "serverPort", "", "8080", "--serverPort=8080")
	cmd.Flags().StringVarP(&dbHost, "dbHost", "", "127.0.0.1", "--dbHost=127.0.0.1")
	cmd.Flags().StringVarP(&dbPort, "dbPort", "", "3306", "--dbPort=3306")
	cmd.Flags().StringVarP(&dbUsername, "dbUsername", "", "root", "--dbUsername=root")
	cmd.Flags().StringVarP(&dbPassword, "dbPassword", "", "root", "--dbPassword=root")
	cmd.Flags().StringVarP(&dbName, "dbName", "", "job", "--dbName=job")
	return cmd
}

func run() error {
	c := server.GenericAPIServerConfig{
		Host: serverHost,
		Port: serverPort,
	}
	s := c.New()
	db, err := mysql.GetMySQLInsOr(dbHost+":"+dbPort, dbUsername, dbPassword, dbName)
	m, err := manager.GetManagerOr()
	if err != nil {
		return err
	}
	store := mysql.NewStore(db)
	svc := service.NewService(store)
	tc := controller.NewTaskController(svc)
	cc := controller.NewCoreController()
	s.Register(constant.TASKCREATEURI, tc.Create)
	s.Register(constant.TASKUPDATEURI, tc.Update)
	s.Register(constant.TASKLISTURI, tc.List)
	s.Register(constant.LIVEURI, cc.LiveProbe)
	go m.RemoveDeadNode()
	s.Start()
	return nil
}
