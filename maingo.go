package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Estrutura para representar um usuário
type Usuario struct {
    CPF            int    `json:"cpf"`
    Nome           string `json:"nome"`
    DataNascimento string `json:"data_nascimento"`
}

// Simulando um banco de dados em memória
var usuarios []Usuario

// Rota para obter todos os contatos
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usuarios)
}

// Rota para obter um usuário específico
func GetUsuario(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, usuario := range usuarios {
        if usuario.CPF == params["cpf"] {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(usuario)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
}

// Rota para criar um novo usuário
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
    var usuario Usuario
    _ = json.NewDecoder(r.Body).Decode(&usuario)
    usuarios = append(usuarios, usuario)
    w.WriteHeader(http.StatusCreated)
}

// Rota para excluir um usuário
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for i, usuario := range usuarios {
        if usuario.CPF == params["cpf"] {
            usuarios = append(usuarios[:i], usuarios[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    w.WriteHeader(http.StatusNotFound)
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/usuarios", GetUsuarios).Methods("GET")
    router.HandleFunc("/usuarios/{cpf}", GetUsuario).Methods("GET")
    router.HandleFunc("/usuarios/{cpf}", CreateUsuario).Methods("POST")
    router.HandleFunc("/usuarios/{cpf}", DeleteUsuario).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}
