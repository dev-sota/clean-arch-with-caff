package icontroller

import "context"

type Administrator interface {
	Post(c context.Context)
	Login(c context.Context)
}
