package main

import (
	"fmt"
	"github.com/PhamDuyKhang/autogermanyVoz/app"
)

func main() {
	cokie,_ := app.LoginVOZ("","")
	fmt.Println(cokie)

}
