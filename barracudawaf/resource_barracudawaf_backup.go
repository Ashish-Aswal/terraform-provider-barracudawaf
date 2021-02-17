package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFBackupCreate,
		Read:   resourceCudaWAFBackupRead,
		Update: resourceCudaWAFBackupUpdate,
		Delete: resourceCudaWAFBackupDelete,

		Schema: map[string]*schema.Schema{
			"source":                {Type: schema.TypeString, Optional: true},
			"backup_encryption_key": {Type: schema.TypeString, Optional: true},
			"exclude_management_interface_configurations": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"day_of_week":                          {Type: schema.TypeString, Required: true},
			"hour_of_day":                          {Type: schema.TypeString, Optional: true},
			"minute_of_hour":                       {Type: schema.TypeString, Optional: true},
			"backup_schedule_location":             {Type: schema.TypeString, Optional: true},
			"backups_to_keep":                      {Type: schema.TypeString, Optional: true},
			"azure_storage_account_name":           {Type: schema.TypeString, Optional: true},
			"azure_storage_container_name":         {Type: schema.TypeString, Optional: true},
			"azure_storage_blob_path":              {Type: schema.TypeString, Optional: true},
			"azure_storage_access_key":             {Type: schema.TypeString, Optional: true},
			"cloud_password":                       {Type: schema.TypeString, Optional: true},
			"cloud_username":                       {Type: schema.TypeString, Optional: true},
			"show_all_backups":                     {Type: schema.TypeString, Optional: true},
			"encrypt_backup":                       {Type: schema.TypeString, Optional: true},
			"file_content":                         {Type: schema.TypeString, Optional: true},
			"filename":                             {Type: schema.TypeString, Optional: true},
			"ftp_address":                          {Type: schema.TypeString, Optional: true},
			"ftp_password":                         {Type: schema.TypeString, Optional: true},
			"ftp_path":                             {Type: schema.TypeString, Optional: true},
			"ftp_port":                             {Type: schema.TypeString, Optional: true},
			"ftp_username":                         {Type: schema.TypeString, Optional: true},
			"ftps_address":                         {Type: schema.TypeString, Optional: true},
			"ftps_password":                        {Type: schema.TypeString, Optional: true},
			"ftps_path":                            {Type: schema.TypeString, Optional: true},
			"ftps_port":                            {Type: schema.TypeString, Optional: true},
			"ftps_username":                        {Type: schema.TypeString, Optional: true},
			"destination":                          {Type: schema.TypeString, Optional: true},
			"amazon_s3_bucket_name":                {Type: schema.TypeString, Optional: true},
			"amazon_s3_directory_path":             {Type: schema.TypeString, Optional: true},
			"server_type":                          {Type: schema.TypeString, Optional: true},
			"smb_address":                          {Type: schema.TypeString, Optional: true},
			"use_ntlm":                             {Type: schema.TypeString, Optional: true},
			"smb_password":                         {Type: schema.TypeString, Optional: true},
			"smb_path":                             {Type: schema.TypeString, Optional: true},
			"smb_port":                             {Type: schema.TypeString, Optional: true},
			"smb_username":                         {Type: schema.TypeString, Optional: true},
			"backup_restore_action":                {Type: schema.TypeString, Optional: true},
			"restore_amazon_s3_bucket_name":        {Type: schema.TypeString, Optional: true},
			"restore_amazon_s3_directory_path":     {Type: schema.TypeString, Optional: true},
			"restore_azure_storage_account_name":   {Type: schema.TypeString, Optional: true},
			"restore_azure_storage_container_name": {Type: schema.TypeString, Optional: true},
			"restore_azure_storage_blob_path":      {Type: schema.TypeString, Optional: true},
			"restore_azure_storage_access_key":     {Type: schema.TypeString, Optional: true},
			"restore_azure_cloud_type":             {Type: schema.TypeString, Optional: true},
			"use_default_restore_location":         {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadBackup(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"source":                d.Get("source").(string),
		"backup-encryption-key": d.Get("backup_encryption_key").(string),
		"exclude-management-interface-configurations": d.Get("exclude_management_interface_configurations").(string),
		"day-of-week":                          d.Get("day_of_week").(string),
		"hour-of-day":                          d.Get("hour_of_day").(string),
		"minute-of-hour":                       d.Get("minute_of_hour").(string),
		"backup-schedule-location":             d.Get("backup_schedule_location").(string),
		"backups-to-keep":                      d.Get("backups_to_keep").(string),
		"azure-storage-account-name":           d.Get("azure_storage_account_name").(string),
		"azure-storage-container-name":         d.Get("azure_storage_container_name").(string),
		"azure-storage-blob-path":              d.Get("azure_storage_blob_path").(string),
		"azure-storage-access-key":             d.Get("azure_storage_access_key").(string),
		"cloud-password":                       d.Get("cloud_password").(string),
		"cloud-username":                       d.Get("cloud_username").(string),
		"show-all-backups":                     d.Get("show_all_backups").(string),
		"encrypt-backup":                       d.Get("encrypt_backup").(string),
		"file-content":                         d.Get("file_content").(string),
		"filename":                             d.Get("filename").(string),
		"ftp-address":                          d.Get("ftp_address").(string),
		"ftp-password":                         d.Get("ftp_password").(string),
		"ftp-path":                             d.Get("ftp_path").(string),
		"ftp-port":                             d.Get("ftp_port").(string),
		"ftp-username":                         d.Get("ftp_username").(string),
		"ftps-address":                         d.Get("ftps_address").(string),
		"ftps-password":                        d.Get("ftps_password").(string),
		"ftps-path":                            d.Get("ftps_path").(string),
		"ftps-port":                            d.Get("ftps_port").(string),
		"ftps-username":                        d.Get("ftps_username").(string),
		"destination":                          d.Get("destination").(string),
		"amazon-s3-bucket-name":                d.Get("amazon_s3_bucket_name").(string),
		"amazon-s3-directory-path":             d.Get("amazon_s3_directory_path").(string),
		"server-type":                          d.Get("server_type").(string),
		"smb-address":                          d.Get("smb_address").(string),
		"use-ntlm":                             d.Get("use_ntlm").(string),
		"smb-password":                         d.Get("smb_password").(string),
		"smb-path":                             d.Get("smb_path").(string),
		"smb-port":                             d.Get("smb_port").(string),
		"smb-username":                         d.Get("smb_username").(string),
		"backup-restore-action":                d.Get("backup_restore_action").(string),
		"restore-amazon-s3-bucket-name":        d.Get("restore_amazon_s3_bucket_name").(string),
		"restore-amazon-s3-directory-path":     d.Get("restore_amazon_s3_directory_path").(string),
		"restore-azure-storage-account-name":   d.Get("restore_azure_storage_account_name").(string),
		"restore-azure-storage-container-name": d.Get("restore_azure_storage_container_name").(string),
		"restore-azure-storage-blob-path":      d.Get("restore_azure_storage_blob_path").(string),
		"restore-azure-storage-access-key":     d.Get("restore_azure_storage_access_key").(string),
		"restore-azure-cloud-type":             d.Get("restore_azure_cloud_type").(string),
		"use-default-restore-location":         d.Get("use_default_restore_location").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{}
		for item := range updatePayloadExceptions {
			delete(resourcePayload, updatePayloadExceptions[item])
		}
	}

	//sanitise the resource payload
	for key, val := range resourcePayload {
		if len(val) <= 0 {
			delete(resourcePayload, key)
		}
	}

	//resourceUpdateData : cudaWAF reource URI update data
	resourceUpdateData := map[string]interface{}{
		"endpoint":  resourceEndpoint,
		"payload":   resourcePayload,
		"operation": resourceOperation,
		"name":      d.Get("name").(string),
	}

	//updateCudaWAFResourceObject : update cudaWAF resource object
	resourceUpdateStatus, resourceUpdateResponseBody := updateCudaWAFResourceObject(
		resourceUpdateData,
	)

	if resourceUpdateStatus == 200 || resourceUpdateStatus == 201 {
		if resourceOperation != "DELETE" {
			d.SetId(resourceUpdateResponseBody["id"].(string))
		}
	} else {
		return fmt.Errorf("some error occurred : %v", resourceUpdateResponseBody["msg"])
	}

	return nil
}

func resourceCudaWAFBackupCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/backup"
	resourceCreateResponseError := makeRestAPIPayloadBackup(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFBackupRead(d, m)
}

func resourceCudaWAFBackupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFBackupUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/backup/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadBackup(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFBackupRead(d, m)
}

func resourceCudaWAFBackupDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/backup/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadBackup(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}