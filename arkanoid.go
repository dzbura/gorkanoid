package main

import (
    
	"image"
	"os"
    _ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
    "golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
    VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
    if err != nil {
		panic(err)
	}

    win.SetSmooth(true)

 	spritesheet, err := loadPicture("sprites.png")
	if err != nil {
		panic(err)
	}
    

    var blockFrames []pixel.Rect
	for y := float64(75); y < spritesheet.Bounds().Max.Y; y += 32 {
		blockFrames = append(blockFrames, pixel.R(0, y, 85, y+32))
	}
    
    ball := pixel.NewSprite(spritesheet, pixel.R(0,0,50,50))
    plank := pixel.NewSprite(spritesheet, pixel.R(0,50,215,75))
    block0 := pixel.NewSprite(spritesheet, blockFrames[0])
    block1 := pixel.NewSprite(spritesheet, blockFrames[1])
    block2 := pixel.NewSprite(spritesheet, blockFrames[2])
    block3 := pixel.NewSprite(spritesheet, blockFrames[3])
    bally:=60.0
    ballx:=512.0
    plankx:=512.0

	for !win.Closed() {
        win.Clear(colornames.Turquoise)
        for y := float64(50); y < 1000; y +=100 {
            block0.Draw(win, pixel.IM.Moved(pixel.V(y, 675)))
            block0.Draw(win, pixel.IM.Moved(pixel.V(y+50, 650)))
            block1.Draw(win, pixel.IM.Moved(pixel.V(y, 625)))
            block1.Draw(win, pixel.IM.Moved(pixel.V(y+50, 600)))
            block2.Draw(win, pixel.IM.Moved(pixel.V(y, 575)))
            block2.Draw(win, pixel.IM.Moved(pixel.V(y+50, 550)))
            block3.Draw(win, pixel.IM.Moved(pixel.V(y, 525)))
            block3.Draw(win, pixel.IM.Moved(pixel.V(y+50, 500)))
        }

        if win.Pressed(pixelgl.KeyRight){
            plankx+=2.5
        }
        if win.Pressed(pixelgl.KeyLeft){
            plankx-=2.5
        }
        plank.Draw(win, pixel.IM.Moved(pixel.V(plankx,20)))

        ball.Draw(win, pixel.IM.Moved(pixel.V(ballx,bally)))
        bally+=1

        win.Update()

	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func main() {
    pixelgl.Run(run)
}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func main() {
    pixelgl.Run(run)
}

