package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tienducitt/go-restful/src/config"
)

// Init init database connection
func Init(config config.Config) (*gorm.DB, error) {
	connStr := ConnStr(config)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("Could not init database connection %v, conn string: %s", err, connStr)
	}

	return db, nil
}

func ConnStr(config config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBName)
}
