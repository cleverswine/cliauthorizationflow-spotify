package main

import (
	"context"
	"fmt"
	"log"
	"os"

	auth "github.com/cleverswine/cliauthorizationflow"
	"github.com/zmb3/spotify"
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
	}

	client, err := auth.NewClient(ctx, config, auth.NewDefaultTokenStorage("spotify-cli"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Persist()

	spotifyClient := spotify.NewClient(client.Client)
	tracks, err := spotifyClient.CurrentUsersTopTracks()
	if err != nil {
		log.Fatal(err)
	}
	for _, track := range tracks.Tracks {
		fmt.Println(track.String())
	}
}
