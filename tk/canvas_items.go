package tk

import "fmt"
import "strings"

type CanvasItem struct {
	canvas *Canvas
	id int
	typ CanvasItemType
}
type CanvasImage struct { *CanvasItem }
type CanvasOval struct {
	*CanvasItem
}
type CanvasLine struct {
	*CanvasItem
}

type CanvasItemType int
const (
	CanvasItemTypeNone CanvasItemType = iota
	CanvasItemTypeArc
//	CanvasItemTypeBitmap
	CanvasItemTypeImage
	CanvasItemTypeLine
	CanvasItemTypeOval
	CanvasItemTypePolygon
	CanvasItemTypeRectangle
	CanvasItemTypeText
	CanvasItemTypeWindow
)

func (w* Canvas) DeleteAllItems() {
	evalAsString(fmt.Sprintf("%v delete all", w.id))
}
func (w* Canvas) Delete(item *CanvasItem) {
	evalAsString(fmt.Sprintf("%v delete %d", w.id, item.id))
}
func (w* Canvas) DeleteLine(line *CanvasLine) {
	w.Delete(line.CanvasItem)
}
func (w* Canvas) DeleteImage(image *CanvasImage) {
	w.Delete(image.CanvasItem)
}
func (w* Canvas) DeleteOval(oval *CanvasOval) {
	w.Delete(oval.CanvasItem)
}
func (w* Canvas) CreateArc(x1, y1, x2, y2 int) {}
//func (w* Canvas) CreateBitmap(x, y int) {}
func (w* Canvas) CreateImage(x, y float64, attributes ...*WidgetAttr) *CanvasImage {
	attrs := buildCanvasItemAttributeScript(CanvasItemTypeImage, attributes)
	id, _ := evalAsInt(fmt.Sprintf("%v create image %f %f %s", w.id, x, y, attrs))
	return &CanvasImage{&CanvasItem{w, id, CanvasItemTypeImage}}
}
func (w* Canvas) CreateLine(coords []CanvasCoordinates, attributes... *WidgetAttr) *CanvasLine {
	attrs := buildCanvasItemAttributeScript(CanvasItemTypeLine, attributes)
	coordStrings := make([]string, 0, len(coords))
	for _, coord := range coords {
		coordStrings = append(coordStrings, coord.String())
	}
	id, _ := evalAsInt(fmt.Sprintf("%v create line %s %s", w.id, strings.Join(coordStrings, " "), attrs))
	return &CanvasLine{&CanvasItem{w, id, CanvasItemTypeLine}}
}

func (w* Canvas) CreateOval(x1, y1, x2, y2 float64, attributes ...*WidgetAttr) *CanvasOval {
	attrs := buildCanvasItemAttributeScript(CanvasItemTypeOval, attributes)
	id, _ := evalAsInt(fmt.Sprintf("%v create oval %f %f %f %f %s", w.id, x1, y1, x2, y2, attrs))
	return &CanvasOval{&CanvasItem{w, id, CanvasItemTypeOval}}
}
func (w* Canvas) CreatePolygon(coords []CanvasCoordinates) {}
func (w* Canvas) CreateRectangle(x1, y1, x2, y2 int) {}
func (w* Canvas) CreateText(x, y int) {}
func (w* Canvas) CreateWindow(x, y int) {}

func (w* CanvasItem) BindEvent(event string, fn func(e *Event)) error {
	if !IsEvent(event) {
		return ErrInvalid
	}
	fnid := makeBindEventId()
	var ev Event
	return w.addCanvasItemEventHelper(event, fnid, &ev, func() {
		fn(&ev)
	})
}

func (w* CanvasItem) addCanvasItemEventHelper(event string, fnid string, ev *Event, fn func()) error {
	mainInterp.CreateAction(fnid, func(args []string) {
		ev.parser(args)
		fn()
	})
	return eval(fmt.Sprintf("%v bind %v %v {+%v %v}", w.canvas.Id(), w.id, event, fnid, ev.params()))
}

