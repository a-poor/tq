package main

var (
  _ dataReader = (*jsonArrReader)(nil)
  _ dataWriter = (*jsonArrWriter)(nil)
)

type jsonArrReader struct {}

func (r *jsonArrReader) getFields() []string {
  return nil
}

func (r *jsonArrReader) readRow() (map[string]any, error) {
  return nil, nil
}

func (r *jsonArrReader) done() bool {
  return false
}

func (r *jsonArrReader) close() error {
  return nil
}

type jsonArrWriter struct {}

func (w *jsonArrWriter) setFields() []string {
  return nil
}

func (w *jsonArrWriter) writeRow(map[string]any) error {
  return nil
}

func (w *jsonArrWriter) done() bool {
  return false
}

func (w *jsonArrWriter) close() error {
  return nil
}

