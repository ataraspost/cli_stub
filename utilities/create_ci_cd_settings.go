package utilities


func CreateCiCdSettings(path string) error {
	path_dir_template := path + "/stub/templates/ci_cd/"

	path_dir_nginx := path + "/stub/"

	err := Copy(path_dir_template+"Dockerfile", path_dir_nginx+"Dockerfile")
	if err != nil {
		return err
	}
	return nil

}
