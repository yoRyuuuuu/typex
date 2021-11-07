package client

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
)

const refreshInterval = 16 * time.Millisecond

type View struct {
	app           *tview.Application
	flex          *tview.Flex
	rightFlex     *tview.Flex
	leftFlex      *tview.Flex
	playerText    map[string]*tview.TextView
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

	view.leftFlex.AddItem(word, 3, 0, false)
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

	view.leftFlex.AddItem(input, 3, 0, true)
}

func setupStatus(view *View) {
	status := tview.NewTextView()
	status.SetTitle("Log").
		SetBorder(true)

	callback := func() {
		status.SetText(view.log)
	}

	view.drawCallbacks = append(view.drawCallbacks, callback)
	view.leftFlex.AddItem(status, 0, 1, false)
}

func setupPlayer(view *View) {
	text := tview.NewTextView()
	text.SetTitle(fmt.Sprintf("YOU")).
		SetBorder(true)
	text.SetText(fmt.Sprintf("score: 0"))
	view.rightFlex.AddItem(text, 3, 0, false)
	view.playerText[view.id] = text

	callback := func() {
		select {
		case player := <-view.playerChannel:
			text := tview.NewTextView()
			text.SetTitle(player.name).
				SetBorder(true)
			text.SetText(fmt.Sprintf("score: 0"))
			view.rightFlex.AddItem(text, 3, 0, false)
			view.playerText[player.id] = text
		case damage := <-view.damageChannel:
			view.playerText[damage.id].SetText(fmt.Sprintf("score: %v", damage.damage))
		default:
		}
	}

	view.drawCallbacks = append(view.drawCallbacks, callback)
}

func NewView(game *Game) *View {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}
	app := tview.NewApplication()

	flex := tview.NewFlex()
	flex.SetBackgroundColor(tcell.ColorBlack)
	leftFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(leftFlex, 0, 2, true)
	rightFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(rightFlex, 0, 1, false)

	view := &View{
		app:           app,
		leftFlex:      leftFlex,
		rightFlex:     rightFlex,
		playerText:    map[string]*tview.TextView{},
		drawCallbacks: []func(){},
		Game:          game,
	}

	setupWord(view)
	setupInput(view)
	setupStatus(view)
	setupPlayer(view)

	view.flex = flex
	return view
}

func (v *View) Start() {
	go v.refresh()

	if err := v.app.SetRoot(v.flex, true).Run(); err != nil {
		panic(err)
	}
}
