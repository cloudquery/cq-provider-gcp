package main

import (
	"github.com/cloudquery/cq-provider-gcp/resources/plugin"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(serve.Options{
		SourcePlugin: plugin.Plugin(),
	})
}
