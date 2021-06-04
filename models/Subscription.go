package models

import (
	"strings"
	"time"

	"github.com/rahulpargaonkar/SubscriptionManagementAPIDemo/dtos"
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	Type      string
	StartDate time.Time
	EndDate   time.Time
	Amount    float64
	UserID    int64
	ProductID int64
}

func GetSubscriptionAmount(productName string, SubscriptionType string, startDate time.Time) (float64, time.Time) {

	amount := 0.0

	var productPricesDTO []dtos.ProductPricesDTO
	query := "select day, price from weekly_product_prices wp , Products p where  p.name='" + productName + "' and p.id=wp.product_id;"
	GetRecordsInStruct(query, &productPricesDTO)
	var endDate time.Time
	if len(productPricesDTO) != 0 {
		dateMap := make(map[string]float64)
		for p := range productPricesDTO {
			dateMap[productPricesDTO[p].Day] = productPricesDTO[p].Price
		}

		if strings.EqualFold(SubscriptionType, "Monthly") {
			endDate = startDate.AddDate(0, 0, 29)
			CalculateSubscriptionAmount(startDate, &amount, endDate, dateMap)

		}
		return amount, endDate
	} else {
		return 0, endDate
	}

}

func CalculateSubscriptionAmount(startDate time.Time, amount *float64, endDate time.Time, dateMap map[string]float64) {
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {

		*amount = *amount + dateMap[d.Weekday().String()]
	}
}
