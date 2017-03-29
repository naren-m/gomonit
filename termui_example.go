package main

import (
	"math"

	"log"

	ui "github.com/gizak/termui"
	"github.com/naren-m/gomonit/utils"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	sinps := (func() []float64 {
		n := 400
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	})()

	rows, cols, err := utils.GetScreenSize()
	if err != nil {
		log.Fatal(err)
	}
	height := rows / 2
	width := cols / 2

	g := ui.NewGauge()
	g.Percent = 50
	g.Width = width
	g.Height = height
	g.BorderLabel = "Gauge"

	par := ui.NewPar("<> This row has 3 columns\n<- Widgets can be stacked up like left side\n<- Stacked widgets are treated as a single widget")
	par.Height = height
	par.Width = width
	par.BorderLabel = "Demonstration"

	ls := ui.NewList()
	ls.Border = false
	ls.Items = []string{
		"[1] Downloading File 1",
		"", // == \newline
		"[2] Downloading File 2",
		"",
		"[3] Uploading File 3",
	}

	ls.Height = height
	ls.Width = width

	lc := ui.NewLineChart()
	lc.BorderLabel = "braille-mode Line Chart"
	lc.Data = sinps
	lc.Height = height
	lc.Width = width
	lc.AxesColor = ui.ColorWhite
	lc.LineColor = ui.ColorYellow | ui.AttrBold

	// build layout
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(6, 0, g),
			ui.NewCol(6, 0, par)),
		ui.NewRow(
			ui.NewCol(6, 0, ls),
			ui.NewCol(6, 0, lc)))
	ui.Body.Align()

	ui.Render(ui.Body)
	// handle key q pressing
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-x", func(ui.Event) {
		// handle Ctrl + x combination
		ui.StopLoop()
	})

	// handle a 1s timer
	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		// t is a EvtTimer
		if t.Count%2 == 0 {
			// do something
			g.Percent += 1
			ui.Render(g)
		}
	})

	ui.Loop() // block until StopLoop is called
}
