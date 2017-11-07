package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

// GetGeneratie genereert wat lijntjes
func GetGeneratie() Generaties {
	// teken de lijn van de eerste generatie
	lijn := new(Lijn)
	lijn.x = 100
	lijn.x2 = 700

	var lijntjes []Lijn
	lijntjes = []Lijn{
		*lijn}

	// stop de lijn in GenerationLines
	generatie := &GenerationLines{
		generatie: 0,
		lijnen:    lijntjes,
		//	lijnen : lijntjes,
	}

	generaties := &Generaties{
		lijnset: []GenerationLines{
			*generatie},
	}
	generaties.lijnset = append(generaties.lijnset, *generatie)
	return *generaties
}

// Lijn representeert begin en eindpunt van lijn
type Lijn struct {
	x  int
	x2 int
}

// GenerationLines bevat alle lijnen van 1 generatie
type GenerationLines struct {
	generatie int
	lijnen    []Lijn
}

// Generaties bevat alle lijnen van alle generaties
type Generaties struct {
	lijnset []GenerationLines
}

func run() {
	// all of our code will be fired up from here
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	generatieLijnen := GetGeneratie()
	win.Clear(colornames.Skyblue)

	for _, generatie := range generatieLijnen.lijnset {
		imd := imdraw.New(nil)
		imd.Color = colornames.Blueviolet
		imd.EndShape = imdraw.RoundEndShape
		imd.Push(pixel.V(float64(generatie.lijnen[generatie.generatie].x), 100), pixel.V(float64(generatie.lijnen[generatie.generatie].x2), 100))
		//imd.EndShape = imdraw.SharpEndShape
		//imd.Push(pixel.V(100, 500), pixel.V(700, 500))
		imd.Line(3)
		imd.Draw(win)

	}

	for !win.Closed() {
		win.Update()
	}
}
