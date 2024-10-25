    package database

    import (
        "log"
        "gorm.io/driver/postgres"
        "gorm.io/gorm"
        "github.com/ABHINAVGARG05/trading-platform/models"
        "time"
        "math/rand"
    )

    var DB *gorm.DB

    func Connect() {
        dsn := "host=db user=postgres password=new_password dbname=trading_platform port=5432 sslmode=disable"
        var err error
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatal("Failed to connect to database:", err)
        }

        DB.AutoMigrate(&models.User{}, &models.Stock{}, &models.Portfolio{},&models.Order{})

    rand.Seed(time.Now().UnixNano())

    for i := 0; i < 100; i++ {
        stock := models.Stock{
            Symbol:    "STK" + string(i+1),
            Price:     rand.Float64() * 100, 
            Volume:    10000,
            Supply:    uint(rand.Intn(500000)), 
            Demand:    uint(rand.Intn(500000)), 
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        }

        if err := DB.Create(&stock).Error; err != nil {
            log.Println("Failed to insert stock:", err)
        }
    }
    }
