// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package b2bi

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_b2bi_capability", capabilityResource)
}

// capabilityResource returns the Terraform awscc_b2bi_capability resource.
// This Terraform resource corresponds to the CloudFormation AWS::B2BI::Capability resource.
func capabilityResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CapabilityArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"capability_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CapabilityId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"capability_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Configuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "Edi": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CapabilityDirection": {
		//	          "enum": [
		//	            "INBOUND",
		//	            "OUTBOUND"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "InputLocation": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "BucketName": {
		//	              "maxLength": 63,
		//	              "minLength": 3,
		//	              "type": "string"
		//	            },
		//	            "Key": {
		//	              "maxLength": 1024,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "OutputLocation": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "BucketName": {
		//	              "maxLength": 63,
		//	              "minLength": 3,
		//	              "type": "string"
		//	            },
		//	            "Key": {
		//	              "maxLength": 1024,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "TransformerId": {
		//	          "maxLength": 64,
		//	          "minLength": 1,
		//	          "pattern": "^[a-zA-Z0-9_-]+$",
		//	          "type": "string"
		//	        },
		//	        "Type": {
		//	          "properties": {
		//	            "X12Details": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "TransactionSet": {
		//	                  "enum": [
		//	                    "X12_110",
		//	                    "X12_180",
		//	                    "X12_204",
		//	                    "X12_210",
		//	                    "X12_211",
		//	                    "X12_214",
		//	                    "X12_215",
		//	                    "X12_259",
		//	                    "X12_260",
		//	                    "X12_266",
		//	                    "X12_269",
		//	                    "X12_270",
		//	                    "X12_271",
		//	                    "X12_274",
		//	                    "X12_275",
		//	                    "X12_276",
		//	                    "X12_277",
		//	                    "X12_278",
		//	                    "X12_310",
		//	                    "X12_315",
		//	                    "X12_322",
		//	                    "X12_404",
		//	                    "X12_410",
		//	                    "X12_417",
		//	                    "X12_421",
		//	                    "X12_426",
		//	                    "X12_810",
		//	                    "X12_820",
		//	                    "X12_824",
		//	                    "X12_830",
		//	                    "X12_832",
		//	                    "X12_834",
		//	                    "X12_835",
		//	                    "X12_837",
		//	                    "X12_844",
		//	                    "X12_846",
		//	                    "X12_849",
		//	                    "X12_850",
		//	                    "X12_852",
		//	                    "X12_855",
		//	                    "X12_856",
		//	                    "X12_860",
		//	                    "X12_861",
		//	                    "X12_864",
		//	                    "X12_865",
		//	                    "X12_869",
		//	                    "X12_870",
		//	                    "X12_940",
		//	                    "X12_945",
		//	                    "X12_990",
		//	                    "X12_997",
		//	                    "X12_999",
		//	                    "X12_270_X279",
		//	                    "X12_271_X279",
		//	                    "X12_275_X210",
		//	                    "X12_275_X211",
		//	                    "X12_276_X212",
		//	                    "X12_277_X212",
		//	                    "X12_277_X214",
		//	                    "X12_277_X364",
		//	                    "X12_278_X217",
		//	                    "X12_820_X218",
		//	                    "X12_820_X306",
		//	                    "X12_824_X186",
		//	                    "X12_834_X220",
		//	                    "X12_834_X307",
		//	                    "X12_834_X318",
		//	                    "X12_835_X221",
		//	                    "X12_837_X222",
		//	                    "X12_837_X223",
		//	                    "X12_837_X224",
		//	                    "X12_837_X291",
		//	                    "X12_837_X292",
		//	                    "X12_837_X298",
		//	                    "X12_999_X231"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Version": {
		//	                  "enum": [
		//	                    "VERSION_4010",
		//	                    "VERSION_4030",
		//	                    "VERSION_5010",
		//	                    "VERSION_5010_HIPAA"
		//	                  ],
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "InputLocation",
		//	        "OutputLocation",
		//	        "TransformerId",
		//	        "Type"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Edi
				"edi": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CapabilityDirection
						"capability_direction": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"INBOUND",
									"OUTBOUND",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: InputLocation
						"input_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BucketName
								"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(3, 63),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(0, 1024),
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
						// Property: OutputLocation
						"output_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BucketName
								"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(3, 63),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(0, 1024),
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
						// Property: TransformerId
						"transformer_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 64),
								stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Type
						"type": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: X12Details
								"x12_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: TransactionSet
										"transaction_set": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.OneOf(
													"X12_110",
													"X12_180",
													"X12_204",
													"X12_210",
													"X12_211",
													"X12_214",
													"X12_215",
													"X12_259",
													"X12_260",
													"X12_266",
													"X12_269",
													"X12_270",
													"X12_271",
													"X12_274",
													"X12_275",
													"X12_276",
													"X12_277",
													"X12_278",
													"X12_310",
													"X12_315",
													"X12_322",
													"X12_404",
													"X12_410",
													"X12_417",
													"X12_421",
													"X12_426",
													"X12_810",
													"X12_820",
													"X12_824",
													"X12_830",
													"X12_832",
													"X12_834",
													"X12_835",
													"X12_837",
													"X12_844",
													"X12_846",
													"X12_849",
													"X12_850",
													"X12_852",
													"X12_855",
													"X12_856",
													"X12_860",
													"X12_861",
													"X12_864",
													"X12_865",
													"X12_869",
													"X12_870",
													"X12_940",
													"X12_945",
													"X12_990",
													"X12_997",
													"X12_999",
													"X12_270_X279",
													"X12_271_X279",
													"X12_275_X210",
													"X12_275_X211",
													"X12_276_X212",
													"X12_277_X212",
													"X12_277_X214",
													"X12_277_X364",
													"X12_278_X217",
													"X12_820_X218",
													"X12_820_X306",
													"X12_824_X186",
													"X12_834_X220",
													"X12_834_X307",
													"X12_834_X318",
													"X12_835_X221",
													"X12_837_X222",
													"X12_837_X223",
													"X12_837_X224",
													"X12_837_X291",
													"X12_837_X292",
													"X12_837_X298",
													"X12_999_X231",
												),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Version
										"version": schema.StringAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.OneOf(
													"VERSION_4010",
													"VERSION_4030",
													"VERSION_5010",
													"VERSION_5010_HIPAA",
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
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstructionsDocuments
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "BucketName": {
		//	        "maxLength": 63,
		//	        "minLength": 3,
		//	        "type": "string"
		//	      },
		//	      "Key": {
		//	        "maxLength": 1024,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"instructions_documents": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: BucketName
					"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(3, 63),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 1024),
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
				listvalidator.SizeBetween(0, 5),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModifiedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 254,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 254),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
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
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
							stringvalidator.LengthBetween(0, 256),
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
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "edi"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"edi",
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
		Description: "Definition of AWS::B2BI::Capability Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::B2BI::Capability").WithTerraformTypeName("awscc_b2bi_capability")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket_name":            "BucketName",
		"capability_arn":         "CapabilityArn",
		"capability_direction":   "CapabilityDirection",
		"capability_id":          "CapabilityId",
		"configuration":          "Configuration",
		"created_at":             "CreatedAt",
		"edi":                    "Edi",
		"input_location":         "InputLocation",
		"instructions_documents": "InstructionsDocuments",
		"key":                    "Key",
		"modified_at":            "ModifiedAt",
		"name":                   "Name",
		"output_location":        "OutputLocation",
		"tags":                   "Tags",
		"transaction_set":        "TransactionSet",
		"transformer_id":         "TransformerId",
		"type":                   "Type",
		"value":                  "Value",
		"version":                "Version",
		"x12_details":            "X12Details",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
