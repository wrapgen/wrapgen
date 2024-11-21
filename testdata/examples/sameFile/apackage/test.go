package apackage

import "context"

type AdminUser struct {
	AdminID  int32
	UserName string
}

//wrapgen:generate -template gomock -package apackage_test -destination gen_test.go -name AuthAdapter
type Adapter interface {
	ATest(ctx context.Context) (AdminUser, error)
}
