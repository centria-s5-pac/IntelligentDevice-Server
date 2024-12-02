package lightbrain

import (
	"time"
)

func Main() {
	initValue()

	for {
		SetValue(GetValue() + 1)
		time.Sleep(time.Second)
	}
}
