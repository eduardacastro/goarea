package goarea

import "math"

//pi é uma variavel publica
const Pi = 3.1416

//responsavel por calcular a area da circunferencia
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//responsavel por calcular a area de um retangulo
func Retangulo(base, altura float64) float64 {
	return base * altura
}

//Nao é visivel
func _TrianguloEquilatero(base, altura float64) float64 {
	return (base * altura) / 2
}
