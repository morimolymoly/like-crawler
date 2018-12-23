package client

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/morimolymoly/like-crawler/config"
)

var singleton *Twitter

// Twitter ... twitter client
type Twitter struct {
	client *twitter.Client
	config *oauth1.Config
	token  *oauth1.Token
}

// GetInstance ... get singleton object
func GetInstance() *Twitter {
	return singleton
}

// GetLikedList ...
func (t *Twitter) GetLikedList(screenName string) {
	singleton.client.Favorites.List(&twitter.FavoriteListParams{
		ScreenName: screenName,
		Count:      200,
	})

	// MaxIDに一番古いlikeのid_strをぶちこんでFavoriteListParamsをつくってリストを取得
	// これでページングできる
}

// Init ... initialize singleton twitter client object
func Init() error {
	c := config.GetInstance()
	if err := config.CheckConfig(); err != nil {
		return err
	}

	singleton := Twitter{}

	singleton.config = oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	singleton.token = oauth1.NewToken(c.AccessToken, c.AccessSecret)
	// oauth client
	httpClient := singleton.config.Client(oauth1.NoContext, singleton.token)
	singleton.client = twitter.NewClient(httpClient)

	return nil
}
