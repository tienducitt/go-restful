package migrate

import (
	"fmt"
	"log"
	"os"

	"github.com/tienducitt/go-restful/src/config"
	"github.com/tienducitt/go-restful/src/model"
	"github.com/tienducitt/go-restful/src/repository/mysql"
	"github.com/tienducitt/goconf"

	"github.com/urfave/cli"
)

//TODO: need improvement using sql scripts
func Migrate(c *cli.Context) (err error) {
	log.Printf("Run db migration")

	var conf = &config.Config{}
	err = goconf.Load(conf, os.Getenv)
	if err != nil {
		return fmt.Errorf("could not load config, %v", err)
	}

	db, err := mysql.Init(*conf)
	if err != nil {
		return err
	}
	db.AutoMigrate(&model.User{})

	return nil
}
