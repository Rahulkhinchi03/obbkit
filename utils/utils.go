package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func utilsCall(repo, name, colorGreen, input string) {
	if err := os.Mkdir(name, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	os.Chdir(name)
	str, err := os.Getwd()
	fmt.Printf("Project Directory: %v\n", str)
	fmt.Printf("Possible error: %v\n", err)
	cmd := exec.Command("git", "clone", repo, ".")
	cmd.Run()
	exec.Command("code", ".").Run()
	fmt.Fprintln(os.Stdout, colorGreen, "Your reactjs project has been successfully created with title name of ->>", input)
}
