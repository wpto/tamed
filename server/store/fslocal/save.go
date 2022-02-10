package fslocal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/h2non/bimg"
	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type MediaContentRepo struct {
	localPath string
}

func NewMediaContentRepo(localPath string) *MediaContentRepo {
	return &MediaContentRepo{localPath}
}

func (rep *MediaContentRepo) Save(contentType string, data []byte) error {
	mediaID, err := NewId(rep.localPath)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

	dirPath := filepath.Join(rep.localPath, mediaID)
	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

	origPath := filepath.Join(dirPath, mediaID)
	err = os.WriteFile(origPath, data, 0644)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

	var mediaInfo *MediaInfo
	if types.MimeToMediaType[contentType] == "vid" {
		mediaInfo, err = VidInfo(origPath)
		if err != nil {
			return errors.Wrap(err, "fslocal.mediacontent.save")
		}

	} else if types.MimeToMediaType[contentType] == "pic" {
		mediaInfo, err = PicInfo(origPath)
		if err != nil {
			return errors.Wrap(err, "fslocal.mediacontent.save")
		}

	}

	meta := &FileMeta{
		Mime:      contentType,
		MediaType: types.MimeToMediaType[contentType],
		Filename:  mediaID,
		ID:        mediaID,
		Width:     mediaInfo.Width,
		Height:    mediaInfo.Height,
		Duration:  mediaInfo.Duration,
	}

	metaData, err := json.Marshal(meta)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

	metaPath := filepath.Join(dirPath, "meta.json")
	err = ioutil.WriteFile(metaPath, metaData, 0644)
	if err != nil {
		return errors.Wrap(err, "fslocal.mediacontent.save")
	}

	return nil
}

type MediaInfo struct {
	Width    int
	Height   int
	Duration int
}

func PicInfo(origPath string) (*MediaInfo, error) {
	buf, err := ioutil.ReadFile(origPath)
	if err != nil {
		return nil, errors.Wrap(err, "picinfo")
	}

	img := bimg.NewImage(buf)
	size, err := img.Size()
	if err != nil {
		return nil, errors.Wrap(err, "picinfo")
	}

	return &MediaInfo{
		Width:    size.Width,
		Height:   size.Height,
		Duration: 0,
	}, nil
}

func VidInfo(origPath string) (*MediaInfo, error) {
	info, err := ffprobeDim(origPath)
	if err != nil {
		return nil, errors.Wrap(err, "vidinfo")
	}
	return info, nil
}

type FFProbeResult struct {
	Streams []struct {
		CodecType string `json:"codec_type"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Duration  string `json:"duration"`
	} `json:"streams"`
}

func ffprobeDim(filePath string) (*MediaInfo, error) {

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
		return nil, errors.Wrap(err,
			fmt.Sprintf("ffprobe dim: %q", errBuf.String()))
	}

	var result FFProbeResult
	err = json.Unmarshal(outBuf.Bytes(), &result)
	if err != nil {
		return nil, errors.Wrap(err, "ffprobe dim")
	}

	for _, stream := range result.Streams {
		if stream.CodecType == "video" {
			dur, err := strconv.ParseFloat(stream.Duration, 32)
			if err != nil {
				return nil, errors.Wrap(err, "ffprobe parse")
			}
			return &MediaInfo{
				Width:    stream.Width,
				Height:   stream.Height,
				Duration: int(math.Ceil(dur)),
			}, nil
		}
	}
	return nil, errors.New("ffprobe: video stream not found")
}
