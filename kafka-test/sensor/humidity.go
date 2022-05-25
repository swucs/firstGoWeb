package sensor

import (
	"math/rand"
	"strconv"
	"time"
)

func Humidity() [10]string {
	var val [10]string
	for i := 0; i < len(val); i++ {
		rand.Seed(time.Now().UnixNano())
		val[i] = "humid : " + strconv.Itoa(rand.Intn(100))
	}
	return val
}
