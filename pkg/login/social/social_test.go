package social

import (
	"fmt"
	"testing"

	"github.com/grafana/grafana/pkg/services/auth"
	. "github.com/smartystreets/goconvey/convey"
)

func newTrue() *bool {
	b := true
	return &b
}

func TestCreateOrganizationMapping(t *testing.T) {
	Convey("Given a static mapping create a SocialGroup array", t, func() {

		authConfig := auth.AuthConfig{
			AuthMappings: []*auth.AuthOrgConfig{
				&auth.AuthOrgConfig{
					Groups: []*auth.GroupToOrgRole{
						&auth.GroupToOrgRole{
							GroupDN: "dn-my-group",
							OrgID:   1,
							OrgRole: "Viewer",
						},

						&auth.GroupToOrgRole{
							GroupDN:        "admin-group",
							OrgID:          1,
							IsGrafanaAdmin: newTrue(),
							OrgRole:        "Admin",
						},
					},
				},
			},
		}

		orgMap := createOrganizationMapping(&authConfig)
		for _, authMap := range authConfig.AuthMappings {
			for _, group := range authMap.Groups {

				Convey(fmt.Sprintf("Groups %s exists in the OrgMap", group.GroupDN), func() {
					grpArray, exists := orgMap[group.GroupDN]
					So(exists, ShouldEqual, true)
					for _, grp := range grpArray {
						So(grp.OrgID, ShouldEqual, 1)
						So(grp.GrafanaAdmin, ShouldEqual, group.IsGrafanaAdmin)
						So(grp.Role, ShouldEqual, group.OrgRole)
					}
				})
			}
		}

	})

}
