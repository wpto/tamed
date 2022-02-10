package fslocal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os/exec"
	"strconv"

	"github.com/h2non/bimg"
	"github.com/pkg/errors"
)

type MediaInfo struct {
	Width    int
	Height   int
	Duration int
}

func PicInfo(origPath string, info *MediaInfo) error {
	buf, err := ioutil.ReadFile(origPath)
	if err != nil {
		return errors.Wrap(err, "picinfo")
	}

	img := bimg.NewImage(buf)
	size, err := img.Size()
	if err != nil {
		return errors.Wrap(err, "picinfo")
	}

	info.Width = size.Width
	info.Height = size.Height
	info.Duration = 0

	return nil
}

func VidInfo(origPath string, info *MediaInfo) error {
	err := probeFFProbe(origPath, info)
	if err != nil {
		return errors.Wrap(err, "vidinfo")
	}
	return nil
}

type FFProbeResult struct {
	Streams []struct {
		CodecType string `json:"codec_type"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Duration  string `json:"duration"`
	} `json:"streams"`
}

func probeFFProbe(filePath string, info *MediaInfo) error {
	ffp := exec.Command("ffprobe",
		"-print_format", "json=compact=1",
		"-v", "error",
		"-show_entries", "stream=codec_type,width,height,duration",
		filePath)

	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	ffp.Stdout = &outBuf
	ffp.Stderr = &errBuf
	err := ffp.Run()
	if err != nil {
		return errors.Wrap(err,
			fmt.Sprintf("ffprobe dim: %q", errBuf.String()))
	}

	var result FFProbeResult
	err = json.Unmarshal(outBuf.Bytes(), &result)
	if err != nil {
		return errors.Wrap(err, "ffprobe dim")
	}

	for _, stream := range result.Streams {
		if stream.CodecType == "video" {
			dur, err := strconv.ParseFloat(stream.Duration, 32)
			if err != nil {
				return errors.Wrap(err, "ffprobe parse")
			}
			info.Width = stream.Width
			info.Height = stream.Height
			info.Duration = int(math.Ceil(dur))

			return nil
		}
	}
	return errors.New("ffprobe: video stream not found")
}

// type MediainfoResult struct {
// }

// func probeMediainfo() {}
