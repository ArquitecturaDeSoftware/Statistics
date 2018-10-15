package service

import (
	"context"

	"github.com/jscastelblancoh/statistic_service/statistic/pkg/db"
	"github.com/jscastelblancoh/statistic_service/statistic/pkg/io"
	"gopkg.in/mgo.v2/bson"
)

// StatisticService describes the service.
type StatisticService interface {
	Get(ctx context.Context) (t []io.Statistic, err error)
	GetbyId(ctx context.Context, id string) (t []io.Statistic, err error)
	Put(ctx context.Context, id string) (t io.Statistic, err error)
	Delete(ctx context.Context, id string) (err error)
	Post(ctx context.Context, statistic io.Statistic) (t io.Statistic, err error)
}

type basicStatisticService struct{}

func (b *basicStatisticService) Get(ctx context.Context) (t []io.Statistic, err error) {
	session, err2 := db.GetMongoSession()
	if err2 != nil {
		return t, err2
	}
	defer session.Close()
	c := session.DB("statistic_service").C("statistics")
	err = c.Find(nil).All(&t)
	return t, err
}
func (b *basicStatisticService) GetbyId(ctx context.Context, id string) (t []io.Statistic, err error) {
	session, err2 := db.GetMongoSession()
	if err2 != nil {
		return t, err2
	}
	defer session.Close()
	c := session.DB("statistic_service").C("statistics")
	err = c.Find(bson.M{"id_restaurant": bson.ObjectIdHex(id)}).All(&t)
	return t, err
}
func (b *basicStatisticService) Put(ctx context.Context, id string) (t io.Statistic, err error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("todo_app").C("todos")
	return t, c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"complete": true}})
}
func (b *basicStatisticService) Delete(ctx context.Context, id string) (err error) {
	session, err := db.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := session.DB("statistic_service").C("statistics")
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func (b *basicStatisticService) Post(ctx context.Context, statistic io.Statistic) (t io.Statistic, err error) {
	statistic.Id = bson.NewObjectId()
	session, err := db.GetMongoSession()
	if err != nil {
		return t, err
	}
	defer session.Close()
	c := session.DB("statistic_service").C("statistics")
	err = c.Insert(&statistic)
	return statistic, err
}

// NewBasicStatisticService returns a naive, stateless implementation of StatisticService.
func NewBasicStatisticService() StatisticService {
	return &basicStatisticService{}
}

// New returns a StatisticService with all of the expected middleware wired in.
func New(middleware []Middleware) StatisticService {
	var svc StatisticService = NewBasicStatisticService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
