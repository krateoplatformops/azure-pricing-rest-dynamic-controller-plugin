package retail

import "k8s.io/apimachinery/pkg/api/resource"

// PricesResponse:
// type: object
// properties:
//   BillingCurrency:
// 	type: string
// 	example: "USD"
//   CustomerEntityId:
// 	type: string
// 	example: "Default"
//   CustomerEntityType:
// 	type: string
// 	example: "Retail"
//   Items:
// 	type: array
// 	items:
// 	  $ref: '#/components/schemas/PriceItem'
//   NextPageLink:
// 	type: string
// 	nullable: true
// 	example: "https://prices.azure.com:443/api/retail/prices?$filter=serviceName%20eq%20%27Virtual%20Machines%27%20and%20armRegionName%20eq%20%27westeurope%27&$skip=1000"
//   Count:
// 	type: integer
// 	example: 3
// required:
//   - BillingCurrency
//   - CustomerEntityId
//   - CustomerEntityType
//   - Items
//   - Count
// PriceItem:
// type: object
// properties:
//   currencyCode:
// 	type: string
// 	example: "USD"
//   tierMinimumUnits:
// 	type: number
// 	format: float
// 	example: 0.0
//   retailPrice:
// 	type: number
// 	format: float
// 	example: 7.2
//   unitPrice:
// 	type: number
// 	format: float
// 	example: 7.2
//   armRegionName:
// 	type: string
// 	example: "westeurope"
//   location:
// 	type: string
// 	example: "EU West"
//   effectiveStartDate:
// 	type: string
// 	format: date-time
// 	example: "2022-06-01T00:00:00Z"
//   meterId:
// 	type: string
// 	example: "0008192f-190b-5119-87b8-96130c21ecfe"
//   meterName:
// 	type: string
// 	example: "L48s v3"
//   productId:
// 	type: string
// 	example: "DZH318Z08NRP"
//   skuId:
// 	type: string
// 	example: "DZH318Z08NRP/001B"
//   productName:
// 	type: string
// 	example: "Virtual Machines Lsv3 Series Windows"
//   skuName:
// 	type: string
// 	example: "Standard_L48s_v3"
//   serviceName:
// 	type: string
// 	example: "Virtual Machines"
//   serviceId:
// 	type: string
// 	example: "DZH313Z7MMC8"
//   serviceFamily:
// 	type: string
// 	example: "Compute"
//   unitOfMeasure:
// 	type: string
// 	example: "1 Hour"
//   type:
// 	type: string
// 	example: "Consumption"
// 	enum:
// 	  - Consumption
// 	  - DevTestConsumption
//   isPrimaryMeterRegion:
// 	type: boolean
// 	example: true
//   armSkuName:
// 	type: string
// 	example: "Standard_L48s_v3"

type PricesResponse struct {
	BillingCurrency    string  `json:"BillingCurrency"`
	CustomerEntityId   string  `json:"CustomerEntityId"`
	CustomerEntityType string  `json:"CustomerEntityType"`
	Items              []Price `json:"Items"`
	NextPageLink       string  `json:"NextPageLink"`
	Count              int     `json:"Count"`
}

type Price struct {
	CurrencyCode         string            `json:"currencyCode"`
	TierMinimumUnits     resource.Quantity `json:"tierMinimumUnits"`
	RetailPrice          resource.Quantity `json:"retailPrice"`
	UnitPrice            resource.Quantity `json:"unitPrice"`
	ArmRegionName        string            `json:"armRegionName"`
	Location             string            `json:"location"`
	EffectiveStartDate   string            `json:"effectiveStartDate"`
	MeterId              string            `json:"meterId"`
	MeterName            string            `json:"meterName"`
	ProductId            string            `json:"productId"`
	ProductName          string            `json:"productName"`
	SkuId                string            `json:"skuId"`
	SkuName              string            `json:"skuName"`
	ServiceName          string            `json:"serviceName"`
	ServiceId            string            `json:"serviceId"`
	ServiceFamily        string            `json:"serviceFamily"`
	UnitOfMeasure        string            `json:"unitOfMeasure"`
	Type                 string            `json:"type"`
	IsPrimaryMeterRegion bool              `json:"isPrimaryMeterRegion"`
	ArmSkuName           string            `json:"armSkuName"`
}
