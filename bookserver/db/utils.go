// This will act as a DB layer for the application. It will have functions to interact with the database. For now, we will use a json
package db

import (
	"encoding/json"
	"log"
	"os"
)

// function to read a json file
func ReadJsonFile(fileName string) ([]byte, error) {
	// Read a json file
	data , err :=  os.ReadFile(fileName)

	if err != nil {
		log.Println(err)
	}
	return data , err
}

// function to write to a json file
func WriteJsonFile(fileName string, data []byte) error {
	// Write to a json file
	err := os.WriteFile(fileName, data, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}


// function to delete a json file
func DeleteJsonFile(fileName string) {
	// Delete a json file
	err := os.Remove(fileName)
	if err != nil {
		log.Println(err)
	}
}

// function to convert a struct to json
func StructToJson(data interface{}) []byte {
	// Convert a struct to json
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return jsonData
}

// function to convert json to a struct
func JsonToStruct(data []byte, structType interface{}) interface{} {
	// Convert json to a struct
	err := json.Unmarshal(data, structType)
	if err != nil {
		log.Println(err)
	}
	return structType
}


