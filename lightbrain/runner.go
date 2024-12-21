package lightbrain

import (
	"fmt"
	"time"
)

func Main() {
	initValue()

	for {
		fmt.Println("Light level:", GetLightLevel())
		SetValue(GetLightLevel() * 2)
		time.Sleep(time.Second)
	}
}
