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

type ThreadsCollector struct {
	*baseCollector

	desc *prometheus.Desc
}

func NewThreadsCollector(db *sql.DB) *ThreadsCollector {
	desc := prometheus.NewDesc("mysql_threads", "mysql threads", []string{"status"}, nil)
	return &ThreadsCollector{
		newBaseCollector(db),
		desc,
	}
}

func (c *ThreadsCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.desc
}

func (c *ThreadsCollector) Collect(metrics chan<- prometheus.Metric) {
	result := c.statusMap("threads")
	for k, v := range result {
		metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, v, k)
	}
}
