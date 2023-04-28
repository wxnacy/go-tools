package gotool

import (
	"fmt"
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var _rand = rand.New(randSource)

func IDGen() string {
	now := time.Now()
	return Md5(fmt.Sprintf("%d%d", now.UnixNano(), _rand.Int63n(now.Unix())))
}
