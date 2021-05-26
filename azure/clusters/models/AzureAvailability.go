// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type AzureAvailability string

const (
	AzureAvailabilitySpotAzure             = "SPOT_AZURE"
	AzureAvailabilityOnDemandAzure         = "ON_DEMAND_AZURE"
	AzureAvailabilitySpotWithFallbackAzure = "SPOT_WITH_FALLBACK_AZURE"
)
