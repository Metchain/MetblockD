package domain

import (
	"github.com/Metchain/Metblock/mconfig"
	"log"
	"os"
)

func DBReset() bool {
	err := os.RemoveAll(mconfig.GetDBDir())
	if err == nil {
		log.Println("Database Successfully Reset")
		return true

	}
	log.Println(err)

	// Add exit code with Manual
	os.Exit(0)
	return false
}
