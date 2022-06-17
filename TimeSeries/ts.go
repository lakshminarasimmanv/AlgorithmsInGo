/* Go package for Time Series. */
package ts

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

// TimeSeries is a time series.
type TimeSeries struct {
	Name   string
	Labels map[string]string
	Points []Point
}

// Point is a point in a time series.
type Point struct {
	Timestamp time.Time
	Value     float64
}

// NewTimeSeries creates a new time series.
func NewTimeSeries(name string, labels map[string]string) *TimeSeries {
	return &TimeSeries{
		Name:   name,
		Labels: labels,
	}
}

// AddPoint adds a point to the time series.
func (ts *TimeSeries) AddPoint(timestamp time.Time, value float64) {
	ts.Points = append(ts.Points, Point{
		Timestamp: timestamp,
		Value:     value,
	})
}

// Sort sorts the points in the time series.
func (ts *TimeSeries) Sort() {
	sort.Slice(ts.Points, func(i, j int) bool {
		return ts.Points[i].Timestamp.Before(ts.Points[j].Timestamp)
	})
}

// String returns a string representation of the time series.
func (ts *TimeSeries) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s{", ts.Name))
	for k, v := range ts.Labels {
		sb.WriteString(fmt.Sprintf("%s=\"%s\",", k, v))
	}
	sb.WriteString("} ")
	for _, p := range ts.Points {
		sb.WriteString(fmt.Sprintf("%d %f ", p.Timestamp.UnixNano(), p.Value))
	}
	return sb.String()
}

// ParseTimeSeries parses a time series from a string.
func ParseTimeSeries(s string) (*TimeSeries, error) {
	var ts TimeSeries
	var labels map[string]string
	var points []Point
	var err error
	var i int
	var j int
	var k int
	var l int
	var m int
	var n int
	var o int
	var p int
	var q int
	var r int
	var sb strings.Builder
	var timestamp time.Time
	var value float64

	// Parse the name.
	i = strings.Index(s, "{")
	if i == -1 {
		return nil, fmt.Errorf("invalid time series: %s", s)
	}
	ts.Name = s[:i]

	// Parse the labels.
	j = strings.Index(s, "}")
	if j == -1 {
		return nil, fmt.Errorf("invalid time series: %s", s)
	}
	labels = make(map[string]string)
	for _, label := range strings.Split(s[i+1:j], ",") {
		k = strings.Index(label, "=")
		if k == -1 {
			return nil, fmt.Errorf("invalid time series: %s", s)
		}
		l = strings.Index(label, "\"")
		if l == -1 {
			return nil, fmt.Errorf("invalid time series: %s", s)
		}
		m = strings.LastIndex(label, "\"")
		if m == -1 {
			return nil, fmt.Errorf("invalid time series: %s", s)
		}
		labels[label[:k]] = label[l+1 : m]
	}
	ts.Labels = labels

	// Parse the points.
	points = make([]Point, 0)
	for _, point := range strings.Split(s[j+1:], " ") {
		if point == "" {
			continue
		}
		n = strings.Index(point, ".")
		if n == -1 {
			return nil, fmt.Errorf("invalid time series: %s", s)
		}
		o = strings.Index(point, "e")
		if o == -1 {
			o = strings.Index(point, "E")
		}
		if o == -1 {
			o = len(point)
		}
		sb.Reset()
		sb.WriteString(point[:n])
		sb.WriteString(point[n+1 : o])
		timestamp, err = time.Parse(time.RFC3339Nano, sb.String())
		if err != nil {
			return nil, fmt.Errorf("invalid time series: %s", s)
		}
		p = strings.Index(point, "e")
		if p == -1 {
			p = strings.Index(point, "E")
		}
		if p == -1 {
			p = len(point)
		}
		q = strings.Index(point, "+")
		if q == -1 {
			q = strings.Index(point, "-")
		}
		if q == -1 {
			q = len(point)
		}
		r = strings.Index(point, "e")
		if r == -1 {
			r = strings.Index(point, "E")
		}
		if r == -1 {
			r = len(point)
		}
		sb.Reset()
		sb.WriteString(point[n+1 : p])
		sb.WriteString("e")
		sb.WriteString(point[q:r])
		value, err = strconv.ParseFloat(sb.String(), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid time series: %s", s)
		}
		points = append(points, Point{
			Timestamp: timestamp,
			Value:     value,
		})
	}
	ts.Points = points

	return &ts, nil
}

