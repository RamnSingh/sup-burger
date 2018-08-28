package form

import(
  "io/ioutil"
  "mime/multipart"
  "path"
)

type FormHelper struct{}

func (formHelper *FormHelper) UploadFile(file multipart.File, filePath string) error{
  data, err := ioutil.ReadAll(file)
  if err == nil {
    err = ioutil.WriteFile(path.Join(filePath), data, 0777)
  }
  return err
}
