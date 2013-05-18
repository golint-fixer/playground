// Package gs implements an unofficial API for grooveshark.com.
package gs

import (
	"errors"
	"net/http"
)

// UserId returns the user id associated with the provided username.
func UserId(username string) (userId int, err error) {
	return 0, errors.New("gs.UserId: not yet implemented.")
}

// UserSongs returns a list of all songs in the provided user's collection.
func UserSongs(userId int) (songs []Song, err error) {
	return nil, errors.New("gs.UserSongs: not yet implemented.")
}

// UserFavorites returns a list of the provided user's favorite songs.
func UserFavorites(userId int) (songs []Song, err error) {
	return nil, errors.New("gs.UserFavorites: not yet implemented.")
}

// UserPlaylists returns a list of the provided user's playlists.
func UserPlaylists(userId int) (playlists []Playlist, err error) {
	return nil, errors.New("gs.UserPlaylists: not yet implemented.")
}

// Sess contains the cookies and token of a grooveshark session.
type Sess struct {
	// Session cookie (PHPSESSID).
	cookie *http.Cookie
	// Communication token based on the user's session.
	commToken string
}

// NewSession creates an unauthenticated session.
func NewSession() (sess *Sess, err error) {
	return nil, errors.New("gs.NewSession: not yet implemented.")
}

// A Song is a music track with associated information.
type Song struct {
	// Song title.
	Title string
	// Artist of song.
	Artist string
	// Song album name.
	Album string
	// Album track number.
	TrackNum int
	// Song id.
	id int
	// Artist id.
	artistId int
}

// A Playlist is an ordered list of songs with an associated name.
type Playlist struct {
	// Playlist name.
	Name string
	// An ordered slice of songs.
	Songs []Song
}
