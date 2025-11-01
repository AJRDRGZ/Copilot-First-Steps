package user

import "sync"

// InMemoryRepo is a thread-safe in-memory user repository implementation.
// It stores users in a map with auto-incrementing integer IDs.
// All operations are protected by a mutex for concurrent access.
type InMemoryRepo struct {
	mu   sync.RWMutex // protects data and next fields
	data map[int]User // maps user ID to User
	next int          // next available ID
}

// NewInMemoryRepo creates and returns a new InMemoryRepo instance
// with an empty data map and starting ID of 1.
func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		data: make(map[int]User),
		next: 1,
	}
}

// Create adds a new user to the repository with an auto-generated ID.
// The user's ID field will be overwritten with the new ID.
// Returns the user with the assigned ID and no error (always succeeds).
// This method is thread-safe.
func (r *InMemoryRepo) Create(u User) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Validate user input before creating
	if err := Validate(u); err != nil {
		return User{}, err
	}

	u.ID = r.next
	r.data[r.next] = u
	r.next++
	return u, nil
}

// Get retrieves a user by ID from the repository.
// Returns the user and true if found, zero-value User and false if not found.
// This method is thread-safe.
func (r *InMemoryRepo) Get(id int) (User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.data[id]
	return u, ok
}

// List returns all users in the repository as a slice.
// The order of users in the returned slice is not guaranteed.
// This method is thread-safe.
func (r *InMemoryRepo) List() []User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]User, 0, len(r.data))
	for _, v := range r.data {
		out = append(out, v)
	}
	return out
}

// Delete removes a user from the repository by ID.
// Returns true if the user was found and deleted, false if not found.
// This method is thread-safe.
func (r *InMemoryRepo) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; ok {
		delete(r.data, id)
		return true
	}
	return false
}
