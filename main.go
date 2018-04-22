package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"flag"
	"os"
	"io"
	"jmcs/core/utils/net/port"
	"jmcs/app/routers"
)

func main() {


	socketServcer()		//socket服务


	routers.Router.Run(":8000")

}

func logConfig()  {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}

func socketServcer()  {
	flag.Parse()
	socketPort := port.Port(*sPort)
	if socketPort.CheckEnabled(nil) {
		fmt.Errorf("error:%s已被占用",socketPort)
	}


}
