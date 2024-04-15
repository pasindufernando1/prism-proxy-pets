package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pet struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner,omitempty"`
}

var pets = []pet{
	{ID: 1, Name: "Elephant", Owner: "Jack"},
	{ID: 2, Name: "Cat", Owner: "Jane"},
	{ID: 3, Name: "Rabbit", Owner: "Jim"},
	{ID: 4, Name: "Hamster", Owner: "Tom"},
	{ID: 5, Name: "Fish", Owner: "Paul"},
}

func getPets(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, pets)
}

func addPets(context *gin.Context) {
	var newPet pet

	if err := context.BindJSON(&newPet); err != nil {
		return
	}
	pets = append(pets, newPet)
	context.IndentedJSON(http.StatusCreated, newPet)
}

func getPetById(id int) (*pet, error) {
	for i, t := range pets {
		if t.ID == id {
			return &pets[i], nil
		}
	}

	return nil, errors.New("Pet not found")
}

func getPet(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	pet, err := getPetById(idInt)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pet not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, pet)
}

func getPetOwner(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	pet, err := getPetById(idInt)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pet not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{
		"id":    pet.ID,
		"owner": pet.Owner})
}

func handleNotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not Implemented"})
}

func main() {
	router := gin.Default()
	// Implemented routes
	router.GET("/pets", getPets)
	router.POST("/pets", addPets)
	router.GET("/pets/:id", getPet)
	// router.GET("/pets/:id/owner",getPetOwner)

	router.NoRoute(handleNotImplemented)
	router.Run("localhost:8080")

}
