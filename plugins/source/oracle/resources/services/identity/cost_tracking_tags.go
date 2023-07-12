package identity

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func CostTrackingTags() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_cost_tracking_tags",
		Resolver:  fetchCostTrackingTags,
		Multiplex: client.TenancyMultiplex,
		Transform: client.TransformWithStruct(&identity.Tag{}),
		Columns:   schema.ColumnList{client.RegionColumn},
	}
}

func fetchCostTrackingTags(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	var page *string
	for {
		request := identity.ListCostTrackingTagsRequest{
			CompartmentId: common.String(cqClient.CompartmentOcid),
			Page:          page,
		}

		response, err := cqClient.OracleClients[cqClient.Region].IdentityIdentityClient.ListCostTrackingTags(ctx, request)

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
