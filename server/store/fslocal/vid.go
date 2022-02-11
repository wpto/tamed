package fslocal

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pgeowng/tamed/types"
	"github.com/pkg/errors"
)

type MediaVidRepo struct {
	localPath string
}

func NewMediaVidRepo(localPath string) *MediaVidRepo {
	return &MediaVidRepo{localPath: localPath}
}

func TempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
}

var CodecPresets = map[string][]string{
	"video/mp4":  []string{"-c:v", "libx264", "-preset", "ultrafast", "-crf", "28"},
	"video/webm": []string{},
}

func (rep *MediaVidRepo) GetContent(opts *types.GetVidOpts) ([]byte, error) {
	dirPath := filepath.Join(rep.localPath, opts.MediaID)
	meta, err := TryFileMeta(dirPath)
	if err != nil {
		return nil, errors.Wrap(err, "fslocal.getvid")
	}

	origPath := filepath.Join(dirPath, meta.Filename)

	data := make([]byte, 0)

	if opts.Width != 0 && opts.Width < meta.Width {
		args := []string{
			"-i", origPath,
		}

		args = append(args, CodecPresets["video/mp4"]...)

		args = append(args, "-vf", fmt.Sprintf("scale=%d:-2", opts.Width),
			"-an")

		fmt.Println(args)

		data, err = ffmpeg(args...)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("fslocal.getvid(%v)", args))
		}
	} else {
		data, err = ioutil.ReadFile(origPath)
		if err != nil {
			return nil, errors.Wrap(err, "fslocal.getvid")
		}
	}

	return data, nil
}

func ffmpeg(args ...string) ([]byte, error) {
	tempPath := TempFileName("videocrop-", ".mp4")
	args = append(args, tempPath)

	ffm := exec.Command("ffmpeg", args...)
	var ffmOut bytes.Buffer
	var ffmErr bytes.Buffer
	ffm.Stdout = &ffmOut
	ffm.Stderr = &ffmErr
	err := ffm.Run()
	if err != nil {
		fmt.Println(args)
		return nil, errors.Wrap(err, fmt.Sprintf("ffmpeg transform: %q,\n stderr: %q", ffmOut.String(), ffmErr.String()))
	}

	data, err := ioutil.ReadFile(tempPath)
	if err != nil {
		return nil, errors.Wrap(err, "ffmpeg read")
	}

	return data, nil
}
