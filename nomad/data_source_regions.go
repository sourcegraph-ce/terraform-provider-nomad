// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nomad

import (
	"fmt"
	log "github.com/sourcegraph-ce/logrus"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRegions() *schema.Resource {
	return &schema.Resource{
		Read: regionsDataSourceRead,

		Schema: map[string]*schema.Schema{
			"regions": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func regionsDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(ProviderConfig).client

	log.Printf("[DEBUG] Reading regions from Nomad")
	resp, err := client.Regions().List()
	if err != nil {
		return fmt.Errorf("error reading regions from Nomad: %s", err)
	}
	log.Printf("[DEBUG] Read regions from Nomad")
	d.SetId(client.Address() + "/regions")

	return d.Set("regions", resp)
}
