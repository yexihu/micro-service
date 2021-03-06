package main

import (
	"github.com/micro/go-micro/v2"
	"log"
	"micro-service/api/handler"

	post "micro-service/proto/post"
	user "micro-service/proto/user"
)

// 将grpc服务转为restful接口
func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&handler.Say{Client: user.NewUserService("go.micro.srv.user", service.Client())},
		),
	)
	service.Server().Handle(
		service.Server().NewHandler(
			&handler.Article{Client: post.NewPostService("go.micro.srv.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
