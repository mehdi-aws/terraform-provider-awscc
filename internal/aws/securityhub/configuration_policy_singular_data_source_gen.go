// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_securityhub_configuration_policy", configurationPolicyDataSource)
}

// configurationPolicyDataSource returns the Terraform awscc_securityhub_configuration_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::SecurityHub::ConfigurationPolicy resource.
func configurationPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the configuration policy.",
		//	  "pattern": "^arn:aws\\S*:securityhub:[a-z0-9-]+:[0-9]{12}:configuration-policy/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the configuration policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that defines how Security Hub is configured.",
		//	  "properties": {
		//	    "SecurityHub": {
		//	      "additionalProperties": false,
		//	      "description": "An object that defines how AWS Security Hub is configured.",
		//	      "properties": {
		//	        "EnabledStandardIdentifiers": {
		//	          "description": "A list that defines which security standards are enabled in the configuration policy.",
		//	          "insertionOrder": true,
		//	          "items": {
		//	            "maxLength": 2048,
		//	            "type": "string"
		//	          },
		//	          "maxItems": 1000,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "SecurityControlsConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "An object that defines which security controls are enabled in an AWS Security Hub configuration policy.",
		//	          "properties": {
		//	            "DisabledSecurityControlIdentifiers": {
		//	              "description": "A list of security controls that are disabled in the configuration policy",
		//	              "insertionOrder": true,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "type": "string"
		//	              },
		//	              "maxItems": 1000,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "EnabledSecurityControlIdentifiers": {
		//	              "description": "A list of security controls that are enabled in the configuration policy.",
		//	              "insertionOrder": true,
		//	              "items": {
		//	                "maxLength": 2048,
		//	                "type": "string"
		//	              },
		//	              "maxItems": 1000,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            },
		//	            "SecurityControlCustomParameters": {
		//	              "description": "A list of security controls and control parameter values that are included in a configuration policy.",
		//	              "insertionOrder": true,
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "description": "An object of security control and control parameter value that are included in a configuration policy.",
		//	                "properties": {
		//	                  "Parameters": {
		//	                    "additionalProperties": false,
		//	                    "description": "An object that specifies parameter values for a control in a configuration policy.",
		//	                    "patternProperties": {
		//	                      "": {
		//	                        "additionalProperties": false,
		//	                        "description": "An object that provides the current value of a security control parameter and identifies whether it has been customized.",
		//	                        "properties": {
		//	                          "Value": {
		//	                            "additionalProperties": false,
		//	                            "description": "An object that includes the data type of a security control parameter and its current value.",
		//	                            "properties": {
		//	                              "Boolean": {
		//	                                "description": "A control parameter that is a boolean.",
		//	                                "type": "boolean"
		//	                              },
		//	                              "Double": {
		//	                                "description": "A control parameter that is a double.",
		//	                                "type": "number"
		//	                              },
		//	                              "Enum": {
		//	                                "description": "A control parameter that is an enum.",
		//	                                "maxLength": 2048,
		//	                                "type": "string"
		//	                              },
		//	                              "EnumList": {
		//	                                "description": "A control parameter that is a list of enums.",
		//	                                "insertionOrder": true,
		//	                                "items": {
		//	                                  "maxLength": 2048,
		//	                                  "type": "string"
		//	                                },
		//	                                "maxItems": 100,
		//	                                "type": "array",
		//	                                "uniqueItems": true
		//	                              },
		//	                              "Integer": {
		//	                                "description": "A control parameter that is an integer.",
		//	                                "type": "integer"
		//	                              },
		//	                              "IntegerList": {
		//	                                "description": "A control parameter that is a list of integers.",
		//	                                "insertionOrder": true,
		//	                                "items": {
		//	                                  "type": "integer"
		//	                                },
		//	                                "maxItems": 100,
		//	                                "type": "array",
		//	                                "uniqueItems": true
		//	                              },
		//	                              "String": {
		//	                                "description": "A control parameter that is a string.",
		//	                                "maxLength": 2048,
		//	                                "type": "string"
		//	                              },
		//	                              "StringList": {
		//	                                "description": "A control parameter that is a list of strings.",
		//	                                "insertionOrder": true,
		//	                                "items": {
		//	                                  "maxLength": 2048,
		//	                                  "type": "string"
		//	                                },
		//	                                "maxItems": 100,
		//	                                "type": "array",
		//	                                "uniqueItems": true
		//	                              }
		//	                            },
		//	                            "type": "object"
		//	                          },
		//	                          "ValueType": {
		//	                            "description": "Identifies whether a control parameter uses a custom user-defined value or subscribes to the default AWS Security Hub behavior.",
		//	                            "enum": [
		//	                              "DEFAULT",
		//	                              "CUSTOM"
		//	                            ],
		//	                            "type": "string"
		//	                          }
		//	                        },
		//	                        "required": [
		//	                          "ValueType"
		//	                        ],
		//	                        "type": "object"
		//	                      }
		//	                    },
		//	                    "type": "object"
		//	                  },
		//	                  "SecurityControlId": {
		//	                    "description": "The ID of the security control.",
		//	                    "maxLength": 2048,
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "maxItems": 1000,
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "ServiceEnabled": {
		//	          "description": "Indicates whether Security Hub is enabled in the policy.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"configuration_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityHub
				"security_hub": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EnabledStandardIdentifiers
						"enabled_standard_identifiers": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "A list that defines which security standards are enabled in the configuration policy.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: SecurityControlsConfiguration
						"security_controls_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DisabledSecurityControlIdentifiers
								"disabled_security_control_identifiers": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "A list of security controls that are disabled in the configuration policy",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: EnabledSecurityControlIdentifiers
								"enabled_security_control_identifiers": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "A list of security controls that are enabled in the configuration policy.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: SecurityControlCustomParameters
								"security_control_custom_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Parameters
											"parameters":              // Pattern: ""
											schema.MapNestedAttribute{ /*START ATTRIBUTE*/
												NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
													Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
														// Property: Value
														"value": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
															Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
																// Property: Boolean
																"boolean": schema.BoolAttribute{ /*START ATTRIBUTE*/
																	Description: "A control parameter that is a boolean.",
																	Computed:    true,
																}, /*END ATTRIBUTE*/
																// Property: Double
																"double": schema.Float64Attribute{ /*START ATTRIBUTE*/
																	Description: "A control parameter that is a double.",
																	Computed:    true,
																}, /*END ATTRIBUTE*/
																// Property: Enum
																"enum": schema.StringAttribute{ /*START ATTRIBUTE*/
																	Description: "A control parameter that is an enum.",
																	Computed:    true,
																}, /*END ATTRIBUTE*/
																// Property: EnumList
																"enum_list": schema.ListAttribute{ /*START ATTRIBUTE*/
																	ElementType: types.StringType,
																	Description: "A control parameter that is a list of enums.",
																	Computed:    true,
																}, /*END ATTRIBUTE*/
																// Property: Integer
																"integer": schema.Int64Attribute{ /*START ATTRIBUTE*/
																	Description: "A control parameter that is an integer.",
																	Computed:    true,
																}, /*END ATTRIBUTE*/
																// Property: IntegerList
																"integer_list": schema.ListAttribute{ /*START ATTRIBUTE*/
																	ElementType: types.Int64Type,
																	Description: "A control parameter that is a list of integers.",
																	Computed:    true,
																}, /*END ATTRIBUTE*/
																// Property: String
																"string": schema.StringAttribute{ /*START ATTRIBUTE*/
																	Description: "A control parameter that is a string.",
																	Computed:    true,
																}, /*END ATTRIBUTE*/
																// Property: StringList
																"string_list": schema.ListAttribute{ /*START ATTRIBUTE*/
																	ElementType: types.StringType,
																	Description: "A control parameter that is a list of strings.",
																	Computed:    true,
																}, /*END ATTRIBUTE*/
															}, /*END SCHEMA*/
															Description: "An object that includes the data type of a security control parameter and its current value.",
															Computed:    true,
														}, /*END ATTRIBUTE*/
														// Property: ValueType
														"value_type": schema.StringAttribute{ /*START ATTRIBUTE*/
															Description: "Identifies whether a control parameter uses a custom user-defined value or subscribes to the default AWS Security Hub behavior.",
															Computed:    true,
														}, /*END ATTRIBUTE*/
													}, /*END SCHEMA*/
												}, /*END NESTED OBJECT*/
												Description: "An object that specifies parameter values for a control in a configuration policy.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: SecurityControlId
											"security_control_id": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The ID of the security control.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Description: "A list of security controls and control parameter values that are included in a configuration policy.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "An object that defines which security controls are enabled in an AWS Security Hub configuration policy.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ServiceEnabled
						"service_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Indicates whether Security Hub is enabled in the policy.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "An object that defines how AWS Security Hub is configured.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that defines how Security Hub is configured.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time, in UTC and ISO 8601 format.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time, in UTC and ISO 8601 format.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the configuration policy.",
		//	  "maxLength": 512,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the configuration policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The universally unique identifier (UUID) of the configuration policy.",
		//	  "pattern": "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$",
		//	  "type": "string"
		//	}
		"configuration_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The universally unique identifier (UUID) of the configuration policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the configuration policy.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the configuration policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the service that the configuration policy applies to is enabled in the policy.",
		//	  "type": "boolean"
		//	}
		"service_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the service that the configuration policy applies to is enabled in the policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair to associate with a resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value pair to associate with a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time, in UTC and ISO 8601 format.",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time, in UTC and ISO 8601 format.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SecurityHub::ConfigurationPolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecurityHub::ConfigurationPolicy").WithTerraformTypeName("awscc_securityhub_configuration_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                   "Arn",
		"boolean":                               "Boolean",
		"configuration_policy":                  "ConfigurationPolicy",
		"configuration_policy_id":               "Id",
		"created_at":                            "CreatedAt",
		"description":                           "Description",
		"disabled_security_control_identifiers": "DisabledSecurityControlIdentifiers",
		"double":                                "Double",
		"enabled_security_control_identifiers":  "EnabledSecurityControlIdentifiers",
		"enabled_standard_identifiers":          "EnabledStandardIdentifiers",
		"enum":                                  "Enum",
		"enum_list":                             "EnumList",
		"integer":                               "Integer",
		"integer_list":                          "IntegerList",
		"name":                                  "Name",
		"parameters":                            "Parameters",
		"security_control_custom_parameters":    "SecurityControlCustomParameters",
		"security_control_id":                   "SecurityControlId",
		"security_controls_configuration":       "SecurityControlsConfiguration",
		"security_hub":                          "SecurityHub",
		"service_enabled":                       "ServiceEnabled",
		"string":                                "String",
		"string_list":                           "StringList",
		"tags":                                  "Tags",
		"updated_at":                            "UpdatedAt",
		"value":                                 "Value",
		"value_type":                            "ValueType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
