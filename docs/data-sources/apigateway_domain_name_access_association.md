---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_apigateway_domain_name_access_association Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ApiGateway::DomainNameAccessAssociation
---

# awscc_apigateway_domain_name_access_association (Data Source)

Data Source schema for AWS::ApiGateway::DomainNameAccessAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `access_association_source` (String) The source of the domain name access association resource.
- `access_association_source_type` (String) The source type of the domain name access association resource.
- `domain_name_access_association_arn` (String) The amazon resource name (ARN) of the domain name access association resource.
- `domain_name_arn` (String) The amazon resource name (ARN) of the domain name resource.
- `tags` (Attributes List) An array of arbitrary tags (key-value pairs) to associate with the domainname access association. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)