package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/config"
	authHandler "github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/auth-service/core/delivery/http"
	authDomain "github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/auth-service/core/domain"
	authRepository "github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/auth-service/core/repository"
	authUsecase "github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/auth-service/core/usecase"
	productHandler "github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/product-service/core/delivery/http"
	productDomain "github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/product-service/core/domain"
	productRepository "github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/product-service/core/repository"
	productUseCase "github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/internal/product-service/core/usecase"
	"github.com/Acad600-TPA/WEB-EJ-NH-JR-KO-WA-261/backend/services/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewRedisClient() redis.UniversalClient {
	password := config.GetEnv("REDIS_PASS", "")
	if addrs := config.GetEnv("REDIS_ADDRS", ""); addrs != "" {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    strings.Split(addrs, ","),
			Password: password,
		})
	}
	return redis.NewClient(&redis.Options{
		Password: password,
		Addr:     fmt.Sprintf("%s:%s", config.GetEnv("REDIS_HOST", "localhost"), config.GetEnv("REDIS_PORT", "6379")),
	})
}

func main() {
	cfg := config.Load()
	var db *gorm.DB
	var err error
	dsn := "host=" + cfg.DBHOST + " user=" + cfg.DBUSER + " password=" + cfg.DBPASS + " dbname=" + cfg.DBNAME + " port=" + cfg.DBPORT + " sslmode=disable"
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
	}

	r := gin.Default()

	// User DB Migration
	if err := db.AutoMigrate(&authDomain.User{}, &productDomain.IceCream{}); err != nil {
		log.Fatalf("gagal migrasi database: %v", err)
	}

	cookieSecure := config.GetEnv("COOKIE_SECURE", "false") == "true"
	redis := NewRedisClient()
	jwtMgr := jwt.NewManager(config.GetEnv("JWT_SECRET", ""))

	// Auth Service
	authCacheRepo := authRepository.NewCacheRepository(redis)
	userRepo := authRepository.NewUserRepository(db)
	authUC := authUsecase.NewAuthUsecase(userRepo, authCacheRepo, *jwtMgr)
	authHand := authHandler.NewAuthHandler(authUC, cookieSecure)
	authHandler.RegisterRouter(r, *authHand, &gin.Context{}, jwtMgr)

	// Product Service
	productRepo := productRepository.NewIceCreamRepository(db)
	productUC := productUseCase.NewIceCreamUsecase(productRepo, cfg)
	productHand := productHandler.NewProductHandler(productUC, cookieSecure)
	productHandler.RegisterRouter(r, *productHand, &gin.Context{}, jwtMgr)

	if err := r.Run(":" + cfg.HTTPPORT); err != nil {
		log.Fatalf("failed to run")
	}
}
