package id

import (
	"regexp"
	"strings"

	"github.com/blang/semver"
)

// MUST NOT CHANGE
const officialPluginIDStr = "reearth"

// OfficialPluginID _
var OfficialPluginID = PluginID{name: officialPluginIDStr, sys: true}

// PluginID is an ID for Plugin.
type PluginID struct {
	name    string
	version string
	sys     bool
	scene   *SceneID
}

var pluginNameRe = regexp.MustCompile("^[a-zA-Z0-9._-]+$")

func validatePluginName(s string) bool {
	if len(s) == 0 || len(s) > 100 || s == "reearth" || strings.Contains(s, "/") {
		return false
	}
	return pluginNameRe.MatchString(s)
}

func NewPluginID(name string, version string, scene *SceneID) (PluginID, error) {
	if name == officialPluginIDStr {
		// official plugin
		return PluginID{name: name, sys: true}, nil
	}

	if !validatePluginName(name) {
		return PluginID{}, ErrInvalidID
	}

	if _, err := semver.Parse(version); err != nil {
		return PluginID{}, ErrInvalidID
	}

	return PluginID{
		name:    name,
		version: version,
		scene:   scene.CopyRef(),
	}, nil
}

// PluginIDFrom generates a new id.PluginID from a string.
func PluginIDFrom(id string) (PluginID, error) {
	if id == officialPluginIDStr {
		// official plugin
		return PluginID{name: id, sys: true}, nil
	}

	// scene id
	var sceneID *SceneID
	scene := strings.SplitN(id, "~", 2)
	if len(scene) == 2 {
		sceneID2, err := SceneIDFrom(scene[0])
		if err != nil {
			return PluginID{}, ErrInvalidID
		}
		sceneID = &sceneID2
		id = scene[1]
	}

	// name and version
	ids := strings.SplitN(id, "#", 2)
	if len(ids) != 2 {
		return PluginID{}, ErrInvalidID
	}

	return NewPluginID(ids[0], ids[1], sceneID)
}

// MustPluginID generates a new id.PluginID from a string, but panics if the string cannot be parsed.
func MustPluginID(id string) PluginID {
	did, err := PluginIDFrom(id)
	if err != nil {
		panic(err)
	}
	return did
}

// PluginIDFromRef generates a new id.PluginID from a string ref.
func PluginIDFromRef(id *string) *PluginID {
	if id == nil {
		return nil
	}
	did, err := PluginIDFrom(*id)
	if err != nil {
		return nil
	}
	return &did
}

// Name returns a name.
func (d PluginID) Name() string {
	return d.name
}

// Version returns a version.
func (d PluginID) Version() semver.Version {
	if d.version == "" {
		return semver.Version{}
	}
	v, err := semver.Parse(d.version)
	if err != nil {
		return semver.Version{}
	}
	return v
}

// System returns if the ID is built-in.
func (d PluginID) System() bool {
	return d.sys
}

// Scene returns a scene ID of the plugin. It indicates this plugin is private and available for only the specific scene.
func (d PluginID) Scene() *SceneID {
	return d.scene
}

// Validate returns true if id is valid.
func (d PluginID) Validate() bool {
	if d.sys {
		return true
	}
	return validatePluginName(d.name)
}

// String returns a string representation.
func (d PluginID) String() (s string) {
	if d.sys {
		return d.name
	}
	if d.scene != nil {
		s = d.scene.String() + "~"
	}
	s += d.name + "#" + d.version
	return
}

// Ref returns a reference.
func (d PluginID) Ref() *PluginID {
	d2 := d
	return &d2
}

// CopyRef _
func (d *PluginID) CopyRef() *PluginID {
	if d == nil {
		return nil
	}
	d2 := *d
	return &d2
}

// StringRef returns a reference of a string representation.
func (d *PluginID) StringRef() *string {
	if d == nil {
		return nil
	}
	id := (*d).String()
	return &id
}

// Equal returns if two IDs are quivarent.
func (d PluginID) Equal(d2 PluginID) bool {
	return d.name == d2.name && d.version == d2.version
}

// MarshalText implements encoding.TextMarshaler interface
func (d *PluginID) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *PluginID) UnmarshalText(text []byte) (err error) {
	*d, err = PluginIDFrom(string(text))
	return
}

// PluginIDToKeys converts IDs into a string slice.
func PluginIDToKeys(ids []PluginID) []string {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, id.String())
	}
	return keys
}

// PluginIDsFrom converts a string slice into a ID slice.
func PluginIDsFrom(ids []string) ([]PluginID, error) {
	dids := make([]PluginID, 0, len(ids))
	for _, id := range ids {
		did, err := PluginIDFrom(id)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}
