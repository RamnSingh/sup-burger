package account

import(
  "../../models"
  "strings"
  "errors"
)

func Register(user models.User, confirmPassword string)  error{
  user.Email = strings.TrimSpace(user.Email)
  user.Username = strings.TrimSpace(user.Username)
  user.Password = strings.TrimSpace(user.Password)
  confirmPassword = strings.TrimSpace(confirmPassword)

  if len(user.Email) == 0{
    return errors.New("Email is required")
  }
  if len(user.Username) == 0{
    return errors.New("Username is required")
  }
  if len(user.Password) == 0{
    return errors.New("Password is required")
  }
  if user.Password != confirmPassword{
    return errors.New("Passwords mismatched")
  }
  return nil
}
