package main

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/nsf/termbox-go"
)

func main() {
	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Set output mode for color support
	termbox.SetOutputMode(termbox.Output256)

	// Clear the screen
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Print a string in the middle of the screen
	width, height := termbox.Size()
	msg := "YOu hIght on MeTHamPhEtamInEs"
	printString((width-len(msg))/2, height/2, msg, termbox.ColorRed, termbox.ColorDefault)

	// Flush to display
	termbox.Flush()

	// Start a goroutine to update mouse position every second
	go func() {
		for {
			robotgo.MoveRelative(1, 0) // Move mouse +1 on the X-axis every second
			time.Sleep(1 * time.Second)
		}
	}()

	// Wait for an event, just to keep the display
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			break loop
		}
	}
}

func printString(x, y int, msg string, fg, bg termbox.Attribute) {
	for _, ch := range msg {
		termbox.SetCell(x, y, ch, fg, bg)
		x++
	}
}
