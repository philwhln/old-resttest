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
$ go run go/balance.go
18377.16
```

Both output the same value luckily, although the Go implementation of obviously much much faster. Also, the Bash implementation is not as robust, less easily expensible and I'm not sure how well `bc` handles rounding errors. I assume its using float for internal representation.

## Additional Features

### Daily balances

Extended the implementation to add daily totals.

```
$ go run go/daily.go
2013-12-12	-227.35
2013-12-13	-1229.58
2013-12-15	-5.39
2013-12-16	-4575.53
2013-12-17	10686.28
2013-12-18	-1841.29
2013-12-19	19753.31
2013-12-20	-4054.60
2013-12-21	-17.98
2013-12-22	-110.71
----------	-----
TOTAL     	18377.16
```

