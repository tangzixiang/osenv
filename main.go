package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

func main() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Printf("wokr dir %v\n", workDir)
	fileName := ".env"
	filePath := path.Join(workDir, fileName)
	fmt.Printf("reading file %v\n", filePath)

	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%v not found, exit", filePath)
			return
		}
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		fmt.Printf("get env: %v\n", scanner.Text())

		splits := strings.Split(scanner.Text(), "=")
		if len(splits) != 2 {
			fmt.Printf("env %v is not valid, pass\n", scanner.Text())
			continue
		}

		key, value := splits[0], splits[1]

		switch runtime.GOOS {
		case `windows`:
			SetEnvWindows(key, value)
		default:
			SetEnvDefault(key, value)
		}
	}
}

func SetEnvDefault(key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		panic(err)
	}
}

func SetEnvWindows(key, value string) {
	err := exec.Command("SETX", key, value).Run()
	if err != nil {
		panic(err)
	}
}
