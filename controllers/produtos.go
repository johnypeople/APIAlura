package controllers

import (
	"alura/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, erro := strconv.ParseFloat(preco, 64)
		if erro != nil {
			log.Println("Erro na conversão de preço para Float")
		}

		quantidadeConvertido, erro := strconv.Atoi(quantidade)
		if erro != nil {
			log.Println("Erro na conversão de quantidade para int")
		}

		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, erro := strconv.Atoi(id)
		if erro != nil {
			log.Println("Erro na conversão do ID para int:", erro)
		}

		precoConvertido, erro := strconv.ParseFloat(preco, 64)
		if erro != nil {
			log.Println("Erro na conversão do preço para float64:", erro)
		}

		quantidadeConvertido, erro := strconv.Atoi(quantidade)
		if erro != nil {
			log.Println("Erro na conversão da quantidade para int:", erro)
		}

		models.AtualizaProduto(idConvertido, quantidadeConvertido, nome, descricao, precoConvertido)

	}
	http.Redirect(w, r, "/", 301)
}
