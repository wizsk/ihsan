package data

const (
	arabicFathatan rune = 0x064B
	arabicDammatan rune = 0x064C
	arabicKasratan rune = 0x064D
	arabicFatha    rune = 0x064E
	arabicDamma    rune = 0x064F
	arabicKasra    rune = 0x0650
	arabicShadda   rune = 0x0651
	arabicSukun    rune = 0x0652
	// arabicMaadAbove rune = 0x0653
)

func removeHarakats(s string) string {
	res := make([]rune, 0, len(s)/2)

	for _, c := range s {
		if c == arabicFathatan || c == arabicFatha ||
			c == arabicKasratan || c == arabicKasra ||
			c == arabicDammatan || c == arabicDamma ||
			c == arabicShadda || c == arabicSukun {
			continue
		}
		res = append(res, c)
	}

	return string(res)
}
