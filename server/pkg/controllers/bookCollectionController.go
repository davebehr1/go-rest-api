package controllers

import (
	"database/sql"
	"encoding/json"
	"lxdAssessmentServer/ent"
	"lxdAssessmentServer/pkg"
	"net/http"
)

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
	json.NewEncoder(w).Encode("api is alive")
}

func (h *Handler) CreateBook(w http.ResponseWriter, req *http.Request) {
	book, err := h.entClient.Book.Create().SetTitle("harry potter").SetDescription("fantasy").SetAuthor("jk").Save(req.Context())
	if err != nil {
		pkg.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.HttpSuccess(w, http.StatusCreated, book)
}
