package writer

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jyisus/notioncli/entity"
	"os"
)

type TerminalWriter struct {
}

func NewTerminalWriter() TerminalWriter {
	return TerminalWriter{}
}

func (w TerminalWriter) PrintTasks(tasks []entity.Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.AppendHeader(table.Row{"#", "Text", "ID"})
	for i, task := range tasks {
		t.AppendRow(table.Row{i, task.Text, task.Id})
	}
	t.Render()
}
