package icontroller

import "context"

type User interface {
	Get(c context.Context)
	List(c context.Context)
	Post(c context.Context)
	Put(c context.Context)
	Delete(c context.Context)
}
