package main

import "github.com/fykzi/go-calculator-api/internal/application"

func main() {
    app := application.New()

    app.RunServer()
}
