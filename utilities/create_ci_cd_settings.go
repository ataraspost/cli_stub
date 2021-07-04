package utilities

func CreateCiCdSettings(path string) error {
	path_template := path + "/stub/templates/ci_cd/"

	path_project := path + "/stub/"

	err := Copy(path_template+".gitlab-ci.yml", path_project+".gitlab-ci.yml")

	if err != nil {
		return err
	}

	return nil


}
