package commonsrv

import (
	"fmt"
	"os"
	"time"

	"github.com/jxskiss/base62"
	"github.com/sony/sonyflake"
)

var lastID uint64 = 0
var machineID uint16 = 0

func UniqID() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return machineID, nil
		},
	})
	newID, err := flake.NextID()
	if err != nil {
		fmt.Printf("srv.art.create.flakeid %v", err)
		os.Exit(1)
	}

	if newID == lastID {
		machineID += 1 // do proper way
		return UniqID()
	}

	lastID = newID
	return string(base62.FormatUint(newID))
}

func TimeNow() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
