/*
 * @Author: ChZheng
 * @Date: 2021-12-17 19:11:35
 * @LastEditTime: 2021-12-23 22:16:10
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/tour/cmd/word.go
 */
package cmd

import (
	"log"
	"strings"
	"tour/internal/word"

	"github.com/spf13/cobra"
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果：%s", content)
	},
}

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换：模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")
var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "word")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "mode")
}
