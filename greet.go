package greeting_proj3ct

import (
	"errors"
	"fmt"
)

func GetGreet(name string) (string, error) {
	if name == "" {
		return "", errors.New("пустое имя")
	} else {
		return fmt.Sprintf("Привет, %s, как жизнь ??", name), nil
	}
}

func Hi(name string) {
	if name == "" {
		panic(errors.New("нет имени"))
	} else {
		fmt.Printf("привет %s", name)
	}
}
