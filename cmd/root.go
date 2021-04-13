package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
	log "github.com/sirupsen/logrus"

	"github.com/esnet/grafana-dashboard-manager/api"
	"github.com/esnet/grafana-dashboard-manager/config"
	"github.com/grafana-tools/sdk"
	"github.com/spf13/cobra"
)

var (
	configProvider config.Provider
	client         *sdk.Client
	tableObj       table.Writer
	debug          bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "generated code example",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//      Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	configProvider = config.Config()
	setupGrafanaClient()
	log.Debug("Creating output locations")
	dir := configProvider.GetString("env.output.datasources")
	os.MkdirAll(dir, 0755)
	dir = configProvider.GetString("env.output.dashboards")
	os.MkdirAll(dir, 0755)
	//Output Renderer
	tableObj = table.NewWriter()
	tableObj.SetOutputMirror(os.Stdout)
	tableObj.SetStyle(table.StyleLight)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	debug = configProvider.GetBool("global.debug")

}

func setupGrafanaClient() {
	grafana := config.GetGrafanaConfig()
	client = api.Login(grafana)

}
