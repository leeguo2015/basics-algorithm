package intToRoman

var (
	RomanMap = map[string]int{}
)

//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/integer-to-roman
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type Roman struct {
	Next *Roman
	Last *Roman
	Str  string
	Int  int
}

func IntToRoman(num int) string {

	var roman = new(Roman)
	roman.Str = "I"
	roman.Int = 1
	RomanStr := map[string]int{
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}
	var NextRoman = new(Roman)
	for key, val := range RomanStr {

		NextRoman.Next = &Roman{Str: key, Int: val}
		NextRoman.Next.Last = NextRoman
		NextRoman = NextRoman.Next
	}
	remnant := num
	Result := ""

	for remnant <= 0 {
		NextRoman = roman
		for {
			if remnant > NextRoman.Int && roman.Next != nil {
				remnant -= NextRoman.Int
				Result += NextRoman.Str

			}

			if NextRoman.Last.Int > remnant && remnant > NextRoman.Int && roman.Next != nil {
				remnant -= NextRoman.Int
				Result += NextRoman.Str

			}

			if remnant == NextRoman.Int {
				remnant -= NextRoman.Int
				Result += NextRoman.Str

			}

			if remnant < NextRoman.Int {
				remnant -= NextRoman.Int
				Result += NextRoman.Str

			}

			NextRoman = NextRoman.Next
		}
		//for key, val := range RomanInt {
		//	if val <= remnant && key+1 < len(RomanMap) {
		//		remnant -= val
		//		Result += RomanStr[key]
		//	}
		//
		//	if val > remnant && key+1 < len(RomanMap) {
		//		continue
		//	} else {
		//		remnant -= val
		//		Result += RomanStr[key]
		//	}
		//
		//}

	}
	return ""
}
