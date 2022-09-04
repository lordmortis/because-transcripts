package main

import (
	"fmt"

	"gopkg.in/errgo.v2/errors"
)

type CreateDatabaseCommand struct {
}

var cdbCommand CreateDatabaseCommand

func init() {
	_, err := parser.AddCommand(
		"createDatabase",
		"Create Database",
		"Create database and user according to config in config file",
		&cdbCommand)
	if err != nil {
		panic(err)
	}
}

func (x *CreateDatabaseCommand) Execute(args []string) error {
	dbCon, err := dbConnect(options.DBRootUser, options.DBRootPw, "postgres")
	if err != nil {
		return err
	}

	username := configFile.DatabaseConfig.Username
	password := configFile.DatabaseConfig.Password
	database := configFile.DatabaseConfig.Database

	var sqlCmd string

	sqlCmd = fmt.Sprintf("DROP DATABASE IF EXISTS \"%s\"", database)
	if _, err = dbCon.Exec(sqlCmd); err != nil {
		return errors.Because(err, nil, "Unable to remove database '"+database+"'")
	}

	sqlCmd = fmt.Sprintf("DROP ROLE IF EXISTS \"%s\"", username)
	if _, err = dbCon.Exec(sqlCmd); err != nil {
		return errors.Because(err, nil, "Unable to remove existing user '"+username+"'")
	}

	sqlCmd = fmt.Sprintf("CREATE ROLE \"%s\" WITH PASSWORD '%s' LOGIN", username, password)
	if _, err = dbCon.Exec(sqlCmd); err != nil {
		return errors.Because(err, nil, "Unable to create user '"+username+"'")
	}

	sqlCmd = fmt.Sprintf("CREATE DATABASE \"%s\" OWNER \"%s\"", database, username)
	if _, err = dbCon.Exec(sqlCmd); err != nil {
		return errors.Because(err, nil, "Unable to create database '"+database+"'")
	}

	return nil
}
