package table

import (
	"code/gen/mapper"
	"code/gen/util/conf"
	"code/gen/util/table"
	"errors"
	"github.com/spf13/cobra"
)

func NewCmdShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show <name>",
		Short: "show table",
		RunE:  showTable,
	}

	return cmd
}

func showTable(cmd *cobra.Command, args []string) error {
	name := args[0]
	database := conf.Database.DatabaseName
	if database == "" {
		return errors.New("database is empty")
	}
	columns,err := mapper.GetColumns(database,name)
	if err != nil {
		return err
	}


	tableLstMap := make([]map[string]interface{}, len(columns))
	for i, s := range columns {
		tableLstMap[i] = table.Struct2Map(s)
	}
	fields := []string{
		"ColumnName",
		"DataType",
		"ColumnType",
		"ColumnComment"}
	table.RenderAsTable(tableLstMap, fields)
	return nil
}
