package utils

import (
	"log"
)

func AuditUserFailedAction(action string, message string) {
	log.Printf("AuditUserFailedAction: action: %s, message: %s", action, message)
}

func AuditUserAction(action string, userID int) {
	log.Printf("AuditUserAction: action: %s, userID: %d", action, userID)
}
