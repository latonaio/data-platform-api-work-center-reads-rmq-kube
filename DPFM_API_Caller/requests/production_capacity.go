package requests

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
