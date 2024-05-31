package util

type Errs struct {
	errors []error
}

func NewErrs() *Errs {
	return &Errs{errors: []error{}}
}

func (e *Errs) Add(err error) {
	if err != nil {
		e.errors = append(e.errors, err)
	}
}

func (e *Errs) Return() []error {
	if len(e.errors) == 0 {
		return nil
	}
	return e.errors
}

func (e *Errs) Has() bool {
  return len(e.errors) > 0
}

func (e *Errs) Print() {
  for _, err := range e.errors {
    println(err.Error())
  }
}
