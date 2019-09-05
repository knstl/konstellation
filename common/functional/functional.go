package functional

type F func() error

func Fs(fs ...F) error {
	for _, f := range fs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
