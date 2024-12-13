package lightbrain

import (
	"time"
)

func Main() {
	initValue()

	for {
		SetValue(GetLightLevel())
		time.Sleep(time.Second)
	}
}
