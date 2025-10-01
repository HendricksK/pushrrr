package helpers

import (
	"fmt"
	"os"
)

func getEnv(key string) string {

	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Printf("%s not set\n", key)
	} else {
		fmt.Printf("%s=%s\n", key, val)
	}

	return val
}
