/*
 * @Author: ChZheng
 * @Date: 2021-12-17 19:18:17
 * @LastEditTime: 2021-12-17 22:59:51
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/tour/cmd/root.go
 */
package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
}
