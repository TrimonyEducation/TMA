package main

import (
	"golang-crud/models"
	"golang-crud/utils"
	"log"
)

func main(){
	utils.ConnectDB()
	utils.DB.AutoMigrate(&models.User{}, &models.Video{}, &models.Class{}, &models.Exercise{}, &models.Playlist{}, &models.Problem{}, &models.Review{}, &models.Take{}, &models.Teacher{},  &models.VideoInstance{} )
	log.Println("Migration succesful!")
}