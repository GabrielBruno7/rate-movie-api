package errs

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

const (
	ErrInvalidCredentials = "4002"
	ErrUserNotFound       = "4003"
	ErrInvalidBody        = "4001"
	ErrInternalServer     = "5000"
)

var errorMessages = map[string]string{
	ErrInvalidCredentials: "Credenciais inválidas",
	ErrUserNotFound:       "Usuário não encontrado",
	ErrInvalidBody:        "Dados inválidos",
	ErrInternalServer:     "Ocorreu um erro inesperado",
}

type DomainError struct {
	Code       string // algo tipo: AUTH_INVALID_CREDENTIALS
	Message    string // mensagem segura
	Err        error  // mensagem real (para logs)
	StackTrace string // stacktrace (só em dev)
}

func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s", e.Code, e.Err.Error())
	}
	return e.Code
}

func (e *DomainError) GetDetails() map[string]interface{} {
	details := map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}

	env := os.Getenv("ENV")
	if env == "" || env == "development" || env == "dev" {
		if e.StackTrace != "" {
			details["stackTrace"] = e.getFilteredStackTrace()
		}
	}

	return details
}

func (e *DomainError) getFilteredStackTrace() []string {
	lines := strings.Split(e.StackTrace, "\n")
	var filtered []string

	for i, line := range lines {
		// Pega apenas linhas que contém nosso código (crud/)
		if strings.Contains(line, "crud/") {
			// Adiciona a linha atual e a próxima (que tem o arquivo:linha)
			filtered = append(filtered, strings.TrimSpace(line))
			if i+1 < len(lines) {
				filtered = append(filtered, strings.TrimSpace(lines[i+1]))
			}
		}
	}

	return filtered
}

func New(code, message string, err error) *DomainError {
	return &DomainError{
		Code:       code,
		Message:    message,
		Err:        err,
		StackTrace: string(debug.Stack()),
	}
}

func NewWithCode(code string, err error) *DomainError {
	message := errorMessages[code]
	if message == "" {
		message = "Erro desconhecido"
	}

	return &DomainError{
		Code:       code,
		Message:    message,
		Err:        err,
		StackTrace: string(debug.Stack()),
	}
}
