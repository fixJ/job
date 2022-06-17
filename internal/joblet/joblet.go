package joblet

import "github.com/spf13/cobra"

var (
	appName = "joblet"
	server  string
)

func NewJobletCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   appName,
		Short: "A good Go practical project",
		Long: `A good Go practical project, used to create user with basic information.
Find more goserver information at:
    https://github.com/marmotedu/goserver/blob/master/docs/master/goserver.md`,

		// stop printing usage when the command errors
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	cmd.Flags().StringVarP(&server, "server", "", "", "server=192.168.0.1:8080")
	return cmd
}

// 定时上报存活
// 定时List task, 运行任务
func run() error {
	return nil
}
