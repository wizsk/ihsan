package data

import (
	"os"
	"sync"
)

// - [x] Open
// - [x] Add
// - [ ] Delete
// - [ ] Update
// - [ ] these functins are needed

// json Databse
type JDB struct {
	path  string // path to json
	data  *vocabs
	mutex *sync.RWMutex
}

func OpenJDB(path string) (*JDB, error) {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		if _, err := os.Create(path); err != nil {
			return nil, err
		}

		return &JDB{
			path:  path,
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

func (db *JDB) Delete(id int) error {
	return nil
}
