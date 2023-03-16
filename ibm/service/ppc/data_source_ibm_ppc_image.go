// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ppc

import (
	"context"

	"github.com/IBM-Cloud/ppc-aas-go-client/helpers"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	//"fmt"
	"github.com/IBM-Cloud/ppc-aas-go-client/clients/instance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPPCImage() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPPCImagesRead,
		Schema: map[string]*schema.Schema{

			helpers.PPCImageName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Imagename Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			helpers.PPCCloudInstanceId: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operatingsystem": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_pool": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPPCImagesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPPCSession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(helpers.PPCCloudInstanceId).(string)

	imageC := instance.NewIBMPPCImageClient(ctx, sess, cloudInstanceID)
	imagedata, err := imageC.Get(d.Get(helpers.PPCImageName).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*imagedata.ImageID)
	d.Set("state", imagedata.State)
	d.Set("size", imagedata.Size)
	d.Set("architecture", imagedata.Specifications.Architecture)
	d.Set("hypervisor", imagedata.Specifications.HypervisorType)
	d.Set("operatingsystem", imagedata.Specifications.OperatingSystem)
	d.Set("storage_type", imagedata.StorageType)
	d.Set("storage_pool", imagedata.StoragePool)
	d.Set("image_type", imagedata.Specifications.ImageType)

	return nil

}
