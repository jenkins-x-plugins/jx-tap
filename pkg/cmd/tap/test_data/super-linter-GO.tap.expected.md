GO Linter
* [main.go](main.go#L30) : `Something` is unused (deadcode)

```go
func Something() {
     ^
```
* [main.go](main.go#L27) : Error return value of `http.ListenAndServe` is not checked (errcheck)

```go
	http.ListenAndServe(" 8080", nil)
	                   ^
```
* [main.go](main.go#L23) : File is not `goimports`-ed (goimports)
