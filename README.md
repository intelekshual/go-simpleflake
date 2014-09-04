# go-simpleflake

[![travis](https://travis-ci.org/intelekshual/go-simpleflake.svg)](https://travis-ci.org/intelekshual/go-simpleflake)

64-bit, roughly-ordered, unique ID generator in Go. Inspired by [this article](http://engineering.custommade.com/simpleflake-distributed-id-generation-for-the-lazy/) and the related [Python library](https://github.com/SawdustSoftware/simpleflake).

## Installation

    $ go get -u github.com/intelekshual/go-snowflake

## Usage

```go
import (
  "github.com/intelekshual/go-snowflake"
)

func main() {
  id, err := snowflake.New() // generates a new ID
  if err != nil {
    snowflake.Parse(id) // parses an ID into its time and sequence components
  }
}
```

You can also customize the ID generation options:
```go
snowflake.SetCustomEpoch(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
snowflake.SetCustomPrecision(48)
```

## License

The MIT License (MIT)

Copyright (c) 2014 Robert Coker

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
