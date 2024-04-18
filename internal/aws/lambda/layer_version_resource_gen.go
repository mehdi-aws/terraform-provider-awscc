// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lambda_layer_version", layerVersionResource)
}

// layerVersionResource returns the Terraform awscc_lambda_layer_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lambda::LayerVersion resource.
func layerVersionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CompatibleArchitectures
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of compatible instruction set architectures.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"compatible_architectures": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of compatible instruction set architectures.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CompatibleRuntimes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of compatible function runtimes. Used for filtering with ListLayers and ListLayerVersions.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"compatible_runtimes": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of compatible function runtimes. Used for filtering with ListLayers and ListLayerVersions.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Content
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The function layer archive.",
		//	  "properties": {
		//	    "S3Bucket": {
		//	      "description": "The Amazon S3 bucket of the layer archive.",
		//	      "type": "string"
		//	    },
		//	    "S3Key": {
		//	      "description": "The Amazon S3 key of the layer archive.",
		//	      "type": "string"
		//	    },
		//	    "S3ObjectVersion": {
		//	      "description": "For versioned objects, the version of the layer archive object to use.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3Bucket",
		//	    "S3Key"
		//	  ],
		//	  "type": "object"
		//	}
		"content": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Bucket
				"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon S3 bucket of the layer archive.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: S3Key
				"s3_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon S3 key of the layer archive.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: S3ObjectVersion
				"s3_object_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "For versioned objects, the version of the layer archive object to use.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The function layer archive.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Content is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the version.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the version.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LayerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name or Amazon Resource Name (ARN) of the layer.",
		//	  "type": "string"
		//	}
		"layer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name or Amazon Resource Name (ARN) of the layer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LayerVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"layer_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LicenseInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The layer's software license.",
		//	  "type": "string"
		//	}
		"license_info": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The layer's software license.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Lambda::LayerVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::LayerVersion").WithTerraformTypeName("awscc_lambda_layer_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"compatible_architectures": "CompatibleArchitectures",
		"compatible_runtimes":      "CompatibleRuntimes",
		"content":                  "Content",
		"description":              "Description",
		"layer_name":               "LayerName",
		"layer_version_arn":        "LayerVersionArn",
		"license_info":             "LicenseInfo",
		"s3_bucket":                "S3Bucket",
		"s3_key":                   "S3Key",
		"s3_object_version":        "S3ObjectVersion",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Content",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}