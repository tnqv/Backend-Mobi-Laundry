package store


type StoreRepository interface{
	Find(id string) (*Store, error)
	FindByStoreName(name string) (*Store, error)
	FindAll() ([]*Store, error)
	Create(store *Store) (string,error)
	Update(store *Store) error
	Delete(id int) (bool,error)
}