package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	auth "github.com/cleverswine/cliauthorizationflow"
)

const (
	SpotifyAuthURL  = "https://accounts.spotify.com/authorize"
	SpotifyTokenURL = "https://accounts.spotify.com/api/token"
	SpotifyAPIBase  = "https://api.spotify.com/v1"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := &auth.Config{
		ClientID:         os.Getenv("SPOTIFY_ID"),
		ClientSecret:     os.Getenv("SPOTIFY_SECRET"),
		AuthorizationURL: SpotifyAuthURL,
		TokenURL:         SpotifyTokenURL,
		Scopes:           []string{"user-top-read"},
		//CallbackPort:     8089,
	}

	client, err := auth.NewClient(ctx, config, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Persist()

	// get my top tracks
	resp, err := client.Get(SpotifyAPIBase + "/me/top/tracks")
	if err != nil {
		log.Fatal(err)
	}
	// write them out to a json file
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("top-tracks.json", body, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// func getConfigFile() string {
// 	osUser, err := user.Current()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	dir := path.Join(osUser.HomeDir, ".config/spot")
// 	_, err = os.Stat(dir)
// 	if os.IsNotExist(err) {
// 		log.Println("creating directory " + dir)
// 		errDir := os.MkdirAll(dir, 0600)
// 		if errDir != nil {
// 			log.Fatal(errDir)
// 		}
// 	}
// 	return path.Join(dir, "t")
// }

// func tokenFromCache() (*oauth2.Token, error) {
// 	fn := getConfigFile()
// 	if _, err := os.Stat(fn); err != nil {
// 		return nil, nil
// 	}
// 	f, err := os.Open(fn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()
// 	t := &oauth2.Token{}
// 	err = json.NewDecoder(f).Decode(t)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return t, nil
// }

// func saveTokenToCache(token *oauth2.Token) error {
// 	f, err := os.Create(getConfigFile())
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	log.Println("saving token to " + f.Name())
// 	return json.NewEncoder(f).Encode(token)
// }
