package main

import (
	"fmt"

	pb "github.com/Nanixel/FirstDownMicro/playerservice/proto"
	micro "github.com/micro/go-micro"
)

func main() {

	db, dberr := CreateConnection()
	defer db.Close()
	if dberr != nil {
		fmt.Println("DB ERROR")
		fmt.Println(dberr.Error())
	}

	service := micro.NewService(
		micro.Name("go.micro.srv.playerservice"),
		micro.Version("latest"),
	)
	service.Init()

	pb.RegisterPlayerServiceHandler(service.Server(), &playerservice{db: db})

	fmt.Println("Starting playerservice....")
	if err := service.Run(); err != nil {
		fmt.Printf("Error occured when starting player service: %v\n", err)
	}

}
