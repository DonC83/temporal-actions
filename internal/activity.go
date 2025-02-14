package internal

import (
	"context"
	"fmt"
)

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("hello %s", name)
	return greeting, nil
}

func AnotherFunction(ctx context.Context) error {
	fmt.Println("doing some other work")
	return nil
}

func ComplainingFunction(ctx context.Context, value string) (string, error) {
	return fmt.Sprintf("I'm really slow doinbg %s", value), nil
}
