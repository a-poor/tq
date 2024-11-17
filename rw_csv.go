package main

var (
  _ dataReader = (*csvReader)(nil)
  _ dataWriter = (*csvWriter)(nil)
)

type csvReader struct {}

func (r *csvReader) getFields() []string {
  return nil
}

func (r *csvReader) readRow() (map[string]any, error) {
  return nil, nil
}

func (r *csvReader) done() bool {
  return false
}

func (r *csvReader) close() error {
  return nil
}

type csvWriter struct {}

func (w *csvWriter) setFields() []string {
  return nil
}

func (w *csvWriter) writeRow(map[string]any) error {
  return nil
}

func (w *csvWriter) done() bool {
  return false
}

func (w *csvWriter) close() error {
  return nil
}

