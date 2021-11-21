package erratum

func Use(open func() (Resource, error), input string) (err error) {
	var res Resource
	for res, err = open(); err != nil; res, err = open() {
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()

	res.Frob(input)
	return
}
