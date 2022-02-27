package mediameta

// type ResultFFProbe struct {
// 	Streams []struct {
// 		CodecType string `json:"codec_type"`
// 		Width     int    `json:"width"`
// 		Height    int    `json:"height"`
// 		Duration  string `json:"duration"`
// 	} `json:"streams"`
// }

// func probeFFProbe(filePath string, info *MediaInfo) error {
// 	ffp := exec.Command("ffprobe",
// 		"-print_format", "json=compact=1",
// 		"-v", "error",
// 		"-show_entries", "stream=codec_type,width,height,duration",
// 		filePath)

// 	var outBuf bytes.Buffer
// 	var errBuf bytes.Buffer
// 	ffp.Stdout = &outBuf
// 	ffp.Stderr = &errBuf
// 	err := ffp.Run()
// 	if err != nil {
// 		return errors.Wrap(err,
// 			fmt.Sprintf("probe.ffprobe %q", errBuf.String()))
// 	}

// 	var result ResultFFProbe
// 	err = json.Unmarshal(outBuf.Bytes(), &result)
// 	if err != nil {
// 		return errors.Wrap(err, "probe.ffprobe")
// 	}

// 	for _, stream := range result.Streams {
// 		if stream.CodecType == "video" {
// 			dur, err := strconv.ParseFloat(stream.Duration, 32)
// 			if err != nil {
// 				return errors.Wrap(err, "probe.ffprobe.parse")
// 			}
// 			info.Width = stream.Width
// 			info.Height = stream.Height
// 			info.Duration = int(math.Ceil(dur))

// 			return nil
// 		}
// 	}
// 	return errors.New("ffprobe: video stream not found")
// }