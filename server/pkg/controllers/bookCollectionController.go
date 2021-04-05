package controllers

import (
	"database/sql"
	"encoding/json"
	"lxdAssessmentServer/ent"
	"lxdAssessmentServer/ent/book"
	"lxdAssessmentServer/ent/collection"
	"lxdAssessmentServer/pkg"
	"net/http"
	"strconv"
	"time"
)

type Book struct {
	ID          int
	Title       string
	Description string
	Author      string
	Collection  string
}

type Handler struct {
	entClient *ent.Client
	db        *sql.DB
}

func NewHandler(client *ent.Client, sqlDB *sql.DB) *Handler {
	return &Handler{
		entClient: client,
		db:        sqlDB,
	}
}

func (h *Handler) Test(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("book collections api is alive")
}

func (h *Handler) CreateBook(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	collectionName := params.Get("collection")
	var p ent.Book
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&p)

	collection, _ := h.entClient.Collection.Query().Where(collection.NameEQ(collectionName)).Only(req.Context())
	bookCreator := h.entClient.Book.Create().SetTitle(p.Title).SetDescription(p.Description).SetAuthor(p.Author)

	if collection != nil {

		bookCreator.SetCollection(collection)
	}

	book, err := bookCreator.Save(req.Context())
	if err != nil {
		pkg.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.HttpSuccess(w, http.StatusCreated, book)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	bookId := params.Get("id")
	Id, _ := strconv.Atoi(bookId)

	b, _ := h.entClient.Book.Query().Where(book.IDEQ(Id)).Only(req.Context())
	if b == nil {
		pkg.HttpError(w, http.StatusInternalServerError, "book does not exist")
	}
	_, err := h.entClient.Book.Delete().Where(book.IDEQ(1)).Exec(req.Context())
	if err != nil {
		pkg.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.HttpSuccess(w, http.StatusCreated, "deleted book")

}

func (h *Handler) UpdateBook(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	bookId := params.Get("id")
	Id, _ := strconv.Atoi(bookId)

	var p ent.Book
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&p)

	if bookId != "" {
		bookUpdater := h.entClient.Book.Update().Where(book.IDEQ(Id))
		if p.Author != "" {
			bookUpdater.SetAuthor(p.Author)
		}
		if p.Description != "" {
			bookUpdater.SetAuthor(p.Description)
		}
		if p.Title != "" {
			bookUpdater.SetAuthor(p.Title)
		}

		_, err := bookUpdater.Save(req.Context())

		if err != nil {
			pkg.HttpError(w, http.StatusInternalServerError, err.Error())
			return
		}
		book, err := h.entClient.Book.Query().Where(book.IDEQ(Id)).First(req.Context())
		if err != nil {
			pkg.HttpError(w, http.StatusInternalServerError, err.Error())
			return
		}

		pkg.HttpSuccess(w, http.StatusCreated, book)
	} else {
		pkg.HttpError(w, http.StatusInternalServerError, "book does not exist")
	}

}

func (h *Handler) DeleteCollection(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	collectionId := params.Get("id")
	Id, _ := strconv.Atoi(collectionId)

	c, _ := h.entClient.Collection.Query().Where(collection.IDEQ(Id)).Only(req.Context())
	if c == nil {
		pkg.HttpError(w, http.StatusInternalServerError, "collection does not exist")
	}

	_, err := h.entClient.Collection.Delete().Where(collection.IDEQ(1)).Exec(req.Context())
	if err != nil {
		pkg.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.HttpSuccess(w, http.StatusCreated, "deleted collection")
}

func (h *Handler) CreateCollection(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	collectionName := params.Get("collection")
	if collectionName != "" {
		collection, err := h.entClient.Collection.Create().SetName(collectionName).Save(req.Context())
		if err != nil {
			pkg.HttpError(w, http.StatusInternalServerError, err.Error())
			return
		}
		pkg.HttpSuccess(w, http.StatusCreated, collection)
	} else {
		pkg.HttpError(w, http.StatusInternalServerError, "collection cant be empty")
	}

}

func (h *Handler) UpdateCollection(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	collectionId := params.Get("id")
	collectionName := params.Get("name")
	Id, _ := strconv.Atoi(collectionId)
	if collectionId != "" {
		collectionUpdater := h.entClient.Collection.Update().Where(collection.IDEQ(Id))

		if collectionName != "" {
			collectionUpdater.SetName(collectionName)
		}
		_, err := collectionUpdater.Save(req.Context())
		if err != nil {
			pkg.HttpError(w, http.StatusInternalServerError, err.Error())
			return
		}
		collection, err := h.entClient.Collection.Query().Where(collection.IDEQ(Id)).First(req.Context())
		if err != nil {
			pkg.HttpError(w, http.StatusInternalServerError, err.Error())
			return
		}
		pkg.HttpSuccess(w, http.StatusCreated, collection)
	} else {
		pkg.HttpError(w, http.StatusInternalServerError, "collection does not exist")
	}

}

func (h *Handler) GetBooks(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	author := params.Get("author")
	title := params.Get("title")
	fromDate := params.Get("from_date")
	toDate := params.Get("to_date")
	from, _ := time.Parse("2010-10-13", fromDate)
	to, _ := time.Parse("2010-10-13", toDate)

	booksBuilder := h.entClient.Book.Query()
	if author != "" {
		booksBuilder.Where(book.AuthorEQ(author))
	}
	if title != "" {
		booksBuilder.Where(book.TitleEQ(title))
	}
	if fromDate != "" {
		booksBuilder.Where(book.PublishedAtGTE(from))
	}
	if toDate != "" {
		booksBuilder.Where(book.PublishedAtLTE(to))
	}

	books, err := booksBuilder.WithCollection().All(req.Context())

	if err != nil {
		pkg.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var booksPayload []Book

	for index, book := range books {
		booksPayload = append(booksPayload, Book{
			ID:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
		})
		if book.Edges.Collection != nil {
			booksPayload[index].Collection = book.Edges.Collection.Name
		}

	}

	pkg.HttpSuccess(w, http.StatusOK, booksPayload)

}

func (h *Handler) GetCollections(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	name := params.Get("name")

	collectionBuilder := h.entClient.Collection.Query()

	if name != "" {
		collectionBuilder.Where(collection.NameEQ(name))
	}

	collections, err := collectionBuilder.All(req.Context())

	if err != nil {
		pkg.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.HttpSuccess(w, http.StatusOK, collections)

}
