package ssh

func (s *Ssh) Excute(command string) (res string, err error) {

	session, err := s.Client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	combo, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(combo), err
}
