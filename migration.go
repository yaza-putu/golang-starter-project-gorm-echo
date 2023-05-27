package main

import (
	"fmt"
	"gitlab.com/putuyaza/antrian/src/core"
	"gitlab.com/putuyaza/antrian/src/database"
	_ "gitlab.com/putuyaza/antrian/src/database/migrations"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// run migration
	migrationRunner()
}

// migrationRunner engine
func migrationRunner() bool {
	if len(os.Args) > 1 {
		// init env
		core.LoadEnv()
		// init database
		database.Connection()
		for i, v := range os.Args {
			if i != 0 {
				switch v {
				case "new":
					if len(os.Args) != 3 {
						fmt.Println("Should be : new name-of-file")
						return true
					}
					// file
					fName := fmt.Sprintf("./src/database/migrations/%s_create_table_%s.go", time.Now().Format("20060102150405"), os.Args[2])
					// from template
					from, err := os.Open("./src/database/template/auto-migration.stub")
					if err != nil {
						log.Fatal(err)
					}
					defer from.Close()
					// to file
					to, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE, 0666)
					if err != nil {
						log.Fatal(err)
					}
					defer to.Close()
					// copy file with template
					_, err = io.Copy(to, from)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("New migration created : %s\n", fName)
				case "up":
					err := database.MigrationUp()
					if err != nil {
						log.Fatal(err)
					} else {
						fmt.Println("Migrating collections successfully")
					}
				case "down":
					err := database.MigrationDown()
					if err != nil {
						log.Fatal(err)
					} else {
						fmt.Println("Drop collections successfully")
					}
				}
			}
		}
		return true
	}
	return false
}
