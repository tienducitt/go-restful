package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tienducitt/go-restful/src/config"
	"github.com/tienducitt/go-restful/src/repository/mysql"
	"github.com/tienducitt/go-restful/src/router"
	"github.com/tienducitt/goconf"
	"github.com/urfave/cli"
)

func StartHttpServer(c *cli.Context) error {
	// load config
	var conf = &config.Config{}
	err := goconf.Load(conf, os.Getenv)
	if err != nil {
		return fmt.Errorf("could not load config, %v", err)
	}

	// init db connection
	db, err := mysql.Init(*conf)
	if err != nil {
		return err
	}

	// init repo & services
	userRepo := mysql.NewUserRepository(db)

	router := router.NewRouter(userRepo).Route()

	log.Println("Starting http server at port " + conf.Port)
	log.Fatal(http.ListenAndServe(":"+conf.Port, router))

	return nil
}
