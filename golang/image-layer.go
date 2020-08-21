package layers

func GenerateImageFromLayers(
	width, height int,
	keys []string,
	layerPaths []string,
	outputPath string,
	saveAsPng, transparent bool,
) error {
	newImage := image.NewNRGBA(image.Rect(0, 0, width, height))

	for i, path := range layerPaths {
		imageFile, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open image at path '%s'.\nkey: %s\nerr: %s",
				path, keys[i], err.Error())
		}

		defer imageFile.Close()

		img, err := png.Decode(imageFile)
		if err != nil {
			return fmt.Errorf("failed to decode image as png at path '%s': %s", path, err.Error())
		}

		if i == 0 && !transparent {
			draw.Draw(newImage, newImage.Bounds(), img, image.Point{}, draw.Src)
		} else {
			draw.Draw(newImage, newImage.Bounds(), img, image.Point{}, draw.Over)
		}
	}

	resultImg, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to save image at path '%s': %s", outputPath, err.Error())
	}

	defer resultImg.Close()

	if saveAsPng {
		encoder := png.Encoder{
			CompressionLevel: png.NoCompression,
		}
		encoder.Encode(resultImg, newImage)
	} else {
		jpeg.Encode(resultImg, newImage, &jpeg.Options{
			Quality: jpeg.DefaultQuality,
		})
	}

	return nil
}
