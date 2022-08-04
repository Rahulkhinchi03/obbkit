/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

const colorRed = "\033[0;31m"
const colorNone = "\033[0m"
const colorGreen = "\033[0;32m"

var (
	// Used for flags.
	project string
	name    string
)

// main logic
func utilsCall(repo, name, colorGreen, input string) {
	if err := os.Mkdir(name, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	os.Chdir(name)
	str, err := os.Getwd()
	fmt.Printf("Project Directory: %v\n", str)
	fmt.Printf("Possible error: %v\n", err)
	bar := progressbar.Default(100, "Creating project")
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
	cmd := exec.Command("git", "clone", repo, ".")
	cmd.Run()
	exec.Command("code", ".").Run()
	fmt.Fprintln(os.Stdout, colorGreen, "Your project has been successfully created with title name of ->>", input)
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Lists available starterkits on onboardbase.",
	Long:    `Lists avaialable starterkits on onboardbase.`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Enter project type: ")
		fmt.Scan(&project)
		fmt.Printf("Enter project name: ")
		fmt.Scan(&name)
		input := name
		switch project {
		case "reactjs":
			var repo = "git@github.com:Onboardbase/Reactjs-Starterkit.git"
			utilsCall(repo, name, colorGreen, input)
		case "vuejs":
			var repo = "git@github.com:Onboardbase/Vuejs-Starterkit.git"
			utilsCall(repo, name, colorGreen, input)
		case "flask":
			var repo = "git@github.com:Onboardbase/Flask-Starterkit.git"
			utilsCall(repo, name, colorGreen, input)
		case "nextjs":
			var repo = "git@github.com:Onboardbase/nextjs-starterkit.git"
			utilsCall(repo, name, colorGreen, input)
		case "nuxtjs":
			var repo = "git@github.com:Onboardbase/nuxtjs-starterkit.git"
			utilsCall(repo, name, colorGreen, input)
		case "gatsby":
			var repo = "git@github.com:Onboardbase/gatsby-starterkit.git"
			utilsCall(repo, name, colorGreen, input)
		default:
			fmt.Fprintln(os.Stdout, "None:", colorRed, "please enter a valid project type.")
			fmt.Fprintln(os.Stdout, "Help:", colorNone, "Use the obbkit list command to see available project types.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	//rootCmd.Flags().StringVarP(&project, "project", "p", "YOUR PROJECT", "project type (required)")
	//rootCmd.MarkFlagRequired("project")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
