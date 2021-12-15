const (
	defaultSize float64 = 45
	defaultSpacing float64 = 1.5
)

func GenerateImage(filename, text string) error {
	m := image.NewRGBA(image.Rect(0, 0, 1600, 900))
	draw.Draw(m, m.Bounds(), image.Black, image.ZP, draw.Src)
	
	// Read the font data.
	fontBytes, err := ioutil.ReadFile("./font/SourceHanSans-Bold.ttf")
	f, err := freetype.ParseFont(fontBytes)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(defaultSize)
	c.SetClip(m.Bounds())
	c.SetDst(m)
	c.SetSrc(image.White)

	pt := freetype.Pt(350, 200+int(c.PointToFixed(defaultSize)>>6))
	_, err := c.DrawString(text, pt)

	fd, err := os.Create(fmt.Sprintf("image/%s.jpeg", filename))

	err = jpeg.Encode(fd, m, nil)

	return nil
}
