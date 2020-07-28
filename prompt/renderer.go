package prompt

type Renderer struct {
	terminal *Terminal
}

func NewRenderer(terminal *Terminal) *Renderer {
	return &Renderer{
		terminal: terminal,
	}
}

func (r *Renderer) Render(scr *Screen) {

}
