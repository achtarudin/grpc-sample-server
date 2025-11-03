package main

import (
	"bytes"
	"time"

	"github.com/fatih/color"
)

type logWriter struct{}

func (writer *logWriter) Write(result []byte) (n int, err error) {
	var c *color.Color
	if bytes.Contains(bytes.ToLower(result), []byte("error")) || bytes.Contains(bytes.ToLower(result), []byte("failed")) {
		c = color.New(color.FgRed)
	} else {
		c = color.New(color.FgGreen)
	}
	return c.Print(time.Now().UTC().Format("02/01/2006 15:04:05") + " " + string(result))
}
