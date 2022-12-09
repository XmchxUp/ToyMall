package main

import (
	"xm-mall/conf"
	"xm-mall/routes"
)

func main() {
	//KEY: Ek1+Ep1==Ek2+Ep2
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
