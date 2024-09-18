// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package connect

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_connect_contact_flow", contactFlowResource)
}

// contactFlowResource returns the Terraform awscc_connect_contact_flow resource.
// This Terraform resource corresponds to the CloudFormation AWS::Connect::ContactFlow resource.
func contactFlowResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ContactFlowArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the contact flow (ARN).",
		//	  "maxLength": 500,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"contact_flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the contact flow (ARN).",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Content
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The content of the contact flow in JSON format.",
		//	  "maxLength": 256000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"content": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The content of the contact flow in JSON format.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256000),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the contact flow.",
		//	  "maxLength": 500,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the contact flow.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(500),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon Connect instance (ARN).",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon Connect instance (ARN).",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the contact flow.",
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the contact flow.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 127),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the contact flow.",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "ARCHIVED"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the contact flow.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ACTIVE",
					"ARCHIVED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the contact flow.",
		//	  "enum": [
		//	    "CONTACT_FLOW",
		//	    "CUSTOMER_QUEUE",
		//	    "CUSTOMER_HOLD",
		//	    "CUSTOMER_WHISPER",
		//	    "AGENT_HOLD",
		//	    "AGENT_WHISPER",
		//	    "OUTBOUND_WHISPER",
		//	    "AGENT_TRANSFER",
		//	    "QUEUE_TRANSFER"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the contact flow.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CONTACT_FLOW",
					"CUSTOMER_QUEUE",
					"CUSTOMER_HOLD",
					"CUSTOMER_WHISPER",
					"AGENT_HOLD",
					"AGENT_WHISPER",
					"OUTBOUND_WHISPER",
					"AGENT_TRANSFER",
					"QUEUE_TRANSFER",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::Connect::ContactFlow",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::ContactFlow").WithTerraformTypeName("awscc_connect_contact_flow")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"contact_flow_arn": "ContactFlowArn",
		"content":          "Content",
		"description":      "Description",
		"instance_arn":     "InstanceArn",
		"key":              "Key",
		"name":             "Name",
		"state":            "State",
		"tags":             "Tags",
		"type":             "Type",
		"value":            "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
