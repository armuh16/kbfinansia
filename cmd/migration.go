package cmd

import (
	"github.com/armuh16/kbfinansia/database/mysql"
	"github.com/armuh16/kbfinansia/package/logger"

	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

var migration = &cobra.Command{
	Use:   "migration",
	Short: "Migration database",

	Run: func(cmd *cobra.Command, args []string) {
		log := logger.NewLogRus()
		mysql := mysql.NewMysql(log)
		if len(args) == 0 {
			log.Errorf("Please insert argument up, down or reset")
			return
		} else if args[0] == "up" {
			if err := Up(mysql); err != nil {
				log.Error(err)
				return
			}
		} else if args[0] == "down" {
			if err := Down(mysql); err != nil {
				log.Error(err)
				return
			}
		} else if args[0] == "reset" {
			if err := Reset(mysql); err != nil {
				log.Error(err)
				return
			}
		} else {
			log.Errorf("Migration argument not found")
			return
		}
	},
}

func Up(db *mysql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}
	if err := goose.Up(db.Sql, "sql"); err != nil {
		return err
	}
	return nil
}

func Down(db *mysql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Up(db.Sql, "sql"); err != nil {
		return err
	}
	return nil
}

func Reset(db *mysql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Reset(db.Sql, "sql"); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(migration)
}
