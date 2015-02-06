package transformer

import (
	"gopkg.in/Clever/optimus.v3"
)

// A Transformer allows you to easily chain multiple transforms on a table.
type Transformer struct {
	table optimus.Table
}

// Table returns the terminating Table in a Transformer chain.
func (t Transformer) Table() optimus.Table {
	return t.table
}

// Apply applies a given TransformFunc to the Transformer.
func (t *Transformer) Apply(transform optimus.TransformFunc) *Transformer {
	// TODO: Should this return a new transformer instead of modifying the existing one?
	t.table = optimus.Transform(t.table, transform)
	return t
}

// Sink consumes all the Rows.
func (t *Transformer) Sink(sink optimus.Sink) error {
	return sink(t.table)
}

// New returns a Transformer that allows you to chain transformations on a Table.
func New(table optimus.Table) *Transformer {
	return &Transformer{table}
}
