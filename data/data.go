package data

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
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

type Vocabs struct {
	Words  []Vocab `json:"words"`
	NextID int     `json:"next_id"`
}

type Vocab struct {
	Id         int       `json:"id"`
	Arabic     string    `json:"arabic"`
	English    string    `json:"english"`
	Created    time.Time `json:"created"`
	LastEdited time.Time `json:"last_edited"`
}

// getNextId keeps track of the id.
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *Vocabs) getNextID() int {
	ni := vo.NextID
	vo.NextID++
	return ni
}

// find finds the arabic word if rmharakats == true then the harakats are removed
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *Vocabs) find(n string, respectHarakts bool) bool {
	if !respectHarakts {
		n = removeHarakats(n)
	}

	for _, v := range vo.Words {
		hey := v.Arabic
		if !respectHarakts {
			hey = removeHarakats(v.Arabic)
		}
		if hey == n {
			return true
		}
	}
	return false
}

// add adds the Vocab to the struct or returns an error.
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *Vocabs) add(ar, eng string, respectHarakts bool) error {
	ar = strings.TrimSpace(ar)
	eng = strings.TrimSpace(eng)

	if ar == "" {
		return ErrAaFieldisEmpty
	} else if eng == "" {
		return ErrEngFieldisEmpty
	}

	if vo.find(ar, respectHarakts) {
		return ErrWordExists
	}

	vo.Words = append(vo.Words, Vocab{Id: vo.getNextID(), Arabic: ar, English: eng, Created: time.Now(), LastEdited: time.Now()})
	return nil
}

// remove removes finds and removes the specefied id or returns an error.
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *Vocabs) remove(id int) error {
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

// edits the Vocab with the given id
//
// it's not threadsafe. I should be called in a threadsafe func.
func (vo *Vocabs) edit(id int, ar, eng string, lastEdited time.Time) error {
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
func (vo *Vocabs) addAndSaveFile(path, ar, eng string, respoectHarakats bool) error {
	if err := vo.add(ar, eng, respoectHarakats); err != nil {
		return err
	}

	return vo.saveToFile(os.WriteFile, path)
}

// removeAndSaveFile appends the ar, eng to the db and saves it to the database
func (vo *Vocabs) removeAndSaveFile(path string, id int) error {
	if err := vo.remove(id); err != nil {
		return err
	}

	return vo.saveToFile(os.WriteFile, path)
}

// editAndSaveFile appends the ar, eng to the db and saves it to the database
func (vo *Vocabs) editAndSaveFile(path string, id int, ar, eng string) error {
	if err := vo.edit(id, ar, eng, time.Now()); err != nil {
		return err
	}

	return vo.saveToFile(os.WriteFile, path)
}

func (vo *Vocabs) saveToFile(write writer, path string) error {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	data, err := json.MarshalIndent(vo, "", "\t")
	if err != nil {
		return err
	}

	return write(path, data, readWritePermission)
}

func ReadFromFile(path string) (*Vocabs, error) {
	return readFromFile(os.ReadFile, path)
}

func readFromFile(read reader, path string) (*Vocabs, error) {
	dataMutex.RLock()
	defer dataMutex.RUnlock()

	data, err := read(path)
	if err != nil {
		return nil, err
	}

	var vo Vocabs
	if err = json.Unmarshal(data, &vo); err != nil {
		return nil, err
	}

	return &vo, nil
}
