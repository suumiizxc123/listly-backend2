package artist_handler

import (
	"kcloudb1/internal/models/artist"
	"kcloudb1/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArtistType(c *gin.Context) {
	var artisType artist.ArtistType

	if err := c.ShouldBindJSON(&artisType); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist type fields required", "Артист төрлийн талбарыг бөглөнө үү"}, err),
		)
		return
	}

	if err := artisType.Create(); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist type creation failed", "Артист төрлийн бүртгэхэд алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(http.StatusCreated, utils.Success([]string{"Artist type created", "Артист төрлийн бүртгэгдлүүд амжилттай боллоо"}, artisType))
}

func GetArtistTypeList(c *gin.Context) {
	var artisType artist.ArtistType

	list, err := artisType.GetList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist type list failed", "Артист төрлийн жагсаалт авахад алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Artist type list", "Артист төрлийн жагсаалт амжилттай боллоо"}, list))
}

func GetArtistTypeListByType(c *gin.Context) {
	var artisType artist.ArtistType

	tp, ok := c.GetQuery("type")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist type type required", "Артист төрлийн төрөл бөглөнө үү"}, nil),
		)
		return
	}

	artisType.Type = tp

	list, err := artisType.GetListByType()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist type list failed", "Артист төрлийн жагсаалт авахад алдаа гарлаа"}, err),
		)
		return
	}

	c.JSON(200, utils.Success([]string{"Artist type list", "Артист төрлийн жагсаалт амжилттай боллоо"}, list))
}

func GetArtistType(c *gin.Context) {

	var artisType artist.ArtistType
	var err error
	tp, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist type id required", "Артист төрлийн ID бөглөнө үү"}, nil),
		)

		return
	}

	artisType.ID, err = strconv.ParseInt(tp, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist type id required", "Артист төрлийн ID бөглөнө үү"}, err),
		)

		return
	}

	if err = artisType.Get(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist type not found", "Артист төрлийн бүртгэлтэй байхад алдаа гарлаа"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist type found", "Артист төрлийн бүртгэлтэй амжилттай боллоо"}, artisType))
}

func UpdateArtistType(c *gin.Context) {

	var artisType artist.ArtistType

	if err := c.ShouldBindJSON(&artisType); err != nil {
		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist type fields required", "Артист төрлийн мэдээлэл бөглөнө үү"}, err),
		)
		return
	}

	if err := artisType.Update(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist type update failed", "Артист төрлийн бүртгэлтэй амжилтгүй боллоо"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist type updated", "Артист төрлийн бүртгэлтэй амжилттай боллоо"}, artisType))

}

func DeleteArtistType(c *gin.Context) {

	var artisType artist.ArtistType
	var err error
	tp, ok := c.GetQuery("id")

	if !ok {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist type id required", "Артист төрлийн ID бөглөнө үү"}, nil),
		)

		return
	}

	artisType.ID, err = strconv.ParseInt(tp, 10, 64)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			utils.Error([]string{"Artist type id required", "Артист төрлийн ID бөглөнө үү"}, err),
		)

		return
	}

	if err = artisType.Delete(); err != nil {

		c.JSON(
			http.StatusInternalServerError,
			utils.Error([]string{"Artist type delete failed", "Артист төрлийн бүртгэлтэй амжилтгүй боллоо"}, err),
		)

		return
	}

	c.JSON(200, utils.Success([]string{"Artist type deleted", "Артист төрлийн бүртгэлтэй амжилттай боллоо"}, struct{}{}))
}
