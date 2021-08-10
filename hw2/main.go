package main

import (
	"flag"
	"fmt"
	"os"

	"best-practice/hw2/pkg/services/duplicate"

	"best-practice/hw2/config"

	log "github.com/sirupsen/logrus"
)

func main() { // линтер whitespace - лишняя строка пустая
	remove := flag.Bool("r", false, "remove duplicate")
	path := flag.String("p", "./test/test_dir", "directory path")
	debug := flag.Bool("debug", true, "set log level to debug")
	flag.Parse()

	c := config.New(path, remove, debug)

	log.SetFormatter(&log.JSONFormatter{})

	if c.GetDebug() {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("start program")
	Run(c.GetPath(), c.GetRemove())
}

func Run(path string, remove bool) {
	l := log.WithField("FuncName", "Run").WithField("path", path)

	duplicates, err := duplicate.GetDuplicateFile(path)

	if err != nil {
		l.Error("get duplicate file: ", err)
	}

	if len(duplicates) == 0 {
		fmt.Println("no duplicates found")
		return
	}

	fmt.Println("Duplicates:")

	for i, item := range duplicates {
		fmt.Printf("%d.  %s", i+1, item)
		fmt.Println()
	}

	if remove {
		fmt.Print("remove duplicates? ", "confirm command: (y/n)  ")

		response := ""
		fmt.Fscan(os.Stdin, &response)

		if response == "y" || response == "Y" {
			err := duplicate.RemoveDuplicate(duplicates)

			if err != nil {
				log.Error("can't remov duplicates :", err)
				return
			}

			fmt.Println("Done!")
		}
	}
}
