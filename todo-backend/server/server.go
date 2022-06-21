package server

import (
  models "todo-backend/models"
)

func Start() {
  router := setRouter()
  models.ConnectDatabase()
  router.Run("localhost:3001")
}