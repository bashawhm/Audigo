package main

import (
	"os"

	"github.com/gordonklaus/portaudio"
	"github.com/nsf/termbox-go"
)

func tbprint(x, y int, ch rune, fg, bg termbox.Attribute) {
	for i := 0; i <= x; i++ {
		termbox.SetCell(i, y, ch, fg, bg)
	}
}

func drawSound(in [][]float32) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	NUMFRAMES := len(in[0])
	for i := 0; i < len(in); i++ {
		for j := 0; j < NUMFRAMES; j++ {
			tbprint(i, j, '#', termbox.ColorBlue, termbox.ColorDefault)
			if in[i][j] > 8 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 7 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 6 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 5 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 4 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 3 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 2 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 1 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 0.9 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 0.8 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 0.7 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 0.6 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 0.5 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else if in[i][j] > 0.25 {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			} else {
				tbprint(int(in[i][j]), j, '#', termbox.ColorMagenta, termbox.ColorDefault)
			}
		}
	}
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	portaudio.Initialize()
	defer portaudio.Terminate()

	hostAPI, err := portaudio.DefaultHostApi()
	if err != nil {
		panic(err)
	}
	streamPeram := portaudio.LowLatencyParameters(hostAPI.DefaultInputDevice, hostAPI.DefaultOutputDevice)
	streamPeram.Input.Channels = 1
	streamPeram.Output.Channels = 1
	_, h := termbox.Size()
	streamPeram.FramesPerBuffer = h

	stream, err := portaudio.OpenStream(streamPeram, processAudio)
	if err != nil {
		panic(err)
	}

	stream.Start()
	defer stream.Stop()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			os.Exit(1)
		}
	}

}

//len(in) == numChannels and len(in[i]) == framesPerBuffer
func processAudio(in, out [][]float32) {
	inter := make([][]float32, len(in))
	for i := range inter {
		inter[i] = make([]float32, len(in[i]))
	}
	for i := range inter {
		for j := range inter[i] {
			inter[i][j] = in[i][j] * 500
			out[i][j] = in[i][j]
		}
	}
	drawSound(inter)
}
