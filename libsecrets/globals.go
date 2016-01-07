package libsecrets

import (
	"fmt"
	"os"
)

// Dir returns the directory where the secrets should be stored
func Dir() string {
	pwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return pwd + "/.secrets"
}

// DirExists does the secret directory exist?
func DirExists() bool {
	_, err := os.Stat(Dir())
	return (err == nil)
}
