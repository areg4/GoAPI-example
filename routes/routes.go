package routes

import (
	"GoAPI/controllers"
	"GoAPI/helpers"
	"GoAPI/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {

	router.HandleFunc("/movie", func(w http.ResponseWriter, r *http.Request) {
		movies, err := controllers.GetMovies()
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusNotFound, "Get Movies error")
			return
		}
		helpers.ResponseWithSuccess(movies, w, http.StatusOK, "movies")
		return
	}).Methods(http.MethodGet)

	router.HandleFunc("/movie/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := helpers.StringToInt(idStr)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusBadRequest, "ID invalid!")
			return
		}
		movie, err := controllers.GetMovieById(id)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusNotFound, "Movie not found!")
			return
		}
		helpers.ResponseWithSuccess(movie, w, http.StatusOK, "Movie")
		return
	}).Methods(http.MethodGet)

	router.HandleFunc("/movie", func(w http.ResponseWriter, r *http.Request) {
		var movie types.MovieResponse
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusBadRequest, "Bad Request")
			return
		}
		movie, err = controllers.CreateMovies(movie)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusInternalServerError, "Cannot save the data!")
			return
		}

		helpers.ResponseWithSuccess(movie, w, http.StatusCreated, "Movie saved!")
		return
	}).Methods(http.MethodPost)

	router.HandleFunc("/movie/{id}", func(w http.ResponseWriter, r *http.Request) {
		var movie types.Movie
		idStr := mux.Vars(r)["id"]
		id, err := helpers.StringToInt(idStr)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusBadRequest, "ID invalid!")
			return
		}

		if _, erro := controllers.GetMovieById(id); erro != nil {
			helpers.ResponseWithError(erro, w, http.StatusNotFound, "Movie not found!")
			return
		}

		err = json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusBadRequest, "Bad request!")
			return
		}

		movieUpdated, err := controllers.UpdateMovieById(id, movie)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusInternalServerError, "Error at updated the record!")
			return
		}
		helpers.ResponseWithSuccess(movieUpdated, w, http.StatusOK, "Data updated!")
		return
	}).Methods(http.MethodPut, http.MethodPatch)

	router.HandleFunc("/movie/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := helpers.StringToInt(idStr)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusBadRequest, "ID invalid!")
			return
		}

		_, err = controllers.GetMovieById(id)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusNotFound, "Movie not found!")
			return
		}

		err = controllers.DeleteMovieById(id)
		if err != nil {
			helpers.ResponseWithError(err, w, http.StatusInternalServerError, "Cannot delete the record!")
			return
		}
		helpers.ResponseWithSuccess(nil, w, 204, "Data deleted!")
		return
	}).Methods(http.MethodDelete)
}
