package data

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

var (
	ErrWordExists = errors.New("word is already in the database")
)

// to be thead safe no race conditions
var dataMutex sync.RWMutex

type writer func(name string, data []byte, perm os.FileMode) error
type reader func(name string) ([]byte, error)

type vocabs struct {
	Words  []vocab `json:"words"`
	NextID uint64  `json:"next_id"`
}

type vocab struct {
	Id      uint64 `json:"id"`
	Arabic  string `json:"arabic"`
	English string `json:"english"`
}

func (v *vocabs) find(n string) bool {
	return v.findEx(n, true)
}

// c = remove harakats
func (vo *vocabs) findEx(n string, rmHarakats bool) bool {
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

func (vo *vocabs) getNextID() uint64 {
	ni := vo.NextID
	vo.NextID++
	return ni
}

// idk if it needs thead safety... or not
// .. but there is none for now
func (vo *vocabs) add(ar, eng string) error {
	// to be thread safe
	// i don't think This func needs to be theard safe..
	// dataMutex.Lock()
	// defer dataMutex.Unlock()

	if vo.find(ar) {
		return ErrWordExists
	}

	vo.Words = append(vo.Words, vocab{Id: vo.getNextID(), Arabic: ar, English: eng})
	return nil
}

func (vo *vocabs) SaveToFile(path string) error {
	return vo.saveToFile(os.WriteFile, path)
}

func (vo *vocabs) saveToFile(write writer, path string) error {
	dataMutex.Lock()
	defer dataMutex.Unlock()

	data, err := json.Marshal(vo)
	if err != nil {
		return err
	}

	return write(path, data, 0677)
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
