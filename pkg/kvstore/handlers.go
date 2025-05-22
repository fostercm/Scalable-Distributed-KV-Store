package kvstore

func (store *KVStore) Set(args *SetArgs, reply *SetReply) error {
	// Lock the mutex to ensure thread-safe access to the map
	store.mu.Lock()
	defer store.mu.Unlock()

	// Set the value in the map
	store.data[args.Key] = args.Value

	// Return nil to indicate success
	return nil
}

func (store *KVStore) Get(args *GetArgs, reply *GetReply) error {
	// Lock the mutex to ensure thread-safe access to the map
	store.mu.Lock()
	defer store.mu.Unlock()

	// Get the value from the map
	value, exists := store.data[args.Key]
	if exists {
		reply.Value = value
	}
	reply.Exists = exists

	// Return nil to indicate success
	return nil
}

func (store *KVStore) Delete(args *DeleteArgs, reply *DeleteReply) error {
	// Lock the mutex to ensure thread-safe access to the map
	store.mu.Lock()
	defer store.mu.Unlock()

	// Delete the key from the map
	delete(store.data, args.Key)

	// Return nil to indicate success
	return nil
}

func (store *KVStore) Exists(args *ExistsArgs, reply *ExistsReply) error {
	// Lock the mutex to ensure thread-safe access to the map
	store.mu.Lock()
	defer store.mu.Unlock()

	// Check if the key exists in the map
	_, exists := store.data[args.Key]
	reply.Exists = exists

	// Return nil to indicate success
	return nil
}

func (store *KVStore) Length(args *LengthArgs, reply *LengthReply) error {
	// Get the length of the map
	reply.Length = len(store.data)

	// Return nil to indicate success
	return nil
}
