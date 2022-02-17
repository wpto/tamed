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
	err := probeMediainfo(origPath, info)
	if err != nil {
		return errors.Wrap(err, "vidinfo")
	}
	return nil
}

type ResultFFProbe struct {
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
			fmt.Sprintf("probe.ffprobe %q", errBuf.String()))
	}

	var result ResultFFProbe
	err = json.Unmarshal(outBuf.Bytes(), &result)
	if err != nil {
		return errors.Wrap(err, "probe.ffprobe")
	}

	for _, stream := range result.Streams {
		if stream.CodecType == "video" {
			dur, err := strconv.ParseFloat(stream.Duration, 32)
			if err != nil {
				return errors.Wrap(err, "probe.ffprobe.parse")
			}
			info.Width = stream.Width
			info.Height = stream.Height
			info.Duration = int(math.Ceil(dur))

			return nil
		}
	}
	return errors.New("ffprobe: video stream not found")
}

type ResultMediainfo struct {
	Media struct {
		Track []struct {
			Type     string `json:"@type"`
			Duration string
			Width    string
			Height   string
		} `json:"track"`
	} `json:"media"`
}

func probeMediainfo(filePath string, info *MediaInfo) error {
	mi := exec.Command("mediainfo", "--Output=JSON", filePath)

	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	mi.Stdout = &outBuf
	mi.Stderr = &errBuf
	err := mi.Run()
	if err != nil {
		return errors.Wrap(err,
			fmt.Sprintf("probe.mediainfo: %q", errBuf.String()))
	}

	var result ResultMediainfo
	err = json.Unmarshal(outBuf.Bytes(), &result)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("probe.mediainfo(input %q)", outBuf.String()))
	}

	for _, track := range result.Media.Track {
		if track.Type == "Video" {
			fmt.Println(track)
			duration, err := strconv.ParseFloat(track.Duration, 32)
			if err != nil {
				return errors.Wrap(err, "probe.mediainfo.result.track.Duration")
			}

			width, err := strconv.ParseInt(track.Width, 10, 0)
			if err != nil {
				return errors.Wrap(err, "probe.mediainfo.result.track.Width")
			}

			height, err := strconv.ParseInt(track.Height, 10, 0)
			if err != nil {
				return errors.Wrap(err, "probe.mediainfo.result.track.Height")
			}

			info.Duration = int(math.Ceil(duration))
			info.Width = int(width)
			info.Height = int(height)

			return nil
		}
	}

	return errors.New("probe.mediainfo.result: parse error")
}
