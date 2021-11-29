package message

import (
	"context"
	"github.com/jucardi/go-titan/components/mongo"
	"github.com/jucardi/go-titan/errors"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	// CollectionName is the name of the collection in mongo this repository uses
	CollectionName = "messages"
)

var (
	// To ensure *IRepository implements IRepository on build
	_ IRepository = (*repository)(nil)
)

type repository struct {
	mongo.IClient
	ctx context.Context
}

func (r *repository) init() *repository {
	r.IClient = mongo.Get()
	return r
}

func (r *repository) WithCtx(ctx context.Context) IRepository {
	return &repository{
		IClient: r.IClient,
		ctx:     ctx,
	}
}
func (r *repository) Create(info *MessageDbe) error {
	session, db, err := r.DB()
	if err != nil {
		return err
	}
	defer session.EndSession(r.context())

	info.ID = mongo.NewObjectId()
	_, err = db.Collection(CollectionName).InsertOne(r.context(), info)
	return err
}

func (r *repository) Update(info *MessageDbe) error {
	session, db, err := r.DB()
	if err != nil {
		return err
	}
	defer session.EndSession(r.context())

	query := bson.M{"name": info.Name}
	_, err = db.Collection(CollectionName).UpdateOne(r.context(), query, info)
	return err
}

func (r *repository) Delete(name string) error {
	session, db, err := r.DB()
	if err != nil {
		return err
	}
	defer session.EndSession(r.context())

	query := bson.M{"name": name}
	_, err = db.Collection(CollectionName).DeleteOne(r.context(), query)
	return err
}

func (r *repository) First(name string) (ret *MessageDbe, err error) {
	session, db, err := r.DB()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(r.context())

	query := bson.M{"name": name}
	cursor, err := db.Collection(CollectionName).Find(r.context(), query)
	if err != nil {
		return nil, err
	}

	if cursor.Next(r.context()) {
		err = cursor.Decode(&ret)
	} else {
		err = errors.New("not found")
	}

	return
}

func (r *repository) context() context.Context {
	if r.ctx == nil {
		return context.Background()
	}
	return r.ctx
}
