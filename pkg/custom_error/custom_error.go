package customError

import (
	"github.com/pkg/errors"
	"log"
)

func CheckError(err error) {
	log.Fatalf("%v", errors.Wrap(err, "Error occurred"))
}
