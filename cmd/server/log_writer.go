package main

import (
	"bytes"
	"time"

	"github.com/fatih/color"
)

type logWriter struct{}

func (writer *logWriter) Write(result []byte) (n int, err error) {

	color.NoColor = false

	var c *color.Color

	byteLower := bytes.ToLower(result)

	if bytes.Contains(byteLower, []byte("error")) || bytes.Contains(byteLower, []byte("failed")) {
		c = color.New(color.FgHiRed)
	} else if bytes.Contains(byteLower, []byte("info")) {
		c = color.New(color.FgHiBlue)
	} else {
		c = color.New(color.FgHiGreen)
	}
	return c.Print(time.Now().UTC().Format("02/01/2006 15:04:05") + " " + string(result))
}
