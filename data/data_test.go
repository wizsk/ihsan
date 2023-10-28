package data

import (
	"fmt"
	"testing"
	"time"
)

func TestData(t *testing.T) {
	md := mockData()
	go func() {
		md.Add("مَرْحَبًا", "Hello")
		fmt.Println("1, I'm done")
	}()
	go func() {
		md.Add("مَرْحَبًا", "Hello")
		fmt.Println("2, I'm done")
	}()
	go func() {
		md.Add("مَرْحَبًا", "Hello")
		fmt.Println("3, I'm done")
	}()
	// err := md.Add("مَرْحَبًا", "Hello")
	// t.Log(err)
	time.Sleep(10*time.Second)
}

func mockData() Vocabs {
	return Vocabs{
		Words: []Vocab{
			{1, "مرحبًا", "Hello"},
			{2, "شكرًا", "Thank you"},
			{3, "ماء", "Water"},
			{4, "سماء", "Sky"},
			{5, "كتاب", "Book"},
		},
	}
}

func _mockDataWithHarakats() Vocabs {
	return Vocabs{
		Words: []Vocab{
			{1, "مَرْحَبًا", "Hello"},
			{2, "شُكْرًا", "Thank you"},
			{3, "مَاء", "Water"},
			{4, "سَمَاء", "Sky"},
			{5, "كِتَاب", "Book"},
		},
	}
}
