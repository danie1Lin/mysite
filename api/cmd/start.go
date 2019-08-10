package cmd

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"database/sql"
	"mysite/pkg/cache"
	"mysite/pkg/user/endpoints"
	"mysite/pkg/user/pb"
	"mysite/pkg/user/service"
	"mysite/pkg/user/transports"

	stdjwt "github.com/dgrijalva/jwt-go"

	"github.com/go-kit/kit/auth/jwt"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = log.With(logger,
			"TS", log.DefaultTimestampUTC,
			"Caller", log.DefaultCaller,
		)

		dbLogger := log.With(logger, "DB", GetSetting("db.name"))

		dataSourceStr := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			GetSetting("db.user"),
			GetSetting("db.password"),
			GetSetting("db.host"),
			GetSetting("db.port"),
			GetSetting("db.name"),
		)
		dbLogger.Log("url", dataSourceStr)
		m, err := migrate.New(
			"file://db/migrations", dataSourceStr)

		if err != nil {
			level.Error(dbLogger).Log("ERR", err)
		}

		m.Log = &MLog{dbLogger}

		err = m.Up()
		if err != nil {
			if err == migrate.ErrNoChange {
				level.Info(dbLogger).Log("migration", "no change")
			} else {
				level.Error(dbLogger).Log("ERR", err)
			}
		}

		db, err := sql.Open("postgres", dataSourceStr)
		{
			if err != nil {
				level.Error(logger).Log("exit", err)
				os.Exit(-1)
			}

			db.SetMaxOpenConns(10)
			db.SetMaxIdleConns(5)
		}

		userLogger := log.With(logger, "svs", "user")
		s := service.NewUserService(userLogger, db, cache.NewRedisCacheStore("localhost:6789", "", "session:"))

		kf := func(token *stdjwt.Token) (interface{}, error) { return []byte("SigningString"), nil }
		Endpoints := endpoints.EndPoints{
			LoginEndPoint:  jwt.NewParser(kf, stdjwt.SigningMethodHS256, jwt.StandardClaimsFactory)(endpoints.MakeLoginEndPoint(s)),
			SignUpEndPoint: endpoints.MakeSinUpEndPoint(s),
		}

		h := transports.NewHTTPHandler(context.WithValue(context.Background(), "JWTToken", ""), Endpoints, logger)
		go func() {
			err := http.ListenAndServe(":8080", h)
			if err != nil {
				level.Error(logger).Log("end", err)
			}
		}()

		grepcServer := grpc.NewServer()
		{
			pb.RegisterUserServiceServer(grepcServer, transports.NewGrpcServer(context.Background(), Endpoints, log.With(logger, "transport", "grpc")))
			listener, err := net.Listen("tcp", ":9000")
			if err != nil {
				level.Error(logger).Log("grpc", err)
			}
			reflection.Register(grepcServer)
			err = grepcServer.Serve(listener)
			if err != nil {
				level.Error(logger).Log("end", err)
			}
		}

	},
}

func GetSetting(name string) string {
	s := viper.GetString(name)
	if s == "" {
		panic(fmt.Sprintf("config %s is not found", name))
	}
	return s
}

type MLog struct {
	log.Logger
}

func (m *MLog) Printf(format string, v ...interface{}) {
	m.Log(v)
}

func (m *MLog) Verbose() bool {
	return true
}
