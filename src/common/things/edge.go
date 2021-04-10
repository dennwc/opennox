package things

func (f *File) skipEDGE() error {
	if err := f.skip(4); err != nil {
		return err
	}
	if err := f.skipBytes8(); err != nil {
		return err
	}
	if err := f.skip(2*4 + 1); err != nil {
		return err
	}
	n1, err := f.readU8()
	if err != nil {
		return err
	}
	if err := f.skip(2 * 1); err != nil {
		return err
	}
	n2, err := f.readU8()
	if err != nil {
		return err
	}
	n3, err := f.readU8()
	if err != nil {
		return err
	}

	n := 2 * int(n1) * (int(n2) + int(n3))
	for i := 0; i < n; i++ {
		if err := f.skipImageRef(); err != nil {
			return err
		}
	}
	return f.checkEND()
}
