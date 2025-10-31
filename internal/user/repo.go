package user

// NOTE: This repository is intentionally naive and NOT concurrency-safe.
// Part of the Kata is to improve its safety and ergonomics.

// InMemoryRepo is a simple in-memory user repository implementation.
// It stores users in a map with auto-incrementing integer IDs.
type InMemoryRepo struct {
	data map[int]User // maps user ID to User
	next int          // next available ID
}

// NewInMemoryRepo creates and returns a new InMemoryRepo instance
// with an empty data map and starting ID of 1.
func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{data: make(map[int]User), next: 1}
}

// Create adds a new user to the repository with an auto-generated ID.
// The user's ID field will be overwritten with the new ID.
// Returns the user with the assigned ID and no error (always succeeds).
func (r *InMemoryRepo) Create(u User) (User, error) {
	u.ID = r.next
	r.data[r.next] = u
	r.next++
	return u, nil
}

// Get retrieves a user by ID from the repository.
// Returns the user and true if found, zero-value User and false if not found.
func (r *InMemoryRepo) Get(id int) (User, bool) {
	u, ok := r.data[id]
	return u, ok
}

// List returns all users in the repository as a slice.
// The order of users in the returned slice is not guaranteed.
func (r *InMemoryRepo) List() []User {
	out := make([]User, 0, len(r.data))
	for _, v := range r.data {
		out = append(out, v)
	}
	return out
}

// Delete removes a user from the repository by ID.
// Returns true if the user was found and deleted, false if not found.
func (r *InMemoryRepo) Delete(id int) bool {
	if _, ok := r.data[id]; ok {
		delete(r.data, id)
		return true
	}
	return false
}
