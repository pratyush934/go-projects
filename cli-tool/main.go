package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	markName  = "GOLANG_CLI_REMAINDER"
	markValue = "1"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <hh:mm> <text message>\n", os.Args[0])
		os.Exit(1)

	}

	now := time.Now()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	parse, err := w.Parse(os.Args[1], now)
	if err != nil {
		fmt.Printf("There is an error in w.Parse %v", err.Error())
		os.Exit(2)

	}

	if parse == nil {
		fmt.Printf("Unable to parse \n")
		os.Exit(2)
	}

	if now.After(parse.Time) {
		fmt.Printf("Set the future time\n")
		os.Exit(3)
	}

	diff := parse.Time.Sub(now)

	if os.Getenv(markName) == markValue {
		fmt.Println("Environment variable is set, sleeping for:", diff)
		time.Sleep(diff)
		err = beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			fmt.Printf("Error in beeep.Alert: %v\n", err)
			os.Exit(4)
		}
	} else {
		fmt.Println("Setting environment variable and re-executing command")
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		err := cmd.Start()
		if err != nil {
			fmt.Printf("Error in cmd.Start: %v\n", err)
			os.Exit(5)
		}
		fmt.Printf("Reminder will be displayed after %v\n", diff.Round(time.Second))
		os.Exit(6)
	}
}
