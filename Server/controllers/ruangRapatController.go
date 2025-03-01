package controllers

import (
	"log"
	"net/http"
	"project-its/initializers"
	"project-its/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RuangRapat struct {
	ID     string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()";json:"id"`
	Title  string `json:"title"`
	Start  string `json:"start"`
	End    string `json:"end"`
	AllDay bool   `json:"allDay"`
}

func generateUUID() string {
	return uuid.New().String()
}

// Create a new event
func GetEvents(c *gin.Context) {
	var events []models.RuangRapat
	if err := initializers.DB.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// Example of using generated UUID
func CreateEvent(c *gin.Context) {
	var event models.RuangRapat
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new UUID if not provided
	if event.ID == uuid.Nil {
		event.ID = uuid.New()
	}

	// Set notification based on NotificationType
	setNotification(&event)

	if err := initializers.DB.Create(&event).Error; err != nil {
		log.Printf("Error creating event: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func GetNotifications(c *gin.Context) {
	var notifications []models.Notification
	if err := initializers.DB.Where("is_read = ?", false).Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

func setNotification(event *models.RuangRapat) {
	// Set lokasi ke WIB
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Printf("Error loading location: %v", err)
		return
	}

	// Parse waktu mulai event ke WIB
	startTime, err := time.ParseInLocation(time.RFC3339, event.Start, loc)
	if err != nil {
		log.Printf("Error parsing start time: %v", err)
		return
	}
	log.Printf("Parsed start time in WIB: %v", startTime)

	// Tentukan waktu notifikasi 24 jam sebelum event
	notificationTime24 := startTime.Add(-24 * time.Hour)
	log.Printf("24-hour notification scheduled for %s", notificationTime24)

	// Tentukan waktu notifikasi 1 jam sebelum event
	notificationTime1 := startTime.Add(-1 * time.Hour)
	log.Printf("1-hour notification scheduled for %s", notificationTime1)

	// Simulasi pengiriman notifikasi 24 jam sebelum event
	go func() {
		time.Sleep(time.Until(notificationTime24))
		log.Printf("24-hour notification sent for event %s at %s", event.Title, notificationTime24)
	}()

	// Simulasi pengiriman notifikasi 1 jam sebelum event
	go func() {
		time.Sleep(time.Until(notificationTime1))
		log.Printf("1-hour notification sent for event %s at %s", event.Title, notificationTime1)
	}()

	notification := models.Notification{
		Title: event.Title,
		Start: startTime.Add(-1 * time.Hour), // assuming you want to notify 1 hour before the event starts
	}
	if err := initializers.DB.Create(&notification).Error; err != nil {
		log.Printf("Error creating notification: %v", err)
	}
}


func MarkNotificationRead(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Model(&models.Notification{}).Where("id = ?", id).Update("is_read", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := initializers.DB.Where("id = ?", id).Delete(&models.Notification{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id") // Menggunakan c.Param jika UUID dikirim sebagai bagian dari URL
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus disertakan"})
		return
	}
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := initializers.DB.Where("id = ?", id).Delete(&RuangRapat{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}