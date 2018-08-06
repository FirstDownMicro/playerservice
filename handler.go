package main

import (
	"context"
	"fmt"

	pb "github.com/Nanixel/FirstDownMicro/playerservice/proto"
	"github.com/jinzhu/gorm"
)

type PlayerModel struct {
	Id       int32 `gorm:"primary_key"`
	Team_id  int32
	Name     string
	Position string
	Rating   int32
}

type playerservice struct {
	db *gorm.DB
}

// func (model *PlayerModel) BeforeCreate(scope *gorm.Scope) error {
// 	uuid := uuid.NewV4()
// 	return scope.SetColumn("id", uuid.String())
// }

func (service *playerservice) Draft(ctx context.Context, req *pb.DraftPlayerRequest, res *pb.PlayerResponse) error {
	fmt.Println(req.TeamId)
	var model = PlayerModel{Team_id: req.TeamId, Name: req.Name, Position: req.Position, Rating: 89}
	fmt.Println(model.Name)
	if err := service.db.Create(&model).Error; err != nil {
		fmt.Println(err)
		return err
	}
	var player = &pb.Player{Id: model.Id, TeamId: req.TeamId, Name: req.Name, Position: req.Position, Rating: model.Rating}

	res.Player = player
	return nil

}

func (service *playerservice) Get(ctx context.Context, req *pb.PlayerRequest, res *pb.PlayerResponse) error {
	var model PlayerModel

	if dberr := service.db.Where("id = ?", req.PlayerId).Find(&model).Error; dberr != nil {
		fmt.Println(dberr)
		return dberr
	}
	var player = &pb.Player{Id: model.Id, TeamId: model.Team_id, Name: model.Name, Position: model.Position, Rating: model.Rating}

	res.Player = player

	return nil
}

func (service *playerservice) Update(ctx context.Context, req *pb.UpdatePlayerRequest, res *pb.PlayerResponse) error {
	var model PlayerModel

	service.db.Where("id = ?", req.Id).Find(&model)

	model.Rating = req.Rating
	model.Team_id = req.TeamId
	service.db.Save(&model)

	var player = &pb.Player{}
	player.Id = model.Id
	player.Rating = model.Rating
	player.TeamId = model.Team_id
	player.Position = model.Position
	player.Name = model.Name

	res.Player = player

	return nil
}

func (service *playerservice) GetAll(ctx context.Context, req *pb.AllPlayersRequest, res *pb.AllPlayersResponse) error {
	var players []PlayerModel
	var playerarr = []*pb.Player{}
	service.db.Find(&players)

	for _, player := range players {
		var p = &pb.Player{Id: player.Id, TeamId: player.Team_id, Name: player.Name, Position: player.Position, Rating: player.Rating}
		playerarr = append(playerarr, p)
	}

	res.Players = playerarr

	return nil
}
