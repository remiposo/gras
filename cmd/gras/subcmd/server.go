package subcmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/remiposo/gras/server"
	"github.com/urfave/cli/v2"
)

func NewServer() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "start server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"P"},
				Usage:   "port to be bound",
				EnvVars: []string{"GRAS_PORT"},
				Value:   8080,
			},
			&cli.StringFlag{
				Name:    "db-user",
				Aliases: []string{"du"},
				Usage:   "user of db",
				EnvVars: []string{"GRAS_DB_USER"},
				Value:   "gras",
			},
			&cli.StringFlag{
				Name:    "db-password",
				Aliases: []string{"dp"},
				Usage:   "password of db",
				EnvVars: []string{"GRAS_DB_PASSWORD"},
				Value:   "gras",
			},
			&cli.StringFlag{
				Name:    "db-host",
				Aliases: []string{"dh"},
				Usage:   "host of db",
				EnvVars: []string{"GRAS_DB_HOST"},
				Value:   "127.0.0.1",
			},
			&cli.IntFlag{
				Name:    "db-port",
				Aliases: []string{"dP"},
				Usage:   "port of db",
				EnvVars: []string{"GRAS_DB_PORT"},
				Value:   3306,
			},
			&cli.StringFlag{
				Name:    "db-name",
				Aliases: []string{"dn"},
				Usage:   "name of db",
				EnvVars: []string{"GRAS_DB_NAME"},
				Value:   "gras",
			},
			&cli.StringFlag{
				Name:    "kvs-host",
				Aliases: []string{"kh"},
				Usage:   "host of kvs",
				EnvVars: []string{"GRAS_KVS_HOST"},
				Value:   "127.0.0.1",
			},
			&cli.IntFlag{
				Name:    "kvs-port",
				Aliases: []string{"kP"},
				Usage:   "port of kvs",
				EnvVars: []string{"GRAS_KVS_PORT"},
				Value:   6379,
			},
		},
		Action: serverAction,
	}
}

func serverAction(cCtx *cli.Context) error {
	ctx, cancel := signal.NotifyContext(cCtx.Context, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// init db
	src := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cCtx.String("db-user"), cCtx.String("db-password"),
		cCtx.String("db-host"), cCtx.Int("db-port"),
		cCtx.String("db-name"),
	)
	db, err := sqlx.ConnectContext(ctx, "mysql", src)
	if err != nil {
		return err
	}
	defer db.Close()

	port := fmt.Sprintf(":%d", cCtx.Int("port"))
	fmt.Fprintf(cCtx.App.Writer, "start server on %v\n", port)

	return server.NewServer(port, db).Run(ctx)
}
