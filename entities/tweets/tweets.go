package tweets

import (
	"log"
	db "not-twitter/db"
	"not-twitter/entities/users"
)

type Tweet struct {
	ID      string      `json:"id"`
	Content string      `json:"content"`
	User    *users.User `json:"user"`
}

func (t Tweet) Save() int64 {
	statement, err := db.Db.Prepare("INSERT INTO Tweets(Content, UserID) VALUE (?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := statement.Exec(t.Content, t.User.ID)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("row inserted")
	return id
}

func GetAll() []Tweet {
	statement, err := db.Db.Prepare("SELECT T.ID, T.Content, T.UserID, U.Username from Tweets T inner join Users U on T.UserID = U.ID ")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var tweets []Tweet
	var username, id string

	for rows.Next() {
		var t Tweet
		if err := rows.Scan(&t.ID, &t.Content, &id, &username); err != nil {
			log.Fatal(err)
		}
		t.User = &users.User{
			ID:       id,
			Username: username,
		}
		tweets = append(tweets, t)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tweets

}
