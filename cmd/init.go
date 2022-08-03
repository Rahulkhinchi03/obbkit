/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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
			if err := os.Mkdir(name, os.ModePerm); err != nil {
				log.Fatal(err)
			}
			os.Chdir(name)
			str, err := os.Getwd()
			fmt.Printf("str: %T, %v\n", str, str)
			fmt.Printf("err: %T, %v\n", err, err)
			var repo = "git@github.com:Onboardbase/Reactjs-Starterkit.git"
			cmd := exec.Command("git", "clone", repo, ".")
			cmd.Run()
			exec.Command("code", ".").Run()
			fmt.Fprintln(os.Stdout, colorGreen, "You reactjs project has been successfully created with title name of ->>", input)
		case "vuejs":
			fmt.Println("You vuejs project has been created with title name of ->>", input)
		case "flask":
			fmt.Println("You flask project has been created with title name of ->>", input)
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
