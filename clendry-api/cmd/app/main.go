package main

import (
	"fmt"
	"github.com/artomsopun/clendry/clendry-api/internal/app"
	"github.com/artomsopun/clendry/clendry-api/internal/config"
	"github.com/artomsopun/clendry/clendry-api/internal/domain"
	"github.com/artomsopun/clendry/clendry-api/pkg/database"
)

const configsDir = "clendry-api/configs"

func main() {
	app.Run(configsDir)
}

func Test() {
	cfg, err := config.Init(configsDir)
	if err != nil {
		return
	}

	// Dependencies
	db := database.NewDB(cfg.MySql.User, cfg.MySql.Password, cfg.MySql.Host, cfg.MySql.Port, cfg.MySql.Name)
	if err != nil {
		return
	}

	var users []domain.User
	err = db.Model(&domain.User{}).Select("users.id, users.nick, users.email").
		Where("users.id != ?", 1).
		Joins("left join block_requests AS b on b.user_id = users.id AND b.def_id = users.id").
		Where("b.user_id != ? AND b.def_id != ?", 1, 1).
		Scan(&users).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}
