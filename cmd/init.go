/*
Copyright © 2022 Daniel Adeboye <adeboyed93@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
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
	if err := os.Chdir(name); err != nil {
		log.Fatal(err)
	}
	if str, err := os.Getwd(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Project Directory: %v\n", str)
	}
	bar := progressbar.Default(100, "Creating project")
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
	if err := exec.Command("git", "clone", repo, ".").Run(); err != nil {
		log.Fatal(err)
	}
	if err := exec.Command("code", ".").Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout, colorGreen, "Your project has been successfully...........")
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
