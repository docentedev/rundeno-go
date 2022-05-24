package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {

	scriptCommand := os.Args[1]
	// read file
	data, err := ioutil.ReadFile("./scripts.json")
	if err != nil {
		fmt.Print(err)
	}

	var m map[string]interface{}
	json.Unmarshal([]byte(data), &m)
	fmt.Println(m[scriptCommand])

	if m[scriptCommand] != nil {

		// execute start script
		cmd := exec.Command("/bin/bash", "-c", m[scriptCommand].(string))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}
}
