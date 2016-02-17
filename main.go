package main

import (
	"time"
	"bytes"
    "io"
    "log"
    "os"
    "os/exec"
)

var (
	moveCursorSeq = []byte{27, '[', ';', 'H'}
	clearScreenSeq = []byte{27, '[', '2', 'J'}
)

func main() {
	cmd := exec.Command("write", "matehegedus")
    stdin, err := cmd.StdinPipe()
    if err != nil {
        log.Panic(err)
    }
    err = cmd.Start()
    if err != nil {
        log.Panic(err)
    }
	for i, frame := range frames {
		// Clear the screen
		/*defer stdin.Close()
		stdin.Write(clearScreenSeq)
		defer stdin.Close()
		stdin.Write(moveCursorSeq)*/
		// Show the movie
		defer stdin.Close()
		io.Copy(stdin, bytes.NewBufferString("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n"))
   		defer stdin.Close()
		io.Copy(stdin, bytes.NewBufferString(frame))
		// os.Stdout.Write([]byte(frame))
		time.Sleep(time.Duration(int64(repeats[i]) * 67e6))
	}
}

func clearScreen() {
	os.Stdout.Write(clearScreenSeq)
	os.Stdout.Write(moveCursorSeq)
}
