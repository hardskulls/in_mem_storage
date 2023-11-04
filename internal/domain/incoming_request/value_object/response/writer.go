package response

func (r Response) Write(str string) error {
	_, err := r.inner.Write([]byte(str))
	return err
}
