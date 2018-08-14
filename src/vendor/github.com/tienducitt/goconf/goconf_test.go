package goconf_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/tienducitt/goconf"
)

type TestConf struct {
	Env       string         `conf:"test_env" required:"true" default:"tada" desc:"Current environment"`
	Version   string         `conf:"test_version" default:"1.0.0" desc:"Version"`
	TestSlice []int          `conf:"test_slice"`
	TestMap   map[string]int `conf:"test_map"`
	Test      int            `conf:"test" default:"1"`
}

var Conf = &TestConf{}

func Test(t *testing.T) {
	err := goconf.Load(Conf, os.Getenv)

	if err != nil {
		panic(err)
	}

	fmt.Print(*Conf)
}
