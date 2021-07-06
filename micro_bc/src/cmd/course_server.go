package main

import (
	"microbc/src/Courses"
	//"github.com/getsentry/sentry-go"
	"github.com/micro/go-micro/v2"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("api.nanshan.com.course"))

	service.Init()

	err := Courses.RegisterCourseServiceHandler(service.Server(), Courses.NewCourseServiceImpl())
	if err != nil {
		log.Fatal(err)
	}

	if err = service.Run(); err != nil {
		log.Println(err)
	}

}
