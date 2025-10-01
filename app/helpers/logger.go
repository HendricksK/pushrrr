package helpers

import (
	"log"
	"os"
	"strconv"
)

var (
	outfile, err = os.Create("storage/logs/error.log") // update path for your needs
	logger       = log.New(outfile, "", 0)
)

func Log(error string, filename string, line int) {
	line_number := strconv.Itoa(line)
	logger.Println(filename + " - line: " + line_number + ": " + error)
}
