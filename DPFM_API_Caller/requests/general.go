package requests

type General struct {
	WorkCenter                   int     `json:"WorkCenter"`
	WorkCenterType               string  `json:"WorkCenterType"`
	WorkCenterName               string  `json:"WorkCenterName"`
	BusinessPartner              int     `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	WorkCenterCategory           *string `json:"WorkCenterCategory"`
	WorkCenterResponsible        *string `json:"WorkCenterResponsible"`
	SupplyArea                   *string `json:"SupplyArea"`
	WorkCenterUsage              *string `json:"WorkCenterUsage"`
	MatlCompIsMarkedForBackflush *bool   `json:"MatlCompIsMarkedForBackflush"`
	WorkCenterLocation           *string `json:"WorkCenterLocation"`
	CapacityInternalID           *string `json:"CapacityInternalID"`
	CapacityCategoryCode         *string `json:"CapacityCategoryCode"`
	ValidityStartDate            *string `json:"ValidityStartDate"`
	ValidityEndDate              *string `json:"ValidityEndDate"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
}
