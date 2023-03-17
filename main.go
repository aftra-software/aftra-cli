/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"context"
	"net/http"

	"github.com/syndis-software/iris-api/cmd"
)

func main() {
	cmd.Execute(context.Background(), &http.Client{})
}
