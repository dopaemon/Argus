package db

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

type Metrics struct {
	ClientIP string `gorm:"primaryKey"`

	CPUName       string
	LogicalCore   string
	PhysicalCore  string
	CPUUsage      string

	TotalRAM string
	UsedRAM  string
	FreeRAM  string
	RAMUsage string

	DiskTotal string
	DiskUsed  string
	DiskFree  string
	DiskUsage string

	Inbound    string
	Outbound   string
	PacketsIn  string
	PacketsOut string

	Hostname      string
	OS            string
	Platform      string
	KernelVersion string
	Uptime        string
	BootTime      string

	UpdatedAt time.Time
}

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"uniqueIndex"`
	Password string
	APIKey   string
}

func generateAPIKey() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic("failed to generate APIKey: " + err.Error())
	}
	return hex.EncodeToString(b)
}

func InitDB() {
	var err error

	dataHome := os.Getenv("XDG_DATA_HOME")
	var home string
	if dataHome == "" {
		home, err = os.UserHomeDir()
		if err != nil {
			log.Fatal("cannot find home dir:", err)
		}
		dataHome = filepath.Join(home, ".local", "share")
	}

	dbDir := filepath.Join(dataHome, "artus")
	dbPath := filepath.Join(dbDir, "database.db")

	if err := os.MkdirAll(dbDir, 0755); err != nil {
		log.Fatal("failed to create db folder:", err)
	}

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	if err := DB.AutoMigrate(&Metrics{}, &User{}); err != nil {
		log.Fatal("failed to migrate schema: ", err)
	}

	DB.Exec("PRAGMA journal_mode = WAL;")

	log.Println("SQLite DB initialized at:", dbPath)
}

func SaveMetrics(metric *Metrics) error {
	metric.UpdatedAt = time.Now()

	return DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "client_ip"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"cpu_name", "logical_core", "physical_core", "cpu_usage",
			"total_ram", "used_ram", "free_ram", "ram_usage",
			"disk_total", "disk_used", "disk_free", "disk_usage",
			"inbound", "outbound", "packets_in", "packets_out",
			"hostname", "os", "platform", "kernel_version",
			"uptime", "boot_time", "updated_at",
		}),
	}).Create(metric).Error
}

func GetAllMetrics() ([]Metrics, error) {
	var metrics []Metrics
	err := DB.Find(&metrics).Error
	return metrics, err
}

func GetMetricsByIP(ip string) (*Metrics, error) {
	var metric Metrics
	err := DB.First(&metric, "client_ip = ?", ip).Error
	if err != nil {
		return nil, err
	}
	return &metric, nil
}

func CreateUser(username, plainPassword string) (*User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		Username: username,
		Password: string(hashed),
		APIKey:   generateAPIKey(),
	}
	if err := DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func Authenticate(username, plainPassword string) bool {
	var user User
	if err := DB.First(&user, "username = ?", username).Error; err != nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainPassword)); err != nil {
		return false
	}
	return true
}

func GetAllUsers() ([]User, error) {
	var users []User
	err := DB.Find(&users).Error
	return users, err
}

func ChangePassword(username, newPlainPassword string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPlainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return DB.Model(&User{}).Where("username = ?", username).Update("password", string(hashed)).Error
}
