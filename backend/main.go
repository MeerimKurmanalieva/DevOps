package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	_ "github.com/go-sql-driver/mysql"
)

var instanceID string
var db *sql.DB // Declare a global DB connection

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request HelloRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Hello, %s!", request.Name)
	response := HelloResponse{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func createEC2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	fmt.Println("AWS Access Key:", awsAccessKey)
	fmt.Println("AWS Secret Key:", awsSecretKey)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"), // Change to your desired region
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})
	if err != nil {
		http.Error(w, "Failed to create AWS session", http.StatusInternalServerError)
		return
	}

	ec2Svc := ec2.New(sess)

	runInput := &ec2.RunInstancesInput{
		ImageId:          aws.String("ami-03a6eaae9938c858c"),
		InstanceType:     aws.String("t2.micro"),
		MinCount:         aws.Int64(1),
		MaxCount:         aws.Int64(1),
		KeyName:          aws.String("myKey"),
		SecurityGroupIds: []*string{aws.String("sg-0b71386f66ac0b359")}, // Add your security group ID
		SubnetId:         aws.String("subnet-091a5ba500af5a208"),        // Add your VPC ID
	}

	// Create the EC2 instance
	result, err := ec2Svc.RunInstances(runInput)
	if err != nil {
		fmt.Println("Error creating EC2 instance:", err)
		http.Error(w, "Error creating EC2 instance", http.StatusInternalServerError)
		return
	}

	// Store the instance ID
	instanceID = *result.Instances[0].InstanceId

	// Insert the instance ID into the database
	_, err = db.Exec("INSERT INTO ec2_instances (instance_id) VALUES (?)", instanceID)
	if err != nil {
		fmt.Println("Error inserting instance ID into the database:", err)
		// You may want to handle this error gracefully
		http.Error(w, "Error inserting instance ID into the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "EC2 instance creation request received.")
}

func terminateEC2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if instanceID == "" {
		http.Error(w, "No EC2 instance to terminate", http.StatusBadRequest)
		return
	}

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	fmt.Println("AWS Access Key:", awsAccessKey)
	fmt.Println("AWS Secret Key:", awsSecretKey)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"), // Change to your desired region
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})
	if err != nil {
		http.Error(w, "Failed to create AWS session", http.StatusInternalServerError)
		return
	}

	ec2Svc := ec2.New(sess)

	terminateInput := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}

	_, err = ec2Svc.TerminateInstances(terminateInput)
	if err != nil {
		fmt.Println("Error terminating EC2 instance:", err)
		http.Error(w, "Error terminating EC2 instance", http.StatusInternalServerError)
		return
	}

	// Remove the instance ID from the database
	_, err = db.Exec("DELETE FROM ec2_instances WHERE instance_id = ?", instanceID)
	if err != nil {
		fmt.Println("Error removing instance ID from the database:", err)
		// You may want to handle this error gracefully
		http.Error(w, "Error removing instance ID from the database", http.StatusInternalServerError)
		return
	}

	// Clear the stored instance ID
	instanceID = ""

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "EC2 instance termination request received.")
}

func main() {
	// Initialize the MySQL database connection
	var err error
	db, err = sql.Open("mysql", "root:Kurmanalieva93@tcp(127.0.0.1:3306)/awsDB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/ec2/create", createEC2Handler)
	http.HandleFunc("/api/ec2/terminate", terminateEC2Handler) // Add termination endpoint

	port := ":8080"
	if len(os.Args) >= 2 && strings.HasPrefix(os.Args[1], ":") {
		port = os.Args[1]
	}

	fmt.Printf("Server is listening on %s\n", port)
	http.ListenAndServe(port, nil)
}
