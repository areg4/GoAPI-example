package controllers

import (
	"GoAPI/db"
	"GoAPI/types"
)

func GetMovies() ([]types.MovieResponse, error) {
	moviesR := []types.MovieResponse{}
	bd, err := db.GetDB()
	if err != nil {
		return moviesR, err
	}
	results := bd.Table("movies").Scan(&moviesR)
	if results.Error != nil {
		return moviesR, results.Error
	}
	return moviesR, nil
}

func CreateMovies(movie types.MovieResponse) (types.MovieResponse, error) {
	bd, err := db.GetDB()
	if err != nil {
		return movie, err
	}
	result := bd.Table("movies").Create(&movie)
	return movie, result.Error
}

func GetMovieById(id int) (types.MovieResponse, error) {
	movie := types.MovieResponse{}
	bd, err := db.GetDB()
	if err != nil {
		return movie, err
	}
	result := bd.First(&types.Movie{}, id).Scan(&movie)
	return movie, result.Error
}

func UpdateMovieById(id int, m types.Movie) (types.MovieResponse, error) {
	movie := types.MovieResponse{}
	bd, err := db.GetDB()
	if err != nil {
		return movie, err
	}
	result := bd.Table("movies").Where("id = ?", id).Updates(&m).Scan(&movie)

	return movie, result.Error

}

func DeleteMovieById(id int) error {
	bd, err := db.GetDB()
	if err != nil {
		return err
	}
	result := bd.Delete(&types.Movie{}, id)
	return result.Error
}
