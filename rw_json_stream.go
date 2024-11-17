package main

var (
  _ dataReader = (*jsonStreamReader)(nil)
  _ dataWriter = (*jsonStreamWriter)(nil)
)

type jsonStreamReader struct {}

func (r *jsonStreamReader) getFields() []string {
  return nil
}

func (r *jsonStreamReader) readRow() (map[string]any, error) {
  return nil, nil
}

func (r *jsonStreamReader) done() bool {
  return false
}

func (r *jsonStreamReader) close() error {
  return nil
}

type jsonStreamWriter struct {}

func (w *jsonStreamWriter) setFields() []string {
  return nil
}

func (w *jsonStreamWriter) writeRow(map[string]any) error {
  return nil
}

func (w *jsonStreamWriter) done() bool {
  return false
}

func (w *jsonStreamWriter) close() error {
  return nil
}

