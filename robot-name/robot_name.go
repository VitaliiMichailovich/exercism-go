package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const testVersion = 1

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Robot string

var robots = make(map[Robot]bool)

func (robot *Robot) Name() string {
	rand.Seed(time.Now().UnixNano())
	for inUse, ok := string(*robot) == "", true; ok && inUse; inUse, ok = robots[*robot] {
		*robot = Robot(name())
	}
	robots[*robot] = true
	return string(*robot)
}

func (robot *Robot) Reset() {
	*robot = Robot("")
}

func name() string {
	let1 := rand.Intn(len(letters))
	let2 := rand.Intn(len(letters))
	return letters[let1:let1+1] + letters[let2:let2+1] + fmt.Sprintf("%03d", rand.Intn(1000))
}