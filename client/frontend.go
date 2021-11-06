package client

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
)

const refreshInterval = 16 * time.Millisecond

type View struct {
	app           *tview.Application
	flex          *tview.Flex
	drawCallbacks []func()
	*Game
}

func (v *View) refresh() {
	tick := time.NewTicker(refreshInterval)
	for {
		for _, callback := range v.drawCallbacks {
			v.app.QueueUpdate(callback)
		}
		v.app.Draw()
		<-tick.C
	}
}

func setupWord(view *View) {
	word := tview.NewTextView()
	word.SetTitle("Word").
		SetBorder(true)

	callback := func() {
		w := view.word
		word.SetText(w)
	}

	view.drawCallbacks = append(view.drawCallbacks, callback)

	view.flex.AddItem(word, 3, 0, false)
}

func setupInput(view *View) {
	input := tview.NewInputField()
	input.SetLabel("input: ")

	input.SetTitle("Terminal").
		SetBorder(true)

	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			if view.checkAnswer(input.GetText()) {
				view.actionChannel <- AnswerAction{
					text: input.GetText(),
				}

				input.SetText("")
			}

			return nil
		}
		return event
	})

	view.flex.AddItem(input, 3, 0, true)
}

func setupStatus(view *View) {
	status := tview.NewTextView()
	status.SetTitle("Log").
		SetBorder(true)

	callback := func() {
		status.SetText(view.log)
	}

	view.drawCallbacks = append(view.drawCallbacks, callback)
	view.flex.AddItem(status, 0, 1, false)
}

func NewView(game *Game) *View {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}
	app := tview.NewApplication()

	flex := tview.NewFlex()
	view := &View{
		app:           app,
		flex:          flex,
		drawCallbacks: []func(){},
		Game:          game,
	}

	flex.SetDirection(tview.FlexRow)
	setupWord(view)
	setupInput(view)
	setupStatus(view)

	view.flex = flex
	return view
}

func (v *View) Start() {
	go v.refresh()

	if err := v.app.SetRoot(v.flex, true).Run(); err != nil {
		panic(err)
	}
}
