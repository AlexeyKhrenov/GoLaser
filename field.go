package main

import (
	"fmt"
	"image/color"
	"sim/models"
	"sim/settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var bg = color.NRGBA{R: 36, G: 36, B: 36, A: 255}
var ta = color.NRGBA{R: 16, G: 82, B: 36, A: 255}
var green = color.NRGBA{R: 36, G: 217, B: 69, A: 255}
var red = color.NRGBA{R: 255, G: 0, B: 0, A: 255}

var secText string

func showField(s settings.Settings, feed <-chan bool) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	game := models.NewGame()
	field := game.Field

	background := canvas.NewRectangle(bg)
	background.Resize(fyne.NewSize(field.Width, field.Height))

	c := myWindow.Canvas()

	go func() {

		for {
			fmt.Println("New round")
			game.NewRound()

			target := formatTarget(&field, &game.Round.Target)
			c.SetContent(container.NewWithoutLayout(background, target))

			for {
				shot := <-feed
				game.Shot(shot)

				if game.Round.IsOver {
					attemptsLabel := formatAttempts(&field, &game)
					elapsedLabel := formatElapsedTime(&field, &game.Round)
					c.SetContent(container.NewWithoutLayout(background, attemptsLabel, elapsedLabel))

					s.OverallScore = game.OverallScore
					settings.SaveSettings(&s)
					break
				}
			}
		}
	}()

	myWindow.Resize(fyne.NewSize(field.Width, field.Height))
	myWindow.ShowAndRun()
}

func formatTarget(f *models.Field, t *models.Target) *canvas.Rectangle {
	target := canvas.NewRectangle(ta)
	target.Resize(fyne.NewSize(t.Width, t.Height))
	target.Move(fyne.NewPos(t.X, t.Y))
	return target
}

func formatAttempts(f *models.Field, game *models.Game) *canvas.Text {
	txt := game.GameScore.ToString()
	label := canvas.NewText(txt, green)
	label.TextSize = 50

	label.Move(fyne.NewPos(100, f.Height-100))
	if !game.Round.IsSuccess {
		label.Color = red
	}

	return label
}

func formatElapsedTime(f *models.Field, round *models.Round) *canvas.Text {
	formatted := format(round.Elapsed)

	label := canvas.NewText(formatted, green)
	label.TextSize = 50

	if !round.IsSuccess {
		label.Color = red
	}

	label.Move(fyne.NewPos(f.Width-300, f.Height-100))
	return label
}
