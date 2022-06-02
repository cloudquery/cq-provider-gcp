package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

// ProjectMultiplex returns a project multiplexer but filters out those projects which don't have all the given services enabled.
func ProjectMultiplex(table string, services ...GcpService) func(schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		cl := meta.(*Client)

		// preallocate all clients just in case
		l := make([]schema.ClientMeta, 0, len(cl.projects))
		for _, projectId := range cl.projects {
			enabled := true
			var missing GcpService
			for _, svc := range services {
				if !cl.EnabledServices[projectId][svc] {
					enabled = false
					missing = svc
					break
				}
			}
			if enabled {
				l = append(l, cl.withProject(projectId))
			} else {
				cl.Logger().Info("skipping fetch for the table due to disabled API", "table", table, "api", missing, "project", projectId)
			}
		}
		return l
	}
}
