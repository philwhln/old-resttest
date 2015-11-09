# Resttest

Two implementations of http://resttest.bench.co/

## Bash

The quick and dirty implementation using `curl`, `jq` and `bc`.

```
$ bash/resttest.sh
18377.16
```

## Golang

Go implentation

```
$ go run go/main.go
18377.16
```

Both output the same value luckily, although the Go implementation of obviously much much faster. Also, the Bash implementation is not as robust, less easily expensible and I'm not sure how well `bc` it handles rounding errors. I assume its using float for internal representation.

