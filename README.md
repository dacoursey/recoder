# recoder

Tool for testing as many encoding methods as possible, as well as some combinations.

Right now the en/decoding is safe, meaning if the input string doesn't generate printable characters nothing will be returned for that coder. I plan on adding unsafe output at a later date.

### Installation

Install through Go:
```
$ go get github.com/dacoursey/recoder
$ $GOPATH/bin/recoder
```

Or, pull the source and install manually:
```
$ git clone https://github.com/dacoursey/recoder
$ cd recoder
$ go install
$ $GOPATH/bin/recoder
```

### Supported Encoding Mechanisms

I'm going to try to put as many as possible, but here is the running list of things that are working:
- Hex
- Base64
- Base32