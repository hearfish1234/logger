package logger

import(
	"testing"
	"fmt"
)

func Test(t *testing.T) {

	logger := SetLogger("./test.log", 1)
	fmt.Println(logger)
	logger.SetLog("test", 1)
	logger.SetLog("test", 2)
	logger.SetLog("test", 3)
	logger.SetLog("test", 4)
	logger.SetLog("test", 5)

}
