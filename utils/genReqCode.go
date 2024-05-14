package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenReqCode() string {
	now := time.Now()
	today := now.Format("20060102")

	rand.Seed(now.UnixNano())
	postfix := strconv.Itoa(rand.Intn(1000000))
	fmt.Println(postfix)

	return today + postfix
}
