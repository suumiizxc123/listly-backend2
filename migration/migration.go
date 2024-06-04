package main

import (
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/artist"
	"kcloudb1/internal/models/common"
	"kcloudb1/internal/models/moscap"
	"kcloudb1/internal/models/org"
	"kcloudb1/internal/models/song"
	"kcloudb1/internal/models/user"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(
		user.SysUser{},
		user.User{},
		song.Song{},
		song.SongCategory{},
		song.SongCategoryCombination{},
		song.SongCounter{},
		song.SongLanguage{},
		org.Org{},
		org.OrgUser{},
		org.OrgAccount{},
		org.OrgAccountTxn{},
		org.OrgAccountTxnLog{},
		moscap.MosCapUser{},
		moscap.MosCapUserLog{},
		common.Language{},
		artist.ArtistProfile{},
		artist.ArtistMember{},
		artist.ArtistMemberSong{},
		artist.ArtistSong{},
		artist.ArtistType{},
		artist.ArtistBand{},
		artist.ArtistBandCombination{},
	)
}