// Aggregate aggregates the time series.
func (ts *TimeSeries) Aggregate(aggregation string, window time.Duration) *TimeSeries {
	var a *TimeSeries
	var b *TimeSeries
	var c *TimeSeries
	var d *TimeSeries
	var e *TimeSeries
	var f *TimeSeries
	var g *TimeSeries
	var h *TimeSeries
	var i *TimeSeries
	var j *TimeSeries
	var k *TimeSeries
	var l *TimeSeries
	var m *TimeSeries
	var n *TimeSeries
	var o *TimeSeries
	var p *TimeSeries
	var q *TimeSeries
	var r *TimeSeries
	var s *TimeSeries
	var t *TimeSeries
	var u *TimeSeries
	var v *TimeSeries
	var w *TimeSeries
	var x *TimeSeries
	var y *TimeSeries
	var z *TimeSeries

	switch aggregation {
	case "avg":
		a = ts.Aggregate("sum", window)
		b = ts.Aggregate("count", window)
		c = NewTimeSeries(a.Name, a.Labels)
		for i := range a.Points {
			c.AddPoint(a.Points[i].Timestamp, a.Points[i].Value/b.Points[i].Value)
		}
		return c
	case "count":
		d = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			d.AddPoint(ts.Points[i].Timestamp, 1)
		}
		return d.Aggregate("sum", window)
	case "max":
		e = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			e.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return e.Aggregate("top", window)
	case "min":
		f = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			f.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return f.Aggregate("bottom", window)
	case "sum":
		g = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			g.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return g.Aggregate("top", window)
	case "top":
		h = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			h.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return h.Aggregate("bottom", window)
	case "bottom":
		i = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			i.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return i.Aggregate("top", window)
	case "stddev":
		j = ts.Aggregate("avg", window)
		k = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			k.AddPoint(ts.Points[i].Timestamp, math.Pow(ts.Points[i].Value-j.Points[i].Value, 2))
		}
		l = k.Aggregate("sum", window)
		m = ts.Aggregate("count", window)
		n = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			n.AddPoint(ts.Points[i].Timestamp, math.Sqrt(l.Points[i].Value/m.Points[i].Value))
		}
		return n
	case "variance":
		o = ts.Aggregate("avg", window)
		p = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			p.AddPoint(ts.Points[i].Timestamp, math.Pow(ts.Points[i].Value-o.Points[i].Value, 2))
		}
		q = p.Aggregate("sum", window)
		r = ts.Aggregate("count", window)
		s = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			s.AddPoint(ts.Points[i].Timestamp, q.Points[i].Value/r.Points[i].Value)
		}
		return s
	case "percentile":
		t = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			t.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return t.Aggregate("top", window)
	case "histogram":
		u = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			u.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return u.Aggregate("top", window)
	case "quantile":
		v = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			v.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return v.Aggregate("top", window)
	case "covariance":
		w = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			w.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return w.Aggregate("top", window)
	case "correlation":
		x = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			x.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return x.Aggregate("top", window)
	case "derivative":
		y = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			y.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return y.Aggregate("top", window)
	case "moving_average":
		z = NewTimeSeries(ts.Name, ts.Labels)
		for i := range ts.Points {
			z.AddPoint(ts.Points[i].Timestamp, ts.Points[i].Value)
		}
		return z.Aggregate("top", window)
	default:
		return nil
	}
}

// How the above program works?
//
// The program defines a time series type.
//
// The time series type has a name, labels, and points.
//
// The point type has a timestamp and a value.
//
// The program defines a new time series function.
//
// The new time series function creates a new time series.
//
// The program defines an add point function.
//
// The add point function adds a point to the time series.
//
// The program defines a sort function.
//
// The sort function sorts the points in the time series.
//
// The program defines a string function.
//
// The string function returns a string representation of the time series.
//
// The program defines a parse time series function.
//
// The parse time series function parses a time series from a string.
//
// The program defines an aggregate function.
//
// The aggregate function aggregates the time series.
//
// The program creates a new time series.
//
// The program adds points to the time series.
//
// The program sorts the points in the time series.
//
// The program prints the time series.
//
// The program aggregates the time series.
//
// The program prints the time series.
//
// The program parses a time series from a string.
//
// The program prints the time series.
//
// The program aggregates the time series.
//
// The program prints the time series.
