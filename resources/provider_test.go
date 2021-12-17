package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/stretchr/testify/assert"
)

func TestTableIdentifiers(t *testing.T) {
	t.Parallel()
	for _, res := range Provider().ResourceMap {
		res := res
		t.Run(res.Name, func(t *testing.T) {
			testTableIdentifiers(t, res)
		})
	}
}

func testTableIdentifiers(t *testing.T, table *schema.Table) {
	const maxIdentifierLength = 63 // maximum allowed identifier length is 63 bytes https://www.postgresql.org/docs/13/limits.html

	assert.NotEmpty(t, table.Name)
	assert.LessOrEqual(t, len(table.Name), maxIdentifierLength, "Table name too long")

	for _, c := range table.Columns {
		assert.NotEmpty(t, c.Name)
		assert.LessOrEqual(t, len(c.Name), maxIdentifierLength, "Column name too long:", c.Name)
	}

	for _, res := range table.Relations {
		res := res
		t.Run(res.Name, func(t *testing.T) {
			testTableIdentifiers(t, res)
		})
	}
}
