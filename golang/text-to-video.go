### internal/util.go
package internal

import "os"

func CreateFolderIfNotExists(folder string) error {
	if dir, err := os.Open(folder); os.IsNotExist(err) {
		return os.MkdirAll(folder, 0700)
	} else {
		_ = dir.Close()
	}
	return nil
}

### main.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

func main() {
	text := `让座,汽车在山路上颠簸，刚上车的老大爷有些站不稳。车子满员，已经无座，
              有几个多余的乘客站在廊道里忍受着游来荡去地颠簸。穿黄色T恤衫的小伙子骂一句：
              “他妈的，刚修的水泥路没几年坏成这样,都他妈只知道偷工减料赚钱了”老大爷手紧紧
              抓着横杠身子游来荡去。天气有些闷热。老人家浑身散发出一股熏人的汗酸臭和泥涩味儿。`
	var data []string
	d := ""
	for _, str :=range strings.Split(text, "\n") {
		for _, t := range str {
			if !unicode.IsPunct(t) {
				d = fmt.Sprintf("%s%s", d, string(t))
			} else if unicode.IsSpace(t) {
			} else {
				if d != "" {
					data = append(data, d)
					d = ""
				}
			}
		}
	}

	// create a input file tell ffmpeg how to contact image
	fd_image, err := os.OpenFile("contact_image.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil { fmt.Println("文件打开失败", err) }
	defer fd_image.Close()

	// create a input file tell ffmpeg how to contact audio
	fd_audio, err := os.OpenFile("contact_audio.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil { fmt.Println("文件打开失败", err) }
	defer fd_audio.Close()

	for i, d := range data {
		fnImage := fmt.Sprintf("image0%d", i)
		fnAudio := fmt.Sprintf("audio0%d", i)
		s := Speech{
			Folder:   "audio",
			FileName: fnAudio,
			Language: "zh",
		}
		_ = s.GenerateAudio(d)
		img := Imager{}
		_ = img.genBaseImage(fnImage, d)

		fd_image.Write([]byte(fmt.Sprintf("file 'image/%s.jpeg'\n", fnImage)))
		duration := 0
		if len(d)/ 3 > 1 {
			duration = len(d)/ 8
		}
		fd_image.Write([]byte(fmt.Sprintf("duration %d\n", duration)))
		fd_audio.Write([]byte(fmt.Sprintf("file 'audio/%s.mp3'\n", fnAudio)))
	}

	// generate video with no
	c1 := exec.Command("ffmpeg", "-f", "concat", "-i", "contact_image.txt", "-vsync", "vfr", "-pix_fmt", "yuv420p", "video.mp4")
	err = c1.Run()
	fmt.Println(err)

	// generate audio
	c2 := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", "contact_audio.txt", "-c", "copy", "output.mp3")
	err = c2.Run()
	fmt.Println(err)

	// merge video and audio
	c3 := exec.Command("ffmpeg", "-i", "video.mp4", "-i", "output.mp3", "-c:v", "copy", "-c:a", "aac", "output.mp4")
	err = c3.Run()
	fmt.Println(err)
}

### tti.go
package main

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/leoython/text-to-video/internal"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"
)

const (
	defaultSize float64 = 45
	defaultSpacing float64 = 1.5
)

type Imager struct {}

func(i *Imager) genBaseImage(filename, text string) error {
	m := image.NewRGBA(image.Rect(0, 0, 1600, 900))
	draw.Draw(m, m.Bounds(), image.Black, image.ZP, draw.Src)
	
	if err := internal.CreateFolderIfNotExists("image"); err != nil {
		return err
	}

	// Read the font data.
	fontBytes, err := ioutil.ReadFile("./front/SourceHanSans-Bold.ttf")
	if err != nil { return err }

	f, err := freetype.ParseFont(fontBytes)
	if err != nil { return err }

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(defaultSize)
	c.SetClip(m.Bounds())
	c.SetDst(m)
	c.SetSrc(image.White)

	pt := freetype.Pt(350, 200+int(c.PointToFixed(defaultSize)>>6))
	if _, err := c.DrawString(text, pt); err != nil { return err }

	fd, err := os.Create(fmt.Sprintf("image/%s.jpeg", filename))
	if err != nil { return err }

	err = jpeg.Encode(fd, m, nil)
	if err != nil { return err }

	return nil
}

### tts.go
package main

import (
	"fmt"
	"github.com/leoython/text-to-video/internal"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Speech struct {
	Folder   string
	FileName string
	Language string
}

func (speech *Speech) GenerateAudio(text string) error {
	var err error

	fileName := speech.Folder + "/" + speech.FileName + ".mp3"

	if err = internal.CreateFolderIfNotExists(speech.Folder); err != nil { return err }

	if err = speech.downloadIfNotExists(fileName, text); err != nil { return err }

	return nil
}

func (speech *Speech) downloadIfNotExists(fileName string, text string) error {
	f, err := os.Open(fileName)
	if err != nil {
		response, err := http.Get(fmt.Sprintf(`http://translate.google.com/translate_tts?
                    ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s`,
                    url.QueryEscape(text), speech.Language))
		if err != nil { return err }
		defer response.Body.Close()

		output, err := os.Create(fileName)
		if err != nil { return err }

		_, err = io.Copy(output, response.Body)
		return err
	}

	_ = f.Close()
	return nil
}
