package main

import (
	"restgolang.com/restGolang/echo-rest/db"
	"restgolang.com/restGolang/echo-rest/routes"
)

func main() {
	db.Init()
	
	e :=	routes.Init();

	e.Logger.Fatal(e.Start(":1234"))
}