package userservice

import (
	"github.com/ebauman/crder"
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	v2 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v2"
	"github.com/hobbyfarm/gargantua/v3/pkg/crd"
)

// UserCRDInstaller is a struct that can generate CRDs for users.
// It implements the CrdInstallerWithServiceReference interface defined in "github.com/hobbyfarm/gargantua/v3/pkg/microservices"
type UserCRDInstaller struct{}

func (ui UserCRDInstaller) GenerateCRDs(caBundle string, reference crd.ServiceReference) []crder.CRD {
	return []crder.CRD{
		crd.HobbyfarmCRD(&v1.User{}, func(c *crder.CRD) {
			c.
				IsNamespaced(true).
				AddVersion("v1", &v1.User{}, func(cv *crder.Version) {
					cv.
						WithColumn("Email", ".spec.email")

					cv.IsServed(true)
					cv.IsStored(false)
				}).
				AddVersion("v2", &v2.User{}, func(cv *crder.Version) {
					cv.WithColumn("Email", ".spec.email")

					cv.IsServed(true)
					cv.IsStored(true)
				}).
				WithConversion(func(cc *crder.Conversion) {
					cc.
						StrategyWebhook().
						WithCABundle(caBundle).
						WithService(reference.Toapiextv1WithPath("/conversion/users.hobbyfarm.io")).
						WithVersions("v2", "v1")
				})
		}),
	}
}
