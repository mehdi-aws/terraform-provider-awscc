// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package vpclattice

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_vpclattice_rule", ruleResource)
}

// ruleResource returns the Terraform awscc_vpclattice_rule resource.
// This Terraform resource corresponds to the CloudFormation AWS::VpcLattice::Rule resource.
func ruleResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Action
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "FixedResponse": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "StatusCode": {
		//	          "maximum": 599,
		//	          "minimum": 100,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "required": [
		//	        "StatusCode"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Forward": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "TargetGroups": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "TargetGroupIdentifier": {
		//	                "maxLength": 2048,
		//	                "minLength": 20,
		//	                "pattern": "^((tg-[0-9a-z]{17})|(arn:[a-z0-9\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:targetgroup/tg-[0-9a-z]{17}))$",
		//	                "type": "string"
		//	              },
		//	              "Weight": {
		//	                "maximum": 999,
		//	                "minimum": 1,
		//	                "type": "integer"
		//	              }
		//	            },
		//	            "required": [
		//	              "TargetGroupIdentifier"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "maxItems": 10,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "TargetGroups"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"action": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FixedResponse
				"fixed_response": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: StatusCode
						"status_code": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.Between(100, 599),
								fwvalidators.NotNullInt64(),
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
				// Property: Forward
				"forward": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: TargetGroups
						"target_groups": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: TargetGroupIdentifier
									"target_group_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(20, 2048),
											stringvalidator.RegexMatches(regexp.MustCompile("^((tg-[0-9a-z]{17})|(arn:[a-z0-9\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:targetgroup/tg-[0-9a-z]{17}))$"), ""),
											fwvalidators.NotNullString(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Weight
									"weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.Int64{ /*START VALIDATORS*/
											int64validator.Between(1, 999),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
											int64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Optional: true,
							Computed: true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 10),
								fwvalidators.NotNullList(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
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
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:service/svc-[0-9a-z]{17}/listener/listener-[0-9a-z]{17}/rule/((rule-[0-9a-z]{17})|(default))$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 22,
		//	  "minLength": 7,
		//	  "pattern": "^((rule-[0-9a-z]{17})|(default))$",
		//	  "type": "string"
		//	}
		"rule_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ListenerIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^((listener-[0-9a-z]{17})|(arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:service/svc-[0-9a-z]{17}/listener/listener-[0-9a-z]{17}))$",
		//	  "type": "string"
		//	}
		"listener_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^((listener-[0-9a-z]{17})|(arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:service/svc-[0-9a-z]{17}/listener/listener-[0-9a-z]{17}))$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// ListenerIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Match
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "HttpMatch": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "HeaderMatches": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "CaseSensitive": {
		//	                "default": false,
		//	                "type": "boolean"
		//	              },
		//	              "Match": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "Contains": {
		//	                    "maxLength": 128,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  },
		//	                  "Exact": {
		//	                    "maxLength": 128,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  },
		//	                  "Prefix": {
		//	                    "maxLength": 128,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "Name": {
		//	                "maxLength": 40,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Match",
		//	              "Name"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "maxItems": 5,
		//	          "type": "array"
		//	        },
		//	        "Method": {
		//	          "enum": [
		//	            "CONNECT",
		//	            "DELETE",
		//	            "GET",
		//	            "HEAD",
		//	            "OPTIONS",
		//	            "POST",
		//	            "PUT",
		//	            "TRACE"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "PathMatch": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "CaseSensitive": {
		//	              "default": false,
		//	              "type": "boolean"
		//	            },
		//	            "Match": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Exact": {
		//	                  "maxLength": 128,
		//	                  "minLength": 1,
		//	                  "pattern": "^\\/[a-zA-Z0-9@:%_+.~#?\u0026\\/=-]*$",
		//	                  "type": "string"
		//	                },
		//	                "Prefix": {
		//	                  "maxLength": 128,
		//	                  "minLength": 1,
		//	                  "pattern": "^\\/[a-zA-Z0-9@:%_+.~#?\u0026\\/=-]*$",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "required": [
		//	            "Match"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "HttpMatch"
		//	  ],
		//	  "type": "object"
		//	}
		"match": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: HttpMatch
				"http_match": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: HeaderMatches
						"header_matches": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: CaseSensitive
									"case_sensitive": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Default:  booldefault.StaticBool(false),
										PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
											boolplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Match
									"match": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Contains
											"contains": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(1, 128),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Exact
											"exact": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(1, 128),
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
													stringvalidator.LengthBetween(1, 128),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Optional: true,
										Computed: true,
										Validators: []validator.Object{ /*START VALIDATORS*/
											fwvalidators.NotNullObject(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
											objectplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Name
									"name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 40),
											fwvalidators.NotNullString(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Optional: true,
							Computed: true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeAtMost(5),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Method
						"method": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"CONNECT",
									"DELETE",
									"GET",
									"HEAD",
									"OPTIONS",
									"POST",
									"PUT",
									"TRACE",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: PathMatch
						"path_match": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CaseSensitive
								"case_sensitive": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Default:  booldefault.StaticBool(false),
									PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
										boolplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Match
								"match": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Exact
										"exact": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.LengthBetween(1, 128),
												stringvalidator.RegexMatches(regexp.MustCompile("^\\/[a-zA-Z0-9@:%_+.~#?&\\/=-]*$"), ""),
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
												stringvalidator.LengthBetween(1, 128),
												stringvalidator.RegexMatches(regexp.MustCompile("^\\/[a-zA-Z0-9@:%_+.~#?&\\/=-]*$"), ""),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Optional: true,
									Computed: true,
									Validators: []validator.Object{ /*START VALIDATORS*/
										fwvalidators.NotNullObject(),
									}, /*END VALIDATORS*/
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
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 63),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Priority
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 100,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 100),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^((svc-[0-9a-z]{17})|(arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:service/svc-[0-9a-z]{17}))$",
		//	  "type": "string"
		//	}
		"service_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^((svc-[0-9a-z]{17})|(arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:service/svc-[0-9a-z]{17}))$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// ServiceIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
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
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
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
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 50),
			}, /*END VALIDATORS*/
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
		Description: "Creates a listener rule. Each listener has a default rule for checking connection requests, but you can define additional rules. Each rule consists of a priority, one or more actions, and one or more conditions.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::VpcLattice::Rule").WithTerraformTypeName("awscc_vpclattice_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                  "Action",
		"arn":                     "Arn",
		"case_sensitive":          "CaseSensitive",
		"contains":                "Contains",
		"exact":                   "Exact",
		"fixed_response":          "FixedResponse",
		"forward":                 "Forward",
		"header_matches":          "HeaderMatches",
		"http_match":              "HttpMatch",
		"key":                     "Key",
		"listener_identifier":     "ListenerIdentifier",
		"match":                   "Match",
		"method":                  "Method",
		"name":                    "Name",
		"path_match":              "PathMatch",
		"prefix":                  "Prefix",
		"priority":                "Priority",
		"rule_id":                 "Id",
		"service_identifier":      "ServiceIdentifier",
		"status_code":             "StatusCode",
		"tags":                    "Tags",
		"target_group_identifier": "TargetGroupIdentifier",
		"target_groups":           "TargetGroups",
		"value":                   "Value",
		"weight":                  "Weight",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ListenerIdentifier",
		"/properties/ServiceIdentifier",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
