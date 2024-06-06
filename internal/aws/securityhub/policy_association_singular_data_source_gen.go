// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_securityhub_policy_association", policyAssociationDataSource)
}

// policyAssociationDataSource returns the Terraform awscc_securityhub_policy_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::SecurityHub::PolicyAssociation resource.
func policyAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociationIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier to indicates if the target has an association",
		//	  "type": "string"
		//	}
		"association_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier to indicates if the target has an association",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssociationStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current status of the association between the specified target and the configuration",
		//	  "enum": [
		//	    "SUCCESS",
		//	    "PENDING",
		//	    "FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"association_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current status of the association between the specified target and the configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssociationStatusMessage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An explanation for a FAILED value for AssociationStatus",
		//	  "type": "string"
		//	}
		"association_status_message": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An explanation for a FAILED value for AssociationStatus",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssociationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the association between the specified target and the configuration was directly applied by the Security Hub delegated administrator or inherited from a parent",
		//	  "enum": [
		//	    "APPLIED",
		//	    "INHERITED"
		//	  ],
		//	  "type": "string"
		//	}
		"association_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the association between the specified target and the configuration was directly applied by the Security Hub delegated administrator or inherited from a parent",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationPolicyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The universally unique identifier (UUID) of the configuration policy or a value of SELF_MANAGED_SECURITY_HUB for a self-managed configuration",
		//	  "pattern": "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$|^SELF_MANAGED_SECURITY_HUB$",
		//	  "type": "string"
		//	}
		"configuration_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The universally unique identifier (UUID) of the configuration policy or a value of SELF_MANAGED_SECURITY_HUB for a self-managed configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the target account, organizational unit, or the root",
		//	  "type": "string"
		//	}
		"target_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the target account, organizational unit, or the root",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the target is an AWS account, organizational unit, or the organization root",
		//	  "enum": [
		//	    "ACCOUNT",
		//	    "ORGANIZATIONAL_UNIT",
		//	    "ROOT"
		//	  ],
		//	  "type": "string"
		//	}
		"target_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the target is an AWS account, organizational unit, or the organization root",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time, in UTC and ISO 8601 format, that the configuration policy association was last updated",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time, in UTC and ISO 8601 format, that the configuration policy association was last updated",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SecurityHub::PolicyAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecurityHub::PolicyAssociation").WithTerraformTypeName("awscc_securityhub_policy_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_identifier":     "AssociationIdentifier",
		"association_status":         "AssociationStatus",
		"association_status_message": "AssociationStatusMessage",
		"association_type":           "AssociationType",
		"configuration_policy_id":    "ConfigurationPolicyId",
		"target_id":                  "TargetId",
		"target_type":                "TargetType",
		"updated_at":                 "UpdatedAt",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
