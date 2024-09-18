// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package resiliencehub

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	registry.AddResourceFactory("awscc_resiliencehub_app", appResource)
}

// appResource returns the Terraform awscc_resiliencehub_app resource.
// This Terraform resource corresponds to the CloudFormation AWS::ResilienceHub::App resource.
func appResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AppArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the App.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"app_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the App.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AppAssessmentSchedule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Assessment execution schedule.",
		//	  "enum": [
		//	    "Disabled",
		//	    "Daily"
		//	  ],
		//	  "type": "string"
		//	}
		"app_assessment_schedule": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Assessment execution schedule.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"Disabled",
					"Daily",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AppTemplateBody
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A string containing full ResilienceHub app template body.",
		//	  "maxLength": 409600,
		//	  "minLength": 0,
		//	  "pattern": "^[\\w\\s:,-\\.'\\/{}\\[\\]:\"]+$",
		//	  "type": "string"
		//	}
		"app_template_body": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A string containing full ResilienceHub app template body.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 409600),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\w\\s:,-\\.'\\/{}\\[\\]:\"]+$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "App description.",
		//	  "maxLength": 500,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "App description.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 500),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DriftStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates if compliance drifts (deviations) were detected while running an assessment for your application.",
		//	  "enum": [
		//	    "NotChecked",
		//	    "NotDetected",
		//	    "Detected"
		//	  ],
		//	  "type": "string"
		//	}
		"drift_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates if compliance drifts (deviations) were detected while running an assessment for your application.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventSubscriptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of events you would like to subscribe and get notification for.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Indicates an event you would like to subscribe and get notification for.",
		//	    "properties": {
		//	      "EventType": {
		//	        "description": "The type of event you would like to subscribe and get notification for.",
		//	        "enum": [
		//	          "ScheduledAssessmentFailure",
		//	          "DriftDetected"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "description": "Unique name to identify an event subscription.",
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      },
		//	      "SnsTopicArn": {
		//	        "description": "Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic.",
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "EventType"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"event_subscriptions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: EventType
					"event_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of event you would like to subscribe and get notification for.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"ScheduledAssessmentFailure",
								"DriftDetected",
							),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Unique name to identify an event subscription.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: SnsTopicArn
					"sns_topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of events you would like to subscribe and get notification for.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the app.",
		//	  "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the app.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PermissionModel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.",
		//	  "properties": {
		//	    "CrossAccountRoleArns": {
		//	      "description": "Defines a list of role Amazon Resource Names (ARNs) to be used in other accounts. These ARNs are used for querying purposes while importing resources and assessing your application.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "InvokerRoleName": {
		//	      "description": "Existing AWS IAM role name in the primary AWS account that will be assumed by AWS Resilience Hub Service Principle to obtain a read-only access to your application resources while running an assessment.",
		//	      "pattern": "",
		//	      "type": "string"
		//	    },
		//	    "Type": {
		//	      "description": "Defines how AWS Resilience Hub scans your resources. It can scan for the resources by using a pre-existing role in your AWS account, or by using the credentials of the current IAM user.",
		//	      "enum": [
		//	        "LegacyIAMUser",
		//	        "RoleBased"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"permission_model": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CrossAccountRoleArns
				"cross_account_role_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Defines a list of role Amazon Resource Names (ARNs) to be used in other accounts. These ARNs are used for querying purposes while importing resources and assessing your application.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: InvokerRoleName
				"invoker_role_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Existing AWS IAM role name in the primary AWS account that will be assumed by AWS Resilience Hub Service Principle to obtain a read-only access to your application resources while running an assessment.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Defines how AWS Resilience Hub scans your resources. It can scan for the resources by using a pre-existing role in your AWS account, or by using the credentials of the current IAM user.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"LegacyIAMUser",
							"RoleBased",
						),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResiliencyPolicyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the Resiliency Policy.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"resiliency_policy_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the Resiliency Policy.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceMappings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of ResourceMapping objects.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Resource mapping is used to map logical resources from template to physical resource",
		//	    "properties": {
		//	      "EksSourceName": {
		//	        "type": "string"
		//	      },
		//	      "LogicalStackName": {
		//	        "type": "string"
		//	      },
		//	      "MappingType": {
		//	        "pattern": "CfnStack|Resource|Terraform|EKS",
		//	        "type": "string"
		//	      },
		//	      "PhysicalResourceId": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "AwsAccountId": {
		//	            "pattern": "^[0-9]{12}$",
		//	            "type": "string"
		//	          },
		//	          "AwsRegion": {
		//	            "pattern": "^[a-z]{2}-((iso[a-z]{0,1}-)|(gov-)){0,1}[a-z]+-[0-9]$",
		//	            "type": "string"
		//	          },
		//	          "Identifier": {
		//	            "maxLength": 255,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Type": {
		//	            "pattern": "Arn|Native",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Identifier",
		//	          "Type"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "ResourceName": {
		//	        "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
		//	        "type": "string"
		//	      },
		//	      "TerraformSourceName": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "MappingType",
		//	      "PhysicalResourceId"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"resource_mappings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: EksSourceName
					"eks_source_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LogicalStackName
					"logical_stack_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: MappingType
					"mapping_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("CfnStack|Resource|Terraform|EKS"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: PhysicalResourceId
					"physical_resource_id": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AwsAccountId
							"aws_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{12}$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: AwsRegion
							"aws_region": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]{2}-((iso[a-z]{0,1}-)|(gov-)){0,1}[a-z]+-[0-9]$"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Identifier
							"identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 255),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("Arn|Native"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: ResourceName
					"resource_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: TerraformSourceName
					"terraform_source_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of ResourceMapping objects.",
			Required:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type Definition for AWS::ResilienceHub::App.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResilienceHub::App").WithTerraformTypeName("awscc_resiliencehub_app")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_arn":                 "AppArn",
		"app_assessment_schedule": "AppAssessmentSchedule",
		"app_template_body":       "AppTemplateBody",
		"aws_account_id":          "AwsAccountId",
		"aws_region":              "AwsRegion",
		"cross_account_role_arns": "CrossAccountRoleArns",
		"description":             "Description",
		"drift_status":            "DriftStatus",
		"eks_source_name":         "EksSourceName",
		"event_subscriptions":     "EventSubscriptions",
		"event_type":              "EventType",
		"identifier":              "Identifier",
		"invoker_role_name":       "InvokerRoleName",
		"logical_stack_name":      "LogicalStackName",
		"mapping_type":            "MappingType",
		"name":                    "Name",
		"permission_model":        "PermissionModel",
		"physical_resource_id":    "PhysicalResourceId",
		"resiliency_policy_arn":   "ResiliencyPolicyArn",
		"resource_mappings":       "ResourceMappings",
		"resource_name":           "ResourceName",
		"sns_topic_arn":           "SnsTopicArn",
		"tags":                    "Tags",
		"terraform_source_name":   "TerraformSourceName",
		"type":                    "Type",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
