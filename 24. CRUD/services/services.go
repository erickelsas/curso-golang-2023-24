package services

import (
	"crud/bd"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func createHash(key string) string {
	// iniciando o modulo de md5
	hasher := md5.New()
	// transformando a string para byte e escrevendo o hash
	hasher.Write([]byte(key))
	// retornando o hash em md5
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var user user

	if erro = json.Unmarshal(body, &user); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para Struct"))
		return
	}
	user.Password = createHash(user.Password)

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO user (name, email, password) values (?, ?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(user.Name, user.Email, user.Password)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o Banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM user")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários!"))
		return
	}
	defer linhas.Close()

	var users []user
	for linhas.Next() {
		var user user
		if erro := linhas.Scan(&user.ID, &user.Email, &user.Password, &user.Name); erro != nil {
			w.Write([]byte("Erro ao Scanear usuário!"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON!"))
		return
	}
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro!"))
		return
	}

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}

	linha, erro := db.Query("SELECT * FROM USER WHERE id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var user user
	if linha.Next() {
		if erro := linha.Scan(&user.ID, &user.Email, &user.Password, &user.Name); erro != nil {
			w.Write([]byte("Erro ao Scanear usuário!"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("Erro ao converter usuário para JSON!"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro!"))
		return
	}

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var user user
	if erro = json.Unmarshal(body, &user); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para Struct"))
		return
	}
	user.Password = createHash(user.Password)

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE user SET name = ?, email = ?, password = ? WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Email, user.Password, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM user WHERE ID = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
