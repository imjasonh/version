# `version`

Simply get the version information from Go [`BuildInfo`](https://pkg.go.dev/runtime/debug#BuildInfo).

Usage:

```
import "github.com/imjasonh/version"

func main() {
    fmt.Println(version.Get())
}
```

This code is not maintained, and I will not accept PRs. Fork it if you want it. It's like 50 lines, just copy it into your codebase.

Please stop using `ldflags` to embed this information. Or if you do, at least embed a reproducible time.

-----

### Rant

Since the early days of Go, if you wanted to embed something like the Git commit that was built in your binary, you had to do something janky like

```
package main

const commit = "unknown"
```

and

```
go build -ldflags="-X 'main.commit=$(git rev-parse HEAD)'"
```

This was kinda gross, but it worked, so people did it. And _boy_ did they do it. This little trick got cargo-culted all over the place! Anecdotally, basically any Go application you use has its version embedded this way.

Sometimes folks also wanted a build date embedded, and okay, sure, we can do that too:

```
go build -ldflags="-X 'main.buildDate=$(date +%Y-%m-%dT%H:%M:%SZ)'"
```

Boom, done. Except...

That's going to cause your build to be non-reproducible. If you build that right now, wait ten seconds, and build it again, you'll get a new binary. You probably don't care that they were built 10 seconds apart from the same source. What you _wanted_ was

```
go build -ldflags="-X 'main.buildDate=$(date -d@$SOURCE_DATE_EPOCH +%Y-%m-%dT%H:%M:%SZ)'"
```

and some sane static value of `SOURCE_DATE_EPOCH` to make that reproducible. Most folks would opt for the date of the commit that it's built from as a sane approach.

_But you don't need to do any of this at all, so stop!_

**Since Go 1.12, back in 2019**, the Go compiler has embedded this exact information for you, without you needing to know what an `ldflag` is. You can just call [`debug.ReadBuildInfo`](https://pkg.go.dev/runtime/debug#ReadBuildInfo) and it'll pop right out, along with a bunch of other stuff.

The `vcs.revision` and `vcs.time` are exactly the same as `git rev-parse HEAD` and "the date of that commit". So stop mucking with `ldflags`, and _especially_ stop embedding a non-reproducible time. You're working too hard.

This package demonstrates this, and encapsulates it in `version.Get`. If you want to change anything about the behavior, copy the code and go for it. I wrapped it in a [`sync.OnceFunc`](https://pkg.go.dev/sync#OnceFunc) so it didn't have to read the info each time, in case you call it multiple times. You're welcome.
