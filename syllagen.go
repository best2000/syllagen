package main

import (
	"fmt"
	"math/rand"
	"time"
)

var alpha = []byte("bcdfghjklmnprstvxyz")
var vowel = []byte("aeiou")
var alphaTh = []rune("กขคงจฉชซดตถทนบปผฝพฟมยรลวสหอฮ")
var vowelTh = []rune("ุูิีึืาโใไะำอเ")
var toneTh = []rune("่้๊๋")

type sylla struct {
	genT      int
	syllabyte []byte
	mem       []byte
}

func (s *sylla) syllagen() {
	s.genT = ranInt(2, 5)
	for i := 0; i < s.genT; i++ {
		fmt.Println(string(s.syllabyte))
		n := ranInt(1, 4)
		if n == 1 {
			temp := av()
			for i := range temp {
				s.syllabyte = append(s.syllabyte, temp[i])
			}
		} else if n == 2 {
			temp := ava()
			for i := range temp {
				s.syllabyte = append(s.syllabyte, temp[i])
			}
		} else if n == 3 {
			temp := avv()
			for i := range temp {
				s.syllabyte = append(s.syllabyte, temp[i])
			}
		} else {
			temp := vav()
			for i := range temp {
				s.syllabyte = append(s.syllabyte, temp[i])
			}
		}
	}
	s.syllabyte[0] = s.syllabyte[0] - 32
	s.mem = s.syllabyte
	s.syllabyte = []byte{}
}

func main() {
	var s sylla
	s.syllagen()
	fmt.Println(string(s.mem))
	s.syllagen()
	fmt.Println(string(s.mem))
	//x := []rune("ใไโ")
	//fmt.Println(x)
	//for i := 1; i < 4; i++ {
	//	fmt.Print(string(ranTh()))
	//}

}

//Thai rune structure => [alpha, vowel, tone(!im.), spellAlpha(!im.)]
func ranTh() []rune {
	var re []rune
	re = append(re, alphaTh[ranInt(0, len(alphaTh)-1)], 0)
	voRune := vowelTh[ranInt(0, len(vowelTh)-1)]
	//vowel check
	//fmt.Println(1, re)
	switch voRune {
	//case 'โ' 'ไ' 'ใ'
	case 3650, 3651, 3652:
		re[1] = re[0]
		re[0] = voRune
	default:
		re[1] = voRune
	}
	//tone chance
	toneChance := ranInt(0, 10)
	switch toneChance % 2 {
	//yes chance
	case 1:
		re = append(re, 0)
		//fmt.Println(2, re)
		switch re[1] {
		case 3635:
			re[2] = 3635
			re[1] = toneTh[ranInt(0, len(toneTh)-1)]
		case 'อ':
			re[2] = 'อ'
			re[1] = toneTh[ranInt(0, len(toneTh)-1)]
		case 'า':
			re[2] = 'า'
			re[1] = toneTh[ranInt(0, len(toneTh)-1)]
		case 'ะ':
			re[2] = 'ะ'
			re[1] = toneTh[ranInt(0, len(toneTh)-1)]
		default:
			re[2] = toneTh[ranInt(0, len(toneTh)-1)]
		}
	}
	//fmt.Println(3, re)
	//spelling alpha check
	switch len(re) {
	case 2:
		if re[1] == 'ะ' || re[1] == 3635 || re[0] == 'ใ' || re[0] == 'ไ' || re[0] == 'โ'{
			return re
		}
	case 3:
		if re[1] == 'ะ' || re[2] == 'ะ' || re[1] == 3635 || re[2] == 3635 || re[0] == 'ใ' || re[0] == 'ไ' || re[0] == 'โ'{
			return re
		}
	}
	//spelling alpha chance
	spellAlphaChance := ranInt(0, 10)
	//fmt.Println(4, re)
	switch spellAlphaChance % 2 {
	//yes chance
	case 1:
		spellAlphaRune := alphaTh[ranInt(0, len(alphaTh)-1)]
		switch spellAlphaRune {
		case 'อ':
			return re
		default:
			re = append(re, spellAlphaRune)
		}
	//no chance
	default:
		switch re[1] {
		//case ' ื'
		case 3639:
			re = append(re, 3629)
		}
	}
	return re
}

//2 alpha
func av() []byte {
	re := make([]byte, 2)
	re[0] = alpha[ranInt(0, len(alpha)-1)]
	re[1] = vowel[ranInt(0, len(vowel)-1)]
	return re
}

//3 alpha
func ava() []byte {
	re := make([]byte, 3)
	re[0] = alpha[ranInt(0, len(alpha)-1)]
	re[1] = vowel[ranInt(0, len(vowel)-1)]
	re[2] = alpha[ranInt(0, len(alpha)-1)]
	return re
}

//3 alpha
func vav() []byte {
	re := make([]byte, 3)
	re[0] = vowel[ranInt(0, len(vowel)-1)]
	re[1] = alpha[ranInt(0, len(alpha)-1)]
	re[2] = vowel[ranInt(0, len(vowel)-1)]
	return re
}

//3 alpha
func avv() []byte {
	re := make([]byte, 3)
	re[0] = alpha[ranInt(0, len(alpha)-1)]
	re[1] = vowel[ranInt(0, len(vowel)-1)]
	re[2] = vowel[ranInt(0, len(vowel)-1)]
	return re
}

func ranInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * 100)
	return min + rand.Intn(max-min)
}
