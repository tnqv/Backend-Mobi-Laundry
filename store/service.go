package store


type StoreService interface {
	GetStoreDetailByName(name string) (*Store,error)
	CreateNewStore(newStore *Store)(*Store,error)
}

