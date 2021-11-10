package frontend

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
	"github.com/yoRyuuuuu/typex/client/backend"
)

const refreshInterval = 16 * time.Millisecond

type View struct {
	app           *tview.Application
	logger        *tview.TextView
	problemView   *tview.TextView
	playerView    *tview.Flex
	inputField    *tview.InputField
	drawCallbacks []func()
	*backend.Game
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

func (v *View) setupProblemView() {
	v.problemView.SetTitle("Problem").
		SetBorder(true).
		SetTitle(v.Problem)
	callback := func() {
		v.problemView.SetText(v.Problem)
	}
	v.drawCallbacks = append(v.drawCallbacks, callback)
}

func (v *View) setupInputField() {
	v.inputField.SetLabel("input: ")
	v.inputField.SetTitle("Terminal").
		SetBorder(true)

	v.inputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			if v.CheckAnswer(v.inputField.GetText()) {
				v.ActionChannel <- backend.Answer{
					Action: nil,
				}
				v.inputField.SetText("")
			}
			return nil
		}
		return event
	})
}

func (v *View) setupLogger() {
	v.logger.SetTitle("Log")
	v.drawCallbacks = append(v.drawCallbacks, v.drawLogger)
}

func (v *View) drawLogger() {
	v.logger.Clear().
		SetBorder(true)
	v.logger.SetText(v.Log)
}

func (v *View) setupPlayerView() {
	v.drawCallbacks = append(v.drawCallbacks, v.drawPlayerView)
}

func (v *View) drawPlayerView() {
	// 描画をリセット
	v.playerView.Clear()
	// 自分のスコアを描画
	mine := tview.NewTextView()
	mine.SetTitle("YOU").
		SetBorder(true)
	mine.SetText(fmt.Sprintf("score: %v", v.Health[v.ID]))
	v.playerView.AddItem(mine, 3, 0, false)
	for _, player := range v.Players {
		// 他プレイヤーのスコアを描画
		text := tview.NewTextView()
		text.SetTitle(player.ID).
			SetBorder(true)
		text.SetText(fmt.Sprintf("score: %v", v.Health[player.ID]))
		v.playerView.AddItem(text, 3, 0, false)
	}
}

func NewView(game *backend.Game) *View {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}

	app := tview.NewApplication()
	root := tview.NewFlex()
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	// 単語のViewを配置
	problemView := tview.NewTextView()
	flex.AddItem(problemView, 3, 0, false)
	// inpuFieldを配置
	inputField := tview.NewInputField()
	flex.AddItem(inputField, 3, 0, true)
	// loggerを配置
	logger := tview.NewTextView()
	flex.AddItem(logger, 0, 1, false)

	root.AddItem(flex, 0, 2, true)

	playerView := tview.NewFlex().SetDirection(tview.FlexRow)
	root.AddItem(playerView, 0, 1, false)
	view := &View{
		app:           app,
		problemView:   problemView,
		logger:        logger,
		inputField:    inputField,
		playerView:    playerView,
		drawCallbacks: []func(){},
		Game:          game,
	}

	// UIのセットアップ
	view.setupProblemView()
	view.setupInputField()
	view.setupLogger()
	view.setupPlayerView()

	view.app.SetRoot(root, true)
	return view
}

func (v *View) Start() {
	go v.refresh()

	if err := v.app.Run(); err != nil {
		panic(err)
	}
}