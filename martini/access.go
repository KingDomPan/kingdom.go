package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/accessflags"
)

const (
	rolePassAll = 0
	roleRobot   = 1
	roleSignOn  = 2
	roleAdmin   = 4 | roleSignOn // 6
)

func main() {
	m := martini.Classic()

	m.Use(judge)
	m.Get("/profile", accessflags.Forbidden(roleSignOn), profileHandler)
	m.Get("/admin", accessflags.Less(roleAdmin, 403), adminHandler)

	m.Run()
}

func judge(c martini.Context) {
	c.Map(roleRobot)
}

func adminHandler() string {
	return "hello"
}
func profileHandler() string {
	return "profile"
}
