/*
 * @Author: ChZheng
 * @Date: 2021-12-15 21:39:20
 * @LastEditTime: 2022-01-01 16:45:38
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/tour/main.go
 */
package main

import (
	"go-programming-tour-book/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
