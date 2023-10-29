package data

import (
	"encoding/json"
	"testing"
)

func TestData(t *testing.T) {
	md := mockData()
	t.Logf("%+v", md)
	d, err := json.Marshal(md)
	t.Logf("%+s, %v", d, err)
}

func mockData() vocabs {
	v := []vocab{
		{1, "مرحبًا", "Hello"},
		{2, "شكرًا", "Thank you"},
		{3, "ماء", "Water"},
		{4, "سماء", "Sky"},
		{5, "كتاب", "Book"},
	}

	vo := vocabs{}
	for _, v := range v {
		err := vo.add(v.Arabic, v.English)
		if err != nil {
			panic(err)
		}
	}

	return vo

}

func _mockDataWithHarakats() vocabs {
	return vocabs{
		Words: []vocab{
			{1, "مَرْحَبًا", "Hello"},
			{2, "شُكْرًا", "Thank you"},
			{3, "مَاء", "Water"},
			{4, "سَمَاء", "Sky"},
			{5, "كِتَاب", "Book"},
		},
	}
}
