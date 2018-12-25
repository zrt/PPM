package PPM

type Renderer interface {
	Render() error
}

type RenderResult struct {
	Pos   PixelPos
	Color V
}
