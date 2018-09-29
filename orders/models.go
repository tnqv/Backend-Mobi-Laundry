package orders

import (
	"github.com/jinzhu/gorm"
	"time"
	"d2d-backend/accounts"
	"d2d-backend/common"
)

type Service struct {
	gorm.Model
	Name        string
	Price       int64
	Description string
	CategoryID  uint
	Categories  Category
}

type Category struct {
	gorm.Model
	Name string
	Description string
	ServiceModels []Service
}

type ServiceOrder struct {
	gorm.Model
	PlacedOrderID uint
	PlacedOrderModel PlacedOrder
	ServiceID uint
	ServiceModel Service
	Description string
	Quantity uint
	Price float32
}

type PlacedOrder struct {
	gorm.Model
	//Store
	StoreID uint
	AssignedStore accounts.Store
	TimePlaced time.Time
	Detail string
	//Order status
	OrderStatusID uint
	OrderStatusModel OrderStatus

	//Customer
	CustomerID uint
	CustomerModel accounts.Customer

	DeliveryAddress string
	DeliveryLatitude float32
	DeliveryLongitude float32
	ServiceTotalPrice float32
	Priority int
	//Review
	ReviewID uint
	OrderReview Review

}

//Status ID
// 1 : Order just Created
// 2 : Accepted orders
// 3 : Delivery verified orders
// 4 : Pick up order to warehouse
// 5 : Store xx received order
// 6 : Order is in process
// 7 : Order is finished laundry
// 8 : Order is delivering
type OrderStatus struct {
	gorm.Model
	StatusID uint
	StatusChangedTime time.Time
	Description string
}

type Review struct {
	gorm.Model
	Content string
	Rate int
	CustomerID uint
	CustomerMode accounts.Customer
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Service{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&OrderStatus{})
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&PlacedOrder{})
	db.AutoMigrate(&ServiceOrder{})
}