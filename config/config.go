/**********************************
* Project:  blob_storage
* Author:   Ray Chen
* Email:    raychen0411@gmail.com
* Time:     11-30-2016
* All rights reserved!
***********************************/

// some global data structures
package config

// config for traffic control
type TrafficControl struct {
	PkgLoss int     // match 1
	Latency int     // match 2
	Reorder int     // match 3
	Init    bool
}

var Tc * TrafficControl
