package main

var (
  _ dataReader = (*xlsxReader)(nil)
  _ dataWriter = (*xlsxWriter)(nil)
)

type xlsxReader struct {}

func (r *xlsxReader) getFields() []string {
  return nil
}

func (r *xlsxReader) readRow() (map[string]any, error) {
  return nil, nil
}

func (r *xlsxReader) done() bool {
  return false
}

func (r *xlsxReader) close() error {
  return nil
}

type xlsxWriter struct {}

func (w *xlsxWriter) setFields() []string {
  return nil
}

func (w *xlsxWriter) writeRow(map[string]any) error {
  return nil
}

func (w *xlsxWriter) done() bool {
  return false
}

func (w *xlsxWriter) close() error {
  return nil
}

