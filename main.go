package main

import (
	"gitlab.com/putuyaza/antrian/src/core"
	"gitlab.com/putuyaza/antrian/src/database"
)

func main() {
	// load env
	core.LoadEnv()
	// init database
	database.Connection()
	// start server & load route
	core.Route()
}
