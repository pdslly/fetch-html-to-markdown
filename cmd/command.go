package cmd

import (
	"fmt"

	"github.com/pdslly/fetch-html-to-markdown/spider"
	"github.com/spf13/cobra"
)

var (
	o string
	v bool
)

var rootCmd = &cobra.Command{
	Use:   "html2md <url>",
	Short: "爬取指定URL文章地址，并解析输出Markdown",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if v {
			fmt.Println(spider.VERSION)
		} else {
			args = append(args, "./")
			spider.Parse(args[0], o)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&v, "version", "v", false, "显示当前版本")
	rootCmd.PersistentFlags().StringVarP(&o, "outfile", "o", "out.md", "输出Markdown文件名")
}

func Execute() error {
	return rootCmd.Execute()
}
