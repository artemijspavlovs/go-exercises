package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type LogEntry struct {
	timestamp string
	logLevel  string
	module    string
	message   string
}

func main() {
	oldF, err := os.Open("example.log")
	if err != nil {
		log.Fatal(err)
	}
	defer oldF.Close()

	newF, err := os.Create("new.log")
	if err != nil {
		log.Fatal(err)
	}
	defer newF.Close()

	scanner := bufio.NewScanner(oldF)
	for scanner.Scan() {
		err := LogParser(scanner.Text(), newF)
		if err != nil {
			log.Printf("WARNING: %s\n", err)
		}
	}
}

// LogParser converts a log entry string into a LogEntry following the new format, [log level] [module] [timestamp]: [message].
func LogParser(s string, writer io.Writer) error {
	// regex to match the old log entry format, [timestamp] [log level] [module]: [message]
	logPattern := `^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z) (\w+) (\w+): (.*)$`
	logRegex := regexp.MustCompile(logPattern)
	matches := logRegex.FindStringSubmatch(s)

	if len(matches) == 0 {
		return fmt.Errorf("log entry '%s' is not in expected format", s)
	} else {
		fields := strings.Fields(s)
		msg := s[strings.Index(s, ": "):]
		msg = strings.Replace(msg, ": ", "", -1)

		newLog := LogEntry{
			timestamp: fields[0],
			logLevel:  fields[1],
			module:    strings.Replace(fields[2], ":", "", -1),
			message:   msg,
		}

		fmt.Fprintf(writer, "%s %s %s: %s\n", newLog.logLevel, newLog.module, newLog.timestamp, newLog.message)
	}
	return nil
}
