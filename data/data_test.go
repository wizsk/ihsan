package data

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

// Testting (*vocab) add(...) is pointless
//
// if add working correnctly then find sould also work tho
// func TestVocabAdd(t *testing.T) {
// 	vo := Vocabs{}
// 	for _, v := range []struct{ ara, eng string }{{"مَرْحَبًا", "Hello"}, {"شُكْرًا", "Thank you"}, {"مَاء", "Water"}, {"سَمَاء", "Sky"}, {"كِتَاب", "Book"}} {
// 		// respect harakts
// 		if err := vo.add(v.ara, v.eng, false); err != nil {
// 			t.Errorf("expted '%v' to be nil", err)
// 			t.FailNow()
// 		}
// 	}

//		for _, x := range []struct{ ara, eng string }{{"مرحبًا", "Hello"}, {"شكرًا", "Thank you"}, {"ماء", "Water"}, {"سماء", "Sky"}, {"كتاب", "Book"}} {
//			err := vo.add(x.ara, x.eng, false)
//			if !errors.Is(err, ErrWordExists) {
//				t.Errorf("expected '%v' to be %v", err, ErrWordExists)
//				t.FailNow()
//			}
//		}
//	}
func TestVocabFind(t *testing.T) {
	vo := Vocabs{}
	for _, v := range []struct{ ara, eng string }{{"مَرْحَبًا", "Hello"}, {"شُكْرًا", "Thank you"}, {"مَاء", "Water"}, {"سَمَاء", "Sky"}, {"كِتَاب", "Book"}} {
		if err := vo.add(v.ara, v.eng, false); err != nil {
			panic(err) // this should not happendd
		}
	}

	// should be eble to find it so, if it returns nil then
	// find didn't find anyting got it!
	for _, x := range []string{"مرحبًا", "شكرًا", "ماء", "سماء", "كتاب"} {
		if vo.find(x, false) == nil {
			t.Errorf("expected '%v' to be in the Vocabs", x)
			t.FailNow()

		}
	}

	// should not find so it should return nil
	for _, x := range []struct{ ara, eng string }{{"مِرْحَبًا", "Hello"}, {"شِكْرًا", "Thank you"}, {"مَاءٌ", "Water"}, {"سَمَاءِ", "Sky"}, {"كِتَابَ", "Book"}} {
		if vo.find(x.ara, true) != nil {
			t.Errorf("expected '%v' not to be in the Vocabs", x)
			t.FailNow()

		}
	}

	// these should not be in the "fonund"
	for _, x := range []string{"محبًا", "شكا", "اء", "مء", "كتب", "foooo", "bar", "হি হি হি"} {
		if vo.find(x, true) != nil {
			t.Errorf("expected '%v' not to be in the Vocabs", x)
			t.FailNow()

		}
	}
}

func TestVocabRemove(t *testing.T) {
	pre := Vocabs{
		Words: []Vocab{
			{Id: 1, Arabic: "مَرْحَبًا", English: "Hello"},
			{Id: 2, Arabic: "شُكْرًا", English: "Thank you"},
			{Id: 3, Arabic: "مَاء", English: "Water"},
			{Id: 4, Arabic: "سَمَاء", English: "Sky"},
			{Id: 5, Arabic: "كِتَاب", English: "Book"},
		},
	}

	err := pre.remove(2)
	if err != nil {
		t.Errorf("expected '%v' to be nil", err)
	}
	err = pre.remove(1)
	if err != nil {
		t.Errorf("expected '%v' to be nil", err)
	}
	err = pre.remove(5)
	if err != nil {
		t.Errorf("expected '%v' to be nil", err)
	}

	post := Vocabs{
		Words: []Vocab{
			{Id: 3, Arabic: "مَاء", English: "Water"},
			{Id: 4, Arabic: "سَمَاء", English: "Sky"},
		},
	}

	if !reflect.DeepEqual(pre.Words, post.Words) {
		t.Error("expeted to be equal")
		t.Errorf("pre: %+v", pre.Words)
		t.Errorf("post: %+v", post.Words)
		t.FailNow()
	}

	err = pre.remove(2000)
	if !errors.Is(err, ErrIdDontExists) {
		t.Errorf("expected '%v' to be non nil", err)
	}
}

func TestVocabEdit(t *testing.T) {
	vo := Vocabs{}
	for _, v := range []struct{ ara, eng string }{{"مَرْحَبًا", "Hello"}, {"شُكْرًا", "Thank you"}, {"مدرسة", "School"}, {"مَاء", "Water"}, {"سَمَاء", "Sky"}, {"كِتَاب", "Book"}} {
		if err := vo.add(v.ara, v.eng, false); err != nil {
			panic(err) // this should not happendd
		}
	}

	// copying it
	res := Vocabs{NextID: vo.NextID}
	res.Words = append(res.Words, vo.Words...)

	for i, v := range []struct{ ara, eng string }{{"قلم", "Pen"}, {"شجرة", "Tree"}, {"كرة", "Ball"}, {"سماء", "Sky"}, {"جمل", "Camel"}} {
		modify_time := time.Now()
		if err := vo.edit(i, v.ara, v.eng, modify_time); err != nil {
			t.Error(err)
			t.FailNow()
		}
		res.Words[i].Arabic = v.ara
		res.Words[i].English = v.eng
		res.Words[i].LastEdited = modify_time
	}

	if !reflect.DeepEqual(vo, res) {
		t.Errorf("expedted to be equeal:\n%+v\n%+v", vo, res)
		t.FailNow()
	}
	if err := vo.edit(1000, "brrr", "fo", time.Now()); !errors.Is(err, ErrIdDontExists) {
		t.Error(err)
		t.FailNow()
	}
}

/*
	 Vocabs{
		Words: []Vocab{
			{1, "مَرْحَبًا", "Hello"},
			{2, "شُكْرًا", "Thank you"},
			{3, "مَاء", "Water"},
			{4, "سَمَاء", "Sky"},
			{5, "كِتَاب", "Book"},
		},
	}

*/
