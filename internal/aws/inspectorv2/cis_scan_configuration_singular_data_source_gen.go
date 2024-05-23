// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package inspectorv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_inspectorv2_cis_scan_configuration", cisScanConfigurationDataSource)
}

// cisScanConfigurationDataSource returns the Terraform awscc_inspectorv2_cis_scan_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::InspectorV2::CisScanConfiguration resource.
func cisScanConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "CIS Scan configuration unique identifier",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "CIS Scan configuration unique identifier",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScanName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the scan",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"scan_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the scan",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Schedule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Choose a Schedule cadence",
		//	  "properties": {
		//	    "Daily": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "StartTime": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "TimeOfDay": {
		//	              "pattern": "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$",
		//	              "type": "string"
		//	            },
		//	            "TimeZone": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "TimeOfDay",
		//	            "TimeZone"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "StartTime"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Monthly": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Day": {
		//	          "enum": [
		//	            "MON",
		//	            "TUE",
		//	            "WED",
		//	            "THU",
		//	            "FRI",
		//	            "SAT",
		//	            "SUN"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "StartTime": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "TimeOfDay": {
		//	              "pattern": "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$",
		//	              "type": "string"
		//	            },
		//	            "TimeZone": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "TimeOfDay",
		//	            "TimeZone"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "StartTime",
		//	        "Day"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "OneTime": {
		//	      "type": "object"
		//	    },
		//	    "Weekly": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Days": {
		//	          "items": {
		//	            "enum": [
		//	              "MON",
		//	              "TUE",
		//	              "WED",
		//	              "THU",
		//	              "FRI",
		//	              "SAT",
		//	              "SUN"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "maxItems": 7,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "StartTime": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "TimeOfDay": {
		//	              "pattern": "^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$",
		//	              "type": "string"
		//	            },
		//	            "TimeZone": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "TimeOfDay",
		//	            "TimeZone"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "StartTime",
		//	        "Days"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  }
		//	}
		"schedule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Daily
				"daily": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: StartTime
						"start_time": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: TimeOfDay
								"time_of_day": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: TimeZone
								"time_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Monthly
				"monthly": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Day
						"day": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: StartTime
						"start_time": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: TimeOfDay
								"time_of_day": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: TimeZone
								"time_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OneTime
				"one_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					CustomType: jsontypes.NormalizedType{},
					Computed:   true,
				}, /*END ATTRIBUTE*/
				// Property: Weekly
				"weekly": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Days
						"days": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: StartTime
						"start_time": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: TimeOfDay
								"time_of_day": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: TimeZone
								"time_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Choose a Schedule cadence",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityLevel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "LEVEL_1",
		//	    "LEVEL_2"
		//	  ],
		//	  "type": "string"
		//	}
		"security_level": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "pattern": "^.{1,255}$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Targets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AccountIds": {
		//	      "items": {
		//	        "pattern": "^\\d{12}|ALL_MEMBERS|SELF$",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 10000,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "TargetResourceTags": {
		//	      "additionalProperties": false,
		//	      "patternProperties": {
		//	        "": {
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "maxItems": 5,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      }
		//	    }
		//	  },
		//	  "required": [
		//	    "AccountIds"
		//	  ]
		//	}
		"targets": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccountIds
				"account_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TargetResourceTags
				"target_resource_tags": // Pattern: ""
				schema.MapAttribute{    /*START ATTRIBUTE*/
					ElementType: types.ListType{ElemType: types.StringType},
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::InspectorV2::CisScanConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::InspectorV2::CisScanConfiguration").WithTerraformTypeName("awscc_inspectorv2_cis_scan_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_ids":          "AccountIds",
		"arn":                  "Arn",
		"daily":                "Daily",
		"day":                  "Day",
		"days":                 "Days",
		"monthly":              "Monthly",
		"one_time":             "OneTime",
		"scan_name":            "ScanName",
		"schedule":             "Schedule",
		"security_level":       "SecurityLevel",
		"start_time":           "StartTime",
		"tags":                 "Tags",
		"target_resource_tags": "TargetResourceTags",
		"targets":              "Targets",
		"time_of_day":          "TimeOfDay",
		"time_zone":            "TimeZone",
		"weekly":               "Weekly",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}