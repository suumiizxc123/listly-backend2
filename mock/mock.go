package mock

import (
	"fmt"
	"kcloudb1/internal/config"
	"kcloudb1/internal/models/org"
	"kcloudb1/internal/models/user"
	"time"

	"github.com/google/uuid"
)

func main() {

	config.ConnectDatabase()

}

func mockSysUser() error {

	var sysUser user.SysUser

	sysUser.UID = uuid.New().String()
	sysUser.RoleID = 1
	sysUser.FirstName = "admin"
	sysUser.LastName = "admin"
	sysUser.Email = "admin@gmail.com"
	sysUser.Phone = "99119911"
	sysUser.Password = "admin"
	sysUser.IsActive = 1
	sysUser.CreatedAt = time.Now()

	return sysUser.Create()
}

func mockOrgAndUser() string {
	var i int
	var errs = ""
	for i = 0; i < 20; i++ {

		var orgAndUser org.OrgAndUserInput

		orgAndUser.KaraokeName = "Karaoke " + fmt.Sprintf("%d", i)
		orgAndUser.Address = "Address " + fmt.Sprintf("%d", i)
		orgAndUser.Picture = "https://s3-eu-west-1.amazonaws.com/prod-ecs-service-web-blog-media/2022/12/tuesday-night-karaoke-room.jpg "
		orgAndUser.Latitude = "10.0"
		orgAndUser.Longitude = "10.0"
		orgAndUser.PhoneNumber = fmt.Sprintf("%d", i+99767870)

		orgAndUser.FirstName = "First Name " + fmt.Sprintf("%d", i)
		orgAndUser.LastName = "Last Name " + fmt.Sprintf("%d", i)
		orgAndUser.Email = "user" + fmt.Sprintf("%d", i) + "@gmail.com"
		orgAndUser.Phone = fmt.Sprintf("%d", i+99767870)
		orgAndUser.Password = fmt.Sprintf("test%v", i)

		_, err := orgAndUser.Create()

		if err != nil {
			errs = errs + err.Error() + "\n"
		}
	}

	return errs

}
