package db

func NewDBStore() Store {
	switch {
	default:
		return newMemoryStore(true)
	}
}
