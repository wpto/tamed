package service

import (
	"fmt"
	"os"
	"time"

	"github.com/jxskiss/base62"
	"github.com/sony/sonyflake"
)

func UniqID() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		fmt.Printf("srv.art.create.flakeid %v", err)
		os.Exit(1)
	}

	return string(base62.FormatUint(id))
}

func TimeNow() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
