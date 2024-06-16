package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	_ "github.com/lib/pq"
	"github.com/millirud/gar-2-db/internal/repo/pg"
	"github.com/millirud/gar-2-db/internal/repo/xmlreader"
	"github.com/millirud/gar-2-db/internal/usecase/xml2db"
	"github.com/rs/zerolog/log"
)

func main() {
	parser := argparse.NewParser("info", "GAR parser from xml to postgres")

	argSubjectDirPath := parser.String("s", "source", &argparse.Options{Required: true, Help: "Path to dir with xml"})
	argDbHost := parser.String("H", "host", &argparse.Options{Required: true, Help: "Host of db"})
	argDbPort := parser.Int("p", "port", &argparse.Options{Required: true, Help: "Port of db"})
	argDbUser := parser.String("u", "user", &argparse.Options{Required: true, Help: "User of db"})
	argDbPassword := parser.String("P", "password", &argparse.Options{Required: true, Help: "Password of db"})
	argDbName := parser.String("n", "name", &argparse.Options{Required: true, Help: "Name of db"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return
	}

	subjectDirPath := *argSubjectDirPath
	dbHost := *argDbHost
	dbPort := *argDbPort
	dbUser := *argDbUser
	dbPassword := *argDbPassword
	dbName := *argDbName

	conn, dbCloser, err := getDb(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		panic(err)
	}

	defer func() {
		dbCloser()
	}()

	reader := xmlreader.New()
	pgRepo := pg.New(conn)

	action := xml2db.New(reader, pgRepo)

	if err := action.Run(subjectDirPath); err != nil {
		log.Error().Err(err).Msgf("failed to process %s", subjectDirPath)
	}

}

func getDb(
	host string,
	port int,
	user string,
	password string,
	dbname string,
) (*sql.DB, func(), error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, nil, err
	}

	return db, func() { db.Close() }, nil
}
