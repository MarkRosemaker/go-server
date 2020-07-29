package tpl

import (
	"net/http"
)

// DataFunc is a function which return is given to a template.
// Since the data for a template depends on the request, that is the input to the function.
type DataFunc func(*http.Request) interface{}

// StdDataFunc is a simple DataFunc that just returns the Request in a struct to the template, i.e. in a template, we can use '{{ .Request }}'.
func StdDataFunc(r *http.Request) interface{} {
	return struct {
		Request *http.Request
	}{r}
}

// Data is the data object that will be given to the template of a page.
// type Data struct {
// 	Req *http.Request
// 	*serverdata.Data
// }

// New creates a new data object from a context and request.
// func New(req *http.Request) *Data {
// 	return &Data{req, serverdata.Get()}
// }

// add

// Neighbors returns all points within radius.
// func (d *Data) Neighbors() point.Points {
// 	list := d.NeighborsWithDistance()
// 	k := len(list)

// 	o := make(point.Points, k)
// 	for i := 0; i < k; i++ {
// 		o[i] = list[i].Point
// 	}
// 	return o
// }

// // NeighborsWithDistance returns all points within radius, with the distance.
// func (d *Data) NeighborsWithDistance() point.PointsWithDistance {
// 	p, err := point.FromRequest(d.Req)
// 	if err != nil {
// 		return nil
// 	}

// 	var r int
// 	if r, err = d.searchRadiusWithError(); err != nil {
// 		return nil
// 	}

// 	return p.NeighborsWithDistance(d.AllPoints, r)
// }

// // SearchOriginPoint returns the point that the user is searching for.
// // It can be used in templates by using 'with' as the error is discarded.
// func (d *Data) SearchOriginPoint() *point.Point {
// 	p, _ := point.FromRequest(d.Req)
// 	return p
// }

// // SearchRadius returns the distance that the user is searching for.
// // It can be used in templates by using 'with' as the error is discarded.
// func (d *Data) SearchRadius() int {
// 	dist, _ := d.searchRadiusWithError()
// 	return dist
// }

// // searchRadiusWithError returns the distance that the user is searching for, or an error.
// func (d *Data) searchRadiusWithError() (int, error) {
// 	dist, err := strconv.Atoi(d.Req.FormValue("distance"))
// 	if err != nil {
// 		return 0, err
// 	}
// 	return dist, nil
// }