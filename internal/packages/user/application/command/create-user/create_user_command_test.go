package createuser_test

import (
	"testing"

	createuser "github.com/rafaelbrunoss/general-server-go/internal/packages/user/application/command/create-user"
)

func TestNewCreateUserCommand(t *testing.T) {
	tests := []struct {
		input     createuser.CreateUserCommandInput
		expected  string
		expectErr bool
	}{
		{
			input:     createuser.CreateUserCommandInput{Name: "John Doe"},
			expected:  "John Doe",
			expectErr: false,
		},
		{
			input:     createuser.CreateUserCommandInput{Name: ""},
			expected:  "",
			expectErr: false, // Assuming empty name is valid
		},
	}

	for _, test := range tests {
		t.Run(test.input.Name, func(t *testing.T) {
			cmd, err := createuser.NewCreateUserCommand(test.input)

			if (err != nil) != test.expectErr {
				t.Fatalf("expected error: %v, got: %v", test.expectErr, err)
			}

			if cmd != nil && cmd.Name != test.expected {
				t.Fatalf("expected name %s, got %s", test.expected, cmd.Name)
			}
		})
	}
}
