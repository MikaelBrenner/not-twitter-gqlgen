package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"not-twitter/auth"
	"not-twitter/entities/tweets"
	"not-twitter/entities/users"
	"not-twitter/graph/generated"
	"not-twitter/graph/model"
	"strconv"
)

func (r *mutationResolver) CreateTweet(ctx context.Context, input model.NewTweet) (*model.Tweet, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Tweet{}, fmt.Errorf("access denied")
	}
	tweet := tweets.Tweet{
		Content: input.Content,
		User: user,
	}

	id := tweet.Save()
	gqlUser := &model.User{
		ID:   user.ID,
		Name: user.Username,
	}
	return &model.Tweet{
		Content: tweet.Content,
		ID:      strconv.FormatInt(id, 10),
		User: gqlUser,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := auth.GenerateToken(user.Username)
	if err != nil{
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		return "", errors.New("wrong username or password")
	}
	token, err := auth.GenerateToken(user.Username)
	if err != nil{
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := auth.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := auth.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Tweets(ctx context.Context) ([]*model.Tweet, error) {
	var responseTweets []*model.Tweet
	var dbTweets []tweets.Tweet
	dbTweets = tweets.GetAll()
	for _, t := range dbTweets {
		gqlUser := &model.User{
			ID:   t.User.ID,
			Name: t.User.Username,
		}
		responseTweets = append(responseTweets, &model.Tweet{
			ID:      t.ID,
			Content: t.Content,
			User: gqlUser,
		})
	}
	return responseTweets, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
