/*
 * @Author: ChZheng
 * @Date: 2021-12-17 19:18:17
 * @LastEditTime: 2021-12-27 01:15:30
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/tour/cmd/root.go
 */
package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
}
