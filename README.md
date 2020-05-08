# gophish

This library is a Go API client for interacting with the [`gophish` API](https://docs.getgophish.com/api-documentation/).

This is currently a WIP. There shouldn't be any backwards incompatible changes, but no promises.

## Using the API client

If you're looking for the API client, this library supports go modules, so 
treat it as a normal library. For interacting with the client, all subgroups
of functionality can be referenced from the main client in a namespaced
manner.

### Creating a new client
To create a client, you need to provide the host where your `gophish`
installation is living and an API token.

```go
package main

import (
    "flag"
    
    "github.com/ttacon/gophish"
)

var (
    host = flag.String("host", "", "gophish host")
    token = flag.String("token", "", "gophish API token")
)

func main() {
    flag.Parse()
    
    client := gophish.NewClient(*host, *token)
    
    //...
}
```

### Using functionality groups
To interact with a specific resource of functionaliy group, just reference
it from the client. As an example, to list all templates, we'd do:

```go
// listGophishTemplates lists all templates in the given gophish installation.
func listGophishTemplates(client *gophish.Client) error  {
    templates, err := client.Templates.ListTemplates()
    if err != nil {
        return err
    }
    
    for _, tmplt := range templates {
        fmt.Println(tmplt)
    }
    return nil
}
```

## Using the CLI

### Installing the CLI
There's a very barebones CLI (it needs some love!), that also shows usage of
the library. To install it from source, assuming that your `$GOBIN` is on your
`$PATH`:

```sh
go get github.com/ttacon/gophish/cmd/guppie
```

### Example: retrieving all templates via the CLI
As with the API client, we need to provide the gophish host and an API token,
and then we're off to the races:

```sh
guppie --host=$host --token=$token templates list
```