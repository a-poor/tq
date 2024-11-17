package main

type dataReader interface {
  // getFields returns the names of the fields in the data
  getFields() []string

  // readRow reads the next row of data
  readRow() (map[string]any, error)

  done() bool

  close() error
}

type dataWriter interface {
  setFields() []string
  
  writeRow(map[string]any) (error)

  done() bool

  close() error
}
