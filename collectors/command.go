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

	"github.com/prometheus/client_golang/prometheus"
)

type ComCollector struct {
	*baseCollector

	ComDelete       *prometheus.Desc
	ComDeleteMulti  *prometheus.Desc
	ComInsert       *prometheus.Desc
	ComInsertSelect *prometheus.Desc
	ComSelect       *prometheus.Desc
	ComUpdate       *prometheus.Desc
	ComUpdateMulti  *prometheus.Desc
}

func NewComCollector(db *sql.DB) *ComCollector {
	comDeleteDesc := prometheus.NewDesc("mysql_global_status_com_delete", "mysql global status com delete", nil, nil)
	comDeleteMultiDesc := prometheus.NewDesc("mysql_global_status_com_delete_multi", "mysql global status com delete multi", nil, nil)
	comInsertDesc := prometheus.NewDesc("mysql_global_status_com_insert", "mysql global status com insert", nil, nil)
	comInsertSelectDesc := prometheus.NewDesc("mysql_global_status_com_insert_select", "mysql global status com insert select", nil, nil)
	comSelectDesc := prometheus.NewDesc("mysql_global_status_com_select", "mysql global status com select", nil, nil)
	comUpdateDesc := prometheus.NewDesc("mysql_global_status_com_update", "mysql global status com update", nil, nil)
	comUpdateMultiDesc := prometheus.NewDesc("mysql_global_status_com_update_multi", "mysql global status com update multi", nil, nil)

	return &ComCollector{
		newBaseCollector(db),
		comDeleteDesc,
		comDeleteMultiDesc,
		comInsertDesc,
		comInsertSelectDesc,
		comSelectDesc,
		comUpdateDesc,
		comUpdateMultiDesc,
	}
}

func (c *ComCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.ComDelete
	desc <- c.ComDeleteMulti
	desc <- c.ComInsert
	desc <- c.ComInsertSelect
	desc <- c.ComSelect
	desc <- c.ComUpdate
	desc <- c.ComUpdateMulti
}

func (c *ComCollector) Collect(metrics chan<- prometheus.Metric) {
	result := c.statusMap("com_")
	metrics <- prometheus.MustNewConstMetric(c.ComDelete, prometheus.CounterValue, result["Com_delete"])
	metrics <- prometheus.MustNewConstMetric(c.ComDeleteMulti, prometheus.CounterValue, result["Com_delete_multi"])
	metrics <- prometheus.MustNewConstMetric(c.ComInsert, prometheus.CounterValue, result["Com_insert"])
	metrics <- prometheus.MustNewConstMetric(c.ComInsertSelect, prometheus.CounterValue, result["Com_insert_select"])
	metrics <- prometheus.MustNewConstMetric(c.ComSelect, prometheus.CounterValue, result["Com_select"])
	metrics <- prometheus.MustNewConstMetric(c.ComUpdate, prometheus.CounterValue, result["Com_update"])
	metrics <- prometheus.MustNewConstMetric(c.ComUpdateMulti, prometheus.CounterValue, result["Com_update_multi"])
}
