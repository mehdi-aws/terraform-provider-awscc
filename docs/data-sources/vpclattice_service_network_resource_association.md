---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_vpclattice_service_network_resource_association Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::VpcLattice::ServiceNetworkResourceAssociation
---

# awscc_vpclattice_service_network_resource_association (Data Source)

Data Source schema for AWS::VpcLattice::ServiceNetworkResourceAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `resource_configuration_id` (String)
- `service_network_id` (String)
- `service_network_resource_association_id` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)