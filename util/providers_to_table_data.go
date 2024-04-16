package util

import "gohtmx/internal/model"

type TableData struct {
	SizeClasses []string
	Header      []string
	Rows        [][]string
}

type TableRowData struct {
	SizeClasses []string
	Values      []string
}

func ProvidersToTableData(providers []*model.ProviderModel) *TableData {
	var result TableData

	result.SizeClasses = []string{"w-56", "w-52", "w-36", "w-56", "w-56"}
	result.Header = []string{"ID", "URL", "Method", "Method Data", "Alias"}

	for _, p := range providers {
		p.SetMethodDataReadableFromSelf()

		result.Rows = append(result.Rows, []string{
			p.Id,
			p.Url,
			string(p.Method),
			p.MethodDataReadable,
			p.Alias,
		})
	}

	return &result
}
