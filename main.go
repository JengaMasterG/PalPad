package main

import (
	"fmt"

	//"palpad/internal/windows"
	"github.com/JengaMasterG/PalPad/app"
)

// TODO: Set mobile layout to HBox for form layouts
//
//	Make layouts for mobile and dekstop

func main() {

	app.Start()

	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
