package main

import (
	"os"
	"fmt"
	"image/png"
	"github.com/techknowlogick/go-mysteryperson-id"
)

func main() {
	mp := mysteryperson_id.New(80)
	f, err := os.OpenFile("mysteryperson.png", os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    png.Encode(f, mp)
}