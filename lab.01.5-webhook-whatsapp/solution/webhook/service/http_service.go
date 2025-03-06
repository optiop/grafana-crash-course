package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
	"whatsapp_service/entity"
	"whatsapp_service/interfaces"
)

var (
	appToken = os.Getenv("APP_TOKEN")
	phone    = os.Getenv("WHATSAPP_PHONE_NUMBER")
)

func init() {
	if phone[0] == '+' {
		phone = phone[1:]
	}
}

func sendNewGrafanaAlertWhatsAppMessage(ws interfaces.WhatsappService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.PathValue("token")
		if token != appToken {
			http.Error(w, "token is required", http.StatusBadRequest)
			return
		}

		var alert GrafanaAlert
		if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
			_, _ = w.Write([]byte("Error decoding alert"))
			return
		}

		if alert.Message == "" {
			http.Error(w, "message is required", http.StatusBadRequest)
			return
		}

		message := entity.Message{
			To:   phone,
			Body: alert.Message,
		}

		ws.SendNewWhatsAppMessageToUser(message)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Message sent to " + message.To))
	}
}

func Run(
	ctx context.Context,
	ws interfaces.WhatsappService,
	wg *sync.WaitGroup,
) {
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("POST /whatsapp/send/grafana-alert/{token}", sendNewGrafanaAlertWhatsAppMessage(ws))

	server := &http.Server{
		Addr:    ":8080",
		Handler: httpMux,
	}

	go func() {
		defer wg.Done()
		fmt.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %v\n", err)
		}
		fmt.Println("HTTP server stopped")
	}()

	go func() {
		defer wg.Done()
		<-ctx.Done()
		fmt.Println("Shutting down server...")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("Error during server shutdown: %v\n", err)
		}
	}()
}
