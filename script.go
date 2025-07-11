package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

type Repo struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    // LanguagesUrl []string `json:"languages"`
    HasPages    bool   `json:"has_pages"`
}

func reposHandler(w http.ResponseWriter, r *http.Request) {
    res, err := http.Get("https://api.github.com/users/LmarDark/repos")
    if err != nil {
        http.Error(w, "Erro ao buscar repositórios", http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        http.Error(w, "Erro ao ler resposta", http.StatusInternalServerError)
        return
    }

    var repositories []Repo
    err = json.Unmarshal(body, &repositories)
    if err != nil {
        http.Error(w, "Erro ao decodificar JSON", http.StatusInternalServerError)
        return
    }

    var filteredRepos []Repo
    for _, repo := range repositories {
        if repo.HasPages {
            filteredRepos = append(filteredRepos, repo)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(filteredRepos)
}

func main() {
    url := "localhost:"
    port := "8080"

    http.HandleFunc("/api/repos", reposHandler)
    fmt.Printf("Aplicação rodando na url: %s%s\n", url, port)
    log.Fatal(http.ListenAndServe(url+port, nil))
}
