package t38c

import (
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

// InwAreaSelector struct
// Intersects Nearby Within
type InwAreaSelector struct {
	client *Client
	cmd    string
	key    string
}

func newInwAreaSelector(client *Client, cmd, key string) InwAreaSelector {
	return InwAreaSelector{
		client: client,
		cmd:    cmd,
		key:    key,
	}
}

// Get any object that already exists in the database.
func (selector InwAreaSelector) Get(objectID string) InwQueryBuilder {
	area := NewCommand("GET", objectID)
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Bounds - a minimum bounding rectangle.
func (selector InwAreaSelector) Bounds(minlat, minlon, maxlat, maxlon float64) InwQueryBuilder {
	area := NewCommand("BOUNDS", floatString(minlat), floatString(minlon), floatString(maxlat), floatString(maxlon))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// FeatureCollection - GeoJSON Feature Collection object.
func (selector InwAreaSelector) FeatureCollection(fc *geojson.FeatureCollection) InwQueryBuilder {
	// TODO: handle error?
	b, _ := fc.MarshalJSON()
	area := NewCommand("OBJECT", string(b))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Feature - GeoJSON Feature object.
func (selector InwAreaSelector) Feature(ft *geojson.Feature) InwQueryBuilder {
	// TODO: handle error?
	b, _ := ft.MarshalJSON()
	area := NewCommand("OBJECT", string(b))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Geometry - GeoJSON Geometry object.
func (selector InwAreaSelector) Geometry(gm *geojson.Geometry) InwQueryBuilder {
	// TODO: handle error?
	b, _ := gm.MarshalJSON()
	area := NewCommand("OBJECT", string(b))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Circle - a circle with the specified center and radius.
func (selector InwAreaSelector) Circle(lat, lon, meters float64) InwQueryBuilder {
	area := NewCommand("CIRCLE", floatString(lat), floatString(lon), floatString(meters))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Tile - an XYZ Tile.
func (selector InwAreaSelector) Tile(x, y, z int) InwQueryBuilder {
	area := NewCommand("TILE", strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Quadkey - a QuadKey.
func (selector InwAreaSelector) Quadkey(quadkey string) InwQueryBuilder {
	area := NewCommand("QUADKEY", quadkey)
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Hash - a Geohash.
func (selector InwAreaSelector) Hash(hash string) InwQueryBuilder {
	area := NewCommand("HASH", hash)
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}
