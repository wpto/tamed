package mediameta

// type ResultMediainfo struct {
// 	Media struct {
// 		Track []struct {
// 			Type     string `json:"@type"`
// 			Duration string
// 			Width    string
// 			Height   string
// 		} `json:"track"`
// 	} `json:"media"`
// }
// func probeMediainfo(filePath string, info *MediaInfo) error {
// 	mi := exec.Command("mediainfo", "--Output=JSON", filePath)

// 	var outBuf bytes.Buffer
// 	var errBuf bytes.Buffer
// 	mi.Stdout = &outBuf
// 	mi.Stderr = &errBuf
// 	err := mi.Run()
// 	if err != nil {
// 		return errors.Wrap(err,
// 			fmt.Sprintf("probe.mediainfo: %q", errBuf.String()))
// 	}

// 	var result ResultMediainfo
// 	err = json.Unmarshal(outBuf.Bytes(), &result)
// 	if err != nil {
// 		return errors.Wrap(err, fmt.Sprintf("probe.mediainfo(input %q)", outBuf.String()))
// 	}

// 	for _, track := range result.Media.Track {
// 		if track.Type == "Video" {
// 			fmt.Println(track)
// 			duration, err := strconv.ParseFloat(track.Duration, 32)
// 			if err != nil {
// 				return errors.Wrap(err, "probe.mediainfo.result.track.Duration")
// 			}

// 			width, err := strconv.ParseInt(track.Width, 10, 0)
// 			if err != nil {
// 				return errors.Wrap(err, "probe.mediainfo.result.track.Width")
// 			}

// 			height, err := strconv.ParseInt(track.Height, 10, 0)
// 			if err != nil {
// 				return errors.Wrap(err, "probe.mediainfo.result.track.Height")
// 			}

// 			info.Duration = int(math.Ceil(duration))
// 			info.Width = int(width)
// 			info.Height = int(height)

// 			return nil
// 		}
// 	}

// 	return errors.New("probe.mediainfo.result: parse error")
// }
