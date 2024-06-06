// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package eks

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_eks_addon", addonResource)
}

// addonResource returns the Terraform awscc_eks_addon resource.
// This Terraform resource corresponds to the CloudFormation AWS::EKS::Addon resource.
func addonResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AddonName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of Addon",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"addon_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of Addon",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AddonVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Version of Addon",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"addon_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Version of Addon",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the add-on",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the add-on",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClusterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of Cluster",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"cluster_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of Cluster",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationValues
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The configuration values to use with the add-on",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"configuration_values": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The configuration values to use with the add-on",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PodIdentityAssociations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of pod identities to apply to this add-on.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A pod identity to associate with an add-on.",
		//	    "properties": {
		//	      "RoleArn": {
		//	        "description": "The IAM role ARN that the pod identity association is created for.",
		//	        "pattern": "^arn:aws(-cn|-us-gov|-iso(-[a-z])?)?:iam::\\d{12}:(role)\\/*",
		//	        "type": "string"
		//	      },
		//	      "ServiceAccount": {
		//	        "description": "The Kubernetes service account that the pod identity association is created for.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ServiceAccount",
		//	      "RoleArn"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"pod_identity_associations": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: RoleArn
					"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The IAM role ARN that the pod identity association is created for.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws(-cn|-us-gov|-iso(-[a-z])?)?:iam::\\d{12}:(role)\\/*"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: ServiceAccount
					"service_account": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Kubernetes service account that the pod identity association is created for.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of pod identities to apply to this add-on.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// PodIdentityAssociations is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: PreserveOnDelete
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "PreserveOnDelete parameter value",
		//	  "type": "boolean"
		//	}
		"preserve_on_delete": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "PreserveOnDelete parameter value",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// PreserveOnDelete is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ResolveConflicts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Resolve parameter value conflicts",
		//	  "enum": [
		//	    "NONE",
		//	    "OVERWRITE",
		//	    "PRESERVE"
		//	  ],
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"resolve_conflicts": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Resolve parameter value conflicts",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtLeast(1),
				stringvalidator.OneOf(
					"NONE",
					"OVERWRITE",
					"PRESERVE",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// ResolveConflicts is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ServiceAccountRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "IAM role to bind to the add-on's service account",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"service_account_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "IAM role to bind to the add-on's service account",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtLeast(1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 255,
		//	        "minLength": 1,
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
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 127),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 255),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Schema for AWS::EKS::Addon",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EKS::Addon").WithTerraformTypeName("awscc_eks_addon")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"addon_name":                "AddonName",
		"addon_version":             "AddonVersion",
		"arn":                       "Arn",
		"cluster_name":              "ClusterName",
		"configuration_values":      "ConfigurationValues",
		"key":                       "Key",
		"pod_identity_associations": "PodIdentityAssociations",
		"preserve_on_delete":        "PreserveOnDelete",
		"resolve_conflicts":         "ResolveConflicts",
		"role_arn":                  "RoleArn",
		"service_account":           "ServiceAccount",
		"service_account_role_arn":  "ServiceAccountRoleArn",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ResolveConflicts",
		"/properties/PreserveOnDelete",
		"/properties/PodIdentityAssociations",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
