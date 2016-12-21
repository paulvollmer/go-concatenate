# go-concatenate [![Build Status](https://travis-ci.org/paulvollmer/go-concatenate.svg?branch=master)](https://travis-ci.org/paulvollmer/go-concatenate) [![Report](https://goreportcard.com/badge/github.com/paulvollmer/go-concatenate)](https://goreportcard.com/report/github.com/paulvollmer/go-concatenate)
golang concatenate package and cli with focus on speed.


## Usage
```
go get github.com/paulvollmer/go-concatenate

install "github.com/paulvollmer/go-concatenate"

data := concatenate.StringsToString("-", "hello", "world")
data, err := concatenate.FilesToBytes("\n", "foo.txt", "bar.txt")
```

## License
Licensed under [MIT-License](LICENSE)
