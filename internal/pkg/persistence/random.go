package persistence

import (
	"fmt"
	"math/rand"
	"time"
)

type RandRepository struct{}

var randRepository *RandRepository

func Rand() *RandRepository {
	if randRepository == nil {
		randRepository = &RandRepository{}
	}
	return randRepository
}

func (r *RandRepository) RandomString(randText string, resultLength int, useDupText bool) string {
	//ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+
	//1234567890
	Result := ""
	s := rand.NewSource(time.Now().UnixNano())
	rs := rand.New(s)
	RandResult := fmt.Sprintf("%d", rs.Int())

	LoopTo := 0
	if resultLength == 0 {
		LoopTo = len(RandResult)
	} else {
		LoopTo = resultLength
	}
	for i := 0; i < LoopTo; i++ {
		SelectDigit := int(RandResult[i] - '0')
		if len(randText) <= SelectDigit {
			SelectDigit = len(randText) - 1
		}

		Result += string(randText[SelectDigit])
		if !useDupText {
			randText = randText[:SelectDigit] + randText[SelectDigit+1:]
		}
	}
	if resultLength == 0 {
		Result += randText
	}

	return Result
}

func (r *RandRepository) RandomTarot() int {
	s := rand.NewSource(time.Now().UnixNano())
	rs := rand.New(s)
	return rs.Intn(78)
}

func (r *RandRepository) RandomTarotNoDup(times int) []int {
	s := rand.NewSource(time.Now().UnixNano())
	rs := rand.New(s)
	numbers := make([]int, 0)
	used := make(map[int]bool)

	for len(numbers) < times {
		num := rs.Intn(78)
		if !used[num] {
			numbers = append(numbers, num)
			used[num] = true
		}
	}

	return numbers
}

func (r *RandRepository) Random01() int {
	s := rand.NewSource(time.Now().UnixNano())
	rs := rand.New(s)
	return rs.Intn(2)
}
