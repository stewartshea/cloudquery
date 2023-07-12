package database

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/database"
)

func ExternalContainerDatabases() *schema.Table {
	return &schema.Table{
		Name:      "oracle_database_external_container_databases",
		Resolver:  fetchExternalContainerDatabases,
		Multiplex: client.RegionCompartmentMultiplex,
		Transform: client.TransformWithStruct(&database.ExternalContainerDatabaseSummary{}),
		Columns:   schema.ColumnList{client.RegionColumn, client.CompartmentIDColumn},
	}
}

func fetchExternalContainerDatabases(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := database.ListExternalContainerDatabasesRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].DatabaseDatabaseClient.ListExternalContainerDatabases(ctx, request)

		if err != nil {
			return err
		}

		res <- response.Items

		if response.OpcNextPage == nil {
			break
		}

		page = response.OpcNextPage
	}

	return nil
}
