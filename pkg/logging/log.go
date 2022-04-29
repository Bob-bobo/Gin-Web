package logging

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F1 *os.File
	F2 *os.File
	F3 *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger1 *log.Logger
	logger2 *log.Logger
	logger3 *log.Logger

	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup initialize the log instance
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()

	F1, err = file.MustOpen(fileName, filePath+"ERROR/")
	F2, err = file.MustOpen(fileName, filePath+"INFO/")
	F3, err = file.MustOpen(fileName, filePath+"WARN/")

	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger1 = log.New(F1, DefaultPrefix, log.LstdFlags)
	logger2 = log.New(F2, DefaultPrefix, log.LstdFlags)
	logger3 = log.New(F3, DefaultPrefix, log.LstdFlags)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger1.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger2.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger1.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger3.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger1.Fatalln(v)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	switch level {
	case 1:
		logger2.SetPrefix(logPrefix)
	case 2:
		logger1.SetPrefix(logPrefix)
	case 3:
		logger3.SetPrefix(logPrefix)
	}
}
