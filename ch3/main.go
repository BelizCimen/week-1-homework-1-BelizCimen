package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	YiB
)

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Printf("%T %[1]v\n", KiB)
	fmt.Printf("%T %[1]v\n", MiB)
	fmt.Printf("%T %[1]v\n", GiB)
	fmt.Printf("%T %[1]v\n", TiB)
	fmt.Printf("%T %[1]v\n", PiB)
	fmt.Printf("%T %[1]v\n", EiB)

	fmt.Printf("%T %[1]v\n", KB)
	fmt.Printf("%T %[1]v\n", MB)
	fmt.Printf("%T %[1]v\n", GB)
	fmt.Printf("%T %[1]v\n", TB)
	fmt.Printf("%T %[1]v\n", PB)
	fmt.Printf("%T %[1]v\n", EB)
}
