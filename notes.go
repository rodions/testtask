package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
)

// Notes - a structure for transfer between json and bytes.
type Notes struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	DateCreate int64  `json:"date_create"`
	DateUpdate int64  `json:"date_update"`
}

func main() {
	e := echo.New()
	e.POST("/notes/", addNotes)
	e.GET("/notes/:id", getNotes)
	e.GET("/notes/", getNotes)
	e.PUT("/notes/:id", modifyNote)
	e.Logger.Fatal(e.Start(":1323"))
}

func addNotes(c echo.Context) error {
	note := new(Notes)
	if err := c.Bind(note); err != nil {
		return err
	}
	ID := fmt.Sprintf("%s", uuid.NewV4())

	note.DateCreate = time.Now().Unix()
	note.DateUpdate = 0
	note.ID = ID

	db, err := bolt.Open("notes.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Notes"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b := tx.Bucket([]byte("Notes"))
		r, _ := json.Marshal(note)
		err = b.Put([]byte(ID), r)
		if err != nil {
			return fmt.Errorf("added record: %s", err)
		}
		return err
	})
	db.Close()
	return c.JSON(http.StatusCreated, note)
}

func getNotes(c echo.Context) error {
	id := c.Param("id")
	db, err := bolt.Open("notes.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var result []Notes

	if id != "" {
		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Notes"))
			v := b.Get([]byte(id))
			ns := new(Notes)
			json.Unmarshal(v, ns)
			if ns.ID != "" {
				result = append(result, *ns)
			}
			return nil
		})

	} else {
		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Notes"))
			b.ForEach(func(k, v []byte) error {
				ns := new(Notes)
				json.Unmarshal(v, ns)
				if ns.ID != "" {
					result = append(result, *ns)
				}
				return nil
			})
			return nil
		})
	}
	db.Close()
	if result == nil {
		return c.JSON(http.StatusNoContent, nil)
	}
	return c.JSON(http.StatusOK, result)
}

func modifyNote(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	note := new(Notes)
	if err := c.Bind(note); err != nil {
		return err
	}
	note.DateUpdate = time.Now().Unix()

	db, err := bolt.Open("notes.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Notes"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		b := tx.Bucket([]byte("Notes"))
		v := b.Get([]byte(id))
		ns := new(Notes)
		json.Unmarshal(v, ns)
		if ns.ID == "" {
			return nil
		}
		note.ID = ns.ID
		note.DateCreate = ns.DateCreate
		r, _ := json.Marshal(note)
		err = b.Put([]byte(fmt.Sprintf("%s", id)), r)
		if err != nil {
			return fmt.Errorf("added record: %s", err)
		}
		return err
	})
	db.Close()
	if note.ID == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, note)
}
