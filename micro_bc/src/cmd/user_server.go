package main

import (
	"context"
	butil "git.aizinger.com/aizinger/go-base/util"
	//"github.com/getsentry/sentry-go"
	"github.com/micro/go-micro/v2"
	"log"
	"microbc/src/Users"
)

type UserService struct {
}

func (g *UserService) Test(ctx context.Context, req *Users.UserRequest, rsp *Users.UserResponse) error {
	rsp.Ret = "users" + req.Id
	return nil
}

func NewUserService() *UserService {
	return &UserService{}
}

func main() {
	butil.Now()
	service := micro.NewService(
		micro.Name("api.nanshan.com.user"))

	service.Init()

	err := Users.RegisterUserServiceHandler(service.Server(), NewUserService())
	if err != nil {
		log.Fatal(err)
	}

	//err = sentry.Init(sentry.ClientOptions{
	//	// Either set your DSN here or set the SENTRY_DSN environment variable.
	//	Dsn: "https://9b3f22de158041a987863435c51d3aff@o575131.ingest.sentry.io/5726966",
	//	// Either set environment and release here or set the SENTRY_ENVIRONMENT
	//	// and SENTRY_RELEASE environment variables.
	//	//Environment: "",
	//	//Release:     "my-project-name@1.0.0",
	//	// Enable printing of SDK debug messages.
	//	// Useful when getting started or trying to figure something out.
	//	Debug: true,
	//})
	//
	//if err != nil {
	//	log.Fatalf("sentry.Init: %s", err)
	//}
	//// Flush buffered events before the program terminates.
	//// Set the timeout to the maximum duration the program can afford to wait.
	//defer sentry.Flush(2 * time.Second)
	//
	//sentry.CaptureMessage("It go works!")

	if err = service.Run(); err != nil {
		log.Println(err)
	}

}
