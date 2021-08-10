package main

import (
	"github.com/djumanoff/amqp"
	common_lib21 "github.com/kirigaikabuto/common-lib21"
	"github.com/kirigaikabuto/lesson20acl"
)

func main() {
	cfg := lesson20acl.PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "setdatauser",
		Password: "123456789",
		Database: "lesson20acl",
		Params:   "sslmode=disable",
	}
	store, err := lesson20acl.NewPostgresRoleStore(cfg)
	if err != nil {
		panic(err)
		return
	}
	service := lesson20acl.NewRoleService(store)
	commandHandler := common_lib21.NewCommandHandler(service)
	roleAmqpEndpoints := lesson20acl.NewRoleAmqpEndpoints(commandHandler)
	config := amqp.Config{
		Host: "localhost",
		Port: 5672,
		LogLevel: 5,
	}
	serverConfig := amqp.ServerConfig{
		ResponseX: "response",
		RequestX:  "request",
	}
	sess := amqp.NewSession(config)
	err = sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	srv, err := sess.Server(serverConfig)
	if err != nil {
		panic(err)
		return
	}
	srv.Endpoint("roles.create", roleAmqpEndpoints.MakeCreateRoleAmqpEndpoint())
	if err := srv.Start(); err != nil {
		panic(err)
		return
	}
}
