package application

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fykzi/go-calculator-api/internal/config"
	"github.com/fykzi/go-calculator-api/pkg/calculator"
	"github.com/fykzi/go-calculator-api/pkg/logger"
)

type App struct {
    Config *config.Config
}

type Request struct {
    Expression string `json:"expression"`
}

type Response struct {
    Result string `json:"result"`
}

type ErrorResponse struct {
    Status int `json:"-"`
    Error string `json:"error"`
}

func New() *App {
    return &App{
        Config: config.LoadConfig(),
    }
}

var log *slog.Logger

func (a *App) RunServer() {
    log = logger.SetupLoger(a.Config.LogLevel)
    log.Info("Setup logger is done") 
    
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/calculate", CalcHandler)

    loggedAndErroreddMux := LoggingMiddleware(ErrorMiddleware(mux))

    serverAddr := fmt.Sprintf("%s:%d", a.Config.Host, a.Config.Port)

    log.Info(fmt.Sprintf("server starts on %s", serverAddr))
    http.ListenAndServe(serverAddr, loggedAndErroreddMux)
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    if r.Method != http.MethodPost {
        response := ErrorResponse{
            Status: http.StatusMethodNotAllowed,
            Error: "Invalid request method",
        }
        SendError(w, response)
        return
    }

    request := new(Request)

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
        response := ErrorResponse{
            Status: http.StatusBadRequest,
            Error: "Bad request",
        }
        SendError(w, response)
		return
	}

	result, err := calculator.Calc(request.Expression)
	if err != nil {
        response := ErrorResponse{
            Status: http.StatusUnprocessableEntity,
            Error: "Expression is not valid",
        }
        SendError(w, response)
        return
	}

    responseJSON, err := json.Marshal(Response{Result: fmt.Sprint(result)})
    if err != nil {
        panic(err)
    }

    w.WriteHeader(http.StatusOK)
    if _, err := w.Write(responseJSON); err != nil {
        panic(err)
    }
}

func SendError(w http.ResponseWriter, response ErrorResponse) { 
    responseJSON, _ := json.Marshal(response)
    w.WriteHeader(response.Status)
    w.Write(responseJSON)
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Info(fmt.Sprintf("Request received %s %s", r.Method, r.URL.Path))
        next.ServeHTTP(w, r)
    })
}

func ErrorMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        defer func(w http.ResponseWriter, r *http.Request) {
            if r := recover(); r != nil {
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
        }(w, r)
        next.ServeHTTP(w, r)
    })
}
