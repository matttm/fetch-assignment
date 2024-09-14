package utilities

import (
	"fmt"
	"os"
)

func Log(s string) {
	if os.Getenv("LOGGING") == "true" {
		fmt.Println(s)
	}
}
