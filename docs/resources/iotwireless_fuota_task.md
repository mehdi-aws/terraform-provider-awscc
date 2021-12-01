---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iotwireless_fuota_task Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Create and manage FUOTA tasks.
---

# awscc_iotwireless_fuota_task (Resource)

Create and manage FUOTA tasks.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **firmware_update_image** (String) FUOTA task firmware update image binary S3 link
- **firmware_update_role** (String) FUOTA task firmware IAM role for reading S3
- **lo_ra_wan** (Attributes) FUOTA task LoRaWAN (see [below for nested schema](#nestedatt--lo_ra_wan))

### Optional

- **associate_multicast_group** (String) Multicast group to associate. Only for update request.
- **associate_wireless_device** (String) Wireless device to associate. Only for update request.
- **description** (String) FUOTA task description
- **disassociate_multicast_group** (String) Multicast group to disassociate. Only for update request.
- **disassociate_wireless_device** (String) Wireless device to disassociate. Only for update request.
- **name** (String) Name of FUOTA task
- **tags** (Attributes Set) A list of key-value pairs that contain metadata for the FUOTA task. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String) FUOTA task arn. Returned after successful create.
- **fuota_task_status** (String) FUOTA task status. Returned after successful read.
- **id** (String) FUOTA task id. Returned after successful create.

<a id="nestedatt--lo_ra_wan"></a>
### Nested Schema for `lo_ra_wan`

Required:

- **rf_region** (String) FUOTA task LoRaWAN RF region
- **start_time** (String) FUOTA task LoRaWAN start time


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iotwireless_fuota_task.example <resource ID>
```