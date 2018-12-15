package main

import (
    
	"image"
	"os"
    _ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
    "golang.org/x/image/colornames"
    "github.com/faiface/pixel/text"
    "fmt"
    "golang.org/x/image/font/basicfont"
    "strconv"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Gorkanoid",
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
    
    basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
    titleTxt := text.New(pixel.V(100, 520), basicAtlas)
    scoreTxt := text.New(pixel.V(250, 535), basicAtlas)
    fmt.Fprintln(titleTxt, "GORKANOID")



    var blockFrames []pixel.Rect
	for y := float64(75); y < spritesheet.Bounds().Max.Y; y += 32 {
		blockFrames = append(blockFrames, pixel.R(0, y, 85, y+32))
	}

    ball := pixel.NewSprite(spritesheet, pixel.R(0,0,50,50))
    plank := pixel.NewSprite(spritesheet, pixel.R(0,50,215,75))
    //block0 := pixel.NewSprite(spritesheet, blockFrames[0])
    //block1 := pixel.NewSprite(spritesheet, blockFrames[1])
    block2 := pixel.NewSprite(spritesheet, blockFrames[2])
    //block3 := pixel.NewSprite(spritesheet, blockFrames[3])
    ballV := pixel.V(512.0, 60.0)
    plankx:=512.0
    motion := pixel.V(0,2)
    var blocksV []pixel.Vec   
    for x := float64(50); x < 1000; x +=100 {
            blocksV = append(blocksV, pixel.V(x, 675))
            blocksV = append(blocksV, pixel.V(x+50, 650))
            blocksV = append(blocksV, pixel.V(x, 625))
            blocksV = append(blocksV, pixel.V(x+50, 600))
            blocksV = append(blocksV, pixel.V(x, 575))
            blocksV = append(blocksV, pixel.V(x+50, 550))
            blocksV = append(blocksV, pixel.V(x, 525))
            blocksV = append(blocksV, pixel.V(x+50, 500))
    }
    score := 0
    highScore := 0


	for !win.Closed() {
        win.Clear(colornames.Turquoise)
        scoreTxt.Clear()
        fmt.Fprintln(scoreTxt, "THE SCORE.")
        fmt.Fprintln(scoreTxt, "Score:" + strconv.Itoa(score))
        fmt.Fprintln(scoreTxt, "High score" + strconv.Itoa(highScore))
        for x:=0; x<len(blocksV); x+=1 {
            block2.Draw(win, pixel.IM.Moved(blocksV[x]))       
        }


        plank.Draw(win, pixel.IM.Moved(pixel.V(plankx, 20)))
        ball.Draw(win, pixel.IM.Moved(ballV))
        titleTxt.Draw(win, pixel.IM.Moved(pixel.V(20,200)))
        scoreTxt.Draw(win, pixel.IM.Moved(pixel.V(20,200)))
        if win.Pressed(pixelgl.KeyRight){
            plankx+=2.5
        }
        if win.Pressed(pixelgl.KeyLeft){
            plankx-=2.5
        }

        ballx, bally := ballV.XY()

        if bally == 768 {
            if motion == pixel.V(2,2){
                motion = pixel.V(2,-2) 
            }  
            if motion == pixel.V(-2,2){
                motion = pixel.V(-2,-2) 
            }
            if motion == pixel.V(0,2){
                motion = pixel.V(2,-2)
            }                  
        } else if ballx == 0 {
            if motion == pixel.V(-2,2){
                motion = pixel.V(2,2) 
            }  
            if motion == pixel.V(-2,-2){
                motion = pixel.V(2,-2) 
            }  
        } else if ballx == 1024 {
            if motion == pixel.V(2,2){
                motion = pixel.V(-2,2) 
            }  
            if motion == pixel.V(2,-2){
                motion = pixel.V(-2,-2) 
            }  
        } else if bally == 36 && ballx < plankx+100 && ballx > plankx-100 {
            if motion == pixel.V(2,-2){
                motion = pixel.V(2, 2)
            }
            if motion == pixel.V(-2,-2){
                motion = pixel.V(-2,2)
            }
        } else if bally == 0 {
            ballV = pixel.V(512.0, 60.0)
            motion = pixel.V(0,2)
            if highScore < score{ 
                highScore = score 
            }


            score = 0
        } else {
   
            for y:=0; y<len(blocksV); y+=1{
                blockx, blocky := blocksV[y].XY()
                if ballx-15 == blockx+35 && blocky+7>bally && bally>blocky-7{
                    if motion == pixel.V(-2,-2){
                        motion = pixel.V(2,-2)
                        score += 10
                        blocksV = removeV(blocksV, y)
                        break
                    }
                    if motion == pixel.V(-2,2){
                        motion = pixel.V(2,2)
                        score += 10
                        blocksV = removeV(blocksV, y)
                        break
                    }
                }
                if ballx+15 == blockx-35 && blocky+7>bally && bally>blocky-7{
                    if motion == pixel.V(2,2){
                        motion = pixel.V(-2,2)
                        score += 10
                        blocksV = removeV(blocksV, y)
                        break
                    }
                    if motion == pixel.V(2,-2){
                        motion = pixel.V(-2,-2)
                        score += 10
                        blocksV = removeV(blocksV, y)
                        break
                    }
                }
		if (bally+15 == blocky-7 || bally+15 == blocky-6) && ballx >= blockx-35 && blockx+35 >= ballx{
                    if motion == pixel.V(0, 2){
                        motion = pixel.V(2, -2)
                        score += 10
                        blocksV = removeV(blocksV, y)
                        break
                    }
                    if motion == pixel.V(-2,2){
                        motion = pixel.V(-2,-2)
                        score += 10
                        blocksV = removeV(blocksV, y)
                        break
                    }
                    if motion == pixel.V(2,2){
                        motion = pixel.V(2,-2)
                        score += 10
                        blocksV = removeV(blocksV, y)
                        break
                    }
                }
                if (bally-15 == blocky+7 || bally-15 == blocky+6) && ballx >= blockx-35 && blockx+35 >= ballx{
                    if motion == pixel.V(-2,-2){
                        motion = pixel.V(-2,2)
                        score += 10
                        blocksV = removeV(blocksV, y)
                        break
                    }
                    if motion == pixel.V(2,-2){
                        motion = pixel.V(2,2)
                        score += 10
                        blocksV = removeV(blocksV, y)
                        break
                    }
                }
            }
        }
        ballV = ballV.Add(motion)
        win.Update()
	}
}

func removeV(list []pixel.Vec, indexV int) []pixel.Vec {
    list[indexV] = list[len(list)-1]
    return list[:len(list)-1]
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
