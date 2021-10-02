// Package errors Marwan Sulaiman GopherCon 2019
package errors

type Kind int64
type Op string

type Error struct {
	Op   Op
	Kind Kind
	Err  error
}

const (
	emptyDetailsMsg = "without details"
)

func (e Error) Error() string {
	return string(e.Kind.String())
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case Kind:
			e.Kind = arg
		case error:
			e.Err = arg
		default:
			panic("unknown behaviour while constructing Error struct")
		}
	}

	return e
}

// Ops Get error's operations stack trace
func Ops(err error) []Op {
	e, ok := err.(*Error)
	if !ok {
		return []Op{}
	}

	res := []Op{e.Op}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)

	return res
}

func SourceDetails(err error) string {
	e, ok := err.(*Error)
	if !ok {
		return emptyDetailsMsg
	}

	if e.Err == nil {
		return emptyDetailsMsg
	}

	return e.Err.Error()
}

func GetKind(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return UnexpectedErr
	}

	return e.Kind
}

func Contains(s []Kind, k Kind) bool {
	isMember := false
	for _, v := range s {
		if k == v {
			isMember = true
		}
	}

	return isMember
}
