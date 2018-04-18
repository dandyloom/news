package dao

import (
	"log"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/dandyloom/news/models"
)

type PostsDAO struct {
	Server		string
	Database	string
}

var db *mgo.Database

const (
	COLLECTION = "posts"
)

func (m *PostsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

func (m *PostsDAO) FindAll() ([]Post, error) {
	var posts []Post
	err := db.C(collection).Find(bson.M{}).All(&posts)
	return movies, err
}

func (m *PostsDAO) FindById(id string) (Post, error) {
	var post Post
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&post)
	return post, err
}

func (m *PostsDAO) Insert(post Post) error {
	err := db.C(COLLECTION).Insert(&post)
	return err
}

func (m *PostsDAO) Delete(post Post) error {
	err := db.C(COLLECTION).Remove(&post)
	return err
}

func (m *PostsDAO) Update(post Post) error {
	err := db.C(COLLECTION).UpdateId(post.ID, &post)
	return err
}