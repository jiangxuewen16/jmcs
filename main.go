package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"jmcs/core"
)

func main() {

	core.Run()

}

func logConfig()  {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}
