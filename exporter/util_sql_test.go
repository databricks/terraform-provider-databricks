package exporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSqlSchemaNames(t *testing.T) {
	sql := "SELECT * FROM cat1.sch1.tbl1 t1\n" +
		"JOIN `cat2`.`sch2`.`tbl2` t2 ON t1.id = t2.id\n" +
		"JOIN Cat3.`sch3`.tbl3 t3 ON t1.id = t3.id\n" +
		"JOIN `my-cat`.sch4.tbl4 t4 ON t1.id = t4.id -- FROM cat5.sch5.tbl5\n" +
		"/* FROM cat6.sch6.tbl6 */ WHERE t1.col = 'cat7.sch7.tbl7' AND t2.col = \"it's cat8.sch8.tbl8\"\n" +
		"AND t3.col = `quoted 'cat9.sch9.tbl9`.x AND t1.a.b.c = 1 AND t1.id IN (SELECT id FROM cat10.sch10.tbl10)\n" +
		"AND t4.id IN (select id from\n  cat11.sch11.tbl11) UNION SELECT * FROM STREAM(cat12.sch12.tbl12)\n" +
		"UNION SELECT * FROM stream cat13.sch13.tbl13 UNION SELECT cat14.sch14.tbl14 FROM x.y"
	names := findSqlSchemaNames(sql)
	var fullNames []string
	for _, name := range names {
		fullNames = append(fullNames, name.FullName())
	}
	// `my-cat` requires backticks, so it isn't substituted. Names outside of FROM & JOIN aren't matched.
	assert.Equal(t, []string{"cat1.sch1", "cat2.sch2", "cat3.sch3", "cat10.sch10", "cat11.sch11",
		"cat12.sch12", "cat13.sch13"}, fullNames)
	assert.Equal(t, "cat1.sch1", sql[names[0].Start:names[0].End])
	assert.Equal(t, "`cat2`.`sch2`", sql[names[1].Start:names[1].End])
	assert.Equal(t, "Cat3.`sch3`", sql[names[2].Start:names[2].End])
}

func TestFindSqlSchemaNamesEdgeCases(t *testing.T) {
	assert.Empty(t, findSqlSchemaNames(""))
	assert.Empty(t, findSqlSchemaNames("SELECT * FROM sch.tbl"))
	assert.Empty(t, findSqlSchemaNames("SELECT 'unterminated cat.sch.tbl"))
	assert.Empty(t, findSqlSchemaNames("SELECT 1 /* unterminated cat.sch.tbl"))
	assert.Empty(t, findSqlSchemaNames("SELECT * FROM `cat``x`.sch.tbl"))
	assert.Empty(t, findSqlSchemaNames("cat.sch.tbl"))
	assert.Empty(t, findSqlSchemaNames("SELECT * FROMcat.sch.tbl"))
	assert.Empty(t, findSqlSchemaNames("SELECT * XFROM cat.sch.tbl"))
	names := findSqlSchemaNames("JOIN cat.sch.tbl")
	assert.Equal(t, 1, len(names))
	assert.Equal(t, "cat.sch", names[0].FullName())
}

func TestSqlSchemaNameCapture(t *testing.T) {
	sql := "SELECT * FROM `cat`.`sch`.tbl -- FROM cat.other.tbl"
	matches := sqlTableReferenceRegex.FindAllStringSubmatchIndex(sql, -1)
	assert.Equal(t, 2, len(matches))
	start, end, lookup, ok := sqlSchemaNameCapture(sql, matches[0])
	assert.True(t, ok)
	assert.Equal(t, "`cat`.`sch`", sql[start:end])
	assert.Equal(t, "cat.sch", lookup)
	_, _, _, ok = sqlSchemaNameCapture(sql, matches[1])
	assert.False(t, ok)
}
