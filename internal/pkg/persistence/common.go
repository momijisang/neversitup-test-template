package persistence

import (
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func SetGridHeader(Text, Align, Value string, Sortable bool) interface{} {
	return map[string]interface{}{
		"text":     Text,
		"align":    Align,
		"sortable": Sortable,
		"value":    Value,
	}
}

func GetDayName(Day time.Time) string {
	DayName := ""
	switch Day.Weekday() {
	case 0:
		DayName = "อาทิตย์"
	case 1:
		DayName = "จันทร์"
	case 2:
		DayName = "อังคาร"
	case 3:
		DayName = "พุธ"
	case 4:
		DayName = "พฤหัสฯ"
	case 5:
		DayName = "ศุกร์"
	case 6:
		DayName = "เสาร์"
	}
	return DayName
}

func GetDayNameEn(Day time.Time) string {
	DayName := ""
	switch Day.Weekday() {
	case 0:
		DayName = "sunday"
	case 1:
		DayName = "monday"
	case 2:
		DayName = "tuesday"
	case 3:
		DayName = "wednesday"
	case 4:
		DayName = "thursday"
	case 5:
		DayName = "friday"
	case 6:
		DayName = "saturday"
	}
	return DayName
}

func GetDayIndex(dayName string) int {
	index := -1
	switch dayName {
	case "sunday":
		index = 0
	case "monday":
		index = 1
	case "tuesday":
		index = 2
	case "wednesday":
		index = 3
	case "thursday":
		index = 4
	case "friday":
		index = 5
	case "saturday":
		index = 6
	}
	return index
}

func FindStringInArray(findString string, Array []string) bool {
	for _, s := range Array {
		if s == findString {
			return true
		}
	}
	return false
}

func SplitToArrayInt(InputString string) []int {
	sList := strings.Split(InputString, ",")
	ary := make([]int, len(sList))
	for i, s := range sList {
		x, _ := strconv.ParseInt(s, 10, 64)
		ary[i] = int(x)
	}
	return ary
}

func GetDataInMap(keys string, m map[string]interface{}) interface{} {
	keysList := strings.Split(keys, ".")
	if len(keysList) == 1 {
		if m[keys] == nil {
			return ""
		} else {
			return m[keys]
		}
	} else {
		return GetDataInMap(strings.Join(keysList[1:], "."), m[keysList[0]].(map[string]interface{}))
	}
}

func ConvertInterfaceToMapStringInterface(data interface{}) interface{} {
	bsonDocument, ok := data.(bson.D)
	if ok {
		bsonByte, _ := bson.MarshalExtJSON(bsonDocument, false, false)
		m := make(map[string]interface{})
		_ = json.Unmarshal(bsonByte, &m)
		return m
	}
	return nil
}

func GenerateRandomNumber(addSeed int) int {
	rand.Seed(time.Now().UnixNano() + int64(addSeed))
	return rand.Intn(1000)
}

func RemoveFile(path, fileName string) {
	if fileName != "" {
		os.Remove(path + fileName)
	}
}
