package main

import (
	"github.com/rivo/tview"
)

type V2 struct {
	X, Y int
}

type Buffer struct {
	name     string
	filename string
}

type BufferView struct {
	buffer *Buffer
	scroll V2
}

type Orientation bool

const (
	Horizontal Orientation = false
	Vertical   Orientation = true
)

type Frame struct {
	bufferViews []*BufferView
	// this is the orientation of the frame and how the frame next to it
	// will be positioned. horizontal means the frame will be placed next to
	// the current frame, vertical means the frame will be placed below the
	// current frame.
	orientation Orientation
	display     *tview.TextView
	frame       *Frame
}

func newFrame() *Frame {
	return &Frame{
		bufferViews: make([]*BufferView, 0),
		orientation: Horizontal,
		display:     tview.NewTextView().SetText("Hello, world!"),
	}
}

func (f *Frame) RenderFrame() *tview.Grid {
	grid := tview.NewGrid()
	grid.AddItem(f.display, 0, 0, 1, 1, 0, 1, true)
	if f.frame != nil {
		if f.orientation == Horizontal {
			grid.SetColumns(0, 0)
			grid.AddItem(f.frame.RenderFrame(), 0, 1, 1, 1, 0, 0, true)
		} else {
			grid.SetRows(0, 0)
			grid.AddItem(f.frame.RenderFrame(), 1, 0, 1, 1, 0, 0, true)
		}
	} else {
		grid.SetColumns(0)
		grid.SetRows(0)
	}
	return grid
}

func main() {
	frame := newFrame()
	subframe := newFrame()
	frame.frame = subframe
	subframe.orientation = Vertical
	subframe.frame = newFrame()
	app := tview.NewApplication().SetRoot(frame.RenderFrame(), true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
