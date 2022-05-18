package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/spf13/cast"
	"github.com/thoas/go-funk"
)

func ResolveProject(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("project_id", client.ProjectId)
}

func resolveLocation(_ context.Context, c schema.ClientMeta, r *schema.Resource) error {
	loc := r.Get("location")
	if loc != nil {
		return nil
	}
	name := r.Get("name")
	if name == nil {
		c.Logger().Warn("missing name for resource", "resource", fmt.Sprintf("%T", r.Item))
		return nil
	}
	nameStr, err := cast.ToStringE(name)
	if err != nil {
		return err
	}
	match := strings.Split(nameStr, "/")
	if len(match) < 3 {
		return nil
	}
	return r.Set("location", match[3])
}

func AddGcpMetadata(ctx context.Context, c schema.ClientMeta, r *schema.Resource) error {
	return resolveLocation(ctx, c, r)
}

func ResolveResourceId(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	data, err := cast.ToStringE(funk.Get(r.Item, "Id", funk.WithAllowZero()))
	if err != nil {
		return err
	}
	return r.Set(c.Name, data)
}

// RequireEnabledServices is a resolver that calls a real resolver passed as an argument
// only if all required services are enabled in the current client.
func RequireEnabledServices(resolver schema.TableResolver, services ...GcpService) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
		cl := meta.(*Client)
		for _, s := range services {
			if _, ok := cl.EnabledServices[cl.ProjectId][s]; !ok {
				cl.Logger().Info("fetch skipped due to disabled API", "service", s)
				return nil
			}
		}
		return resolver(ctx, meta, parent, res)
	}
}
