package commons

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func GetenvInt(key string) (int, error) {
	v, err := strconv.Atoi(key)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func CreateUuID() string {
	id := uuid.New()
	fmt.Println(id.String())
	return id.String()
}
