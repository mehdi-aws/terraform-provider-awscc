// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package fis

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	registry.AddResourceFactory("awscc_fis_experiment_template", experimentTemplateResource)
}

// experimentTemplateResource returns the Terraform awscc_fis_experiment_template resource.
// This Terraform resource corresponds to the CloudFormation AWS::FIS::ExperimentTemplate resource.
func experimentTemplateResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Actions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The actions for the experiment.",
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies an action for the experiment template.",
		//	      "properties": {
		//	        "ActionId": {
		//	          "description": "The ID of the action.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "Description": {
		//	          "description": "A description for the action.",
		//	          "maxLength": 512,
		//	          "type": "string"
		//	        },
		//	        "Parameters": {
		//	          "additionalProperties": false,
		//	          "description": "The parameters for the action, if applicable.",
		//	          "patternProperties": {
		//	            "": {
		//	              "maxLength": 1024,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "StartAfter": {
		//	          "description": "The names of the actions that must be completed before the current action starts.",
		//	          "items": {
		//	            "maxLength": 64,
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "Targets": {
		//	          "additionalProperties": false,
		//	          "description": "One or more targets for the action.",
		//	          "patternProperties": {
		//	            "": {
		//	              "maxLength": 64,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "ActionId"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"actions":                 // Pattern: ""
		schema.MapNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ActionId
					"action_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the action.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(64),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A description for the action.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(512),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Parameters
					"parameters":        // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The parameters for the action, if applicable.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
							mapplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: StartAfter
					"start_after": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The names of the actions that must be completed before the current action starts.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.ValueStringsAre(
								stringvalidator.LengthAtMost(64),
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Targets
					"targets":           // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "One or more targets for the action.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
							mapplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The actions for the experiment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the experiment template.",
		//	  "maxLength": 512,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the experiment template.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(512),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ExperimentOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AccountTargeting": {
		//	      "description": "The account targeting setting for the experiment template.",
		//	      "enum": [
		//	        "multi-account",
		//	        "single-account"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "EmptyTargetResolutionMode": {
		//	      "description": "The target resolution failure mode for the experiment template.",
		//	      "enum": [
		//	        "fail",
		//	        "skip"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"experiment_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccountTargeting
				"account_targeting": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The account targeting setting for the experiment template.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"multi-account",
							"single-account",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EmptyTargetResolutionMode
				"empty_target_resolution_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The target resolution failure mode for the experiment template.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"fail",
							"skip",
						),
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
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"experiment_template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LogConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CloudWatchLogsConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "LogGroupArn": {
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "LogGroupArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "LogSchemaVersion": {
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "S3Configuration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "BucketName": {
		//	          "maxLength": 63,
		//	          "minLength": 3,
		//	          "type": "string"
		//	        },
		//	        "Prefix": {
		//	          "maxLength": 1024,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "BucketName"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "LogSchemaVersion"
		//	  ],
		//	  "type": "object"
		//	}
		"log_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CloudWatchLogsConfiguration
				"cloudwatch_logs_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LogGroupArn
						"log_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(20, 2048),
								fwvalidators.NotNullString(),
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
				// Property: LogSchemaVersion
				"log_schema_version": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(1),
						fwvalidators.NotNullInt64(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: S3Configuration
				"s3_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketName
						"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(3, 63),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Prefix
						"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 1024),
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
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.",
		//	  "maxLength": 1224,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1224),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: StopConditions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more stop conditions.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Source": {
		//	        "maxLength": 64,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 2048,
		//	        "minLength": 20,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Source"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"stop_conditions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Source
					"source": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(64),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(20, 2048),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more stop conditions.",
			Required:    true,
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
			Required:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Targets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The targets for the experiment.",
		//	  "patternProperties": {
		//	    "": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies a target for an experiment.",
		//	      "properties": {
		//	        "Filters": {
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "Describes a filter used for the target resource input in an experiment template.",
		//	            "properties": {
		//	              "Path": {
		//	                "description": "The attribute path for the filter.",
		//	                "maxLength": 256,
		//	                "type": "string"
		//	              },
		//	              "Values": {
		//	                "description": "The attribute values for the filter.",
		//	                "items": {
		//	                  "maxLength": 128,
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "required": [
		//	              "Path",
		//	              "Values"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "Parameters": {
		//	          "additionalProperties": false,
		//	          "patternProperties": {
		//	            "": {
		//	              "maxLength": 1024,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "ResourceArns": {
		//	          "description": "The Amazon Resource Names (ARNs) of the target resources.",
		//	          "items": {
		//	            "maxLength": 2048,
		//	            "minLength": 20,
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "ResourceTags": {
		//	          "additionalProperties": false,
		//	          "patternProperties": {
		//	            "": {
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "ResourceType": {
		//	          "description": "The AWS resource type. The resource type must be supported for the specified action.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        },
		//	        "SelectionMode": {
		//	          "description": "Scopes the identified resources to a specific number of the resources at random, or a percentage of the resources.",
		//	          "maxLength": 64,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ResourceType",
		//	        "SelectionMode"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"targets":                 // Pattern: ""
		schema.MapNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Filters
					"filters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Path
								"path": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The attribute path for the filter.",
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
								// Property: Values
								"values": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "The attribute values for the filter.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.ValueStringsAre(
											stringvalidator.LengthAtMost(128),
										),
										fwvalidators.NotNullList(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Parameters
					"parameters":        // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
							mapplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ResourceArns
					"resource_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "The Amazon Resource Names (ARNs) of the target resources.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.ValueStringsAre(
								stringvalidator.LengthBetween(20, 2048),
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ResourceTags
					"resource_tags":     // Pattern: ""
					schema.MapAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
							mapplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ResourceType
					"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The AWS resource type. The resource type must be supported for the specified action.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(64),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: SelectionMode
					"selection_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Scopes the identified resources to a specific number of the resources at random, or a percentage of the resources.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(64),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The targets for the experiment.",
			Required:    true,
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
		Description: "Resource schema for AWS::FIS::ExperimentTemplate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FIS::ExperimentTemplate").WithTerraformTypeName("awscc_fis_experiment_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_targeting":             "AccountTargeting",
		"action_id":                     "ActionId",
		"actions":                       "Actions",
		"bucket_name":                   "BucketName",
		"cloudwatch_logs_configuration": "CloudWatchLogsConfiguration",
		"description":                   "Description",
		"empty_target_resolution_mode":  "EmptyTargetResolutionMode",
		"experiment_options":            "ExperimentOptions",
		"experiment_template_id":        "Id",
		"filters":                       "Filters",
		"log_configuration":             "LogConfiguration",
		"log_group_arn":                 "LogGroupArn",
		"log_schema_version":            "LogSchemaVersion",
		"parameters":                    "Parameters",
		"path":                          "Path",
		"prefix":                        "Prefix",
		"resource_arns":                 "ResourceArns",
		"resource_tags":                 "ResourceTags",
		"resource_type":                 "ResourceType",
		"role_arn":                      "RoleArn",
		"s3_configuration":              "S3Configuration",
		"selection_mode":                "SelectionMode",
		"source":                        "Source",
		"start_after":                   "StartAfter",
		"stop_conditions":               "StopConditions",
		"tags":                          "Tags",
		"targets":                       "Targets",
		"value":                         "Value",
		"values":                        "Values",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
