package internal

type ZimtRootEntry struct {
	entries []ZimtEntry `yaml:entries`
}

type ZimtEntry struct {
	Type string `yaml:type`
	Date string `yaml:date`
	Time string `yaml:time`
}
