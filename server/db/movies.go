package db

import (
	"context"
	"encoding/json"
	"strconv"
)

const MOVIES = "movies"
const REQUESTS = "requests"

type Movie struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Runtime     string `json:"runtime"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
}

type MovieRequest struct {
	Title       string `json:"title" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description"`
}

func (db *Database) AddMovie(movie *Movie) error {
	var movieList []Movie

	list, _ := db.Client.Get(context.Background(), MOVIES).Result()
	currentList, err := strconv.Unquote(list)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(currentList), &movieList)
	if err != nil {
		return err
	}
	movieList = append(movieList, *movie)

	err = db.Client.Set(context.Background(), MOVIES, movieList, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) AddMovieRequest(request *MovieRequest) error {
	var requestList []MovieRequest

	list, _ := db.Client.Get(context.Background(), REQUESTS).Result()
	currentList, err := strconv.Unquote(list)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(currentList), &requestList)
	if err != nil {
		return err
	}
	requestList = append(requestList, *request)

	err = db.Client.Set(context.Background(), REQUESTS, requestList, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) FindMovieByTitle(title string) error {
	var movieList []Movie
	movies, err := db.Client.Get(context.Background(), MOVIES).Result()
	if err != nil {
		return err
	}

	currentList, _ := strconv.Unquote(movies)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(currentList), &movieList)
	if err != nil {
		return err
	}

	for _, movie := range movieList {
		if movie.Title == title {
			// Implement fuzzy search - WIP
		}
	}

  return nil
}
