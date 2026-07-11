//go:build ((gles || arm || arm64) && !android && !ios && !mobile && !darwin && !wasm && !test_web_driver) || ((android || ios || mobile) && (!wasm || !test_web_driver)) || wasm || test_web_driver

package gl

import (
	_ "embed"

	"fyne.io/fyne/v2/canvas"
)

var (
	//go:embed shaders/blur_es.frag
	shaderBlurFrag []byte

	//go:embed shaders/blur_es.vert
	shaderBlurVert []byte

	//go:embed shaders/line_es.frag
	shaderLineFrag []byte

	//go:embed shaders/line_es.vert
	shaderLineVert []byte

	//go:embed shaders/rectangle_es.frag
	shaderRectangleFrag []byte

	//go:embed shaders/rectangle_es.vert
	shaderRectangleVert []byte

	//go:embed shaders/round_rectangle_es.frag
	shaderRoundrectangleFrag []byte

	//go:embed shaders/simple_es.frag
	shaderSimpleFrag []byte

	//go:embed shaders/simple_es.vert
	shaderSimpleVert []byte

	//go:embed shaders/regular_polygon_es.frag
	shaderPolygonFrag []byte

	//go:embed shaders/arc_es.frag
	shaderArcFrag []byte

	//go:embed shaders/bezier_curve_es.frag
	shaderBezierCurveFrag []byte

	//go:embed shaders/arbitrary_polygon_es.frag
	shaderArbitraryPolygonFrag []byte

	//go:embed shaders/ellipse_es.frag
	shaderEllipseFrag []byte
)

// userShaderFragment returns the fragment shader source to use for the given
// shader object on this build target.
func userShaderFragment(s *canvas.Shader) []byte {
	return s.SourceES
}
