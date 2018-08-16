package account_validator

import(
  "../models"
  "strings"
  "errors"
)

func Register(user *models.User, confrmPassword string)  *errors.error{
  user.Email = strings.TrimSpace(user.Email)
  user.Username = strings.TrimSpace(user.Username)
  user.Password = strings.TrimSpace(user.Password)
  confirm_password = strings.TrimSpace(confrm_password)

  if len(user.Email) > 0{
    return errors.New("Email is required")
  }
  if len(user.Username) > 0{
    return errors.New("Username is required")
  }
  if len(user.Password) > 0{
    return errors.New("Password is required")
  }
  if user.Password != confrmPassword{
    err = errors.New("Passwords mismatched")
  }
}
