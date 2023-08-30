package controllers

import (
	database "bytive-task/dbConn"
	"bytive-task/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var projectCollection *mongo.Collection = database.OpenCollection(database.Client, "project")
var validate = validator.New()

func CreateProject() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var project models.ProjectInfo

		defer cancel()
		if err := c.ShouldBind(&project); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(project)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		project.ID = primitive.NewObjectID()
		project.ProjectID = project.ID.Hex()
		project.BillableTime = CalculateBillableTime(project.Start_Time.Time, project.End_Time.Time)
		project.Hours = CalculateHours(project.Start_Time.Time, project.End_Time.Time)
		project.Time_Spent.Hours = project.Time_Spent.Hours + int(project.End_Time.Time.Sub(project.Start_Time.Time).Hours())
		project.Time_Spent.Minutes = project.Time_Spent.Minutes + int(project.End_Time.Time.Sub(project.Start_Time.Time).Minutes())%60
		if project.BillableTime != "" {
			project.PaymentStatus.Billable = true
		}
		if project.PaymentStatus.Billed {
			project.PaymentStatus.Billable = false
		}
		if project.Who == "my list" {
			project.TaskList = []models.TaskType{{Type: "My List"}}
		}
		result, insertErr := projectCollection.InsertOne(ctx, project)
		if insertErr != nil {
			msg := "Project item was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, result)

	}
}
func GetProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID := c.Param("id")

		var project models.ProjectInfo
		objID, err := primitive.ObjectIDFromHex(projectID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}

		filter := bson.M{"_id": objID}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = projectCollection.FindOne(ctx, filter).Decode(&project)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, project)
	}
}

func GetProjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		var projects []models.ProjectInfo

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cursor, err := projectCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching projects"})
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var project models.ProjectInfo
			err := cursor.Decode(&project)
			if err != nil {
				continue
			}
			projects = append(projects, project)
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, projects)
	}
}

func DeleteProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID := c.Param("id")

		objID, err := primitive.ObjectIDFromHex(projectID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}

		filter := bson.M{"_id": objID}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err = projectCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting project"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
	}
}

func UpdateAllProjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeToAddStr := c.Query("timeToAdd") // Get the time to add from query parameter

		// Parse the time string to a duration
		timeToAdd, parseErr := time.ParseDuration(timeToAddStr)
		if parseErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
			return
		}

		var updatedProjects []models.ProjectInfo

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cursor, err := projectCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching projects"})
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var project models.ProjectInfo
			err := cursor.Decode(&project)
			if err != nil {
				continue
			}

			// Update end time of each project by adding the specified time
			updatedEndTime := project.End_Time.Time.Add(timeToAdd)

			// Calculate updated billable time based on the updated end time
			updatedBillableTime := CalculateBillableTime(project.Start_Time.Time, updatedEndTime)

			// Calculate updated time spent based on the start time and updated end time
			updatedTimeSpent := models.TimeSpent{
				Hours:   project.Time_Spent.Hours + int(updatedEndTime.Sub(project.End_Time.Time).Hours()),
				Minutes: project.Time_Spent.Minutes + int(updatedEndTime.Sub(project.End_Time.Time).Minutes())%60,
			}

			// Update the project's end time, billable time, and time spent
			project.End_Time.Time = updatedEndTime
			project.BillableTime = updatedBillableTime
			project.Time_Spent = updatedTimeSpent

			// Update the project in the database
			update := bson.M{
				"$set": bson.M{
					"end_time":      project.End_Time,
					"billable_time": updatedBillableTime,
					"time_spent":    updatedTimeSpent,
				},
			}
			filter := bson.M{"_id": project.ID}

			_, updateErr := projectCollection.UpdateOne(ctx, filter, update)
			if updateErr != nil {
				continue
			}

			updatedProjects = append(updatedProjects, project)
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, updatedProjects)
	}
}

func CalculateBillableTime(start, end time.Time) string {
	timeDuration := end.Sub(start)
	hours := int(timeDuration.Hours())
	minutes := int(timeDuration.Minutes()) - hours*60
	return fmt.Sprintf("%dh %dm", hours, minutes)
}

func CalculateHours(start, end time.Time) float64 {
	timeDuration := end.Sub(start)
	return timeDuration.Hours()
}
