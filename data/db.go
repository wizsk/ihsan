package data

import (
	"fmt"
	"os"
	"sync"
)

// - [x] Open
// - [x] Add
// - [x] Remove
// - [x] Edit

// json Databse
type JDB struct {
	path  string // path to json
	data  *Vocabs
	mutex *sync.RWMutex
}

func OpenJDB(path string) (*JDB, error) {
	if path == "" {
		return nil, fmt.Errorf("path can not be empty")
	}

	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		vo := &Vocabs{}
		if err := vo.saveToFile(os.WriteFile, path); err != nil {
			return nil, err
		}

		return &JDB{
			path:  path,
			data:  vo,
			mutex: new(sync.RWMutex),
		}, nil

	} else if err != nil {
		// this is some other err so ya chill && return it
		return nil, err
	}

	// read it
	v, err := readFromFile(os.ReadFile, path)
	if err != nil {
		return nil, err
	}

	db := &JDB{
		path:  path,
		data:  v,
		mutex: new(sync.RWMutex),
	}
	return db, nil
}

func (db *JDB) Add(ar, eng string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.data.addAndSaveFile(db.path, ar, eng)
}

func (db *JDB) Remove(id int) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.data.removeAndSaveFile(db.path, id)
}

func (db *JDB) Edit(id int, arabic, english string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	return db.data.editAndSaveFile(db.path, id, arabic, english)
}

// copies the data and Returns it
func (db *JDB) GetVocabs() Vocabs {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	vo := Vocabs{NextID: db.data.NextID}
	vo.Words = append(vo.Words, db.data.Words...)

	return vo
}
