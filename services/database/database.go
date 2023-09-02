package database

import (
	"go-api-fiber-starterkit/prisma/db"
)

type Database struct{
	Client *db.PrismaClient
}


func New() (*Database,  error) {
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	return &Database{
		Client: client,
	}, nil
}

func (d *Database) Disconnect(client db.PrismaClient){
	if err := client.Prisma.Disconnect(); err != nil {
		panic(err)
	}
}