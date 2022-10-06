package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"hermannvincent/deliveryservice/ecs"
	"log"
	"os"
)

type Database struct {
	Entities []ecs.ModelEntity `json:"entities"`
}

func (db *Database) LoadReturnEntities() (modelEntities []ecs.ModelEntity, err error) {
	file, err := os.Open("assets/entities.json")
	if err != nil {
		return nil, errors.New("cannot open assets/entities.json")
	}
	defer file.Close()

	var fileData []byte
	readedFile, err := file.Read(fileData)
	if err != nil {
		return nil, errors.New("cannot read assets/entities.json")
	}

	log.Printf("Read assets/entities.json of length(%v)\n", readedFile)

	if err := json.NewDecoder(file).Decode(&modelEntities); err != nil {
		return nil, fmt.Errorf("cannot unmarshal assets/entities.json to db.Entities : %v", err.Error())
	}

	return modelEntities, nil
}

func (db *Database) LoadStoreEntities() (err error) {
	file, err := os.Open("assets/entities.json")
	if err != nil {
		return errors.New("cannot open assets/entities.json")
	}
	defer file.Close()

	var fileData []byte
	readedFile, err := file.Read(fileData)
	if err != nil {
		return errors.New("cannot read assets/entities.json")
	}

	log.Printf("Read assets/entities.json of length(%v)\n", readedFile)

	if err := json.NewDecoder(file).Decode(&db.Entities); err != nil {
		return fmt.Errorf("cannot unmarshal assets/entities.json to db.Entities : %v", err.Error())
	}

	return nil
}
