package db

import "github.com/google/uuid"

type Indentifiable interface {
	GetID() string
	SetID(string)
}

type DBModel[T Indentifiable] struct {
	// Model is a struct that contains the model name
	// and the model instance
	tableName     string
	modelInstance T
}

// NewModel is a function that creates a new Model struct
func NewModel[T Indentifiable](tableName string, modelInstance T) *DBModel[T] {
	return &DBModel[T]{
		tableName:     tableName,
		modelInstance: modelInstance,
	}
}

// GetModelInstance is a function that returns the model instance
func (m *DBModel[T]) GetModelInstance() T {
	return m.modelInstance
}

// GetTableName is a function that returns the table name
func (m *DBModel[T]) GetTableName() string {
	return m.tableName
}

// CreateObject is a method that creates an object in the database
func (m *DBModel[T]) Create(object T) error {
	_id := m.GetTableName() + "-" + uuid.New().String()
	object.SetID(_id)
	db := GetDBServerInstance().GetDB()
	result := db.Table(m.GetTableName()).Create(object)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetObject is a method that gets an object from the database
func (m *DBModel[T]) Get(id string) (T, error) {
	var object T
	db := GetDBServerInstance().GetDB()
	result := db.Table(m.GetTableName()).Where("id = ?", id).First(&object)
	if result.Error != nil {
		var zeroValue T
		return zeroValue, result.Error
	}
	return object, nil
}

func (m *DBModel[T]) GetByField(field string, value string) (T, error) {
	var object T
	db := GetDBServerInstance().GetDB()
	result := db.Table(m.GetTableName()).Where(field+" = ?", value).First(&object)
	if result.Error != nil {
		var zeroValue T
		return zeroValue, result.Error
	}
	return object, nil
}

// GetObjects is a method that gets all objects from the database
func (m *DBModel[T]) GetObjects() ([]T, error) {
	var objects []T
	db := GetDBServerInstance().GetDB()
	result := db.Table(m.GetTableName()).Find(&objects)
	if result.Error != nil {
		return nil, result.Error
	}
	return objects, nil
}

// UpdateObject is a method that updates an object in the database
func (m *DBModel[T]) Update(object T, id string) error {
	db := GetDBServerInstance().GetDB()
	result := db.Table(m.GetTableName()).Where("id = ?", id).Updates(&object)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteObject is a method that deletes an object from the database
func (m *DBModel[T]) Delete(id string) error {
	db := GetDBServerInstance().GetDB()
	result := db.Table(m.GetTableName()).Where("id = ?", id).Delete(nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteObjects is a method that deletes all objects from the database
func (m *DBModel[T]) DeleteObjects() error {
	db := GetDBServerInstance().GetDB()
	result := db.Table(m.GetTableName()).Delete(nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
