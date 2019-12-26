# github.com/GolubAlexander/config

### Description.
Library-helper to unmarshal a configuration file or a slice of bytes to the structure.

This library supports next types:
```go
config.TypeJson
config.TypeYaml
```

Also the library has next types of errors:
```go
ErrDataNull       = errors.New("bytes must not be a null")
ErrDataEmpty      = errors.New("bytes must not be an empty slice")
ErrNotImplemented = errors.New("file's format is not implemented")
ErrNotPointer     = errors.New("param must be a pointer to a struct or a map")
ErrReadFile       = errors.New("read config file")
ErrDecodeData     = errors.New("decode config")
```

### How's to use.
Imports the library.
```go
import "github.com/GolubAlexander/config"
```
Describes type.
```go
type conf struct {
    Test string `json:"test" yaml:"test"`
}
```
Makes function's calls to load from a file.
```go
fileName:= "./example.json"
var cfg conf
if err := config.FromFile(&cfg, fileName); err != nil {
	panic(err)
}
```
Or to load from a slice of bytes:
```go
confBytes := []byte("test: 'test'")
if err := config.FromBytes(&cfg, confBytes, config.TypeYaml); err != nil {
    panic(err)
}
```