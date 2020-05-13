package main

import (
	"fmt"

	"github.com/ciiiii/sync-image/convert"
	"github.com/spf13/cobra"
)

var (
	source  string
	target  string
	rootCmd = &cobra.Command{
		Use:   "imageConverter",
		Short: "Docker image replace tool",
		Long: `A Tool to replace docker images with the mirror images in yaml, 
it should be used with the synchronous mirroring feature of this project together.
Complete docs is available at https://github.com/ciiiii/sync-image.`,
		Run: func(cmd *cobra.Command, args []string) {
			converter := convert.Converter{
				Source: source,
				Target: target,
			}
			if err := converter.Parse(); err != nil {
				fmt.Errorf("get source image list failed:\n%s", err)
			}
			if err := converter.Replace(converter.StringMapper()); err != nil {
				fmt.Errorf("Replace images failed:\n%s", err)
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&source, "source", "https://raw.githubusercontent.com/ciiiii/sync-image/master/images.json", "url or local file path")
	rootCmd.PersistentFlags().StringVar(&target, "target", ".", "yaml file or yaml directory")
}

func main() {
	rootCmd.Execute()
}
