package dal

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/Pick-Down/BTC_API/models"
)

type FileRepository struct {
	//Mutex to protect io operations from concurrency errors
	mu    sync.Mutex
	users []models.User
}

func (repo *FileRepository) Add(u models.User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	//Lazy loading - if slice hasn't been initialised yet, read the file.
	if err := repo.readUsers(); err != nil {
		return err
	}

	//Open file for read, write, append and create file if it doesn't exist
	f, err := os.OpenFile("users.data", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	//Append formatted string to the end of file
	if _, err := f.WriteString(fmt.Sprintf("%s %s\n", u.Email, u.Password)); err != nil {
		return err
	}

	repo.users = append(repo.users, u)

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func (repo *FileRepository) Get(email string) (*models.User, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	//Lazy loading - if slice hasn't been initialised yet, read the file.
	if err := repo.readUsers(); err != nil {
		return nil, err
	}

	//Get user with the specified email
	for _, user := range repo.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil
}

func (repo *FileRepository) readUsers() error {
	//If already read - return
	if repo.users != nil {
		return nil
	}

	f, err := os.OpenFile("users.data", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	c := make([]byte, 1)
	var user []byte

	//Resets the file's offset in order to read every user.
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	//While file hasn't been read to the end(Foreach byte)
	for {
		_, err := f.Read(c)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		//If not end of line(each line is single user)
		if c[0] != '\n' {
			user = append(user, c[0])
		} else {
			//Convert slice of bytes to string and split by empty space
			userData := strings.Split(string(user), " ")
			repo.users = append(repo.users, models.User{Email: userData[0], Password: userData[1]})
			user = nil
		}
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
