package template

import (
	"encoding/json"
	"fmt"
)

func RenderJSON() string {
	marshalled, err := json.Marshal(originalOutput)
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(string(marshalled))
}
