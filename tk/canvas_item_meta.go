package tk

var (
	canvasItemAttributeMap = make(map[CanvasItemType][]string)
)

func (t CanvasItemType) HasAttribute(attr string) bool {
	if attr == "" {
		return false
	}
	for _, v := range canvasItemAttributeMap[t] {
		if v == attr {
			return true
		}
	}
	return false
}

func init() {
	canvasItemAttributeMap[CanvasItemTypeArc] =
		[]string{"dash",
			"activedash",
			"disableddash",
			"dashoffset",
			"fill",
			"activefill",
			"disabledfill",
			"outline",
			"activeoutline",
			"disabledoutline",
			"offset",
			//"outlinestipple",
			//"activeoutlinestipple",
			//"disabledoutlinestipple",
			"outlineoffset",
			"stipple",
			//"activestipple",
			//"disabledstipple",
			"state",
			"tags",
			"width",
			"activewidth",
			"disabledwidth",
			"extent",
			"start",
			"style"}
	// canvasItemAttributeMap[CanvasItemTypeBitmap]
	canvasItemAttributeMap[CanvasItemTypeImage] =
		[]string{"dash",
			"activedash",
			"disableddash",
			"dashoffset",
			"fill",
			"activefill",
			"disabledfill",
			"outline",
			"activeoutline",
			"disabledoutline",
			"offset",
			//"outlinestipple",
			//"activeoutlinestipple",
			//"disabledoutlinestipple",
			"outlineoffset",
			"stipple",
			//"activestipple",
			//"disabledstipple",
			"state",
			"tags",
			"width",
			"activewidth",
			"disabledwidth",
			"anchor",
			"image",
			"activeimage",
			"disabledimage"}
	canvasItemAttributeMap[CanvasItemTypeLine] =
		[]string{"dash",
			"activedash",
			"disableddash",
			"dashoffset",
			"fill",
			"activefill",
			"disabledfill",
			"outline",
			"activeoutline",
			"disabledoutline",
			"offset",
			//"outlinestipple",
			//"activeoutlinestipple",
			//"disabledoutlinestipple",
			"outlineoffset",
			"stipple",
			//"activestipple",
			//"disabledstipple",
			"state",
			"tags",
			"width",
			"activewidth",
			"disabledwidth",
			"anchor",
			"arrow",
			"arrowshape",
			"capstyle",
	"joinstyle",
	"smooth",
	"splinesteps"}
	canvasItemAttributeMap[CanvasItemTypeOval] =
		[]string{"dash",
			"activedash",
			"disableddash",
			"dashoffset",
			"fill",
			"activefill",
			"disabledfill",
			"outline",
			"activeoutline",
			"disabledoutline",
			"offset",
			//"outlinestipple",
			//"activeoutlinestipple",
			//"disabledoutlinestipple",
			"outlineoffset",
			"stipple",
			//"activestipple",
			//"disabledstipple",
			"state",
			"tags",
			"width",
			"activewidth",
			"disabledwidth"}

}
