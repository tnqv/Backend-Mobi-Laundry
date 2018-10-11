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
	ImageUrl	string		`json:"url"`
	//Categories  category `gorm:"PRELOAD:false"`
}

type Category struct {
	gorm.Model 				 `json:"-"`
	Name string				 `json:"name"`
	Description string		 `json:"description"`
	Services []Service		 `json:"service"`
}

type ServiceOrder struct {
	gorm.Model				 		`json:"-"`
	PlacedOrderID uint				`json:"placed_order_id"`
	PlacedOrderModel PlacedOrder	`json:"-"`
	ServiceID uint					`json:"service_id"`
	ServiceModel Service			`json:"-"`
	Description string				`json:"description"`
	Quantity uint					`json:"quantity"`
	Price float32					`json:"price"`
}

type PlacedOrder struct {
	gorm.Model							`json:"-"`
	//Store
	StoreID uint						`json:"-"`
	AssignedStore accounts.Store		`json:"-"`
	TimePlaced time.Time				`json:"time_placed"`
	Detail string						`json:"note"`
	//Order status
	OrderStatusID uint					`json:"status"`
	OrderStatusModel OrderStatus		`json:"-"`

	//Customer
	UserID uint					    	`json:"user_id"`
	UserModel accounts.User				`json:"-"`

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
	OrderCode string					`json:"order_code"`
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
	UserID uint
	UserModel			uint		`json:"-"`
	StatusChangedTime time.Time
	Description string
}

type Review struct {
	gorm.Model						`json:"-"`
	Content string
	Rate int
	UserID uint
	UserModel accounts.User
}

type Notification struct {
	gorm.Model         						`json:"-"`
	NotificationTypeID 		uint   			`json:"notification_type_id"`
	Read               		bool   			`json:"read"`
	Content            		string 			`json:"content"`
	//Customer
	UserID    				uint          	`json:"user_id"`
	UserModel 				accounts.User 	`json:"-"`
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
	db.AutoMigrate(&Notification{})
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

/*func createOrderStatus(orderstatus *OrderStatus) {
	db := common.GetDB()
	db.Create(&orderstatus)
}*/

func getCustomerInformations(accountID uint) (accounts.User) {
	db := common.GetDB()
	var customer accounts.User
	db.Find(&customer, "account_id = ?", accountID)
	return customer
}

/*func updateOrderStatus(orderID uint, orderStatusID uint) (PlacedOrder, error) {
	db := common.GetDB()
	var order PlacedOrder
	err := db.Model(&order).Where("id = ?", orderID).Update("order_status_id", orderStatusID).Error
	return order, err
}*/

//Minh's function
func getAllOrdersBasedOnAccountID(accountid uint)([]PlacedOrder,error){
	db := common.GetDB()
	var order []PlacedOrder
	//var customer accounts.User
	//db.Find(&customer, "user_id = ?", accountid)
	err := db.Set("gorm:auto_preload", true).Find(&order, "user_id = ?", accountid).Error
	return order,err
}

func getOrders()([]PlacedOrder,error){
	db := common.GetDB()
	var order []PlacedOrder
	err := db.Set("gorm:auto_preload", true).Find(&order).Error
	return order,err
}

func createOrderService(orderservice *ServiceOrder){
	db := common.GetDB()
	db.Create(&orderservice)
}

//category
func createCategory(category *Category){
	db := common.GetDB()
	db.Create(&category)
}

func updateCategory(category *Category)(error){
	db := common.GetDB()
	err := db.Model(&category).Update(map[string]interface{}{"name":category.Name,"description":category.Description}).Error
	return err
}

func getCategory(cateId uint)(Category, error){
	db := common.GetDB()
	var cate Category
	err := db.First(&cate, cateId).Error
	return cate, err
}

func deleteCategory(cateId uint) (error) {
	db := common.GetDB()
	err := db.Delete(&Category{}, "id = ?", cateId).Error
	return err
}

//Order Status
func createOrderStatus(orderstatus *OrderStatus){
	db := common.GetDB()
	db.Create(&orderstatus)
}

func updateOrderStatus(orderstatus *OrderStatus)(error){
	db := common.GetDB()
	err := db.Model(&orderstatus).Update(map[string]interface{}{"description":orderstatus.Description}).Error
	return err
}

func getOrderStatus(orderstatusId uint)(OrderStatus, error){
	db := common.GetDB()
	var orderstatus OrderStatus
	err := db.First(&orderstatus, orderstatusId).Error
	return orderstatus, err
}

func getListOrderStatuses() ([]OrderStatus, error) {
	db := common.GetDB()
	var list []OrderStatus
	err := db.Find(&list).Error
	return list, err
}

func deleteOrderStatus(orderstatusId uint) (error) {
	db := common.GetDB()
	err := db.Delete(&OrderStatus{}, "id = ?", orderstatusId).Error
	return err
}

//SERVICE_ORDERS ENTITY
func getListServiceOrders() ([]ServiceOrder, error) {
	db := common.GetDB()
	var list []ServiceOrder
	err := db.Find(&list).Error
	return list, err
}

func getServiceOrder(serviceOrderId uint) (ServiceOrder, error) {
	db := common.GetDB()
	var serviceOrder ServiceOrder
	err := db.First(&serviceOrder, serviceOrderId).Error
	return serviceOrder, err
}

func createServiceOrder(serviceOrder *ServiceOrder) (error) {
	db := common.GetDB()
	err := db.Create(&serviceOrder).Error
	return err
}

func updateServiceOrder(serviceOrder *ServiceOrder) (error) {
	db := common.GetDB()
	err := db.Model(&ServiceOrder{}).Where("id = ?", serviceOrder.ID).Update("quantity", serviceOrder.Quantity).Error
	return err
}

func deleteServiceOrder(serviceOrderID uint) (error) {
	db := common.GetDB()
	err := db.Delete(&ServiceOrder{}, "id = ?", serviceOrderID).Error
	return err
}
//END SERVICE_ORDERS ENTITY

//SERVICE ENTITY
func getListServices() ([]Service, error) {
	db := common.GetDB()
	var list []Service
	err := db.Find(&list).Error
	return list, err
}

func getService(serviceId uint) (Service, error) {
	db := common.GetDB()
	var service Service
	err := db.First(&service, serviceId).Error
	return service, err
}

func createService(service *Service) (error) {
	db := common.GetDB()
	err := db.Create(&service).Error
	return err
}

func updateService(service *Service) (error) {
	db := common.GetDB()
	err := db.Save(&service).Error
	//err := db.Model(&service).Updates(map[string]interface{}{"name": service.Name, "description": service.Description}).Error
	return err
}

func deleteService(serviceId uint) (error) {
	db := common.GetDB()
	err := db.Delete(&Service{}, "id = ?", serviceId).Error
	return err
}

//Notification
//func getListNotifications() ([]Notification, error) {
//	db := common.GetDB()
//	var list []Notification
//	err := db.Find(&list).Error
//	return list, err
//}
//
//func getNotifications(notificationId uint) (Notification, error) {
//	db := common.GetDB()
//	var serviceOrder ServiceOrder
//	err := db.First(&serviceOrder, notificationId).Error
//	return serviceOrder, err
//}