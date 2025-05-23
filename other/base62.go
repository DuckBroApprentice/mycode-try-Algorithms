package other

func Base62(randomNum int) string {
	trans := make(map[int]string)
	trans = map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
		16: "g",
		17: "h",
		18: "i",
		19: "j",
		20: "k",
		21: "l",
		22: "m",
		23: "n",
		24: "o",
		25: "p",
		26: "q",
		27: "r",
		28: "s",
		29: "t",
		30: "u",
		31: "v",
		32: "w",
		33: "x",
		34: "y",
		35: "z",
		37: "A",
		38: "B",
		39: "C",
		40: "D",
		41: "E",
		42: "F",
		43: "G",
		44: "H",
		45: "I",
		46: "J",
		47: "K",
		48: "L",
		49: "M",
		50: "N",
		51: "O",
		52: "P",
		53: "Q",
		54: "R",
		55: "S",
		56: "T",
		57: "U",
		58: "V",
		59: "W",
		60: "X",
		61: "Y",
		62: "Z",
	}
	if randomNum/62 == 0 {
		num := randomNum % 62
		if num == 0 {
			return ""
		}
		return trans[num]
	}

	return Base62(randomNum/62) + trans[randomNum%62]
}
