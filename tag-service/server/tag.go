/*
 * @Author: ChZheng
 * @Date: 2022-01-23 00:44:26
 * @LastEditTime: 2022-01-23 01:46:32
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /tag-service/server/tag.go
 */
package server

import (
	"context"
	"encoding/json"
	"go-programming-tour-book/tag-service/pkg/bapi"
	"go-programming-tour-book/tag-service/pkg/errcode"
	pb "go-programming-tour-book/tag-service/proto"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}
func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return &tagList, nil
}
