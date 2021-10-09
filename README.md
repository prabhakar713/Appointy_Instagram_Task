

# Appointy <h1>

# Task  | Instagram Backend API <h1>

# Create a user <h1>

```
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"passwd" bson:"passwd"`
}
```
<br>


# Create a Post <h1>
```
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Caption     string        `json:"name" bson:"Caption"`
	Image_url   string        `json:"email" bson:"Image_url"`
	Timestamp   string        `json:"passwd" bson:"Timestamp"`
}
```

# Create a User controller <h2>
```
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
```

<br>

# Create a Post controller <h2>

```
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
```





