package main

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
)

// func main( {
// arr:= []int{2, }
// log.Printf("Hello",arr, "WWW)
// // log.Pancf("this)
// log.Printf( "This is")
// }

func main() {
	logFile, err := os.Create("/" + time.Now().Format("2006-01-02") + ".log")

	if err != nil {
		fmt.Println(err)

	}
}

func SugarLogss() {
	logger, _ := zap.NewProduction()

}
