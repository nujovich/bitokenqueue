//customlogger provides a customized logger to print execution results
package customlogger

import (
	"log"
	"os"
	"sync"
)

type CustomLogger struct {
	filename string
	*log.Logger
}

var (
	customLog CustomLogger
	once      sync.Once
)

//GetInstance() creates an indtance of CustomLogger struct
func GetInstance() CustomLogger {
	once.Do(func() {
		customLog = createLogger("mylogger.log")
	})
	return customLog
}

//createLogger() takes as a param the name of the file that will prompt the INFO messages
//and returns an instance of CustomLogger struct
func createLogger(fname string) CustomLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	customLog := CustomLogger{
		filename: fname,
		Logger:   log.New(file, "INFO: ", log.Lshortfile),
	}
	return customLog
}
