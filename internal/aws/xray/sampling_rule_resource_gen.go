// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package xray

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_xray_sampling_rule", samplingRuleResource)
}

// samplingRuleResource returns the Terraform awscc_xray_sampling_rule resource.
// This Terraform resource corresponds to the CloudFormation AWS::XRay::SamplingRule resource.
func samplingRuleResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: RuleARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	  "type": "string"
		//	}
		"rule_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RuleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	  "maxLength": 32,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 32),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SamplingRule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Attributes": {
		//	      "additionalProperties": false,
		//	      "$comment": "String to string map",
		//	      "description": "Matches attributes derived from the request.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "FixedRate": {
		//	      "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
		//	      "maximum": 1,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "HTTPMethod": {
		//	      "description": "Matches the HTTP method from a request URL.",
		//	      "maxLength": 10,
		//	      "type": "string"
		//	    },
		//	    "Host": {
		//	      "description": "Matches the hostname from a request URL.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "Priority": {
		//	      "description": "The priority of the sampling rule.",
		//	      "maximum": 9999,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "ReservoirSize": {
		//	      "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "ResourceARN": {
		//	      "description": "Matches the ARN of the AWS resource on which the service runs.",
		//	      "maxLength": 500,
		//	      "type": "string"
		//	    },
		//	    "RuleARN": {
		//	      "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	      "type": "string"
		//	    },
		//	    "RuleName": {
		//	      "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	      "maxLength": 32,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "ServiceName": {
		//	      "description": "Matches the name that the service uses to identify itself in segments.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "ServiceType": {
		//	      "description": "Matches the origin that the service uses to identify its type in segments.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "URLPath": {
		//	      "description": "Matches the path from a request URL.",
		//	      "maxLength": 128,
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of the sampling rule format (1)",
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "FixedRate",
		//	    "Host",
		//	    "HTTPMethod",
		//	    "Priority",
		//	    "ReservoirSize",
		//	    "ResourceARN",
		//	    "ServiceName",
		//	    "ServiceType",
		//	    "URLPath"
		//	  ],
		//	  "type": "object"
		//	}
		"sampling_rule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Attributes
				"attributes":        // Pattern: ""
				schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Matches attributes derived from the request.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: FixedRate
				"fixed_rate": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Float64{ /*START VALIDATORS*/
						float64validator.Between(0.000000, 1.000000),
						fwvalidators.NotNullFloat64(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: HTTPMethod
				"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the HTTP method from a request URL.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(10),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Host
				"host": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the hostname from a request URL.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(64),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Priority
				"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The priority of the sampling rule.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(1, 9999),
						fwvalidators.NotNullInt64(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ReservoirSize
				"reservoir_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(0),
						fwvalidators.NotNullInt64(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ResourceARN
				"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the ARN of the AWS resource on which the service runs.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(500),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RuleARN
				"rule_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RuleName
				"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 32),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ServiceName
				"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the name that the service uses to identify itself in segments.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(64),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ServiceType
				"service_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the origin that the service uses to identify its type in segments.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(64),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: URLPath
				"url_path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the path from a request URL.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(128),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The version of the sampling rule format (1)",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
						int64planmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SamplingRuleRecord
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CreatedAt": {
		//	      "description": "When the rule was created, in Unix time seconds.",
		//	      "type": "string"
		//	    },
		//	    "ModifiedAt": {
		//	      "description": "When the rule was modified, in Unix time seconds.",
		//	      "type": "string"
		//	    },
		//	    "SamplingRule": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Attributes": {
		//	          "additionalProperties": false,
		//	          "$comment": "String to string map",
		//	          "description": "Matches attributes derived from the request.",
		//	          "patternProperties": {
		//	            "": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "FixedRate": {
		//	          "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
		//	          "maximum": 1,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        },
		//	        "HTTPMethod": {
		//	          "description": "Matches the HTTP method from a request URL.",
		//	          "maxLength": 10,
		//	          "type": "string"
		//	        },
		//	        "Host": {
		//	          "description": "Matches the hostname from a request URL.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "Priority": {
		//	          "description": "The priority of the sampling rule.",
		//	          "maximum": 9999,
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        },
		//	        "ReservoirSize": {
		//	          "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
		//	          "minimum": 0,
		//	          "type": "integer"
		//	        },
		//	        "ResourceARN": {
		//	          "description": "Matches the ARN of the AWS resource on which the service runs.",
		//	          "maxLength": 500,
		//	          "type": "string"
		//	        },
		//	        "RuleARN": {
		//	          "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	          "type": "string"
		//	        },
		//	        "RuleName": {
		//	          "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	          "maxLength": 32,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "ServiceName": {
		//	          "description": "Matches the name that the service uses to identify itself in segments.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "ServiceType": {
		//	          "description": "Matches the origin that the service uses to identify its type in segments.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "URLPath": {
		//	          "description": "Matches the path from a request URL.",
		//	          "maxLength": 128,
		//	          "type": "string"
		//	        },
		//	        "Version": {
		//	          "description": "The version of the sampling rule format (1)",
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "FixedRate",
		//	        "Host",
		//	        "HTTPMethod",
		//	        "Priority",
		//	        "ReservoirSize",
		//	        "ResourceARN",
		//	        "ServiceName",
		//	        "ServiceType",
		//	        "URLPath"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sampling_rule_record": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CreatedAt
				"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "When the rule was created, in Unix time seconds.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ModifiedAt
				"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "When the rule was modified, in Unix time seconds.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SamplingRule
				"sampling_rule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Attributes
						"attributes":        // Pattern: ""
						schema.MapAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Matches attributes derived from the request.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
								mapplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: FixedRate
						"fixed_rate": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(0.000000, 1.000000),
								fwvalidators.NotNullFloat64(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
								float64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: HTTPMethod
						"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the HTTP method from a request URL.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(10),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Host
						"host": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the hostname from a request URL.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(64),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Priority
						"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The priority of the sampling rule.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(1, 9999),
								fwvalidators.NotNullInt64(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ReservoirSize
						"reservoir_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.AtLeast(0),
								fwvalidators.NotNullInt64(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ResourceARN
						"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the ARN of the AWS resource on which the service runs.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(500),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: RuleARN
						"rule_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: RuleName
						"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 32),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ServiceName
						"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the name that the service uses to identify itself in segments.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(64),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ServiceType
						"service_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the origin that the service uses to identify its type in segments.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(64),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: URLPath
						"url_path": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Matches the path from a request URL.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(128),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Version
						"version": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The version of the sampling rule format (1)",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.AtLeast(1),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
								int64planmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SamplingRuleUpdate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Attributes": {
		//	      "additionalProperties": false,
		//	      "$comment": "String to string map",
		//	      "description": "Matches attributes derived from the request.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "FixedRate": {
		//	      "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
		//	      "maximum": 1,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "HTTPMethod": {
		//	      "description": "Matches the HTTP method from a request URL.",
		//	      "maxLength": 10,
		//	      "type": "string"
		//	    },
		//	    "Host": {
		//	      "description": "Matches the hostname from a request URL.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "Priority": {
		//	      "description": "The priority of the sampling rule.",
		//	      "maximum": 9999,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "ReservoirSize": {
		//	      "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "ResourceARN": {
		//	      "description": "Matches the ARN of the AWS resource on which the service runs.",
		//	      "maxLength": 500,
		//	      "type": "string"
		//	    },
		//	    "RuleARN": {
		//	      "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	      "type": "string"
		//	    },
		//	    "RuleName": {
		//	      "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
		//	      "maxLength": 32,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "ServiceName": {
		//	      "description": "Matches the name that the service uses to identify itself in segments.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "ServiceType": {
		//	      "description": "Matches the origin that the service uses to identify its type in segments.",
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    },
		//	    "URLPath": {
		//	      "description": "Matches the path from a request URL.",
		//	      "maxLength": 128,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sampling_rule_update": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Attributes
				"attributes":        // Pattern: ""
				schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Matches attributes derived from the request.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: FixedRate
				"fixed_rate": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The percentage of matching requests to instrument, after the reservoir is exhausted.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Float64{ /*START VALIDATORS*/
						float64validator.Between(0.000000, 1.000000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: HTTPMethod
				"http_method": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the HTTP method from a request URL.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(10),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Host
				"host": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the hostname from a request URL.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(64),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Priority
				"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The priority of the sampling rule.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(1, 9999),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ReservoirSize
				"reservoir_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(0),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ResourceARN
				"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the ARN of the AWS resource on which the service runs.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(500),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RuleARN
				"rule_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RuleName
				"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 32),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ServiceName
				"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the name that the service uses to identify itself in segments.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(64),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ServiceType
				"service_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the origin that the service uses to identify its type in segments.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(64),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: URLPath
				"url_path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Matches the path from a request URL.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(128),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "This schema provides construct and validation rules for AWS-XRay SamplingRule resource parameters.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::XRay::SamplingRule").WithTerraformTypeName("awscc_xray_sampling_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attributes":           "Attributes",
		"created_at":           "CreatedAt",
		"fixed_rate":           "FixedRate",
		"host":                 "Host",
		"http_method":          "HTTPMethod",
		"key":                  "Key",
		"modified_at":          "ModifiedAt",
		"priority":             "Priority",
		"reservoir_size":       "ReservoirSize",
		"resource_arn":         "ResourceARN",
		"rule_arn":             "RuleARN",
		"rule_name":            "RuleName",
		"sampling_rule":        "SamplingRule",
		"sampling_rule_record": "SamplingRuleRecord",
		"sampling_rule_update": "SamplingRuleUpdate",
		"service_name":         "ServiceName",
		"service_type":         "ServiceType",
		"tags":                 "Tags",
		"url_path":             "URLPath",
		"value":                "Value",
		"version":              "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
