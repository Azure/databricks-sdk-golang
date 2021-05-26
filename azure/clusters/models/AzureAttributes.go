package models

type AzureAttributes struct {
	FirstOnDemand   int32             `json:"first_on_demand,omitempty" url:"first_on_demand,omitempty"`
	Availability    AzureAvailability `json:"availability,omitempty" url:"availability,omitempty"`
	SpotBidMaxPrice float64           `json:"spot_bid_max_price,omitempty" url:"spot_bid_max_price,omitempty"`
}
