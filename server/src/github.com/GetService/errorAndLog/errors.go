package errorAndLog

import "log"

func ErrorHandler(err error, str string) {
	if err != nil {
		log.Print("Error at: " + str)
		panic(err)
	}
}
