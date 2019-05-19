package urlmaker

type URLArgument struct {
	Name  string
	Value interface{}
}

// IsValid Method checks Argument's Name.
func (arg URLArgument) IsValid() bool {

	if len(arg.Name) == 0 {
		return false
	}

	return true
}
