/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"io"
	"path/filepath"
	"strings"
	"github.com/spf13/cobra"
)

var fileExt string
var dirName string

// duplicateCmd represents the duplicate command
var duplicateCmd = &cobra.Command{
	Use:   "duplicate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("duplicate called")
		fmt.Println(fileExt)
		files := listDir(".")
		foundedFile, matchFiles := checkFiles(fileExt, files)
		if(foundedFile){
		  duplicateFile(dirName, matchFiles)
		}
	},
}

func init() {
	fmt.Println("init called")
	duplicateCmd.Flags().StringVarP(&fileExt, "extension", "e", "", "file extension is required")
	duplicateCmd.MarkFlagRequired("extension")
	duplicateCmd.PersistentFlags().StringVarP(&dirName, "dirname", "d", "copied_files", "dir to copie")
	duplicateCmd.MarkFlagRequired("dirname")
	rootCmd.AddCommand(duplicateCmd)

}

func Duplicate(){
	if err := duplicateCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createDir(dirName string) {
	os.MkdirAll(dirName, os.ModePerm)
  }

  func listDir(dirName string) []string{
	  var files []string

	  _ = filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		  files = append(files, path)
		  return nil
	  })

	  return files
  }

  func checkFiles(fileExt string, files []string) (bool, []string) {
	  result := false
	  var matchFiles []string
	  for _, file := range files {
		  if(strings.Index(file, fileExt) > 0){
			  matchFiles = append(matchFiles, file)
			  result = true
		  }
	  }

	  return result, matchFiles
  }

  func duplicateFile(newDirName string, matchFiles []string) {
	  fmt.Println("Creating " + newDirName + " dir.")
	  createDir(newDirName)

	  for _, fileName := range matchFiles {
		  file, _ := os.Open(fileName)
		  dest, _ := os.Create(newDirName + "/" + fileName)
		  io.Copy(dest, file)
	  }
  }
