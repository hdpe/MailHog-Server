package config

type Constraints interface {
	AllowsSender(sender string) bool
}

type StaticConstraints struct {
	AllowedSenders []string
}

func (s StaticConstraints) AllowsSender(sender string) bool {
	for i := 0; i < len(s.AllowedSenders); i ++ {
		if s.AllowedSenders[i] == sender {
			return true
		}
	}
	return false
}


