// This will act as a DB layer for the application. It will have functions to interact with the database. For now, we will use a json
package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Constants for the path and extension of the file
const path = "./data/"
const ext = ".json"

type FileDB struct {
	Name string
}

var FileDatabaseMap map[string]FileDB

func CreateFileDB(name string) FileDB {
	// Create a DB
	// Check if DBs is nil
	if FileDatabaseMap == nil {
		FileDatabaseMap = make(map[string]FileDB)
	}

	// Check if the DB already exists
	if _, ok := FileDatabaseMap[name]; ok {
		return FileDatabaseMap[name]
	}

	// Create a new DB
	db := FileDB{Name: name}
	FileDatabaseMap[name] = db
	return db
}

func DeleteDB(name string) {
	// Delete a DB
	delete(FileDatabaseMap, name)
}

func GetDB(name string) (FileDB, error) {
	// Get a DB
	for _, db := range FileDatabaseMap {
		if db.Name == name {
			return db, nil
		}
	}
	return FileDB{}, fmt.Errorf("DB not found")
}

// Method to initialize a DB
func (db FileDB) InitDB() error {

	fileName := filepath.Join(GetCurrentWorkingDirectory(), path, db.Name+ext)

	// Check if the file exists or not and create a new file if it does not exist with an empty array of data
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Create a new file
		err = WriteJsonFile(fileName, []byte("[]"))
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// Method to add data to a DB
func (db FileDB) AddData(data interface{}) error {
	// Add data to a DB
	structType := data
	jsonData, err := json.Marshal(structType)
	if err != nil {
		return fmt.Errorf("error marshalling data: %s", err)
	}

	// Write to a json file
	fileName := filepath.Join(GetCurrentWorkingDirectory(), path, db.Name+ext)

	err = WriteJsonFile(fileName, jsonData)

	if err != nil {
		fmt.Printf("Error Writing to file: %v\n", err)
		return err
	}
	return nil
}

// GetAll retrieves all records from the database and unmarshals them into the provided slice.
func (db FileDB) GetAll(data interface{}) error {
	// Get data from a DB

	fileName := filepath.Join(GetCurrentWorkingDirectory(), path, db.Name+ext)
	jsonData, err := ReadJsonFile(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	errr := json.Unmarshal(jsonData, &data)
	if errr != nil {
		fmt.Println(errr)
		return errr
	}
	return nil
}

// Method to delete data from a DB
func (db FileDB) DeleteData(data *interface{}) error {
	// Delete data from a DB

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	DeleteJsonFile(db.Name)
	WriteJsonFile(db.Name, jsonData)
	return nil
}

func GetCurrentWorkingDirectory() string {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return dir
}
