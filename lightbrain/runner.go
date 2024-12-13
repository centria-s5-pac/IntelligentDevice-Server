package lightbrain

import (
	"fmt"
	"time"
)

func Main() {
	initValue()

	for {
		SetValue(GetValue() + 1)
		time.Sleep(time.Second)

		data := GetJson()

		if data == nil {
			continue
		}

		fmt.Println("Data:", data)

	}
}
