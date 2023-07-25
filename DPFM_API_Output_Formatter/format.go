package dpfm_api_output_formatter

import (
	"data-platform-api-work-center-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	general := make([]General, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.General{}

		err := rows.Scan(
			&pm.WorkCenter,
			&pm.WorkCenterType,
			&pm.WorkCenterName,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.WorkCenterCategory,
			&pm.WorkCenterResponsible,
			&pm.SupplyArea,
			&pm.WorkCenterUsage,
			&pm.ComponentIsMarkedForBackflush,
			&pm.WorkCenterLocation,
			&pm.CapacityCategory,
			&pm.CapacityQuantityUnit,
			&pm.CapacityQuantity,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		general = append(general, General{
			WorkCenter:                   	data.WorkCenter,
			WorkCenterType:               	data.WorkCenterType,
			WorkCenterName:               	data.WorkCenterName,
			BusinessPartner:              	data.BusinessPartner,
			Plant:                        	data.Plant,
			WorkCenterCategory:           	data.WorkCenterCategory,
			WorkCenterResponsible:        	data.WorkCenterResponsible,
			SupplyArea:                   	data.SupplyArea,
			WorkCenterUsage:              	data.WorkCenterUsage,
			ComponentIsMarkedForBackflush:	data.ComponentIsMarkedForBackflush,
			WorkCenterLocation:           	data.WorkCenterLocation,
			CapacityCategory:         		data.CapacityCategory,
			CapacityQuantityUnit:         	data.CapacityQuantityUnit,
			CapacityQuantity:         		data.CapacityQuantity,
			ValidityStartDate:            	data.ValidityStartDate,
			ValidityEndDate:              	data.ValidityEndDate,
			CreationDate:            		data.CreationDate,
			LastChangeDate:              	data.LastChangeDate,
			IsMarkedForDeletion:          	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &general, nil
}

func ConvertToProductionCapacity(rows *sql.Rows) (*[]ProductionCapacity, error) {
	defer rows.Close()
	productionCapacity := make([]ProductionCapacity, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ProductionCapacity{}

		err := rows.Scan(
			&pm.WorkCenter,
			&pm.WorkCenterProductionCapacityID,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.Product,
			&pm.CapacityFormula,
			&pm.CalculatedCapacityQuantityInProductionUnit,	
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &productionCapacity, err
		}

		data := pm
		productionCapacity = append(productionCapacity, ProductionCapacity{
			WorkCenter:                   					data.WorkCenter,
			WorkCenterProductionCapacityID:         		data.WorkCenterProductionCapacityID,
			BusinessPartner:              					data.BusinessPartner,
			Plant:                        					data.Plant,
			Product:                        				data.Product,
			CapacityFormula:                        		data.CapacityFormula,
			CalculatedCapacityQuantityInProductionUnit:		data.CalculatedCapacityQuantityInProductionUnit,
			CreationDate:            						data.CreationDate,
			LastChangeDate:              					data.LastChangeDate,
			IsMarkedForDeletion:          					data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &productionCapacity, nil
	}

	return &productionCapacity, nil
}
