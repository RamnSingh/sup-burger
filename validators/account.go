package validators

import(
  "../models"
  "strings"
  "errors"
  "strconv"
)

func RegisterUser(user models.User, confirmPassword string, roleId string)  error{
  user.Email = strings.TrimSpace(user.Email)
  user.Username = strings.TrimSpace(user.Username)
  user.Password = strings.TrimSpace(user.Password)
  user.Street = strings.TrimSpace(user.Street)
  confirmPassword = strings.TrimSpace(confirmPassword)
  _, err := strconv.Atoi(roleId)

  if err != nil {
    return errors.New("City error")
  }

  if len(user.Email) == 0{
    return errors.New("Email is required")
  }
  if len(user.Username) == 0{
    return errors.New("Username is required")
  }
  if len(user.Password) == 0{
    return errors.New("Password is required")
  }
  if len(user.Street) == 0{
    return errors.New("Street is required")
  }
  if user.Password != confirmPassword{
    return errors.New("Passwords mismatched")
  }
  return nil
}

func LoginUser(user models.User)  error{
  user.Username = strings.TrimSpace(user.Username)
  user.Password = strings.TrimSpace(user.Password)

  if len(user.Username) == 0{
    return errors.New("Username is required")
  }
  if len(user.Password) == 0{
    return errors.New("Password is required")
  }

  return nil
}
