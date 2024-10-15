// This will act as a DB layer for the application. It will have functions to interact with the database. For now, we will use a json
package db

import (
	"encoding/json"
	"fmt"
	"os"
)

const path = "./db/data/"
const ext = ".json"

type DB struct {
	Name string
}

var DBs map[string]DB

func CreateDB(name string) DB {
	// Create a DB
	// Check if DBs is nil
	if DBs == nil {
		DBs = make(map[string]DB)
	}

	// Check if the DB already exists
	if _, ok := DBs[name]; ok {
		return DBs[name]
	}

	// Create a new DB
	db := DB{Name: name}
	DBs[name] = db
	return db
}

func DeleteDB(name string) {
	// Delete a DB
	delete(DBs, name)
}

func GetDB(name string) (DB, error) {
	// Get a DB
	for _, db := range DBs {
		if db.Name == name {
			return db, nil
		}
	}
	return DB{}, fmt.Errorf("DB not found")
}

// Method to initialize a DB
func (db DB) InitDB() error {
	fileName := path + db.Name + ext

	// Check if the file exists or not and create a new file if it does not exist with an empty array of data
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err = WriteJsonFile(fileName, []byte("[]"))
		if err != nil {
			fmt.Printf("Error Writing to file: %v\n", err)
			return err
		}
	}
	return nil
}

// Method to add data to a DB
func (db DB) AddData(data interface{}) error {
	// Add data to a DB
	structType := data
	jsonData, err := json.Marshal(structType)
	if err != nil {
		return fmt.Errorf("error marshalling data: %s", err)
	}

	// Write to a json file
	fileName := path + db.Name + ext

	err = WriteJsonFile(fileName, jsonData)

	if err != nil {
		fmt.Printf("Error Writing to file: %v\n", err)
		return err
	}
	return nil
}

// GetAll retrieves all records from the database and unmarshals them into the provided slice.
func (db DB) GetAll(data interface{}) (error) {
	// Get data from a DB
	fileName := path + db.Name + ext

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
func (db DB) DeleteData(data *interface{}) error {
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