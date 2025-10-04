package commonUtils

import (
	"encoding/json"
	"fmt"
)

const maxLength = 1024

func MaxLength(s string) string {
	if len(s) > maxLength {
		s = fmt.Sprintf("%s...", s[:maxLength])
	}
	return s
}

func ForceString(model interface{}) string {
	data, err := json.Marshal(model)
	if err != nil {
		return fmt.Sprintf("%+v", model)
	}
	return string(data)
}

func ForceStringMaxLength(model interface{}) string {
	return MaxLength(ForceString(model))
}
