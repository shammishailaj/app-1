package ui

import "fmt"

// Carousel is slideshow component for cycling through elements, images or
// slides of text like a carousel.
type Carousel struct {
	Data  []string
	Compo string
}

func (c *Carousel) OnMount() {
	fmt.Println("data:", c.Data)
}

// Render returns the html representation of the component.
func (c *Carousel) Render() string {
	return `
<div>
	{{range .Data}}
		<div style="
			width: 100%;
			height: 100%;
		">
			{{.}}
		</div>
	{{end}}
</div>
	`
}
