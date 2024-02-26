package helpers

import (
	"github.com/sirupsen/logrus"
)

func PanicIfError(err error, msg string) {
	if err != nil {
		logrus.WithField("Message: ", msg).Error(err)
		panic(err)
	}
}
