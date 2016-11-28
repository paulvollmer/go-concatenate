# go-concatenate [![Build Status](https://travis-ci.com/paulvollmer/go-concatenate.svg?token=Wck9khUHkFjieXJypmaZ&branch=master)](https://travis-ci.com/paulvollmer/go-concatenate)
go concatenate utility lib


## Usage
```
go get github.com/paulvollmer/go-concatenate

install "github.com/paulvollmer/go-concatenate"

data := concatenate.StringsToString("-", "hello", "world")
data, err := concatenate.FilesToBytes("\n", "foo.txt", "bar.txt")
```

## License
Licensed under [MIT-License](LICENSE)
