package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Task struct {
	TaskID      string `json:"task_id"`     // Unique ID for the task
	Title       string `json:"title"`       // Task title
	Description string `json:"description"` // Task description
	Priority    string `json:"priority"`    // Task priority (e.g., High, Medium, Low)
	Status      string `json:"status"`      // Task status (e.g., Pending, Completed)
	DueDate     string `json:"due_date"`    // Task due date in ISO8601 format
}

var (
	db *dynamodb.DynamoDB
)

func main() {
	// Initialize AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8002"),
	})
	if err != nil {
		log.Fatalf("Failed to initialize AWS session: %v", err)
	}
	db = dynamodb.New(sess)

	// http.HandleFunc("/tasks", handleTasks)
	// http.HandleFunc("/tasks/", handleTaskByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createTask(w, r)
	} else if r.Method == http.MethodGet {
		listTasks(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	task.TaskID = fmt.Sprintf("task-%d", time.Now().UnixNano()) // Generate a unique TaskID

	item, err := dynamodbattribute.MarshalMap(task)
	if err != nil {
		http.Error(w, "Error marshalling task", http.StatusInternalServerError)
		return
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Tasks"),
		Item:      item,
	}

	_, err = db.PutItem(input)
	if err != nil {
		http.Error(w, "Error saving task to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Tasks"),
	}

	result, err := db.Scan(input)
	if err != nil {
		http.Error(w, "Error fetching tasks from database", http.StatusInternalServerError)
		return
	}

	var tasks []Task
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &tasks); err != nil {
		http.Error(w, "Error unmarshalling tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
