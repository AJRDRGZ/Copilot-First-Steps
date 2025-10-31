package user

import (
	"errors"
	"testing"
)

func Test_Validate_ValidUser_ReturnsNil(t *testing.T) {
	tests := []struct {
		name string
		user User
	}{
		{
			name: "typical valid user",
			user: User{
				ID:    1,
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
		},
		{
			name: "user with spaces in name",
			user: User{
				ID:    2,
				Name:  "  Jane Smith  ",
				Email: "jane@test.org",
			},
		},
		{
			name: "user with complex email",
			user: User{
				ID:    3,
				Name:  "Bob Wilson",
				Email: "bob.wilson+test@sub.domain.co.uk",
			},
		},
		{
			name: "user with hyphenated name",
			user: User{
				ID:    4,
				Name:  "Mary-Jane Watson",
				Email: "mj@spider.net",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange - test case data is provided above

			// Act
			err := Validate(tt.user)

			// Assert
			if err != nil {
				t.Errorf("Validate() returned error for valid user: %v", err)
			}
		})
	}
}

func Test_Validate_InvalidName_ReturnsError(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr string
	}{
		{
			name: "empty name",
			user: User{
				ID:    1,
				Name:  "",
				Email: "valid@example.com",
			},
			wantErr: "name is required",
		},
		{
			name: "whitespace only name",
			user: User{
				ID:    2,
				Name:  "   ",
				Email: "valid@example.com",
			},
			wantErr: "name is required",
		},
		{
			name: "tabs and spaces name",
			user: User{
				ID:    3,
				Name:  "\t  \n  ",
				Email: "valid@example.com",
			},
			wantErr: "name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange - test case data is provided above

			// Act
			err := Validate(tt.user)

			// Assert
			if err == nil {
				t.Fatal("Validate() expected error but got nil")
			}
			if err.Error() != tt.wantErr {
				t.Errorf("Validate() error = %q, want %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func Test_Validate_InvalidEmail_ReturnsError(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr string
	}{
		{
			name: "empty email",
			user: User{
				ID:    1,
				Name:  "Valid Name",
				Email: "",
			},
			wantErr: "invalid email",
		},
		{
			name: "missing @ symbol",
			user: User{
				ID:    2,
				Name:  "Valid Name",
				Email: "invalidemail.com",
			},
			wantErr: "invalid email",
		},
		{
			name: "missing domain",
			user: User{
				ID:    3,
				Name:  "Valid Name",
				Email: "user@",
			},
			wantErr: "invalid email",
		},
		{
			name: "missing username",
			user: User{
				ID:    4,
				Name:  "Valid Name",
				Email: "@example.com",
			},
			wantErr: "invalid email",
		},
		{
			name: "multiple @ symbols",
			user: User{
				ID:    5,
				Name:  "Valid Name",
				Email: "user@@example.com",
			},
			wantErr: "invalid email",
		},
		{
			name: "spaces in email",
			user: User{
				ID:    6,
				Name:  "Valid Name",
				Email: "user @example.com",
			},
			wantErr: "invalid email",
		},
		{
			name: "email with whitespace only",
			user: User{
				ID:    7,
				Name:  "Valid Name",
				Email: "   ",
			},
			wantErr: "invalid email",
		},
		{
			name: "missing TLD",
			user: User{
				ID:    8,
				Name:  "Valid Name",
				Email: "user@domain",
			},
			wantErr: "invalid email",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange - test case data is provided above

			// Act
			err := Validate(tt.user)

			// Assert
			if err == nil {
				t.Fatal("Validate() expected error but got nil")
			}
			if err.Error() != tt.wantErr {
				t.Errorf("Validate() error = %q, want %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func Test_Validate_BothNameAndEmailInvalid_ReturnsNameError(t *testing.T) {
	// Arrange
	user := User{
		ID:    1,
		Name:  "",
		Email: "invalid-email",
	}

	// Act
	err := Validate(user)

	// Assert
	if err == nil {
		t.Fatal("Validate() expected error but got nil")
	}
	// Should return name error first (order dependency test)
	if err.Error() != "name is required" {
		t.Errorf("Validate() error = %q, want %q", err.Error(), "name is required")
	}
}

func Test_Validate_EdgeCaseEmails_ReturnsExpectedResult(t *testing.T) {
	tests := []struct {
		name      string
		user      User
		wantError bool
	}{
		{
			name: "single character domain",
			user: User{
				ID:    1,
				Name:  "Test User",
				Email: "test@a.b",
			},
			wantError: false,
		},
		{
			name: "single character username",
			user: User{
				ID:    2,
				Name:  "Test User",
				Email: "a@example.com",
			},
			wantError: false,
		},
		{
			name: "email with numbers",
			user: User{
				ID:    3,
				Name:  "Test User",
				Email: "user123@domain123.com",
			},
			wantError: false,
		},
		{
			name: "email with special chars",
			user: User{
				ID:    4,
				Name:  "Test User",
				Email: "user-name_test@sub-domain.example.com",
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange - test case data is provided above

			// Act
			err := Validate(tt.user)

			// Assert
			if tt.wantError && err == nil {
				t.Error("Validate() expected error but got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("Validate() unexpected error: %v", err)
			}
		})
	}
}

// Test the exported error for comparison in other packages
func Test_Validate_ErrorType_IsRegularError(t *testing.T) {
	// Arrange
	user := User{
		ID:    1,
		Name:  "",
		Email: "valid@example.com",
	}

	// Act
	err := Validate(user)

	// Assert
	if err == nil {
		t.Fatal("Validate() expected error but got nil")
	}

	// Verify it's a regular error that can be compared
	var targetErr error = errors.New("name is required")
	if err.Error() != targetErr.Error() {
		t.Errorf("Validate() error message = %q, want %q", err.Error(), targetErr.Error())
	}
}
