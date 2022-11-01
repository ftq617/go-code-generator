package table

import (
	"code/gen/auto"
	"code/gen/mapper"
	"code/gen/util/conf"
	"code/gen/util/table"
	"errors"
	"github.com/spf13/cobra"
	"strings"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List table",
		RunE:  listTable,
	}

	cmd.Flags().StringP("name", "n", "", "Filter with table name")
	return cmd
}

func listTable(cmd *cobra.Command, args []string) error {
	name,_ := cmd.Flags().GetString("name")
	database := conf.Database.DatabaseName
	if database == "" {
		return errors.New("database is empty")
	}
	tableInfo,err := mapper.GetTables(database)
	if err != nil {
		return err
	}

	var tables []auto.TableInfo
	if name != "" {
		for _,t := range tableInfo {
			if strings.Contains(t.TableName,name) {
				tables = append(tables,t)
			}
		}
	} else {
		tables = tableInfo
	}

	tableLstMap := make([]map[string]interface{}, len(tables))
	for i, s := range tables {
		tableLstMap[i] = table.Struct2Map(s)
	}
	fields := []string{
		"TableName",
		"TableComment",
		"CreateTime"}
	table.RenderAsTable(tableLstMap, fields)

	return nil
}
