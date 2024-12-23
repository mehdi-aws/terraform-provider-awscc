---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_s3tables_table_bucket Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::S3Tables::TableBucket
---

# awscc_s3tables_table_bucket (Data Source)

Data Source schema for AWS::S3Tables::TableBucket



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `table_bucket_arn` (String) The Amazon Resource Name (ARN) of the specified table bucket.
- `table_bucket_name` (String) A name for the table bucket.
- `unreferenced_file_removal` (Attributes) Settings governing the Unreferenced File Removal maintenance action. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots. (see [below for nested schema](#nestedatt--unreferenced_file_removal))

<a id="nestedatt--unreferenced_file_removal"></a>
### Nested Schema for `unreferenced_file_removal`

Read-Only:

- `noncurrent_days` (Number) S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.
- `status` (String) Indicates whether the Unreferenced File Removal maintenance action is enabled.
- `unreferenced_days` (Number) For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.