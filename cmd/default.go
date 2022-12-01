package cmd

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	specifyFile string
	outFile     string
)

var FileToBase64Cmd = &cobra.Command{
	Use:   "fileTobase64",
	Short: "File to Base64",
	Long:  "Specify a file to convert to Base64",
	Run: func(cmd *cobra.Command, args []string) {
		if strings.ReplaceAll(specifyFile, " ", "") == "" {
			fmt.Printf("please fill in specify a file")
			return
		}

		if strings.ReplaceAll(outFile, " ", "") == "" {
			outFile = "out.txt"
		}

		data, err := ioutil.ReadFile(specifyFile)
		if err != nil {
			fmt.Printf("read specify a file fail: %v", err)
			return
		}
		base := base64.StdEncoding.EncodeToString(data)
		file, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Printf("open out file fail: %v", err)
			return
		}
		defer file.Close()
		file.Write([]byte(base))
	},
}

func init() {
	FileToBase64Cmd.Flags().StringVarP(&specifyFile, "sfile", "s", "", "specify a file")
	FileToBase64Cmd.Flags().StringVarP(&outFile, "ofile", "o", "", "out file")
}
