package main

import (
	"flag"
	"github.com/bufsnake/ldap-server/api"
	"github.com/bufsnake/ldap-server/config"
	"github.com/bufsnake/ldap-server/pkg/datas"
	ldap_server "github.com/bufsnake/ldap-server/pkg/ldap-server"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	terminal := config.Terminal{}
	flag.StringVar(&terminal.Sign, "sign", "a72720eb2bcedb900b83233720659e52", "api auth")
	flag.StringVar(&terminal.LDAPServer, "ldap", "0.0.0.0:10389", "ldap listen address")
	flag.StringVar(&terminal.HTTPServer, "http", "0.0.0.0:9200", "http listen address")
	flag.Parse()
	data := datas.NewData()
	api_ := api.NewAPI(data, terminal.Sign)
	ldap := ldap_server.NewLDAPServer(data, terminal.LDAPServer)
	go func() {
		err := ldap.Listen()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		defer ldap.Stop()
	}()
	engine := gin.Default()
	engine.GET("/verify", api_.Verify)
	err := engine.Run(terminal.HTTPServer)
	if err != nil {
		log.Println(err)
		return
	}
}
