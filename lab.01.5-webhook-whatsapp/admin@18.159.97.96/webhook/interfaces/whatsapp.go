package interfaces

import "whatsapp_service/entity"

type WhatsappService interface {
	SendNewWhatsAppMessageToUser(msg entity.Message)
	SendNewWhatsAppMessageToGroup(msg entity.Message)
}
