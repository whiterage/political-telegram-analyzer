package telegram

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

type terminalAuth struct{}

func newAuthFlow() auth.Flow {
	return auth.NewFlow(terminalAuth{}, auth.SendCodeOptions{})
}

func (terminalAuth) Phone(_ context.Context) (string, error) {
	fmt.Print("Enter Telegram phone number: ")

	value, err := readLine()
	if err != nil {
		return "", nil
	}

	return value, nil
}

func (terminalAuth) Code(_ context.Context, _ *tg.AuthSentCode) (string, error) {
	fmt.Print("Enter Telegram code: ")

	value, err := readLine()
	if err != nil {
		return "", err
	}

	return value, nil
}

func (terminalAuth) Password(_ context.Context) (string, error) {
	fmt.Print("Enter Telegram password 2FA: ")

	value, err := readLine()
	if err != nil {
		return "", nil
	}

	return value, nil
}

func readLine() (string, error) {
	value, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(value), nil
}

func (terminalAuth) AcceptTermsOfService(_ context.Context, tos tg.HelpTermsOfService) error {
	fmt.Println("Telegram Terms of Service:")
	fmt.Println(tos.Text)
	fmt.Print("Accept Telegram Terms of Service? Type YES to continue: ")

	value, err := readLine()
	if err != nil {
		return err
	}

	if value != "YES" {
		return errors.New("telegram terms of service were not accepted")
	}

	return nil
}

func (terminalAuth) SignUp(_ context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, errors.New("telegram sign up is not supported")
}
