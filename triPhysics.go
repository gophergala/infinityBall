
import (
	"math"
	"github.com/go-gl/mathgl/mgl64"
)

func closestPoint (tri1, tri2, tri3 mgl64.Vec3, object Object)
mgl64.Vec3 {

	// not correct but good enough
	var centroid mgl64.Vec3
	centroid := tri1

	// norm of the triangle, make sure that this is right-oriented
	norm := normTriangle(tri1-tri3, tri2-tri3)

	// projection of vector u on the normal
	proj := object.Position.Projection(norm)

	if object.Size > distDiff(proj, centroid) {
		return impact(object)
	}

// projection method
func (v1 mgl64.Vec3) Projection(v2 mgl64.Vec3) mgl64.Vec3 {
	return mgl64.Vec3{v1.Dot(v2)/(math.Pow(v2[0],2),
	math.Pow(v2[1],2),math.Pow(v2[2],2))*v2}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func distDiff(object Object, centroid mgl64.Vec3) float64 { 
	diff := (centroid-object.Position).Len()
	return diff
}

func impact (object Object) Object {
	// just changing the z - direction
	// if it is x,y,z
	object.Direction[2] = -object.Direction[2]
	return object
}

func normTriangle (triPoint1, triPoint2 mgl64.Vec3){
	// make sure that this is right-oriented
	surfaceNorm := triPoint1.Cross(triPoint2)
	return surfaceNorm
}
