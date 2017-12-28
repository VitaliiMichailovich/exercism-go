package erratum

func Use(ro ResourceOpener, in string) (result error) {
	resource, err := open(ro)
	if err != nil {
		return err
	}
	defer resource.Close()

	defer func() {
		if recover := recover(); recover != nil {
			if asFrob, ok := recover.(FrobError); ok {
				resource.Defrob(asFrob.defrobTag)
			}
			if asError, ok := recover.(error); ok {
				result = asError
			}
		}
	}()
	resource.Frob(in)
	return
}

func open(ro ResourceOpener) (resource Resource, err error) {
	for {
		resource, err = ro()
		if _, ok := err.(TransientError); ok {
			continue
		} else {
			break
		}
	}
	return
}
