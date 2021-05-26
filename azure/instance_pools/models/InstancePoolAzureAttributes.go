package models

import (
	clusterModels "github.com/polar-rams/databricks-sdk-golang/azure/clusters/models"
)

type InstancePoolAzureAttributes struct {
	Availability    clusterModels.AzureAvailability `json:"availability,omitempty" url:"availability,omitempty"`
	SpotBidMaxPrice float64                         `json:"spot_bid_max_price,omitempty" url:"spot_bid_max_price,omitempty"`
}
