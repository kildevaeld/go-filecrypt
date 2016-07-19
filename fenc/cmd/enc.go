// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"io"
	"os"

	"github.com/kildevaeld/go-filecrypt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// encCmd represents the enc command
var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		if len(args) == 0 {
			printError(errors.New("usage: fenc enc <path>"))
		}

		keyr := viper.GetString("key")
		if keyr == "" {
			printError(errors.New("no key"))
		}

		var reader io.ReadCloser
		var writer io.WriteCloser
		var err error

		if reader, err = os.Open(args[0]); err != nil {
			printError(err)
		}
		defer reader.Close()
		if writer, err = os.Create(outputFlag); err != nil {
			printError(err)
		}
		defer writer.Close()

		key := filecrypt.Key([]byte(keyr))

		if _, err := filecrypt.Encrypt(writer, reader, &key); err != nil {
			printError(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(encCmd)

	//decCmd.Flags().StringVarP(&outputFlag, "output", "o", "output.out", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
