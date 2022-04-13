package logs

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func LogSetting() []func() error {
	// ----------------------------------------------
	close := []func() error{}
	logLevel := viper.GetInt("LOG_LEVEL")
	isFile, isPrint := viper.GetBool("LOG_FILEOUT"), viper.GetBool("LOG_STDOUT")
	path := viper.GetString("LOG_FILE_PATH")
	if path == "" {
		path = "log"
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// Default
	Error = log.New(ioutil.Discard, "", log.LstdFlags)
	Warning = log.New(ioutil.Discard, "", log.LstdFlags)
	Info = log.New(ioutil.Discard, "", log.LstdFlags)
	// Setting
	if viper.GetBool("LOG_DEFAULT") {
		output, cls := checkOutput("Default", isPrint, isFile)
		close = append(close, cls)
		log.SetOutput(output)
	}
	if logLevel >= 1 {
		close = append(close, createLog("Error", &Error, isPrint, isFile))
	}
	if logLevel >= 2 {
		close = append(close, createLog("Warning", &Warning, isPrint, isFile))
	}
	if logLevel >= 3 {
		close = append(close, createLog("Info", &Info, isPrint, isFile))
	}
	return close
}

func createLog(name string, logParameters **log.Logger, isPrint, isFile bool) func() error {
	print, close := checkOutput(name, isPrint, isFile)
	if print == nil {
		print = ioutil.Discard
	}
	*logParameters = log.New(print, fmt.Sprintf("【%s】", name), log.Ldate|log.Ltime|log.Lshortfile)
	return close
}

func checkOutput(name string, isPrint, isFile bool) (io.Writer, func() error) {
	var close func() error
	var print io.Writer
	if isFile {
		path := viper.GetString("LOG_PATH")
		if path == "" {
			path = "./log"
		}
		save, err := os.OpenFile(
			fmt.Sprintf(
				"%s/%s_%s.log",
				path,
				name,
				time.Now().UTC().Format("2006-01-02"),
			),
			os.O_RDWR|os.O_CREATE|os.O_APPEND,
			0666)
		if err != nil {
			panic(err)
		}
		close = save.Close
		print = save
	}
	if isPrint {
		if print != nil {
			print = io.MultiWriter(print, os.Stderr)
		} else {
			print = os.Stderr
		}
	}
	return print, close
}
