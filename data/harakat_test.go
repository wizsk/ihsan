package data

import "testing"

func TestRemoveHarakats(t *testing.T) {
	tests := []struct {
		inp, out string
	}{
		{
			inp: "الم",
			out: "الم",
		},
		{
			inp: "الَمِ مَمُمٍكُتُ",
			out: "الم مممكت",
		},
		{
			inp: "الم",
			out: "الم",
		},
	}

	for _, tt := range tests {
		out := removeHarakats(tt.inp)
		if tt.out != out {
			t.Errorf("tt.out:%s != out:%s", tt.out, out)
			t.FailNow()
		}
	}
}
