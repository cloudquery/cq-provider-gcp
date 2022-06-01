package provider

import (
	_ "embed"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/views"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegration(t *testing.T) {
	providertest.TestView(t, providertest.ViewTestCase{
		Provider: Provider(),
		SQLView:  views.ResourceView,
	})
}
