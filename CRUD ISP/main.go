package main

import (
	"fmt"
	"net/http"
)

// Interfaces Segregadas: Cada una tiene una sola misión
type Reporter interface {
	CreateReport() string
}

type eater interface {
	EatLunch() string
}

type GoodRobot struct{}

func (r GoodRobot) CreateReport() string { return "Reporte de precisión robótica" }

type GoodHuman struct{}

func (h GoodHuman) CreateReport() string { return "Reporte humano" }
func (h GoodHuman) EatLunch() string     { return "Almorzando pizza" }

// --- HANDLERS ---

func handleGoodReport(w http.ResponseWriter, r *http.Request) {
	tipo := r.URL.Query().Get("tipo")
	var worker Reporter

	if tipo == "robot" {
		worker = GoodRobot{}
	} else {
		worker = GoodHuman{}
	}
	fmt.Fprintf(w, "Resultado: %s", worker.CreateReport())
}

func handleGoodEat(w http.ResponseWriter, r *http.Request) {
	// Aquí ni siquiera podemos asignar un Robot a algo que 'Come'
	// El sistema es seguro por diseño.
	h := GoodHuman{}
	fmt.Fprintf(w, "Resultado: %s", h.EatLunch())
}

func main() {
	http.HandleFunc("/good/report", handleGoodReport)
	http.HandleFunc("/good/eat", handleGoodEat)
	fmt.Println("✅ Servidor CON ISP corriendo en :8082")
	http.ListenAndServe(":8082", nil)
}
