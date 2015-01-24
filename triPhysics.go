
import (
	"math"
)

type Vector struct {
	    x, y, z float64
}


func validateImpact (tri1, tri2, tri3 Vector, object Object) Vector {

	// calculate the cross-product of the norm and direction
	if object.ObjectSpace > distDiff(object) {
		return impact(object)
	}

func distDiff(object) float64 { 

	tri1, tri2, height = object.Position[0],object.Position[2],object.Position[2]

	centroid Vector := {(tri1[0]+tri2[0]+tri3[0])/3,(tri1[1]+tri2[1]+tri3[1])/3,(tri1[2]+tri2[2]+tri3[2])/3}

	diff := math.Abs(centroid)-math.Abs(object.Position)

	return diff
}
func impact (object Object) Object {
	// just changing the z - direction
	object.Direction[2] = -object.Direction[2]
	return
}


func normTriangle (){

	point := triCentroid()

	surfaceNorm := crossProduct(a,b)
}
func makeVertex () {

}

func triCentroid () {
	

}
