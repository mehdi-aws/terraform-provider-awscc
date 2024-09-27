---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_events_connection Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Events::Connection
---

# awscc_events_connection (Data Source)

Data Source schema for AWS::Events::Connection



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The arn of the connection resource.
- `auth_parameters` (Attributes) (see [below for nested schema](#nestedatt--auth_parameters))
- `authorization_type` (String)
- `description` (String) Description of the connection.
- `name` (String) Name of the connection.
- `secret_arn` (String) The arn of the secrets manager secret created in the customer account.

<a id="nestedatt--auth_parameters"></a>
### Nested Schema for `auth_parameters`

Read-Only:

- `api_key_auth_parameters` (Attributes) (see [below for nested schema](#nestedatt--auth_parameters--api_key_auth_parameters))
- `basic_auth_parameters` (Attributes) (see [below for nested schema](#nestedatt--auth_parameters--basic_auth_parameters))
- `invocation_http_parameters` (Attributes) (see [below for nested schema](#nestedatt--auth_parameters--invocation_http_parameters))
- `o_auth_parameters` (Attributes) (see [below for nested schema](#nestedatt--auth_parameters--o_auth_parameters))

<a id="nestedatt--auth_parameters--api_key_auth_parameters"></a>
### Nested Schema for `auth_parameters.api_key_auth_parameters`

Read-Only:

- `api_key_name` (String)
- `api_key_value` (String)


<a id="nestedatt--auth_parameters--basic_auth_parameters"></a>
### Nested Schema for `auth_parameters.basic_auth_parameters`

Read-Only:

- `password` (String)
- `username` (String)


<a id="nestedatt--auth_parameters--invocation_http_parameters"></a>
### Nested Schema for `auth_parameters.invocation_http_parameters`

Read-Only:

- `body_parameters` (Attributes List) (see [below for nested schema](#nestedatt--auth_parameters--invocation_http_parameters--body_parameters))
- `header_parameters` (Attributes List) (see [below for nested schema](#nestedatt--auth_parameters--invocation_http_parameters--header_parameters))
- `query_string_parameters` (Attributes List) (see [below for nested schema](#nestedatt--auth_parameters--invocation_http_parameters--query_string_parameters))

<a id="nestedatt--auth_parameters--invocation_http_parameters--body_parameters"></a>
### Nested Schema for `auth_parameters.invocation_http_parameters.body_parameters`

Read-Only:

- `is_value_secret` (Boolean)
- `key` (String)
- `value` (String)


<a id="nestedatt--auth_parameters--invocation_http_parameters--header_parameters"></a>
### Nested Schema for `auth_parameters.invocation_http_parameters.header_parameters`

Read-Only:

- `is_value_secret` (Boolean)
- `key` (String)
- `value` (String)


<a id="nestedatt--auth_parameters--invocation_http_parameters--query_string_parameters"></a>
### Nested Schema for `auth_parameters.invocation_http_parameters.query_string_parameters`

Read-Only:

- `is_value_secret` (Boolean)
- `key` (String)
- `value` (String)



<a id="nestedatt--auth_parameters--o_auth_parameters"></a>
### Nested Schema for `auth_parameters.o_auth_parameters`

Read-Only:

- `authorization_endpoint` (String)
- `client_parameters` (Attributes) (see [below for nested schema](#nestedatt--auth_parameters--o_auth_parameters--client_parameters))
- `http_method` (String)
- `o_auth_http_parameters` (Attributes) (see [below for nested schema](#nestedatt--auth_parameters--o_auth_parameters--o_auth_http_parameters))

<a id="nestedatt--auth_parameters--o_auth_parameters--client_parameters"></a>
### Nested Schema for `auth_parameters.o_auth_parameters.client_parameters`

Read-Only:

- `client_id` (String)
- `client_secret` (String)


<a id="nestedatt--auth_parameters--o_auth_parameters--o_auth_http_parameters"></a>
### Nested Schema for `auth_parameters.o_auth_parameters.o_auth_http_parameters`

Read-Only:

- `body_parameters` (Attributes List) (see [below for nested schema](#nestedatt--auth_parameters--o_auth_parameters--o_auth_http_parameters--body_parameters))
- `header_parameters` (Attributes List) (see [below for nested schema](#nestedatt--auth_parameters--o_auth_parameters--o_auth_http_parameters--header_parameters))
- `query_string_parameters` (Attributes List) (see [below for nested schema](#nestedatt--auth_parameters--o_auth_parameters--o_auth_http_parameters--query_string_parameters))

<a id="nestedatt--auth_parameters--o_auth_parameters--o_auth_http_parameters--body_parameters"></a>
### Nested Schema for `auth_parameters.o_auth_parameters.o_auth_http_parameters.body_parameters`

Read-Only:

- `is_value_secret` (Boolean)
- `key` (String)
- `value` (String)


<a id="nestedatt--auth_parameters--o_auth_parameters--o_auth_http_parameters--header_parameters"></a>
### Nested Schema for `auth_parameters.o_auth_parameters.o_auth_http_parameters.header_parameters`

Read-Only:

- `is_value_secret` (Boolean)
- `key` (String)
- `value` (String)


<a id="nestedatt--auth_parameters--o_auth_parameters--o_auth_http_parameters--query_string_parameters"></a>
### Nested Schema for `auth_parameters.o_auth_parameters.o_auth_http_parameters.query_string_parameters`

Read-Only:

- `is_value_secret` (Boolean)
- `key` (String)
- `value` (String)