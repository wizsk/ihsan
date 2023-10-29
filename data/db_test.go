package data

import (
	"os"
	"testing"
)

// db test will be derectly written to disk!

func TestOpenDb(t *testing.T) {
	file := "tmp.json"
	db, err := OpenJDB(file)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err := db.Add("نسبsd`f", "idk"); err != nil {
		t.Error(err)
		t.FailNow()

	}

	if err := db.Edit(db.data.NextID-1, "نسبsd`f", "idk"); err != nil {
		t.Error(err)
		t.FailNow()

	}

	if err := db.Remove(db.data.NextID - 1); err != nil {
		t.Error(err)
		t.FailNow()

	}

	os.Remove(file)
}
