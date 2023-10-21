package golangembed

import (
	"embed" // jika lebih gunakan yang ini.
	_ "embed" // digunakan jika anda hanya memanggil //go:embed <filename> saja
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed hello.txt
var Hello string // tipe datanya harus menyesuaikan dengan isi filenya.

//go:embed image/said.jpeg
var Image []byte // karena gambar bertipe []byte

//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

//go:embed files/*.txt
var path embed.FS

func TestHello(t *testing.T) {
	fmt.Println(Hello)
}

func TestImage(t *testing.T) {
	fmt.Println(Image)
}

func TestCopyImage(t *testing.T) {
	err := os.WriteFile("image/said.copy.jpeg", Image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestMultipleFiles(t *testing.T) {
	a, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}

	b, err := files.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(a))
	fmt.Println(string(b))
}

func TestPathMatcher(t *testing.T) {
	directories, _ := path.ReadDir("files")
	for _, file := range directories {
		if !file.IsDir() {
			text, _ := path.ReadFile("files/" + file.Name())
			fmt.Println(string(text))
		}
	}
}

