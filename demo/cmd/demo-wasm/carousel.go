package main

import (
	"syscall/js"

	"github.com/maxence-charriere/app/pkg/app"
)

type carousel struct {
	Images []string
}

func (c *carousel) OnMount() {
	c.Images = []string{
		"beijing.jpg",
		"paris.jpg",
		"sf.jpg",
		"space.jpg",
	}
	app.Render(c)
}

func (c *carousel) Render() string {
	return `
<div class="carousel">
	<button class="Menu" onclick="OnMenuClick" oncontextmenu="OnMenuClick">☰</button>

	<main>
		<ui.carousel data="{{json .Images}}">
	</main>
</div>
	`
}

func (c *carousel) OnMenuClick(s, e js.Value) {
	app.NewContextMenu(
		app.MenuItem{
			Label: "Reload",
			Keys:  "cmdorctrl+r",
			OnClick: func(s, e js.Value) {
				app.Reload()
			},
		},
		app.MenuItem{Separator: true},
		app.MenuItem{
			Label: "Go to repository",
			OnClick: func(s, e js.Value) {
				app.Navigate("https://github.com/maxence-charriere/app")
			}},
		app.MenuItem{
			Label: "Source code",
			OnClick: func(s, e js.Value) {
				app.Navigate("https://github.com/maxence-charriere/app/blob/master/demo/cmd/demo-wasm/hello.go")
			}},
		app.MenuItem{Separator: true},
		app.MenuItem{
			Label: "Hello example",
			OnClick: func(s, e js.Value) {
				app.Navigate("hello")
			}},
		app.MenuItem{
			Label: "City example",
			OnClick: func(s, e js.Value) {
				app.Navigate("city")
			}},
	)
}

type image struct {
	Path string
}

func (i *image) Render() string {
	return `<img src="{{.Path}}">`
}
