package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	General 				*[]General 				`json:"General"`
	ProductionCapacity 		*[]ProductionCapacity	`json:"ProductionCapacity"`
}

type General struct {
	WorkCenter                   	int     `json:"WorkCenter"`
	WorkCenterType               	string  `json:"WorkCenterType"`
	WorkCenterName               	string  `json:"WorkCenterName"`
	BusinessPartner              	int     `json:"BusinessPartner"`
	Plant                        	string  `json:"Plant"`
	WorkCenterCategory           	*string `json:"WorkCenterCategory"`
	WorkCenterResponsible        	*string `json:"WorkCenterResponsible"`
	SupplyArea                   	*string `json:"SupplyArea"`
	WorkCenterUsage              	*string `json:"WorkCenterUsage"`
	ComponentIsMarkedForBackflush	*bool   `json:"ComponentIsMarkedForBackflush"`
	WorkCenterLocation           	*string `json:"WorkCenterLocation"`
	CapacityCategory         		string  `json:"CapacityCategory"`
	CapacityQuantityUnit         	string  `json:"CapacityQuantityUnit"`
	CapacityQuantity         		float32 `json:"CapacityQuantity"`
	ValidityStartDate            	string  `json:"ValidityStartDate"`
	ValidityEndDate              	string  `json:"ValidityEndDate"`
	CreationDate            		string  `json:"CreationDate"`
	LastChangeDate              	string  `json:"LastChangeDate"`
	IsMarkedForDeletion          	*bool   `json:"IsMarkedForDeletion"`
}

type ProductionCapacity struct {
	WorkCenter                   	            int     `json:"WorkCenter"`
	WorkCenterProductionCapacityID              int     `json:"WorkCenterProductionCapacityID"`
	BusinessPartner              	            int     `json:"BusinessPartner"`
	Plant                        	            string  `json:"Plant"`
	Product                       	            string  `json:"Product"`
	CapacityFormula           	                string  `json:"CapacityFormula"`
	CalculatedCapacityQuantityInProductionUnit  float32 `json:"CalculatedCapacityQuantityInProductionUnit"`
	CreationDate            		            string  `json:"CreationDate"`
	LastChangeDate              	            string  `json:"LastChangeDate"`
	IsMarkedForDeletion          	            *bool   `json:"IsMarkedForDeletion"`
}
