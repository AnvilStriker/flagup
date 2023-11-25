# flagup
The beginnings of an exploration of a (possibly!) low-impact way of incrementally upgrading a large codebase from `codegangsta/cli` v1 flags to `urfave/cli/v2`.

Not sure if this approach may be more harmful than helpful, nor what kinds of blind alleys it may lead into.

## Contents
- `appv1/main.go` - The "original" program using v1-style flags.
- `utils/flags.go` - The v1-style flag definitions.
- `utils/const.go` - Some constants that the flags reference.
- `appv2/main.go` - A revised version of the original program using v2-style flags, synthesized from v1-style flags.
- `utils/flagup.go` - The v2-style flag synthesizer.
