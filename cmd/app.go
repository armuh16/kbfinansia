package cmd

import (
	"context"
	"github.com/armuh16/kbfinansia/module"

	"github.com/armuh16/kbfinansia/database/mysql"
	//"github.com/armuh16/kbfinansia/module"
	"github.com/armuh16/kbfinansia/package/logger"
	"github.com/armuh16/kbfinansia/router"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// appCmd represents the app command
var app = &cobra.Command{
	Use:   "start",
	Short: "Running service",

	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			fx.Provide(router.NewRouter),
			fx.Provide(mysql.NewMysql),
			fx.Provide(logger.NewLogRus),
			module.BundleRepository,
			module.BundleLogic,
			module.BundleRoute,
			fx.Invoke(registerHooks),
		).Run()
	},
}

func init() {
	rootCmd.AddCommand(app)
}

func registerHooks(
	lifecycle fx.Lifecycle,
	e *router.Router,
	msql *mysql.DB,
	log *logger.LogRus) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go e.Start(":8090")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				if err := e.Shutdown(ctx); err != nil {
					log.Fatal(err.Error())
					return err
				}
				defer func() {
					if err := msql.Sql.Close(); err != nil {
						log.Fatal(err.Error())
					}
				}()
				return nil
			},
		},
	)
}
