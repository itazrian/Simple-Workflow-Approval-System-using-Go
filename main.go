package main

import (
	"log"
	"workflow-approval/internal/handler"
	"workflow-approval/internal/model"
	"workflow-approval/internal/repository"
	"workflow-approval/internal/router"
	"workflow-approval/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/tes_medela_030226?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Workflow{}, &model.WorkflowStep{}, &model.Request{})

	ws := service.NewWorkflowService(db)
	ss := service.NewStepService(db)
	rs := service.NewRequestService(db, repository.NewRequestRepository(db))

	h := handler.NewRequestHandler(ws, ss, rs)

	e := echo.New()
	router.Register(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
