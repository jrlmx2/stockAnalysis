package mariadb

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/jrlmx2/stockAnalysis/utils/config"
	"github.com/jrlmx2/stockAnalysis/utils/term"
)

func (p *Pool) BackupDB(dump config.Dump, database config.Database) {
	interval, err := time.ParseDuration(dump.Interval)
	if err != nil {
		panic(fmt.Sprintf("Database dump errored with %s, %+v", err, dump))
	}
	fmt.Printf("Establishing Comand %s, %s, %s, %s, %s\n", dump.CommandFile, database.User, database.Password, database.Schema, fmt.Sprintf(dump.Out, time.Now().UnixNano()))
	for {
		if term.WasTerminated() {
			return
		}
		fmt.Println("Running Command")
		err = exec.Command(dump.CommandFile, database.User, database.Password, database.Schema, fmt.Sprintf(dump.Out, time.Now().UnixNano())).Run()
		if err != nil {
			panic(fmt.Sprintf("Database dump command errored with %s", err))
		}
		time.Sleep(interval)
	}
}
