//go:build tools
// +build tools

package tools

import (
	_ "github.com/gofiber/fiber/3"
	_ "gorm.io/gorm"
	_ "gorm.io/driver/mysql"
	_ "github.com/go-playground/validator/v10"
	_ "github.com/spf13/viper"
	_ "github.com/golang-jwt/jwt/v5"
	_ "github.com/sirupsen/logrus"
	_ "github.com/stretchr/testify"
)
