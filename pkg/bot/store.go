package bot

type data map[string]any
type usersStore map[int64]data

type Store struct {
	usersStore usersStore
}

func NewStore() *Store {
	return &Store{
		usersStore: usersStore{},
	}
}

func (s *Store) AddData(id int64, key string, value any) {
	userStore, ok := s.usersStore[id]
	if !ok {
		s.usersStore[id] = data{}
		userStore = s.usersStore[id]
	}
	userStore[key] = value
}

func (s *Store) GetData(id int64, key string) any {
	userStore := s.usersStore[id]
	return userStore[key]
}

func (s *Store) ResetData(id int64) {
	delete(s.usersStore, id)
}
