package notion

import (
	"context"
	"fmt"
)

type UserService service

type ListUserResponse struct {
	Results    []*User `json:"results"`
	NextCursor string  `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
}

type User struct {
	// Object (will always be "user")
	Object string `json:"object"`
	// ID is the users UUID (unique indentifier)
	ID string `json:"id"`
	// Type of user (person or a bot) (could be empty)
	Type string `json:"type"`
	// Name of the user (could be empty)
	Name string `json:"name"`
	// AvatarURL is the url of the avatar (could be empty)
	AvatarURL string `json:"avatar_url"`

	// Person represents a person. It will only be populated if the "type" is set to "person"
	Person Person `json:"person"`
	// Bot represents a bot. It will only be populated if the "type" is set to "bot"
	Bot Bot `json:"bot"`
}

// Person represents a notion person/user
type Person struct {
	// Email is the email of the person
	Email string `json:"email"`
}

// Bot represents a notion bot
type Bot struct {
}

// Get will get a single user by their id
func (s *UserService) Get(ctx context.Context, userID string) (*User, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("users/%s", userID), nil)
	if err != nil {
		return nil, err
	}

	var user *User
	_, err = s.client.Do(ctx, req, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// List will list all of the users
func (s *UserService) List(ctx context.Context) ([]*User, error) {
	// TODO: pagination
	req, err := s.client.NewRequest("GET", "users", nil)
	if err != nil {
		return nil, err
	}

	var users ListUserResponse
	_, err = s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, err
	}

	return users.Results, nil
}
