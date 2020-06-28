package validate

import (
	"dam/config"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CheckRepoName(name string) error {
	l := len(name)
	if l<=3 && l>=9 {
		return fmt.Errorf("Repository name '%s' is bad. It must have lenght 3-9 symbols", name)
	}

	regexPattern := "[A-Za-z0-9_]"
	matched, err := regexp.Match(regexPattern, []byte(name))
	if err != nil {
		return fmt.Errorf("Cannot match regex patern '%s' with registry name '%s'", regexPattern, name)
	}
	if !matched {
		return fmt.Errorf("Repository name '%s' is bad. It must have only letters, numbers and '_'", name)
	}

	_, err = strconv.ParseInt(name, 10, 32)
	if err == nil {
		return fmt.Errorf("Repository name '%s' is bad. It cannot be a registry number (ID)", name)
	}

	return nil
}

func CheckServer(server string) error {
	if len(server) > 120 {
		return fmt.Errorf("Server URL '%s' is bad. It must have lenght '<' or '=' 120 symbols", server)
	}

	if len(server) == 0 {
		return fmt.Errorf("Server URL '%s' is not valid. It cannot be an empty string", server)
	}

	return nil
}

func CheckLogin(login string) error {
	if len(login) > 24 {
		return fmt.Errorf("Login '%s' is bad. It must have lenght '<' or '=' 24 symbols", login)
	}
	return nil
}

func CheckPassword(pass string) error {
	if len(pass) > 120 {
		return fmt.Errorf("Password is bad. It must have lenght '<' or '=' 120 symbols")
	}
	return nil
}

func CheckID(id string) error {
	id32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return fmt.Errorf("ID '%s' is bad. It must be a number (ID)", id)
	}

	if id32 == 0 {
		return fmt.Errorf("ID '%s' is bad. It cannot be '0'", id)
	}

	return nil
}

func CheckAppName(name string) error {
	l := len(name)
	if l<=4 && l>=32 {
		return fmt.Errorf("App name '%s' is bad. It must have lenght 4-32 symbols", name)
	}

	regexPattern := "[A-Za-z0-9_]"
	matched, err := regexp.Match(regexPattern, []byte(name))
	if err != nil {
		return fmt.Errorf("Cannot match regex patern '%s' with registry name '%s'", regexPattern, name)
	}
	if !matched {
		return fmt.Errorf("App name '%s' is bad. It must have only letters, numbers and '_'", name)
	}

	_, err = strconv.ParseInt(name, 10, 32)
	if err == nil {
		return fmt.Errorf("App name '%s' is bad. It cannot be a registry number (ID)", name)
	}

	return nil
}

func CheckVersion(version string) error {
	l := len(version)
	if l == 0 {
		return fmt.Errorf("App version '%s' is bad. It cannot be an empty string", version)
	}

	if l<=1 && l>=32 {
		return fmt.Errorf("App version '%s' is bad. It must have lenght 1-32 symbols", version)
	}

	regexPattern := "[A-Za-z0-9_.]"
	matched, err := regexp.Match(regexPattern, []byte(version))
	if err != nil {
		return fmt.Errorf("Cannot match regex patern '%s' with app version '%s'", regexPattern, version)
	}
	if !matched {
		return fmt.Errorf("App version '%s' is bad. It must have only letters, numbers, '_' and '.'", version)
	}

	return nil
}

func CheckApp(app string) error {
	arr := strings.Split(app, ":")
	if len(arr) != 2 {
		return fmt.Errorf("'%s' is not <app>:<version>. It must be one symbol ':'", app)
	}
	if CheckAppName(arr[0]) != nil {
		return fmt.Errorf("'%s' is not <app> in <app>:<version> option", app)
	}
	if CheckVersion(arr[1]) != nil {
		return fmt.Errorf("'%s' is not <version> in <app>:<version> option", app)
	}

	return nil
}

func ProjectDir(path string) error {
	l := len(path)
	if l == 0 {
		return fmt.Errorf("Path '%s' is not valid. It cannot be an empty string", path)
	}

	return nil
}

func FilePath(path string) error {
	l := len(path)
	if l == 0 {
		return fmt.Errorf("Path '%s' is not valid. It cannot be an empty string", path)
	}

	return nil
}

func CheckMask(mask string) error {
	regexPattern := "[A-Za-z0-9-_.]"
	matched, err := regexp.Match(regexPattern, []byte(mask))
	if err != nil {
		return fmt.Errorf("Cannot match regex patern '%s' with search mask '%s'", regexPattern, mask)
	}
	if !matched {
		return fmt.Errorf("Mask '%s' is bad. It must have only letters, numbers, '_', '-' and '.'", mask)
	}

	return nil
}

func CheckDockerID(id string) error {
	l := len(id)
	if l != 12 {
		return fmt.Errorf("Docker id '%s' is bad. It must have lenght only 12 symbols", id)
	}

	regexPattern := "[a-z0-9]"
	matched, err := regexp.Match(regexPattern, []byte(id))
	if err != nil {
		return fmt.Errorf("Cannot match regex patern '%s' with docker id '%s'", regexPattern, id)
	}
	if !matched {
		return fmt.Errorf("Docker id '%s' is bad. It must have only lowercase letters and numbers", id)
	}

	return nil
}

func CheckBool(b string) error {
	if b != config.FILES_DB_BOOL_FLAG && b != ""  {
		return fmt.Errorf("Bool flag with value '%s' is bad", b)
	}

	return nil
}