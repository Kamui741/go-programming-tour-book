/*
 * @Author: ChZheng
 * @Date: 2021-12-15 21:39:20
 * @LastEditTime: 2021-12-17 23:36:53
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/tour/main.go
 */
package main

import (
	"log"
	"tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
