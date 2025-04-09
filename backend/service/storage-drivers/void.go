package storage_drivers

type VoidDriver struct{}

func (d *VoidDriver) Init(targetStorageName string, targetDir string) error {
	return nil
}

func (d *VoidDriver) Upload(uuid string, sourceFilePath string) error {
	return nil
}

func (d *VoidDriver) Exists(uuid string) (bool, error) {
	return true, nil
}
