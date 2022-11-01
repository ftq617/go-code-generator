package code

import (
	"code/gen/auto"
	"code/gen/mapper"
	"code/gen/service"
	"code/gen/util/conf"
	"errors"
	"github.com/spf13/cobra"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create code",
		RunE:  createCode,
	}
	cmd.Flags().StringSliceP("name", "n", []string{}, "Filter with table name")
	return cmd
}

func createCode(cmd *cobra.Command, args []string) error {
	names,_ := cmd.Flags().GetStringSlice("name")
	database := conf.Database.DatabaseName
	if database == "" {
		return errors.New("database is empty")
	}
	if len(names) == 0 {
		return errors.New("请选择表名字")
	}
	tables,err := mapper.GetTables(database)
	if err != nil {
		return err
	}
	nameMap := make(map[string]struct{})
	for _,name := range names{
		nameMap[name] = struct{}{}
	}
	var ts []auto.TableInfo
	for _,t := range tables {
		if _,ok := nameMap[t.TableName]; ok {
			ts = append(ts,t)
		}
	}
	if len(ts) == 0{
		return errors.New("表名不正确！")
	}

	err = service.Autocode(ts)
	if err != nil {
		return err
	}
	return nil
}
