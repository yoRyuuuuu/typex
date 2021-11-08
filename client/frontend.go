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
	mainView      *tview.Flex
	playerView    *tview.Flex
	playerStatus  map[string]*tview.TextView
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

func (v *View) setupWordView() {
	word := tview.NewTextView()
	word.SetTitle("Word").
		SetBorder(true)

	callback := func() {
		w := v.word
		word.SetText(w)
	}

	v.drawCallbacks = append(v.drawCallbacks, callback)

	v.mainView.AddItem(word, 3, 0, false)
}

func (v *View) setupInputField() {
	input := tview.NewInputField()
	input.SetLabel("input: ")

	input.SetTitle("Terminal").
		SetBorder(true)

	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			if v.checkAnswer(input.GetText()) {
				v.actionChannel <- AnswerAction{
					text: input.GetText(),
				}

				input.SetText("")
			}

			return nil
		}
		return event
	})

	v.mainView.AddItem(input, 3, 0, true)
}

func (v *View) setupLogger() {
	status := tview.NewTextView()
	status.SetTitle("Log").
		SetBorder(true)

	callback := func() {
		status.SetText(v.log)
	}

	v.drawCallbacks = append(v.drawCallbacks, callback)
	v.mainView.AddItem(status, 0, 1, false)
}

func (v *View) setupPlayerView() {
	text := tview.NewTextView()
	text.SetTitle(fmt.Sprintf("YOU")).
		SetBorder(true)
	text.SetText(fmt.Sprintf("score: 0"))
	v.playerView.AddItem(text, 3, 0, false)
	v.playerStatus[v.id] = text
	v.drawCallbacks = append(v.drawCallbacks, v.drawPlayerView)
}

func (v *View) drawPlayerView() {
	select {
	case player := <-v.playerChannel:
		text := tview.NewTextView()
		text.SetTitle(player.name).
			SetBorder(true)
		text.SetText(fmt.Sprintf("score: 0"))
		v.playerView.AddItem(text, 3, 0, false)
		v.playerStatus[player.id] = text
	case damage := <-v.damageChannel:
		v.playerStatus[damage.id].SetText(fmt.Sprintf("score: %v", damage.damage))
	default:
	}

	return
}

func NewView(game *Game) *View {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}
	app := tview.NewApplication()

	flex := tview.NewFlex()
	mainView := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(mainView, 0, 2, true)
	playerView := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(playerView, 0, 1, false)
	view := &View{
		app:           app,
		mainView:      mainView,
		playerView:    playerView,
		playerStatus:  map[string]*tview.TextView{},
		drawCallbacks: []func(){},
		Game:          game,
	}

	view.setupWordView()
	view.setupInputField()
	view.setupLogger()
	view.setupPlayerView()

	view.app.SetRoot(flex, true)
	return view
}

func (v *View) Start() {
	go v.refresh()

	if err := v.app.Run(); err != nil {
		panic(err)
	}
}
