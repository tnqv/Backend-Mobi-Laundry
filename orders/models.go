package orders

import (
	"github.com/jinzhu/gorm"
	"time"
	"d2d-backend/accounts"
	"d2d-backend/common"
)

type Service struct {
	gorm.Model 				`json:"-"`
	Name        string		`json:"name"`
	Price       int64		`json:"price"`
	Description string		`json:"description"`
	CategoryID  uint		`json:"-"`
	//Categories  Category `gorm:"PRELOAD:false"`
}

type Category struct {
	gorm.Model 				 `json:"-"`
	Name string				 `json:"name"`
	Description string		 `json:"description"`
	Services []Service		 `json:"services"`
}

type ServiceOrder struct {
	gorm.Model				 `json:",omitempty"`
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

	Capacity float32
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

func getAllServicesBasedOnCategory()([]Category,error){
	db := common.GetDB()
	var category []Category
	err := db.Set("gorm:auto_preload", true).Find(&category).Error
	return category,err
}

func getAllOrdersBasedOnCustomerID(userid *uint)([]PlacedOrder,error){
	db := common.GetDB()
	var order []PlacedOrder
	err := db.Set("gorm:auto_preload", true).Find(&order, "customer_id = ?", userid).Error
	return order,err
}

func getTenOrders()([]PlacedOrder,error){
	db := common.GetDB()
	var order []PlacedOrder
	err := db.Limit(10).Set("gorm:auto_preload", true).Find(&order).Error
	return order,err
}