package main

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/artist"
	"kcloudb1/internal/models/common"
	"kcloudb1/internal/models/org"
	"kcloudb1/internal/models/song"
	"kcloudb1/internal/models/user"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(
		user.SysUser{},
		user.User{},
		user.ServiceLog{},
		song.Song{},
		song.SongCategory{},
		song.SongCategoryCombination{},
		org.Org{},
		common.Language{},
		artist.ArtistProfile{},
		artist.ArtistMember{},
		artist.ArtistMemberSong{},
		artist.ArtistSong{},
		artist.ArtistType{},
	)
}
