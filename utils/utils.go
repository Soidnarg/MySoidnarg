package utils

import (
	"encoding/json"
	"os"
)

func GetToken(idToken string) (strToken string) {
	m := make(map[string]string)

	file, err := os.ReadFile(".vscode/tokens.json")
	if err != nil {
		return ""
	}

	err = json.Unmarshal(file, &m)
	if err != nil {
		return ""
	}

	return m[idToken]
}
