package persistence

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

type TestRepository struct{}

var testRepository *TestRepository

func Test() *TestRepository {
	if testRepository == nil {
		testRepository = &TestRepository{}
	}
	return testRepository
}

func (r *TestRepository) Test2(input string) []string {
	var list []string
	r.Permutations(input, 0, &list)
	sort.Strings(list)
	beforeValue := ""
	var result []string
	for _, v := range list {
		if beforeValue != v {
			result = append(result, v)
		}
		beforeValue = v
	}
	return result
}

func (r *TestRepository) Permutations(input string, index int, result *[]string) {
	sLen := utf8.RuneCountInString(input)
	if index == sLen-1 {
		*result = append(*result, input)
		return
	}

	for i := index; i < sLen; i++ {
		s := strings.Split(input, "")
		s[index], s[i] = s[i], s[index]
		r.Permutations(strings.Join(s, ""), index+1, result)
		s[index], s[i] = s[i], s[index]
	}
}

func (r *TestRepository) Test3(input string) string {
	arrInt := SplitToArrayInt(input)
	sort.Ints(arrInt)
	if len(arrInt) == 1 {
		return fmt.Sprintf("%v", arrInt)
	}

	var oddNumber []int
	dupCount := 0
	beforeValue := arrInt[0]
	for i := 1; i < len(arrInt); i++ {
		if beforeValue == arrInt[i] {
			dupCount++
		} else {
			if dupCount%2 == 0 {
				oddNumber = append(oddNumber, beforeValue)
			}
			dupCount = 0
		}
		beforeValue = arrInt[i]
	}
	if dupCount%2 == 0 {
		oddNumber = append(oddNumber, beforeValue)
	}
	return fmt.Sprintf("%v", oddNumber)
}

func (r *TestRepository) Test4(input interface{}) ([]string, int) {
	faceList := input.([]string)
	countFaceOfSmileys := 0
	var smileys []string
	for _, face := range faceList {
		if r.CheckSmileys(strings.Trim(face, " ")) {
			countFaceOfSmileys++
			smileys = append(smileys, face)
		}
	}

	return smileys, countFaceOfSmileys
}

func (r *TestRepository) CheckSmileys(input string) bool {
	input = strings.Replace(input, "'", "", 1)
	eyeOfSmile := []string{":", ";"}
	nose := []string{"-", "~"}
	mouthOfSmile := []string{")", "D"}

	s := strings.Split(input, "")
	trueEye := FindStringInArray(s[0], eyeOfSmile)
	trueNose := (len(s) == 2) || FindStringInArray(s[1], nose)
	trueMouth := FindStringInArray(s[len(s)-1], mouthOfSmile)
	return trueEye && trueNose && trueMouth
}
