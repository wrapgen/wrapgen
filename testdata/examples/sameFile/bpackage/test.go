package bpackage

import "context"

type AdminUser struct {
	AdminID  int32
	UserName string
}

//wrapgen:generate -template gomock -package bpackage_test -destination gen_test.go -name AuthAdapter
type Adapter interface {
	BTest(ctx context.Context) (AdminUser, error)
}
