/*
Copyright (c) 2022 xiexianbin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package collectors

import (
	"database/sql"
	"log"
)

type baseCollector struct {
	db *sql.DB
}

func newBaseCollector(db *sql.DB) *baseCollector {
	return &baseCollector{db}
}

func (c *baseCollector) status(name string) float64 {
	row := c.db.QueryRow("show global status where variable_name=?", name)
	var (
		key   string
		value float64
	)
	if err := row.Scan(&key, &value); err != nil {
		return value
	}
	return -1
}

func (c *baseCollector) statusMap(key string) map[string]float64 {
	result := map[string]float64{}
	rows, err := c.db.Query("show global status where variable_name like ?", key+"%")
	if err != nil {
		log.Printf("show status %s db: %#v", key, err.Error())
		return result
	}
	for rows.Next() {
		var (
			key   string
			value float64
		)
		_ = rows.Scan(&key, &value)
		result[key] = value
	}
	return result
}

func (c *baseCollector) variables(name string) float64 {
	row := c.db.QueryRow("show global variables where variable_name=?", name)
	var (
		key   string
		value float64
	)
	if err := row.Scan(&key, &value); err != nil {
		return value
	}
	return -1
}

func (c *baseCollector) variablesString(name string) string {
	row := c.db.QueryRow("show global variables where variable_name=?", name)
	var (
		key   string
		value string
	)
	if err := row.Scan(&key, &value); err == nil {
		return value
	} else {
		log.Printf("show version: %#v", err.Error())
	}
	return ""
}
