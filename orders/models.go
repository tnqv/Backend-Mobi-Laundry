package orders

import (
	"d2d-backend/accounts"
	"d2d-backend/common"
	"github.com/jinzhu/gorm"
	"time"
)

type Service struct {
	gorm.Model 				`json:"-"`
	Name        string		`json:"name"`
	Price       int64		`json:"price"`
	Description string		`json:"description"`
	CategoryID  uint		`json:"-"`
	ImageUrl			string		`json:"url"`
	//Categories  Category `gorm:"PRELOAD:false"`
}

type Category struct {
	gorm.Model 				 `json:"-"`
	Name string				 `json:"name"`
	Description string		 `json:"description"`
	Services []Service		 `json:"services"`
}

type ServiceOrder struct {
	gorm.Model				 `json:"-"`
	PlacedOrderID uint
	PlacedOrderModel PlacedOrder
	ServiceID uint
	ServiceModel Service
	Description string
	Quantity uint
	Price float32
}

type PlacedOrder struct {
	gorm.Model							`json:"-"`
	//Store
	StoreID uint						`json:"-"`
	AssignedStore accounts.Store		`json:"store"`
	TimePlaced time.Time				`json:"time_placed"`
	Detail string						`json:"note"`
	//Order status
	OrderStatusID uint					`json:"status"`
	OrderStatusModel OrderStatus		`json:"-"`

	//Customer
	CustomerID uint						`json:"customer_id"`
	CustomerModel accounts.Customer		`json:"-"`

	Capacity float32					`json:"capacity"`
	EstimatedCapacity float32			`json:"estimated_capacity"`
	DeliveryAddress string				`json:"delivery_address"`
	DeliveryLatitude float32			`json:"delivery_latitude"`
	DeliveryLongitude float32			`json:"delivery_longitude"`
	ServiceTotalPrice float32			`json:"total"`
	Priority int						`json:"priority"`
	//Review
	ReviewID uint						`json:"-"`
	OrderReview Review					`json:"-"`

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
// 9 : Order completed
type OrderStatus struct {
	gorm.Model						`json:"-"`
	StatusID uint
	AccountID uint
	StatusChangedTime time.Time
	Description string
}

type Review struct {
	gorm.Model						`json:"-"`
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

//DuyNQ's function
func createPlaceOrder(order *PlacedOrder)  {
	db := common.GetDB()
	db.Create(&order)
}

func createOrderStatus(orderstatus *OrderStatus) {
	db := common.GetDB()
	db.Create(&orderstatus)
}

func getCustomerInformations(accountID uint) (accounts.Customer) {
	db := common.GetDB()
	var customer accounts.Customer
	db.Find(&customer, "account_id = ?", accountID)
	return customer
}

//Minh's function
func getAllOrdersBasedOnAccountID(accountid uint)([]PlacedOrder,error){
	db := common.GetDB()
	var order []PlacedOrder
	var customer accounts.Customer
	db.Find(&customer, "account_id = ?", accountid)
	err := db.Set("gorm:auto_preload", true).Find(&order, "customer_id = ?", customer.ID).Error
	return order,err
}

func getOrders()([]PlacedOrder,error){
	db := common.GetDB()
	var order []PlacedOrder
	err := db.Set("gorm:auto_preload", true).Find(&order).Error
	return order,err
}