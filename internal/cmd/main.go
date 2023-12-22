package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kiper01/Post-CRUD/internal/app/migrations"
	"github.com/kiper01/Post-CRUD/internal/app/repository"
	"github.com/kiper01/Post-CRUD/internal/app/service"
	pb "github.com/kiper01/Post-CRUD/internal/proto"
	"github.com/kiper01/Post-CRUD/pkg/database"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var dbConfig database.Config

func init() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	viper.SetConfigName(env)
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	dbConfig = database.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
		PoolSize: viper.GetInt("database.poolsize"),
	}
}

func main() {

	fmt.Println("gRPC server running ...")

	port := viper.GetInt("server.port")
	if port == 0 {
		log.Fatal("Server port is not set in the config file")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	dbPool, err := database.ConnectDB(dbConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	defer database.CloseDB(dbPool)

	if _, err := dbPool.Exec(context.Background(), migrations.CreateTablePost); err != nil {
		log.Fatalf("Failed to execute migration: %v", err)
	}

	repo := repository.NewPostInfoRepository(dbPool)
	postInfoService := service.NewPostInfoService(repo)

	s := grpc.NewServer()
	pb.RegisterPostInfoServiceServer(s, postInfoService)

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}

// 	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Connection successful")

// 	// Применение миграции от gpt
// 	err = migrations.ApplyMigration(db)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }
