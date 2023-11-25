package version

import "fmt"

var (
	version = "ad-hoc"
)

func ShowVersion() {
	fmt.Println(version)
}
