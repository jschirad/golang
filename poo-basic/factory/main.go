package main

import "fmt"

// Clase Abstracta - Abstrac Factory - GET & PRINT
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

// Clase Abstracta - Abstrac Factory - GET
type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// PRIMER TIPO DE MENSAJE
type SMSNotification struct {
}

// MENSAJE - SMS
func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

// GET SENDER - SMS
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// SMS SENDER
type SMSNotificationSender struct {
}

// METODO - SMS
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// SEGUNDO TIPO DE MENSAJE
type EmailNotification struct {
}

// MENSAJE - EMAIL
func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// EMAIL SENDER
type EmailNotificationSender struct {
}

// METODO - EMAIL
func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No Notification Type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}
func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
