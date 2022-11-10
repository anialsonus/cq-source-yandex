// Code generated by codegen; DO NOT EDIT.

package containerregistry

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func ScanResults() *schema.Table {
	return &schema.Table{
		Name:      "yandex_containerregistry_scan_results",
		Resolver:  fetchScanResults,
		Multiplex: client.MultiplexBy(client.Folders),
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
				Description: `Resource ID`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "folder_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveMultiplexedResourceID,
				Description: `Folder ID`,
			},
			{
				Name:     "image_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageId"),
			},
			{
				Name:     "scanned_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("ScannedAt"),
			},
			{
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "vulnerabilities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Vulnerabilities"),
			},
		},
	}
}