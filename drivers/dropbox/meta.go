package dropbox

import (
	"github.com/OpenListTeam/OpenList/v4/internal/driver"
	"github.com/OpenListTeam/OpenList/v4/internal/op"
)

type Addition struct {
	driver.RootPath
	UseOnlineAPI    bool   `json:"use_online_api" default:"false"`
	APIAddress      string `json:"api_url_address" default:"https://api.oplist.org/dropboxs/renewapi"`
	ClientID        string `json:"client_id" required:"false" help:"Keep it empty if you don't have one"`
	ClientSecret    string `json:"client_secret" required:"false" help:"Keep it empty if you don't have one"`
	UseAppFolder    bool   `json:"use_app_folder" default:"false" help:"Use app folder instead of root folder"`
	AccessToken     string
	RefreshToken    string `json:"refresh_token" required:"true"`
	RootNamespaceId string
}

var config = driver.Config{
	Name:              "Dropbox",
	NoOverwriteUpload: true,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &Dropbox{
			base:        "https://api.dropboxapi.com",
			contentBase: "https://content.dropboxapi.com",
		}
	})
}
