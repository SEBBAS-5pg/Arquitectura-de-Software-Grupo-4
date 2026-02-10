package main

import (
	"fmt"
	"net/http"
)

// Interfaz Gigante: Obliga a todos a hacer todo
type CRUDWorker interface {
	CreateReport() string
	EatLunch() string
}

type RobotWorker struct{}

func (r RobotWorker) CreateReport() string { return "Reporte generado por IA" }
func (r RobotWorker) EatLunch() string {
	panic("Error de Hardware: Los robots no consumen materia orgánica")
}

type HumanWorker struct{}

func (h HumanWorker) CreateReport() string { return "Reporte escrito a mano" }
func (h HumanWorker) EatLunch() string     { return "Almorzando ensalada" }

func handleBadRobot(w http.ResponseWriter, r *http.Request) {
	var worker CRUDWorker = RobotWorker{}
	// Esto causará que el servidor se caiga (Panic) porque obligamos al robot a comer
	fmt.Fprintf(w, "Acción: %s", worker.EatLunch())
}

func main() {
	http.HandleFunc("/bad/robot-eat", handleBadRobot)
	fmt.Println("❌ Servidor SIN ISP corriendo en :8081")
	http.ListenAndServe(":8081", nil)
}
