package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/apcera/termtables"
	"github.com/fatih/structs"
	"github.com/gosuri/uitable"
	"github.com/iris-contrib/color"
)

type AuthServer struct {
	SID     string
	AppName string
	Env     string
	Ip      string
	Role    string
}

var authServers = []AuthServer{
	{"a1", "vlbhost", "test", "10.1.101.22", "www"},
	{"a2", "vlb-tengine-test", "dev", "10.1.101.103", "www"},
}

func main() {
	table := uitable.New()
	table.MaxColWidth = 100

	info := color.New(nil, color.FgRed).SprintFunc()
	table.AddRow(info("[a[n]]"), "APPNAME", "ENV", "IP", "ROLE")
	for _, authServer := range authServers {
		table.AddRow(info(authServer.SID), authServer.AppName, authServer.Env, authServer.Ip, authServer.Role)
	}
	fmt.Println(table)

	defaultKey := []interface{}{"SID", "AppName", "Env", "TargetIp", "Role"}

	table2 := termtables.CreateTable()
	table2.AddHeaders(defaultKey...)

	for _, authServer := range authServers {
		s := structs.New(&authServer)
		m := s.Map()
		row := make([]interface{}, 0)
		for _, key := range defaultKey {
			if _, ok := m[key.(string)]; !ok {
				row = append(row, "")
			} else {
				row = append(row, m[key.(string)].(string))
			}
		}
		table2.AddRow(row...)
	}

	//fmt.Println(table2.Render())

	t := terminal.NewTerminal(os.Stdout, "")
	t.Write([]byte(table2.Render()))
}