func buildCanvasItemAttributeScript(typ CanvasItemType, attributes []*WidgetAttr) string {
	var list []string
	for _, attr := range attributes {
		if attr == nil {
			continue
		}
		if !typ.HasAttribute(attr.key) {
			continue
		}
		if im, ok := attr.value.(*Image); ok {
			pname := "atk_tmp_" + attr.key
			setObjText(pname, im.Id())
			list = append(list, fmt.Sprintf("-%v $%v", attr.key, pname))
			continue
		}
		if strs, ok := attr.value.([]string); ok {
			pname := "atk_tmp_" + attr.key
			setObjTextList(pname, strs)
			list = append(list, fmt.Sprintf("-%v $%v", attr.key, pname))
			continue
		}
		if s, ok := attr.value.(string); ok {
			pname := "atk_tmp_" + attr.key
			setObjText(pname, s)
			list = append(list, fmt.Sprintf("-%v $%v", attr.key, pname))
			continue
		}
		list = append(list, fmt.Sprintf("-%v {%v}", attr.key, attr.value))
	}
	return strings.Join(list, " ")
}

type CanvasCoordinates struct {
	x, y float64
	unit string
}
func (c* CanvasCoordinates) String() string {
	return fmt.Sprintf("%f%s %f%s", c.x, c.unit, c.y, c.unit)
}
func CanvasPixelCoords(x, y float64) CanvasCoordinates {
	return CanvasCoordinates{x, y, ""}
}
func CanvasMillimeterCoords(x, y float64) CanvasCoordinates {
	return CanvasCoordinates{x, y, "m"}
}

func CanvasItemAttrDash(pattern string) *WidgetAttr {
	return &WidgetAttr{"dash", pattern}
}
func CanvasItemAttrActiveDesh(pattern string) *WidgetAttr {
	return &WidgetAttr{"activedash", pattern}
}
func CanvasItemAttrDisabledDash(pattern string) *WidgetAttr {
	return &WidgetAttr{"disableddash", pattern}
}
func CanvasItemAttrDashOffset(coords CanvasCoordinates) *WidgetAttr {
	return &WidgetAttr{"dashoffset", coords}
}
func CanvasItemAttrFill(color string) *WidgetAttr {
	return &WidgetAttr{"fill", color}
}
func CanvasItemAttrOutline(color string) *WidgetAttr {
	return &WidgetAttr{"outline", color}
}
func CanvasItemAttrTags(tags []string) *WidgetAttr {
	return &WidgetAttr{"tags", tags}
}

func CanvasItemAttrWidth(width float64) *WidgetAttr {
	return &WidgetAttr{"width", width}
}

func CanvasImageAttrAnchor(anchor Anchor) *WidgetAttr {
	return &WidgetAttr{"anchor", anchor}
}
func CanvasImageAttrImage(image *Image) *WidgetAttr {
	return &WidgetAttr{"image", image}
}
func CanvasImageAttrActiveImage(image *Image) *WidgetAttr {
	return &WidgetAttr{"activeimage", image}
}
func CanvasImageAttrDisabledImage(image *Image) *WidgetAttr {
	return &WidgetAttr{"disabledimage", image}
}

func (w* Canvas) LowerItems(tags string) error {
	return eval(fmt.Sprintf("%v lower %s", w.id, tags))
}
func (w* Canvas) LowerItemsBelow(tags, below string) error {
	return eval(fmt.Sprintf("%v lower %s %s", w.id, tags, below))
}
func (w* Canvas) RaiseItems(tags string) error {
	return eval(fmt.Sprintf("%v raise %s", w.id, tags))
}
func (w* Canvas) RaiseItemsAbove(tags, above string) error {
	return eval(fmt.Sprintf("%v raise %s %s", w.id, tags, above))
}

func (c* CanvasItem) Lower() error {
	return eval(fmt.Sprintf("%v lower %d", c.canvas.id, c.id))
}
func (c* CanvasItem) Move(dx, dy float64) error {
	return eval(fmt.Sprintf("%v move %d %f %f", c.canvas.id, c.id, dx, dy))
}
func (c* CanvasItem) MoveTo(x, y float64) error {
	return eval(fmt.Sprintf("%v moveto %d %f %f", c.canvas.id, c.id, x, y))
}
func (c* CanvasItem) Raise() error {
	return eval(fmt.Sprintf("%v raise %d", c.canvas.id, c.id))
}
func (c* CanvasItem) SetFill(color string) error {
	return eval(fmt.Sprintf("%v itemconfigure %d -fill {%v}", c.canvas.id, c.id, color))
}
func (c* CanvasItem) SetOutline(color string) error {
	return eval(fmt.Sprintf("%v itemconfigure %d -outline {%v}", c.canvas.id, c.id, color))
}
func (c* CanvasItem) Type() CanvasItemType {
	return c.typ
}
