package main

import (
	"log"

	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/controller/cmd/bootstrap"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/controller/cmd/config"
	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/controller/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	signalHandler := bootstrap.SignalHandler()
	api := bootstrap.Api(rabbitMqChannel)

	log.Println("Bootstrap finished. Event API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Event API finished gracefully")
}
