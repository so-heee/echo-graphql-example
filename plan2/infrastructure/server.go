package infrastructure

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/so-heee/graphql-example/plan2/interfaces/database"
)

type Config struct {
	MysqlUser     string `required:"true" split_words:"true"`
	MysqlPassword string `required:"true" split_words:"true"`
	MysqlDatabase string `required:"true" split_words:"true"`
	MysqlHost     string `required:"true" split_words:"true"`
}

var config Config

func Run() {

	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}

	dns := fmt.Sprintf(`%v:%v@tcp(%v:3306)/%v?parseTime=True`,
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlDatabase,
	)
	repo, err := database.NewRepository(dns)
	if err != nil {
		fmt.Errorf("failed to create repository: %+v", err)
	}
	r := Router(repo)

	log.Printf("connect to http://localhost:8080/playground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", r))
}
