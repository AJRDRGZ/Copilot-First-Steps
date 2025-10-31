package user

import (
	"reflect"
	"testing"
)

func Test_NewInMemoryRepo_CreatesEmptyRepo_WithCorrectInitialState(t *testing.T) {
	// Arrange - no setup needed

	// Act
	repo := NewInMemoryRepo()

	// Assert
	if repo == nil {
		t.Fatal("NewInMemoryRepo() returned nil")
	}
	if repo.data == nil {
		t.Error("NewInMemoryRepo() data map is nil")
	}
	if len(repo.data) != 0 {
		t.Errorf("NewInMemoryRepo() data length = %d, want 0", len(repo.data))
	}
	if repo.next != 1 {
		t.Errorf("NewInMemoryRepo() next = %d, want 1", repo.next)
	}
}

func Test_Create_SingleUser_ReturnsUserWithID(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Act
	result, err := repo.Create(user)

	// Assert
	if err != nil {
		t.Errorf("Create() returned error: %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Create() ID = %d, want 1", result.ID)
	}
	if result.Name != user.Name {
		t.Errorf("Create() Name = %q, want %q", result.Name, user.Name)
	}
	if result.Email != user.Email {
		t.Errorf("Create() Email = %q, want %q", result.Email, user.Email)
	}
}

func Test_Create_MultipleUsers_ReturnsIncrementingIDs(t *testing.T) {
	tests := []struct {
		name        string
		users       []User
		expectedIDs []int
	}{
		{
			name: "two users",
			users: []User{
				{Name: "User One", Email: "one@example.com"},
				{Name: "User Two", Email: "two@example.com"},
			},
			expectedIDs: []int{1, 2},
		},
		{
			name: "three users",
			users: []User{
				{Name: "Alice", Email: "alice@test.com"},
				{Name: "Bob", Email: "bob@test.com"},
				{Name: "Charlie", Email: "charlie@test.com"},
			},
			expectedIDs: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			repo := NewInMemoryRepo()

			// Act & Assert
			for i, user := range tt.users {
				result, err := repo.Create(user)
				if err != nil {
					t.Errorf("Create() returned error for user %d: %v", i, err)
				}
				if result.ID != tt.expectedIDs[i] {
					t.Errorf("Create() user %d ID = %d, want %d", i, result.ID, tt.expectedIDs[i])
				}
			}
		})
	}
}

func Test_Create_UserWithExistingID_OverwritesID(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()
	user := User{
		ID:    999, // This should be overwritten
		Name:  "Test User",
		Email: "test@example.com",
	}

	// Act
	result, err := repo.Create(user)

	// Assert
	if err != nil {
		t.Errorf("Create() returned error: %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Create() ID = %d, want 1 (should overwrite existing ID)", result.ID)
	}
}

func Test_Create_EmptyUser_StillAssignsID(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()
	user := User{} // Empty user

	// Act
	result, err := repo.Create(user)

	// Assert
	if err != nil {
		t.Errorf("Create() returned error: %v", err)
	}
	if result.ID != 1 {
		t.Errorf("Create() ID = %d, want 1", result.ID)
	}
}

func Test_Get_ExistingUser_ReturnsUserAndTrue(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()
	originalUser := User{Name: "Test User", Email: "test@example.com"}
	createdUser, _ := repo.Create(originalUser)

	// Act
	result, found := repo.Get(createdUser.ID)

	// Assert
	if !found {
		t.Error("Get() found = false, want true")
	}
	if !reflect.DeepEqual(result, createdUser) {
		t.Errorf("Get() result = %+v, want %+v", result, createdUser)
	}
}

func Test_Get_NonExistentUser_ReturnsZeroValueAndFalse(t *testing.T) {
	tests := []struct {
		name string
		id   int
	}{
		{name: "positive non-existent ID", id: 999},
		{name: "zero ID", id: 0},
		{name: "negative ID", id: -1},
		{name: "large ID", id: 2147483647}, // max int32
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			repo := NewInMemoryRepo()

			// Act
			result, found := repo.Get(tt.id)

			// Assert
			if found {
				t.Error("Get() found = true, want false for non-existent user")
			}
			expectedZero := User{}
			if !reflect.DeepEqual(result, expectedZero) {
				t.Errorf("Get() result = %+v, want zero value %+v", result, expectedZero)
			}
		})
	}
}

func Test_Get_AfterDelete_ReturnsZeroValueAndFalse(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()
	user := User{Name: "To Be Deleted", Email: "delete@example.com"}
	created, _ := repo.Create(user)
	repo.Delete(created.ID)

	// Act
	result, found := repo.Get(created.ID)

	// Assert
	if found {
		t.Error("Get() found = true, want false for deleted user")
	}
	expectedZero := User{}
	if !reflect.DeepEqual(result, expectedZero) {
		t.Errorf("Get() result = %+v, want zero value %+v", result, expectedZero)
	}
}

func Test_List_EmptyRepo_ReturnsEmptySlice(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()

	// Act
	result := repo.List()

	// Assert
	if result == nil {
		t.Error("List() returned nil, expected empty slice")
	}
	if len(result) != 0 {
		t.Errorf("List() length = %d, want 0", len(result))
	}
}

func Test_List_WithUsers_ReturnsAllUsers(t *testing.T) {
	tests := []struct {
		name      string
		userCount int
	}{
		{name: "one user", userCount: 1},
		{name: "three users", userCount: 3},
		{name: "five users", userCount: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			repo := NewInMemoryRepo()
			expectedUsers := make([]User, 0, tt.userCount)

			for i := 0; i < tt.userCount; i++ {
				user := User{
					Name:  "User " + string(rune('A'+i)),
					Email: "user" + string(rune('1'+i)) + "@example.com",
				}
				created, _ := repo.Create(user)
				expectedUsers = append(expectedUsers, created)
			}

			// Act
			result := repo.List()

			// Assert
			if len(result) != tt.userCount {
				t.Errorf("List() length = %d, want %d", len(result), tt.userCount)
			}

			// Check all expected users are present (order doesn't matter)
			for _, expected := range expectedUsers {
				found := false
				for _, actual := range result {
					if reflect.DeepEqual(actual, expected) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("List() missing expected user: %+v", expected)
				}
			}
		})
	}
}

func Test_List_AfterDeletion_ReturnsRemainingUsers(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()
	user1, _ := repo.Create(User{Name: "User 1", Email: "user1@example.com"})
	user2, _ := repo.Create(User{Name: "User 2", Email: "user2@example.com"})
	user3, _ := repo.Create(User{Name: "User 3", Email: "user3@example.com"})

	// Delete the middle user
	repo.Delete(user2.ID)

	// Act
	result := repo.List()

	// Assert
	if len(result) != 2 {
		t.Errorf("List() length = %d, want 2", len(result))
	}

	// Verify only user1 and user3 are present
	expectedUsers := []User{user1, user3}
	for _, expected := range expectedUsers {
		found := false
		for _, actual := range result {
			if reflect.DeepEqual(actual, expected) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("List() missing expected user: %+v", expected)
		}
	}

	// Verify user2 is not present
	for _, actual := range result {
		if reflect.DeepEqual(actual, user2) {
			t.Errorf("List() contains deleted user: %+v", user2)
		}
	}
}

func Test_Delete_ExistingUser_ReturnsTrueAndRemovesUser(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()
	user := User{Name: "To Delete", Email: "delete@example.com"}
	created, _ := repo.Create(user)

	// Act
	result := repo.Delete(created.ID)

	// Assert
	if !result {
		t.Error("Delete() = false, want true for existing user")
	}

	// Verify user is actually removed
	_, found := repo.Get(created.ID)
	if found {
		t.Error("Delete() did not remove user from repository")
	}
}

func Test_Delete_NonExistentUser_ReturnsFalse(t *testing.T) {
	tests := []struct {
		name string
		id   int
	}{
		{name: "positive non-existent ID", id: 999},
		{name: "zero ID", id: 0},
		{name: "negative ID", id: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			repo := NewInMemoryRepo()

			// Act
			result := repo.Delete(tt.id)

			// Assert
			if result {
				t.Error("Delete() = true, want false for non-existent user")
			}
		})
	}
}

func Test_Delete_EmptyRepo_ReturnsFalse(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()

	// Act
	result := repo.Delete(1)

	// Assert
	if result {
		t.Error("Delete() = true, want false for empty repository")
	}
}

func Test_Delete_SameUserTwice_SecondCallReturnsFalse(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()
	user := User{Name: "Delete Twice", Email: "twice@example.com"}
	created, _ := repo.Create(user)

	// Act
	firstDelete := repo.Delete(created.ID)
	secondDelete := repo.Delete(created.ID)

	// Assert
	if !firstDelete {
		t.Error("Delete() first call = false, want true")
	}
	if secondDelete {
		t.Error("Delete() second call = true, want false")
	}
}

// Integration tests combining multiple operations
func Test_InMemoryRepo_ComplexWorkflow_MaintainsCorrectState(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()

	// Act & Assert - Create users
	user1, _ := repo.Create(User{Name: "Alice", Email: "alice@example.com"})
	user2, _ := repo.Create(User{Name: "Bob", Email: "bob@example.com"})
	user3, _ := repo.Create(User{Name: "Charlie", Email: "charlie@example.com"})

	if user1.ID != 1 || user2.ID != 2 || user3.ID != 3 {
		t.Errorf("Create() IDs = [%d, %d, %d], want [1, 2, 3]", user1.ID, user2.ID, user3.ID)
	}

	// Verify all users can be retrieved
	_, found1 := repo.Get(1)
	_, found2 := repo.Get(2)
	_, found3 := repo.Get(3)

	if !found1 || !found2 || !found3 {
		t.Error("Get() failed to find created users")
	}

	// Delete middle user
	deleted := repo.Delete(2)
	if !deleted {
		t.Error("Delete() failed to delete existing user")
	}

	// Verify list shows only remaining users
	remaining := repo.List()
	if len(remaining) != 2 {
		t.Errorf("List() after delete length = %d, want 2", len(remaining))
	}

	// Create another user - should get next ID
	user4, _ := repo.Create(User{Name: "David", Email: "david@example.com"})
	if user4.ID != 4 {
		t.Errorf("Create() after delete ID = %d, want 4", user4.ID)
	}

	// Final verification
	finalList := repo.List()
	if len(finalList) != 3 {
		t.Errorf("Final List() length = %d, want 3", len(finalList))
	}
}

func Test_InMemoryRepo_BoundaryValues_HandlesEdgeCases(t *testing.T) {
	// Arrange
	repo := NewInMemoryRepo()

	// Act - Create many users to test ID increment
	const numUsers = 1000
	for i := 0; i < numUsers; i++ {
		user := User{
			Name:  "User",
			Email: "user@example.com",
		}
		created, err := repo.Create(user)

		// Assert
		if err != nil {
			t.Errorf("Create() user %d returned error: %v", i, err)
		}
		if created.ID != i+1 {
			t.Errorf("Create() user %d ID = %d, want %d", i, created.ID, i+1)
		}
	}

	// Verify final state
	allUsers := repo.List()
	if len(allUsers) != numUsers {
		t.Errorf("List() after creating %d users length = %d, want %d", numUsers, len(allUsers), numUsers)
	}
}
