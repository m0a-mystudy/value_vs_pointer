# int型のフィールドいくつまでなら値渡しで良いのか？

* 結論


```
type ValueVsPointer04 struct {
	Filed1 int
	Filed2 int
	Filed3 int
	Filed4 int
}
```
上記くらいなら値渡しでもよし

* 気になる点

フィールド数が４つまでだと同じ時間しかかからず
5つ目で突然変化したのが気になる。



## result

```
go version
go version go1.11.1 darwin/amd64

$ go test -bench . -benchtime=10s
goos: darwin
goarch: amd64
pkg: github.com/m0a-mystudy/value_vs_pointer
BenchmarkUsePointerFieldCount01-8   	10000000000	         0.28 ns/op
BenchmarkUseValueFieldCount01-8     	10000000000	         0.28 ns/op
BenchmarkUsePointerFieldCount02-8   	10000000000	         0.28 ns/op
BenchmarkUseValueFieldCount02-8     	10000000000	         0.28 ns/op
BenchmarkUsePointerFieldCount03-8   	10000000000	         0.28 ns/op
BenchmarkUseValueFieldCount03-8     	10000000000	         0.28 ns/op
BenchmarkUsePointerFieldCount04-8   	10000000000	         0.28 ns/op
BenchmarkUseValueFieldCount04-8     	10000000000	         0.28 ns/op
BenchmarkUsePointerFieldCount05-8   	10000000000	         2.30 ns/op
BenchmarkUseValueFieldCount05-8     	3000000000	         4.74 ns/op
BenchmarkUsePointerFieldCount08-8   	3000000000	         4.18 ns/op
BenchmarkUseValueFieldCount08-8     	2000000000	         7.00 ns/op
```
