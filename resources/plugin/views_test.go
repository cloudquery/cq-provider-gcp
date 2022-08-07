package plugin

import (
	_ "embed"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/views"
	providertest "github.com/cloudquery/cq-provider-sdk/testing"
)

func TestViews(t *testing.T) {
	providertest.HelperTestView(t, providertest.ViewTestCase{
		Provider: Plugin(),
		SQLView:  views.ResourceView,
	})
}
