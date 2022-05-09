package serviceusage

import (
	"context"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource services --config gen.hcl --output .
func Services() *schema.Table {
	return &schema.Table{
		Name:         "gcp_serviceusage_services",
		Description:  "A service that is available for use by the consumer",
		Resolver:     fetchServiceusageServices,
		Multiplex:    client.ProjectMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteProjectFilter,
		Columns: []schema.Column{
			{
				Name:        "config_documentation_documentation_root_url",
				Description: "The URL to the root of documentation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Documentation.DocumentationRootUrl"),
			},
			{
				Name:        "config_documentation_overview",
				Description: "Declares a single overview page",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Documentation.Overview"),
			},
			{
				Name:        "config_documentation_service_root_url",
				Description: "Specifies the service root url if the default one (the service name from the yaml file) is not suitable",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Documentation.ServiceRootUrl"),
			},
			{
				Name:        "config_documentation_summary",
				Description: "A short description of what the service does",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Documentation.Summary"),
			},
			{
				Name:        "config_name",
				Description: "The DNS address at which this service is available",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Name"),
			},
			{
				Name:        "config_title",
				Description: "The product title for this service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Title"),
			},
			{
				Name:        "config_usage_producer_notification_channel",
				Description: "The full resource name of a channel used for sending notifications to the service producer",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Usage.ProducerNotificationChannel"),
			},
			{
				Name:        "config_usage_requirements",
				Description: "Requirements that must be satisfied before a consumer project can use the service",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Config.Usage.Requirements"),
			},
			{
				Name:        "name",
				Description: "The resource name of the consumer and service",
				Type:        schema.TypeString,
			},
			{
				Name:        "parent",
				Description: "The resource name of the consumer",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "\"STATE_UNSPECIFIED\" - The default value, which indicates that the enabled state of the service is unspecified or not meaningful Currently, all consumers other than projects (such as folders and organizations) are always in this state   \"DISABLED\" - The service cannot be used by this consumer",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_serviceusage_service_config_apis",
				Description: "Api is a light-weight descriptor for an API Interface Interfaces are also described as \"protocol buffer services\" in some contexts, such as by the \"service\" keyword in a proto file, but they are different from API Services, which represent a concrete implementation of an interface as opposed to simply a description of methods and bindings",
				Resolver:    fetchServiceusageServiceConfigApis,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The fully qualified name of this interface, including package name followed by the interface's simple name",
						Type:        schema.TypeString,
					},
					{
						Name:        "source_context_file_name",
						Description: "The path-qualified name of the proto file that contained the associated protobuf element",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SourceContext.FileName"),
					},
					{
						Name:        "syntax",
						Description: "\"SYNTAX_PROTO2\" - Syntax `proto2`   \"SYNTAX_PROTO3\" - Syntax `proto3`",
						Type:        schema.TypeString,
					},
					{
						Name:        "version",
						Description: "A version string for this interface",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "gcp_serviceusage_service_config_api_methods",
						Description: "Method represents a method of an API interface",
						Resolver:    fetchServiceusageServiceConfigApiMethods,
						Columns: []schema.Column{
							{
								Name:        "service_config_api_cq_id",
								Description: "Unique CloudQuery ID of gcp_serviceusage_service_config_apis table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The simple name of this method",
								Type:        schema.TypeString,
							},
							{
								Name:        "request_streaming",
								Description: "If true, the request is streamed",
								Type:        schema.TypeBool,
							},
							{
								Name:        "request_type_url",
								Description: "A URL of the input message type",
								Type:        schema.TypeString,
							},
							{
								Name:        "response_streaming",
								Description: "If true, the response is streamed",
								Type:        schema.TypeBool,
							},
							{
								Name:        "response_type_url",
								Description: "The URL of the output message type",
								Type:        schema.TypeString,
							},
							{
								Name:        "syntax",
								Description: "\"SYNTAX_PROTO2\" - Syntax `proto2`   \"SYNTAX_PROTO3\" - Syntax `proto3`",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "gcp_serviceusage_service_config_api_method_options",
								Description: "A protocol buffer option, which can be attached to a message, field, enumeration, etc",
								Resolver:    fetchServiceusageServiceConfigApiMethodOptions,
								Columns: []schema.Column{
									{
										Name:        "service_config_api_method_cq_id",
										Description: "Unique CloudQuery ID of gcp_serviceusage_service_config_api_methods table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "name",
										Description: "The option's name",
										Type:        schema.TypeString,
									},
									{
										Name:        "value",
										Description: "The option's value packed in an Any message",
										Type:        schema.TypeByteArray,
									},
								},
							},
						},
					},
					{
						Name:        "gcp_serviceusage_service_config_api_mixins",
						Description: "- If after comment and whitespace stripping, the documentation string of the redeclared method is empty, it will be inherited from the original method",
						Resolver:    fetchServiceusageServiceConfigApiMixins,
						Columns: []schema.Column{
							{
								Name:        "service_config_api_cq_id",
								Description: "Unique CloudQuery ID of gcp_serviceusage_service_config_apis table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The fully qualified name of the interface which is included",
								Type:        schema.TypeString,
							},
							{
								Name:        "root",
								Description: "If non-empty specifies a path under which inherited HTTP paths are rooted",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "gcp_serviceusage_service_config_api_options",
						Description: "A protocol buffer option, which can be attached to a message, field, enumeration, etc",
						Resolver:    fetchServiceusageServiceConfigApiOptions,
						Columns: []schema.Column{
							{
								Name:        "service_config_api_cq_id",
								Description: "Unique CloudQuery ID of gcp_serviceusage_service_config_apis table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The option's name",
								Type:        schema.TypeString,
							},
							{
								Name:        "value",
								Description: "The option's value packed in an Any message",
								Type:        schema.TypeByteArray,
							},
						},
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_authentication_providers",
				Description: "Configuration for an authentication provider, including support for JSON Web Token (JWT) (https://toolsietforg/html/draft-ietf-oauth-json-web-token-32)",
				Resolver:    fetchServiceusageServiceConfigAuthenticationProviders,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "audiences",
						Description: "The list of JWT audiences (https://toolsietforg/html/draft-ietf-oauth-json-web-token-32#section-413) that are allowed to access",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorization_url",
						Description: "Redirect URL if JWT token is required but not present or is expired",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The unique identifier of the auth provider",
						Type:        schema.TypeString,
					},
					{
						Name:        "issuer",
						Description: "Identifies the principal that issued the JWT",
						Type:        schema.TypeString,
					},
					{
						Name:        "jwks_uri",
						Description: "URL of the provider's public key set to validate signature of the JWT",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "gcp_serviceusage_service_config_authentication_provider_jwt_locations",
						Description: "Specifies a location to extract JWT from an API request",
						Resolver:    fetchServiceusageServiceConfigAuthenticationProviderJwtLocations,
						Columns: []schema.Column{
							{
								Name:        "service_config_authentication_provider_cq_id",
								Description: "Unique CloudQuery ID of gcp_serviceusage_service_config_authentication_providers table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "header",
								Description: "Specifies HTTP header name to extract JWT token",
								Type:        schema.TypeString,
							},
							{
								Name:        "query",
								Description: "Specifies URL query parameter name to extract JWT token",
								Type:        schema.TypeString,
							},
							{
								Name:        "value_prefix",
								Description: "The value prefix",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_authentication_rules",
				Description: "Authentication rules for the service",
				Resolver:    fetchServiceusageServiceConfigAuthenticationRules,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allow_without_credential",
						Description: "If true, the service accepts API keys without any other credential",
						Type:        schema.TypeBool,
					},
					{
						Name:        "oauth_canonical_scopes",
						Description: "The list of publicly documented OAuth scopes that are allowed access",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Oauth.CanonicalScopes"),
					},
					{
						Name:        "selector",
						Description: "Selects the methods to which this rule applies",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "gcp_serviceusage_service_config_authentication_rule_requirements",
						Description: "User-defined authentication requirements, including support for JSON Web Token (JWT) (https://toolsietforg/html/draft-ietf-oauth-json-web-token-32)",
						Resolver:    fetchServiceusageServiceConfigAuthenticationRuleRequirements,
						Columns: []schema.Column{
							{
								Name:        "service_config_authentication_rule_cq_id",
								Description: "Unique CloudQuery ID of gcp_serviceusage_service_config_authentication_rules table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "audiences",
								Description: "This will be deprecated soon, once AuthProvideraudiences is implemented and accepted in all the runtime components",
								Type:        schema.TypeString,
							},
							{
								Name:        "provider_id",
								Description: "id from authentication provider",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_documentation_pages",
				Description: "Represents a documentation page",
				Resolver:    fetchServiceusageServiceConfigDocumentationPages,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "content",
						Description: "The Markdown content of the page",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of the page",
						Type:        schema.TypeString,
					},
					{
						Name:        "subpages",
						Description: "Subpages of this page",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_documentation_rules",
				Description: "A documentation rule provides information about individual API elements",
				Resolver:    fetchServiceusageServiceConfigDocumentationRules,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "deprecation_description",
						Description: "Deprecation description of the selected element(s)",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "Description of the selected proto element (eg",
						Type:        schema.TypeString,
					},
					{
						Name:        "selector",
						Description: "The selector is a comma-separated list of patterns for any element such as a method, a field, an enum value",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_endpoints",
				Description: "`Endpoint` describes a network address of a service that serves a set of APIs",
				Resolver:    fetchServiceusageServiceConfigEndpoints,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allow_cors",
						Description: "Allowing CORS (https://enwikipediaorg/wiki/Cross-origin_resource_sharing), aka cross-domain traffic, would allow the backends served from this endpoint to receive and respond to HTTP OPTIONS requests",
						Type:        schema.TypeBool,
					},
					{
						Name:        "name",
						Description: "The canonical name of this endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "target",
						Description: "The specification of an Internet routable address of API frontend that will handle requests to this API Endpoint (https://cloudgooglecom/apis/design/glossary)",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_monitored_resources",
				Description: "An object that describes the schema of a MonitoredResource object using a type name and a set of labels",
				Resolver:    fetchServiceusageServiceConfigMonitoredResources,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "description",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "display_name",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "launch_stage",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Required",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "gcp_serviceusage_service_config_monitored_resource_labels",
						Description: "A description of a label",
						Resolver:    fetchServiceusageServiceConfigMonitoredResourceLabels,
						Columns: []schema.Column{
							{
								Name:        "service_config_monitored_resource_cq_id",
								Description: "Unique CloudQuery ID of gcp_serviceusage_service_config_monitored_resources table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "description",
								Description: "A human-readable description for the label",
								Type:        schema.TypeString,
							},
							{
								Name:        "key",
								Description: "The label key",
								Type:        schema.TypeString,
							},
							{
								Name:        "value_type",
								Description: "\"STRING\" - A variable-length string",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_monitoring_consumer_destinations",
				Description: "Configuration of a specific monitoring destination (the producer project or the consumer project)",
				Resolver:    fetchServiceusageServiceConfigMonitoringConsumerDestinations,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "metrics",
						Description: "Types of the metrics to report to this monitoring destination",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "monitored_resource",
						Description: "The monitored resource type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_monitoring_producer_destinations",
				Description: "Configuration of a specific monitoring destination (the producer project or the consumer project)",
				Resolver:    fetchServiceusageServiceConfigMonitoringProducerDestinations,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "metrics",
						Description: "Types of the metrics to report to this monitoring destination",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "monitored_resource",
						Description: "The monitored resource type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_quota_limits",
				Description: "`QuotaLimit` defines a specific limit that applies over a specified duration for a limit type",
				Resolver:    fetchServiceusageServiceConfigQuotaLimits,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "default_limit",
						Description: "Default number of tokens that can be consumed during the specified duration",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "description",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "display_name",
						Description: "User-visible display name for this limit",
						Type:        schema.TypeString,
					},
					{
						Name:        "duration",
						Description: "Duration of this limit in textual notation",
						Type:        schema.TypeString,
					},
					{
						Name:        "free_tier",
						Description: "Free tier value displayed in the Developers Console for this limit",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "max_limit",
						Description: "Maximum number of tokens that can be consumed during the specified duration",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "metric",
						Description: "The name of the metric this quota limit applies to",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Name of the quota limit",
						Type:        schema.TypeString,
					},
					{
						Name:        "unit",
						Description: "Specify the unit of the quota limit",
						Type:        schema.TypeString,
					},
					{
						Name:        "values",
						Description: "Tiered limit values",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_quota_metric_rules",
				Description: "Bind API methods to metrics",
				Resolver:    fetchServiceusageServiceConfigQuotaMetricRules,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "metric_costs",
						Description: "Metrics to update when the selected methods are called, and the associated cost applied to each metric",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "selector",
						Description: "Selects the methods to which this rule applies",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_config_usage_rules",
				Description: "Usage configuration rules for the service",
				Resolver:    fetchServiceusageServiceConfigUsageRules,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allow_unregistered_calls",
						Description: "If true, the selected method allows unregistered calls, eg",
						Type:        schema.TypeBool,
					},
					{
						Name:        "selector",
						Description: "Selects the methods to which this rule applies",
						Type:        schema.TypeString,
					},
					{
						Name:        "skip_service_control",
						Description: "If true, the selected method should skip service control and the control plane features, such as quota and billing, will not be available",
						Type:        schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchServiceusageServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigApiMethods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigApiMethodOptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigApiMixins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigApiOptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigAuthenticationProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigAuthenticationProviderJwtLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigAuthenticationRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigAuthenticationRuleRequirements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigDocumentationPages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigDocumentationRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigMonitoredResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigMonitoredResourceLabels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigMonitoringConsumerDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigMonitoringProducerDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigQuotaLimits(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigQuotaMetricRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchServiceusageServiceConfigUsageRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
