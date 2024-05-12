// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_networkmanager_vpc_attachment", vpcAttachmentResource)
}

// vpcAttachmentResource returns the Terraform awscc_networkmanager_vpc_attachment resource.
// This Terraform resource corresponds to the CloudFormation AWS::NetworkManager::VpcAttachment resource.
func vpcAttachmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the attachment.",
		//	  "type": "string"
		//	}
		"attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AttachmentPolicyRuleNumber
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The policy rule number associated with the attachment.",
		//	  "type": "integer"
		//	}
		"attachment_policy_rule_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The policy rule number associated with the attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AttachmentType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Attachment type.",
		//	  "type": "string"
		//	}
		"attachment_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Attachment type.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of a core network for the VPC attachment.",
		//	  "type": "string"
		//	}
		"core_network_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of a core network for the VPC attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a core network for the VPC attachment.",
		//	  "type": "string"
		//	}
		"core_network_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a core network for the VPC attachment.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Creation time of the attachment.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Creation time of the attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EdgeLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Region where the edge is located.",
		//	  "type": "string"
		//	}
		"edge_location": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Region where the edge is located.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Options
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Vpc options of the attachment.",
		//	  "properties": {
		//	    "ApplianceModeSupport": {
		//	      "default": false,
		//	      "description": "Indicates whether to enable ApplianceModeSupport Support for Vpc Attachment. Valid Values: true | false",
		//	      "type": "boolean"
		//	    },
		//	    "Ipv6Support": {
		//	      "default": false,
		//	      "description": "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ApplianceModeSupport
				"appliance_mode_support": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether to enable ApplianceModeSupport Support for Vpc Attachment. Valid Values: true | false",
					Optional:    true,
					Computed:    true,
					Default:     booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Ipv6Support
				"ipv_6_support": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
					Optional:    true,
					Computed:    true,
					Default:     booldefault.StaticBool(false),
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Vpc options of the attachment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OwnerAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Owner account of the attachment.",
		//	  "type": "string"
		//	}
		"owner_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Owner account of the attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProposedSegmentChange
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The attachment to move from one segment to another.",
		//	  "properties": {
		//	    "AttachmentPolicyRuleNumber": {
		//	      "description": "The rule number in the policy document that applies to this change.",
		//	      "type": "integer"
		//	    },
		//	    "SegmentName": {
		//	      "description": "The name of the segment to change.",
		//	      "type": "string"
		//	    },
		//	    "Tags": {
		//	      "description": "The key-value tags that changed for the segment.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A key-value pair to associate with a resource.",
		//	        "properties": {
		//	          "Key": {
		//	            "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"proposed_segment_change": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AttachmentPolicyRuleNumber
				"attachment_policy_rule_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The rule number in the policy document that applies to this change.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SegmentName
				"segment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the segment to change.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Tags
				"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The key-value tags that changed for the segment.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The attachment to move from one segment to another.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Resource.",
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Resource.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SegmentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the segment attachment..",
		//	  "type": "string"
		//	}
		"segment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the segment attachment..",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "State of the attachment.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "State of the attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Subnet Arn list",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"subnet_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Subnet Arn list",
			Required:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags for the attachment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags for the attachment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Last update time of the attachment.",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Last update time of the attachment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the VPC.",
		//	  "type": "string"
		//	}
		"vpc_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the VPC.",
			Required:    true,
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
		Description: "AWS::NetworkManager::VpcAttachment Resoruce Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::VpcAttachment").WithTerraformTypeName("awscc_networkmanager_vpc_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"appliance_mode_support":        "ApplianceModeSupport",
		"attachment_id":                 "AttachmentId",
		"attachment_policy_rule_number": "AttachmentPolicyRuleNumber",
		"attachment_type":               "AttachmentType",
		"core_network_arn":              "CoreNetworkArn",
		"core_network_id":               "CoreNetworkId",
		"created_at":                    "CreatedAt",
		"edge_location":                 "EdgeLocation",
		"ipv_6_support":                 "Ipv6Support",
		"key":                           "Key",
		"options":                       "Options",
		"owner_account_id":              "OwnerAccountId",
		"proposed_segment_change":       "ProposedSegmentChange",
		"resource_arn":                  "ResourceArn",
		"segment_name":                  "SegmentName",
		"state":                         "State",
		"subnet_arns":                   "SubnetArns",
		"tags":                          "Tags",
		"updated_at":                    "UpdatedAt",
		"value":                         "Value",
		"vpc_arn":                       "VpcArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(60).WithDeleteTimeoutInMinutes(60)

	opts = opts.WithUpdateTimeoutInMinutes(60)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
