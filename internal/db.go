package internal

import "os"

func init() {
	os.Setenv("TOY_ENV", "test")

}
