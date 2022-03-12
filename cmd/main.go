package main

import (
	v1 "gee-example/api/v1"
	"github.com/xylong/gee"
)

func main() {
	gee.
		Ignite().
		Bean(new(gee.MysqlFactory).Create(gee.G)).
		Mount("v1", v1.NewUser()).
		Launch()
}
