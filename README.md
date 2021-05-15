# NotionGo

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/google/go-github/v35/github)
[![Test Status](https://github.com/google/go-github/workflows/tests/badge.svg)](https://github.com/google/go-github/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/google/go-github/branch/master/graph/badge.svg)](https://codecov.io/gh/google/go-github)


NotionGo is a Go client library for accessing the [Notion API v1]("https://developers.notion.com/").


## Installation ##

NotionGo is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/tempor1s/notiongo/
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/tempor1s/notiongo/notion"
```

## Usage ##

```go
import "github.com/tempor1s/notiongo/v1/notion"	// with go modules enabled (GO111MODULE=on or outside GOPATH)
import "github.com/tempor1s/notiongo/notion" // with go modules disabled
```

Construct a new Notion client, then use the various services on the client to
access different parts of the Notion API. For example:

```go
client := notion.NewClient(nil)

// Retrieves a Page object using the ID specified.
page, err := client.page(id)
```

### Integration Tests ###

You can run integration tests from the `test` directory. See the integration tests [README](test/README.md).

## Contributing ##
I would like to cover the entire GitHub API and contributions are of course always welcome. The
calling pattern is pretty well established, so adding new methods is relatively
straightforward. See [`CONTRIBUTING.md`](CONTRIBUTING.md) for details.


## License ##

Library distributed under the Apache-2.0 license found in the [LICENSE](./LICENSE)
file.