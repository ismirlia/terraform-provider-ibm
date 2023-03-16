// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ppc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPPCInstanceIPDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPPCInstanceIPDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_ppc_instance_ip.testacc_ds_instance_ip", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPPCInstanceIPDataSourceConfig() string {
	return fmt.Sprintf(`
	data "ibm_ppc_instance_ip" "testacc_ds_instance_ip" {
		ppc_network_name = "%[1]s"
		ppc_instance_name = "%[2]s"
		ppc_cloud_instance_id = "%[3]s"
	}
	`, acc.Ppc_network_name, acc.Ppc_instance_name, acc.Ppc_cloud_instance_id)
}
