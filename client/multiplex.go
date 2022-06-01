package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ProjectMultiplex(table string, required ...GcpService) func(schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		client := meta.(*Client)

		l := make([]schema.ClientMeta, 0, len(client.projects))
		for _, projectId := range client.projects {
			enabled := true
			var disabledApi GcpService
			for _, svc := range required {
				if _, ok := client.EnabledServices[projectId][svc]; !ok {
					enabled = false
					disabledApi = svc
					break
				}
			}
			if enabled {
				l = append(l, client.withProject(projectId))
			} else {
				client.Logger().Info("skipping fetch for the table due to disabled API", "table", table, "api", disabledApi, "project", projectId)
			}
		}
		return l
	}
}
