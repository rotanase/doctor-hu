package main

import (
	"encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	// "time"
)

type User struct {
	Id        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
	Sex       string `json:"sex, omitempty"`
	Diseases  string `json:"diseases,omitempty"`
	// TODO
	// Questions []*Question `json:"questions,omitempty"`
	// Responses []*Response `json:"responses,omitempty"`
}

type Doctor struct {
	Id         string `json:"id,omitempty"`
	Firstname  string `json:"firstname,omitempty"`
	Lastname   string `json:"lastname,omitempty"`
	Email      string `json:"email,omitempty"`
	Age        int    `json:"age,omitempty"`
	Sex        string `json:"sex, omitempty"`
	University string `json:"university,omitempty"`
	Specialty  string `json:"specialty, omitempty"`
	Title      string `json:"title, omitempty"`
	// TODO
	// Responses []*Response `json:"responses,omitempty"`
}

type Response struct {
	Id          string `json:"id,omitempty"`
	User_id     string `json:"user_id,omitempty"`
	Doctor_id   string `json:"doctor_id,omitempty"`
	Question_id string `json:"question_id,omitempty"`
	Content     string `json:"content,omitempty"`
	// TODO
	// Date_of_creation time   `json:"date_of_creation,omitempty"`
}

type Question struct {
	Id      string `json:"id,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Content string `json:"content,omitempty"`
	// TODO
	// Date_of_creation time   `json:"date_of_creation,omitempty"`
}

var users []User
var doctors []Doctor
var questions []Question
var responses []Response

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Id = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.Id == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(users)
	}
}

func GetDoctors(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(doctors)
}

func GetDoctor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range doctors {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Doctor{})
}

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var doctor Doctor
	_ = json.NewDecoder(r.Body).Decode(&doctor)
	doctor.Id = params["id"]
	doctors = append(doctors, doctor)
	json.NewEncoder(w).Encode(doctors)
}

func DeleteDoctor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range doctors {
		if item.Id == params["id"] {
			doctors = append(doctors[:index], doctors[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(doctors)
	}
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(questions)
}

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range questions {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Question{})
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var question Question
	_ = json.NewDecoder(r.Body).Decode(&question)
	question.Id = params["id"]
	questions = append(questions, question)
	json.NewEncoder(w).Encode(questions)
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range questions {
		if item.Id == params["id"] {
			questions = append(questions[:index], questions[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(questions)
	}
}

func GetResponses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(responses)
}

func GetResponse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range responses {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Response{})
}

func CreateResponse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var response Response
	_ = json.NewDecoder(r.Body).Decode(&response)
	response.Id = params["id"]
	responses = append(responses, response)
	json.NewEncoder(w).Encode(responses)
}

func DeleteResponse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range responses {
		if item.Id == params["id"] {
			responses = append(responses[:index], responses[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(responses)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/doctors", GetDoctors).Methods("GET")
	router.HandleFunc("/doctors/{id}", GetDoctor).Methods("GET")
	router.HandleFunc("/doctors/{id}", CreateDoctor).Methods("POST")
	router.HandleFunc("/doctors/{id}", DeleteDoctor).Methods("DELETE")

	router.HandleFunc("/questions", GetQuestions).Methods("GET")
	router.HandleFunc("/questions/{id}", GetQuestion).Methods("GET")
	router.HandleFunc("/questions/{id}", CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", DeleteQuestion).Methods("DELETE")

	router.HandleFunc("/responses", GetResponses).Methods("GET")
	router.HandleFunc("/responses/{id}", GetResponse).Methods("GET")
	router.HandleFunc("/responses/{id}", CreateResponse).Methods("POST")
	router.HandleFunc("/responses/{id}", DeleteResponse).Methods("DELETE")

	users = append(users, User{Id: "1", Firstname: "John", Lastname: "Doe", Email: "john.doe@ymail.com", Age: 24, Sex: "male", Diseases: "me doare in gat"})
	users = append(users, User{Id: "2", Firstname: "Koko", Lastname: "Doe", Email: "koko.doe@gmail.com", Age: 23, Sex: "femele", Diseases: "me supera barbatu"})
	users = append(users, User{Id: "3", Firstname: "Francis", Lastname: "Sunday", Email: "fsunday@gmail.com", Age: 59, Sex: "male", Diseases: "me doare spatele"})

	doctors = append(doctors, Doctor{Id: "1", Firstname: "Andrew", Lastname: "Tanenbaum", Email: "a.tanenbaum@minix.net", Age: 44, University: "Carolina Davila", Specialty: "chirurg", Title: "domnu'"})

	questions = append(questions, Question{Id: "1", User_id: "1", Content: "Ma doare in gat ca aseara prea multa bere rece. Ce ma fac?"})
	questions = append(questions, Question{Id: "2", User_id: "2", Content: "Barbatul nu ma mai satisface :(. E vina mea?"})

	responses = append(responses, Response{Id: "1", Doctor_id: "1", Question_id: "1", Content: "Un coldrex si o frecie cu spirt si esti gigiuc!"})
	responses = append(responses, Response{Id: "2", Doctor_id: "1", Question_id: "2", Content: "Vine si el obosit acasa si tu il bati la cap. Da, e vina ta!"})
	responses = append(responses, Response{Id: "2", User_id: "1", Question_id: "2", Content: "Koko, cum adica nu te mai satisfac!?!?!"})

	log.Fatal(http.ListenAndServe(":8080", router))
}
