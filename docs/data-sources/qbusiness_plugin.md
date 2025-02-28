---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_qbusiness_plugin Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::QBusiness::Plugin
---

# awscc_qbusiness_plugin (Data Source)

Data Source schema for AWS::QBusiness::Plugin



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `application_id` (String)
- `auth_configuration` (Attributes) (see [below for nested schema](#nestedatt--auth_configuration))
- `build_status` (String)
- `created_at` (String)
- `custom_plugin_configuration` (Attributes) (see [below for nested schema](#nestedatt--custom_plugin_configuration))
- `display_name` (String)
- `plugin_arn` (String)
- `plugin_id` (String)
- `server_url` (String)
- `state` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `type` (String)
- `updated_at` (String)

<a id="nestedatt--auth_configuration"></a>
### Nested Schema for `auth_configuration`

Read-Only:

- `basic_auth_configuration` (Attributes) (see [below for nested schema](#nestedatt--auth_configuration--basic_auth_configuration))
- `no_auth_configuration` (String)
- `o_auth_2_client_credential_configuration` (Attributes) (see [below for nested schema](#nestedatt--auth_configuration--o_auth_2_client_credential_configuration))

<a id="nestedatt--auth_configuration--basic_auth_configuration"></a>
### Nested Schema for `auth_configuration.basic_auth_configuration`

Read-Only:

- `role_arn` (String)
- `secret_arn` (String)


<a id="nestedatt--auth_configuration--o_auth_2_client_credential_configuration"></a>
### Nested Schema for `auth_configuration.o_auth_2_client_credential_configuration`

Read-Only:

- `authorization_url` (String)
- `role_arn` (String)
- `secret_arn` (String)
- `token_url` (String)



<a id="nestedatt--custom_plugin_configuration"></a>
### Nested Schema for `custom_plugin_configuration`

Read-Only:

- `api_schema` (Attributes) (see [below for nested schema](#nestedatt--custom_plugin_configuration--api_schema))
- `api_schema_type` (String)
- `description` (String)

<a id="nestedatt--custom_plugin_configuration--api_schema"></a>
### Nested Schema for `custom_plugin_configuration.api_schema`

Read-Only:

- `payload` (String)
- `s3` (Attributes) (see [below for nested schema](#nestedatt--custom_plugin_configuration--api_schema--s3))

<a id="nestedatt--custom_plugin_configuration--api_schema--s3"></a>
### Nested Schema for `custom_plugin_configuration.api_schema.s3`

Read-Only:

- `bucket` (String)
- `key` (String)




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
