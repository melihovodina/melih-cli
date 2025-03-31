package commands

import (
	"fmt"
	"log"
	"melihCli/pkg/store"
	"os"
	"os/exec"
	"path/filepath"
)

type fileToCreate struct {
	name    string
	content string
}

type defineResult struct {
	funcs    []func()
	commands [][]string
	files    []fileToCreate
}

func defineCommand() defineResult {
	if len(os.Args) < 2 {
		fmt.Println("Expected some command")
		os.Exit(1)
	}

	command := os.Args[1]
	var funcs []func()
	var commands [][]string
	var files []fileToCreate

	switch command {
	case "nest-cds-init":
		store.DefineName(os.Args[2])
		commands = store.NestCds

		files = append(files, fileToCreate{name: "tsconfig.json", content: store.TsConf})
		files = append(files, fileToCreate{name: "src/main.ts", content: store.MainTs})

		funcs = append(funcs, func() { runCommands(commands) })
		funcs = append(funcs, func() { createFiles(files) })
	default:
		fmt.Println("Неизвестная команда:", command)
		os.Exit(1)
	}

	return defineResult{
		funcs:    funcs,
		commands: commands,
		files:    files,
	}
}

func runCommand(name string, args ...string) {
	switch name {
	case "cd":
		err := os.Chdir(args[0])
		if err != nil {
			log.Fatalf("Failed to change directory to %s: %v", args[0], err)
		}
	case "mkdir":
		err := os.Mkdir(args[0], 0755)
		if err != nil {
			log.Fatalf("Failed to create directory %s: %v", args[0], err)
		}
	default:
		cmd := exec.Command(name, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to run command %s %v: %v", name, args, err)
		}
	}
}

func runCommands(commands [][]string) {
	for _, command := range commands {
		runCommand(command[0], command[1:]...)
	}
}

func createFiles(files []fileToCreate) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(currentDir, file.name)

		err := os.WriteFile(filePath, []byte(file.content), 0644)
		if err != nil {
			log.Fatalf("Failed to create %s: %v", filePath, err)
		}
	}
}

func Run() {
	params := defineCommand()

	for _, function := range params.funcs {
		function()
	}
}
