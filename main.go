package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
    cpf, err := strconv.Atoi(params["cpf"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    for _, usuario := range usuarios {
        if usuario.CPF == cpf {
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
    cpf, err := strconv.Atoi(params["cpf"])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    for i, usuario := range usuarios {
        if usuario.CPF == cpf {
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
    router.HandleFunc("/usuarios", CreateUsuario).Methods("POST")
    router.HandleFunc("/usuarios/{cpf}", DeleteUsuario).Methods("DELETE")
    fmt.Printf("Porta 8000 pronta para ser usada\n")
    log.Fatal(http.ListenAndServe(":8000", router))
}
