package main

import (
	"task-manager/controller"
	"task-manager/db"
	"task-manager/repository"
	"task-manager/router"
	"task-manager/usecase"
	"task-manager/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskREpository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskREpository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
