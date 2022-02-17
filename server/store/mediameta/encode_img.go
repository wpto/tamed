package mediameta

// func (t *Trnasform) EncodeImg(srcPath, destPath string, targetType string, targetWidth int) error {

// 	buf, err = ioutil.ReadFile(srcPath)
// 	if err != nil {
// 		return errors.Wrap(err, "encode.imgread")
// 	}

// 	img := bimg.NewImage(buf)
// 	size, err := img.Size()
// 	if err != nil {
// 		return errors.Wrap(err, "encode.imgsize")
// 	}

// 	if targetType != img.Type() {
// 		buf, err = formatPic(img, newFormat)
// 		if err != nil {
// 			return errors.Wrap(err, "encode.imgconv")
// 		}
// 		img = bimg.NewImage(buf)
// 	}

// 	if targetWidth != 0 && targetWidth < size.Width {
// 		buf, err = resizePic(img, opts.Width)
// 		if err != nil {
// 			return nil, errors.Wrap(err, "encode.imgresize")
// 		}
// 	}

// 	err = ioutil.WriteFile(destPath, buf, 0644)
// 	if err != nil {
// 		return errors.Wrap(err, "encode.imgwrite")
// 	}

// 	return nil
// }

// var TypeToBimg = map[string]bimg.ImageType{
// 	"jpeg": bimg.JPEG,
// 	"png":  bimg.PNG,
// 	"webp": bimg.WEBP,
// }

// func formatPic(img *bimg.Image, imgType string) ([]byte, error) {
// 	newFormat, ok := TypeToBimg[imgType]
// 	if !ok {
// 		return errors.Errorf("%s is not supported", targetType)
// 	}
// 	buf, err := img.Convert(newFormat)
// 	if err != nil {
// 		return nil, errors.Wrap(err, fmt.Sprintf("formatpic(%s to %s)", img.Type(), imgType))
// 	}

// 	return buf, nil
// }

// func resizePic(img *Image, width int) ([]byte, error) {
// 	size, err := img.Size()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "encode.resize")
// 	}

// 	coef := float64(width) / float64(size.Width)

// 	newWidth := int(math.Floor(float64(size.Width) * coef))
// 	newHeight := int(math.Floor(float64(size.Height) * coef))

// 	buf, err = img.ResizeAndCrop(newWidth, newHeight)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "resizepic")
// 	}

// 	return buf, nil
// }
