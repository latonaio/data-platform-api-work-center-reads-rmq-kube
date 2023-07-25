package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-work-center-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-work-center-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *[]dpfm_api_output_formatter.General
	var productionCapacity *[]dpfm_api_output_formatter.ProductionCapacity
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		case "Generals":
			func() {
				general = c.Generals(mtx, input, output, errs, log)
			}()
		case "ProductionCapacity":
			func() {
				productionCapacity = c.ProductionCapacity(mtx, input, output, errs, log)
			}()
		case "ProductionCapacities":
			func() {
				productionCapacity = c.ProductionCapacities(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		General: 			general,
		ProductionCapacity: productionCapacity,
	}

	return data
}

func (c *DPFMAPICaller) General(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	where := fmt.Sprintf("WHERE WorkCenter = %d ", input.General.WorkCenter)

	if input.General.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.General.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_work_center_general_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, WorkCenter DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Generals(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	where := "WHERE 1 = 1"
	if input.General.BusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND BusinessPartner = %v", where, *input.General.BusinessPartner)
	}
	if input.General.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.General.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_work_center_general_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ProductionCapacities(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductionCapacity {
	productionCapacity := &dpfm_api_input_reader.ProductionCapacity{}
	if len(input.General.ProductionCapacity) > 0 {
		productionCapacity = &input.General.ProductionCapacity[0]
	}
	where := "WHERE 1 = 1"

	if productionCapacity != nil {
		if productionCapacity.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND productionCapacity.IsMarkedForDeletion = %v", where, *productionCapacity.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT 
			*
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_work_center_production_capacity_data as productionCapacity
		` + where + ` ORDER BY productionCapacity.IsMarkedForDeletion ASC, productionCapacity.WorkCenter DESC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProductionCapacity(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
