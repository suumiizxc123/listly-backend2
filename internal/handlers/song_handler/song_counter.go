package song_handler

import (
	"kcloudb1/internal/models/song"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSongCounter(c *gin.Context) {
	var song song.SongCounter
	var err error

	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"fields required", "Дууны мэдээлэл дутуу байна"}, err),
		)
		return
	}

	if err = song.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"creation failed", "Дууны мэдээлэл хадгалахад алдаа үүслээ"}, err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success([]string{"created", "Дууны мэдээлэл хадгаллаа"}, song))
}

// bayaraa number 99267774
func UpdateSongCounter(c *gin.Context) {
    var song song.SongCounter
    var err error

    if err := c.ShouldBindJSON(&song); err != nil {
        c.JSON(
            http.StatusBadRequest,
            utils.Error([]string{"fields required", "Дууны мэдээлэл дутуу байна"}, err),
        )
        return
    }

    if err = song.Update(); err != nil {
        c.JSON(
            http.StatusInternalServerError,
            utils.Error([]string{"update failed", "Дууны мэдээлэл хадгалахад алдаа үүслээ"}, err),
        )
        return
    }

    c.JSON(200, utils.Success([]string{"updated", "Дууны мэдээлэл хадгаллаа"}, song))
}

func GetSongCounter(c *gin.Context) {

	var song song.SongCounter
	var err error
	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"id required", "Дууны id мэдээлэл дутуу байна"}, "id must be required"),
		)

		return
	}

	song.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"id cannot be parsed", "Дууны id буруу байна"}, err.Error()),
		)

		return
	}

	if err = song.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"not found", "Дуу олдсонгүй"}, err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"found", "Дууны мэдээлэл олдлоо"}, song))
}

func DeleteSongCounter(c *gin.Context) {

	var song song.SongCounter
	var err error

	id, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"id required", "Дууны id мэдээлэл дутуу байна"}, "id must be required"),
		)

		return
	}

	song.ID, err = strconv.ParseInt(id, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"id cannot be parsed", "Дууны id буруу байна"}, err.Error()),
		)

		return
	}

	if err = song.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"delete failed", "Дуу устгаж чадсангүй"}, err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"deleted", "Дууны мэдээлэл устгагдлаа"}, struct{}{}))
}

func GetSongCounterList(c *gin.Context) {
	var song song.SongCounter

	songs, err := song.GetList()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"list not found", "Дуу олдсонгүй"}, err.Error()),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"list", "Дууны жагсаалт"}, songs))
}
