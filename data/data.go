package data

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
)

const (
	readWritePermission = 0666
)

var (
	ErrWordExists      = errors.New("word is already in the database")
	ErrIdDontExists    = errors.New("id don't exists")
	ErrAaFieldisEmpty  = errors.New("arabic field is empty")
	ErrEngFieldisEmpty = errors.New("english field is empty")
)

// to be thead safe no race conditions
var dataMutex sync.RWMutex

type writer func(name string, data []byte, perm os.FileMode) error
type reader func(name string) ([]byte, error)

type vocabs struct {
	Words  []vocab `json:"words"`
	NextID int     `json:"next_id"`
}

type vocab struct {
	Id         int       `json:"id"`
	Arabic     string    `json:"arabic"`
	English    string    `json:"english"`
	Created    time.Time `json:"created"`
	LastEdited time.Time `json:"last_edited"`
}

// getNextId keeps track of the id.
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *vocabs) getNextID() int {
	ni := vo.NextID
	vo.NextID++
	return ni
}

// find finds the arabic word if rmharakats == true then the harakats are removed
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *vocabs) find(n string, rmHarakats bool) bool {
	if rmHarakats {
		n = removeHarakats(n)
	}

	for _, v := range vo.Words {
		hey := v.Arabic
		if rmHarakats {
			hey = removeHarakats(v.Arabic)
		}
		if hey == n {
			return true
		}
	}
	return false
}

// add adds the vocab to the struct or returns an error.
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *vocabs) add(ar, eng string) error {
	if ar == "" {
		return ErrAaFieldisEmpty
	} else if eng == "" {
		return ErrEngFieldisEmpty
	}

	if vo.find(ar, true) {
		return ErrWordExists
	}

	vo.Words = append(vo.Words, vocab{Id: vo.getNextID(), Arabic: ar, English: eng, Created: time.Now(), LastEdited: time.Now()})
	return nil
}

// remove removes finds and removes the specefied id or returns an error.
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *vocabs) remove(id int) error {
	for i, v := range vo.Words {
		if v.Id == id {
			// remove it here
			for i := i; i < len(vo.Words)-1; i++ {
				vo.Words[i] = vo.Words[i+1]
			}
			vo.Words = vo.Words[:len(vo.Words)-1]
			return nil
		}
	}

	return ErrIdDontExists
}

// edits the vocab with the given id
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *vocabs) edit(id int, ar, eng string, lastEdited time.Time) error {
	if ar == "" {
		return ErrAaFieldisEmpty
	} else if eng == "" {
		return ErrEngFieldisEmpty
	}

	for i := 0; i < len(vo.Words); i++ {
		if vo.Words[i].Id == id {
			vo.Words[i].Arabic = ar
			vo.Words[i].English = eng
			vo.Words[i].LastEdited = lastEdited
			return nil
		}
	}

	return ErrIdDontExists
}

// saveToFile appends the ar, eng to the db and saves it to the database
func (vo *vocabs) addAndSaveFile(path, ar, eng string) error {
	if err := vo.add(ar, eng); err != nil {
		return err
	}

	return vo.saveToFile(os.WriteFile, path)
}

func (vo *vocabs) saveToFile(write writer, path string) error {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	data, err := json.Marshal(vo)
	if err != nil {
		return err
	}

	return write(path, data, readWritePermission)
}

func ReadFromFile(path string) (*vocabs, error) {
	return readFromFile(os.ReadFile, path)
}

func readFromFile(read reader, path string) (*vocabs, error) {
	dataMutex.RLock()
	defer dataMutex.RUnlock()

	data, err := read(path)
	if err != nil {
		return nil, err
	}

	var vo vocabs
	if err = json.Unmarshal(data, &vo); err != nil {
		return nil, err
	}

	return &vo, nil
}
