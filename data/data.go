package data

import (
	"errors"
	"sync"
)

var (
	ErrWordExists = errors.New("word is already in the database")
)

var dataMutex sync.RWMutex

type Vocabs struct {
	Words  []Vocab
	NextID uint64 `json:"next_id"`
}

type Vocab struct {
	Id      uint64
	Arabic  string
	English string
}

// func MarshalVocabs(v *Vocabs) ([]byte, error) {
// 	return json.Marshal(v)
// }

// func UnmarshalVocabs(v []byte) (*Vocabs, error) {
// 	var vocabs Vocabs
// 	if err := json.Unmarshal(v, &vocabs); err != nil {
// 		return nil, err
// 	}
// 	return &vocabs, nil
// }

func (v *Vocabs) find(n string) bool {
	return v.findEx(n, true)
}

// c = remove harakats
func (vo *Vocabs) findEx(n string, rmHarakats bool) bool {
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

func (vo *Vocabs) getNextID() uint64 {
	ni := vo.NextID
	vo.NextID++
	return ni
}

func (vo *Vocabs) Add(ar, eng string) error {
	// to be thread safe
	dataMutex.Lock()
	defer dataMutex.Unlock()

	if vo.find(ar) {
		return ErrWordExists
	}

	vo.Words = append(vo.Words, Vocab{Id: vo.getNextID(), Arabic: ar, English: eng})
	return nil
}
