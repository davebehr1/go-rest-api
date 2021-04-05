package controllers

import (
	"database/sql"
	"encoding/json"
	"lxdAssessmentServer/ent"
	"lxdAssessmentServer/ent/book"
	"lxdAssessmentServer/ent/collection"
	"lxdAssessmentServer/pkg"
	"net/http"
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
	b, _ := h.entClient.Book.Query().Where(book.IDEQ(1)).Only(req.Context())
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

func (h *Handler) DeleteCollection(w http.ResponseWriter, req *http.Request) {
	c, _ := h.entClient.Collection.Query().Where(collection.IDEQ(1)).Only(req.Context())
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

func (h *Handler) GetBooks(w http.ResponseWriter, req *http.Request) {

	books, err := h.entClient.Book.Query().WithCollection().All(req.Context())

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

	collections, err := h.entClient.Collection.Query().All(req.Context())

	if err != nil {
		pkg.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.HttpSuccess(w, http.StatusOK, collections)

}
