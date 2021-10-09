package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/prabhakar/Appointy_instagram/models"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("Appointy").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content_type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Body)
	path, err := url.Parse(r.URL.Path)
	id := path.Path

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.session.DB("Appointy").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

type PostController struct {
	session *mgo.Session
}

func NewPostController(s *mgo.Session) *PostController {
	return &PostController{s}
}

func (po PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	p := models.Post{}

	json.NewDecoder(r.Body).Decode(&p)

	p.Id = bson.NewObjectId()

	p.session.DB("Appointy").C("Post").Insert(p)

	pos, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content_type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", pos)
}

func (po PostController) GetUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Body)
	path, err := url.Parse(r.URL.Path)
	id := path.Path

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := po.session.DB("Appointy").C("Post").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	pos, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", pos)
}
