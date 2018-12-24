package client

import (
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/morimolymoly/like-crawler/config"
)

var singleton *Twitter

// Twitter ... twitter client
type Twitter struct {
	client     *twitter.Client
	config     *oauth1.Config
	httpClient *http.Client
	token      *oauth1.Token
}

// GetInstance ... get singleton object
func GetInstance() *Twitter {
	return singleton
}

func (t *Twitter) getPictureURLs(screenName string, maxID int64, sinceID int64) ([]string, int64, int64, error) {
	urls := []string{}
	param := twitter.FavoriteListParams{
		ScreenName: screenName,
		Count:      200,
	}
	if maxID != 0 {
		param.MaxID = maxID
	}
	if sinceID != 0 {
		param.SinceID = sinceID
	}

	likes, _, err := t.client.Favorites.List(&param)
	if err != nil {
		return urls, 0, 0, err
	}

	for _, l := range likes {
		if l.ExtendedEntities == nil {
			continue
		}
		if l.ExtendedEntities.Media == nil {
			continue
		}
		for _, m := range l.ExtendedEntities.Media {
			if m.Type == typePICTURE {
				urls = append(urls, m.MediaURLHttps)
			}
		}
	}
	if len(likes) == 0 {
		return urls, 0, 0, nil
	}
	return urls, likes[len(likes)-1].ID, likes[0].ID, nil
}

// GetLikedAllPictureURLs ...
func (t *Twitter) GetLikedAllPictureURLs(screenName string) ([]string, error) {
	var maxID int64
	var sinceID int64
	lurl := []string{}

	for {
		urls, mid, sid, err := t.getPictureURLs(screenName, maxID, 0)
		if err != nil {
			return nil, err
		}
		if len(urls) == 0 {
			break
		}
		lurl = append(lurl, urls...)
		maxID = mid
		if sinceID == 0 {
			sinceID = sid
		}
	}
	c := config.GetInstance()
	c.UpdateSinceID(sinceID)
	return lurl, nil
}

// GetLatestLikedPictureURLs ...
func (t *Twitter) GetLatestLikedPictureURLs(screenName string) ([]string, error) {
	c := config.GetInstance()
	sinceID := c.SinceID
	lurl := []string{}

	for {
		urls, _, sid, err := t.getPictureURLs(screenName, 0, sinceID)
		if err != nil {
			return nil, err
		}
		if len(urls) == 0 {
			break
		}
		lurl = append(lurl, urls...)
		sinceID = sid
	}
	c.UpdateSinceID(sinceID)
	return lurl, nil
}

// GetLikedList ...
func (t *Twitter) GetLikedList(screenName string) ([]twitter.Tweet, error) {
	param := twitter.FavoriteListParams{
		ScreenName: screenName,
		Count:      200,
	}
	likes, _, err := t.client.Favorites.List(&param)

	if err != nil {
		return nil, err
	}

	return likes, nil
}

// Init ... initialize singleton twitter client object
func Init() error {
	c := config.GetInstance()
	if err := config.CheckConfig(); err != nil {
		return err
	}

	singleton = &Twitter{}
	singleton.config = oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	singleton.token = oauth1.NewToken(c.AccessToken, c.AccessSecret)
	// oauth client
	singleton.httpClient = singleton.config.Client(oauth1.NoContext, singleton.token)
	singleton.client = twitter.NewClient(singleton.httpClient)
	return nil
}
