package whatsapp

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"

	"whatsapp_service/entity"
	"whatsapp_service/interfaces"
)

type WhatsappService struct {
	client *whatsmeow.Client

	cUserMessage,
	cGroupMessage chan entity.Message
}

func New(ctx context.Context, wg *sync.WaitGroup) interfaces.WhatsappService {
	if _, err := os.Stat("data"); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir("data", os.ModePerm); err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
	}

	wa := &WhatsappService{
		cUserMessage:  make(chan entity.Message, 1024),
		cGroupMessage: make(chan entity.Message, 1024),
	}

	wa.setupWhatsappService(ctx, wg)

	return wa
}

func (*WhatsappService) eventHandler(evt any) {
	switch v := evt.(type) {
	case *events.Message:
		log.Println("Received a message: ", v.Message.GetConversation())
	}
}

func (ws *WhatsappService) setupWhatsappService(
	ctx context.Context,
	wg *sync.WaitGroup,
) {
	debug := strings.ToLower(os.Getenv("APP_DEBUG")) == "true"
	level := "INFO"

	if debug {
		level = "DEBUG"
	}

	dbLog := waLog.Stdout("Database", level, true)
	container, err := sqlstore.New("sqlite3", "data/sqlite3.db?_foreign_keys=on", dbLog)
	if err != nil {
		log.Panic(err)
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		log.Panic(err)
	}

	clientLog := waLog.Stdout("Client", level, true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(ws.eventHandler)

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			log.Panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				log.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			log.Panic(err)
		}
	}

	groups, err := client.GetJoinedGroups()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Joined groups:")
	for _, group := range groups {
		log.Println("Name: ", group.Name)
		log.Println("Jid: ", group.JID)
		log.Println("----------------")
	}

	go ws.handelSendUserMessages(ctx)
	go ws.disconnect(ctx, wg)

	ws.client = client
}

func (ws *WhatsappService) SendNewWhatsAppMessageToUser(msg entity.Message) {
	ws.cUserMessage <- msg
}

func (ws *WhatsappService) SendNewWhatsAppMessageToGroup(msg entity.Message) {
	ws.cGroupMessage <- msg
}

func (ws *WhatsappService) handelSendUserMessages(ctx context.Context) {
	for msg := range ws.cUserMessage {
		to := types.NewJID(msg.To, "s.whatsapp.net")
		message := &waE2E.Message{
			ExtendedTextMessage: &waE2E.ExtendedTextMessage{
				Text: &msg.Body,
			},
		}

		_, err := ws.client.SendMessage(ctx, to, message)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (ws *WhatsappService) disconnect(
	ctx context.Context,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	<-ctx.Done()
	close(ws.cUserMessage)
	close(ws.cGroupMessage)

	<-time.After(3 * time.Second)

	ws.client.Disconnect()
}
