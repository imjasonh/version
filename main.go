package version

import (
	"fmt"
	"log/slog"
	"runtime/debug"
	"sync"
)

type Version struct {
	Revision string
	Time     string
	Dirty    bool
}

var ver = Version{
	Revision: "unknown",
	Time:     "unknown",
	Dirty:    false,
}
var once sync.Once

func (v Version) String() string {
	return fmt.Sprintf(`Revision: %s
BuildTime: %s
Dirty: %t`, v.Revision, v.Time, v.Dirty)
}

func Get() Version {
	once.Do(func() {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			slog.Warn("version: no build info detected")
			return
		}
		for _, setting := range bi.Settings {
			if setting.Value == "" {
				continue
			}
			switch setting.Key {
			case "vcs.revision":
				ver.Revision = setting.Value
			case "vcs.time":
				ver.Time = setting.Value
			case "vcs.modified":
				ver.Dirty = setting.Value == "true"
			}
		}
	})
	return ver
}
