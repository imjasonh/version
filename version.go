package version

import (
	"fmt"
	"log/slog"
	"runtime/debug"
	"sync"
)

type Version struct {
	Revision, Version, Time string
	Dirty                   bool
}

var ver = Version{
	Revision: "unknown",
	Version:  "unknown",
	Time:     "unknown",
	Dirty:    false,
}

func (v Version) String() string {
	return fmt.Sprintf(`Revision: %s
Version: %s
BuildTime: %s
Dirty: %t`, v.Revision, v.Version, v.Time, v.Dirty)
}

func Get() Version {
	sync.OnceFunc(func() {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			slog.Warn("version: no build info detected")
			return
		}
		if bi.Main.Version != "" && bi.Main.Version != "(devel)" {
			ver.Version = bi.Main.Version
		}
		for _, setting := range bi.Settings {
			switch setting.Key {
			case "vcs.revision":
				ver.Revision = setting.Value
			case "vcs.time":
				ver.Time = setting.Value
			case "vcs.modified":
				ver.Dirty = setting.Value == "true"
			}
		}
	})()
	return ver
}
