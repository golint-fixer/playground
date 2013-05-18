// tank is a backup utility for grooveshark users.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mewmew/playground/gs"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: tank USERNAME")
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	err := tank(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
}

func tank(username string) (err error) {
	gs.Verbose = true
	sess, err := gs.NewSession()
	if err != nil {
		return err
	}
	userId, err := sess.UserId(username)
	if err != nil {
		return err
	}

	now := time.Now()
	date := now.Format("2006-01-02")
	base := fmt.Sprintf("%s - %s", username, date)
	dir := fmt.Sprintf("%s/playlists", base)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Songs.
	err = songs(sess, userId, base)
	if err != nil {
		return err
	}

	// Favorites.
	err = favorites(sess, userId, base)
	if err != nil {
		return err
	}

	// Playlists.
	err = playlists(sess, userId, dir)
	if err != nil {
		return err
	}

	return nil
}

func writeSongs(w io.Writer, songs []*gs.Song) (err error) {
	for _, song := range songs {
		_, err = fmt.Fprintln(w, song)
		if err != nil {
			return err
		}
	}
	return nil
}

func songs(sess *gs.Session, userId int, dir string) (err error) {
	filePath := fmt.Sprintf("%s/songs.txt", dir)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	songs, err := sess.UserSongs(userId)
	if err != nil {
		return err
	}
	err = writeSongs(f, songs)
	if err != nil {
		return err
	}

	return nil
}

func favorites(sess *gs.Session, userId int, dir string) (err error) {
	filePath := fmt.Sprintf("%s/favorites.txt", dir)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	songs, err := sess.UserFavorites(userId)
	if err != nil {
		return err
	}
	err = writeSongs(f, songs)
	if err != nil {
		return err
	}

	return nil
}

func playlists(sess *gs.Session, userId int, dir string) (err error) {
	playlists, err := sess.UserPlaylists(userId)
	if err != nil {
		return err
	}

	for _, playlist := range playlists {
		name := strings.Replace(playlist.Name, "/", ",", -1)
		filePath := fmt.Sprintf("%s/%s.txt", dir, name)
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
		err = writeSongs(f, playlist.Songs)
		if err != nil {
			return err
		}
	}

	return nil
}
