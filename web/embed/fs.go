package embed

import "io/fs"

var Public fs.FS

func init() {
	var err error

	Public, err = fs.Sub(Files, "public")
	if err != nil {
		panic(err)
	}
}
