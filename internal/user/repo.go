package user

// NOTE: This repository is intentionally naive and NOT concurrency-safe.
// Part of the Kata is to improve its safety and ergonomics.

type InMemoryRepo struct {
	data map[int]User
	next int
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{data: make(map[int]User), next: 1}
}

func (r *InMemoryRepo) Create(u User) (User, error) {
	u.ID = r.next
	r.data[r.next] = u
	r.next++
	return u, nil
}

func (r *InMemoryRepo) Get(id int) (User, bool) {
	u, ok := r.data[id]
	return u, ok
}

func (r *InMemoryRepo) List() []User {
	out := make([]User, 0, len(r.data))
	for _, v := range r.data {
		out = append(out, v)
	}
	return out
}

func (r *InMemoryRepo) Delete(id int) bool {
	if _, ok := r.data[id]; ok {
		delete(r.data, id)
		return true
	}
	return false
}
