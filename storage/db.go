package storage

import "fmt"

type KeyType interface{}

type DB[ValueType any] struct {
	entries map[KeyType]ValueType
}

func (db *DB[ValueType]) AssignEntries() {
	if assigned := db.entries; assigned == nil {
		db.entries = make(map[KeyType]ValueType)
	}
}

func (db *DB[ValueType]) HasKey(key KeyType) bool {
	_, found := db.entries[key]
	return found
}

func (db *DB[ValueType]) Insert(key KeyType, value ValueType) {
	db.entries[key] = value
}

func (db *DB[ValueType]) Find(key KeyType) (ValueType, error) {
	ok := db.HasKey(key)
	if ok {
		return db.entries[key], nil
	} else {
		var empty ValueType
		return empty, fmt.Errorf("value with key=%d not found", key)
	}
}

func (db *DB[ValueType]) Delete(key KeyType) error {
	ok := db.HasKey(key)
	if ok {
		delete(db.entries, key)
		return nil
	} else {
		return fmt.Errorf("value with key=%d not found", key)
	}
}

func (db *DB[ValueType]) Values() []ValueType {
	var values []ValueType
	for _, value := range db.entries {
		values = append(values, value)
	}
	return values
}
