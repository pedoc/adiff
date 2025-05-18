package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: adiff str1 str2 splitstr")
		fmt.Println("Sample: tool \"a,b,c\" \"b,c,d\" \",\"")
		os.Exit(1)
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	sep := ","
	if len(os.Args) >= 4 {
		sep = os.Args[3]
	}

	set1StringArray := strings.Split(str1, sep)
	set2StringArray := strings.Split(str2, sep)
	set1 := sliceToSet(set1StringArray)
	set2 := sliceToSet(set2StringArray)

	diff1 := difference(set1, set2)
	diff2 := difference(set2, set1)
	inter := intersection(set1, set2)

	printTable(str1, str2, sep, set1StringArray, set2StringArray, diff1, diff2, inter)
}

func printTable(str1 string, str2 string, sep string, set1, set2, diff1, diff2, inter []string) {
	sort.Strings(set1)
	sort.Strings(set2)
	sort.Strings(diff1)
	sort.Strings(diff2)
	sort.Strings(inter)

	data := [][]string{
		{"raw", str1, str2},
		{"raw(ordered)", strings.Join(set1, sep), strings.Join(set2, sep)},
		//{"", "", "", ""},
		{"diff", strings.Join(diff1, sep), strings.Join(diff2, sep)},
		//{"", "", "", ""},
		{"inter", strings.Join(inter, sep), strings.Join(inter, sep)},
	}

	table := tablewriter.NewTable(os.Stdout,
		//tablewriter.WithRenderer(renderer.NewColorized()),
		tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{
			Settings: tw.Settings{
				Separators: tw.Separators{BetweenRows: tw.On},
			},
		})),
		tablewriter.WithConfig(tablewriter.Config{
			Row: tw.CellConfig{
				Formatting: tw.CellFormatting{
					AutoWrap:  tw.WrapNormal, // Wrap long content
					Alignment: tw.AlignLeft,  // Left-align rows
					MergeMode: tw.MergeHorizontal,
				},
				ColMaxWidths: tw.CellWidth{Global: 60},
			},
			Footer: tw.CellConfig{
				Formatting: tw.CellFormatting{Alignment: tw.AlignRight},
			},
		}),
	)
	table.Header([]string{"name" + "(sep=\"" + sep + "\")", "str1", "str2"})
	table.Bulk(data)
	//table.Footer([]string{"", "Total", strconv.Itoa(len(data))})
	table.Render()
}

func sliceToSet(slice []string) map[string]struct{} {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		if s == "" {
			continue
		}
		set[s] = struct{}{}
	}
	return set
}

func difference(a, b map[string]struct{}) []string {
	var diff []string
	for k := range a {
		if _, found := b[k]; !found {
			diff = append(diff, k)
		}
	}
	return diff
}

func intersection(a, b map[string]struct{}) []string {
	var inter []string
	for k := range a {
		if _, found := b[k]; found {
			inter = append(inter, k)
		}
	}
	return inter
}
