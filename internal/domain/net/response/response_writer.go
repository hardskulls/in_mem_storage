package response

func (r Response) Write(bytes []byte) error {
	_, err := r.inner.Write(bytes)
	return err
}
