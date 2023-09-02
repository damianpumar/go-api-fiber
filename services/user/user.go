package user

import (
	"context"
	"go-api-fiber-starterkit/prisma/db"
	"go-api-fiber-starterkit/services/database"
)

type UserService struct {
	database *database.Database
}

func (s *UserService) GetUsers() ([]db.UserModel, error) {
	ctx := context.Background()
	
	defer func ()  {
		s.database.Disconnect(*s.database.Client)
	}()

	users, err := s.database.Client.User.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func New() *UserService{
	client, err := database.New()

	if(err != nil){
		panic(err)
	}

	return &UserService{
		database: client,
	}
}