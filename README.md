# `version`

Simply get the version information from Go [`BuildInfo`](https://pkg.go.dev/runtime/debug#BuildInfo).

Usage:

```
import "github.com/imjasonh/version"

func main() {
    fmt.Println(version.Get())
}
```

This code is not maintained, and I will not accept PRs. Fork it if you want it.

Please stop using `ldflags` to embed this information. Or if you do, at least embed a reproducible time.