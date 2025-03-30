/*
Copyright © 2025 Vidyasagar Gopi vidyasagar0405@gmail.com

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
	"wallslider/internals/wallslider"

	"github.com/spf13/cobra"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Set the next wallpaper in the sequence",
	Long:  `Move to the next wallpaper in your collection and set it as your desktop background.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := wallslider.NewConfig(imagePath)
		if err := config.NextWallpaper(); err != nil {
			return fmt.Errorf("failed to set next wallpaper: %w", err)
		}

		if err := config.IndexToJson(); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
