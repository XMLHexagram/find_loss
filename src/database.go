package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type lost struct {
	gorm.Model
	Stuno string
	Name  string
}

func (s *Service) DBInit()  {
	strDb := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.Config.DB.User,
		s.Config.DB.Password,
		s.Config.DB.Address,
		s.Config.DB.DBName)
	db,err := gorm.Open("mysql", strDb)
	DealError(err)
	fmt.Println("success connect to database")

	db.AutoMigrate(&lost{})
	s.DB = db
	//fmt.Println(s.DB)
}
