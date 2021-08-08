package mongodoc

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

type SceneWidgetDocument struct {
	ID           string
	Plugin       string
	Extension    string
	Property     string
	Enabled      bool
	WidgetLayout *WidgetLayoutDocument
}

type ScenePluginDocument struct {
	Plugin   string
	Property *string
}

type SceneDocument struct {
	ID          string
	Project     string
	Team        string
	RootLayer   string
	Widgets     []SceneWidgetDocument
	AlignSystem SceneAlignSystemDocument
	Plugins     []ScenePluginDocument
	UpdateAt    time.Time
	Property    string
}

type SceneConsumer struct {
	Rows []*scene.Scene
}

func (c *SceneConsumer) Consume(raw bson.Raw) error {
	if raw == nil {
		return nil
	}

	var doc SceneDocument
	if err := bson.Unmarshal(raw, &doc); err != nil {
		return err
	}
	scene, err := doc.Model()
	if err != nil {
		return err
	}
	c.Rows = append(c.Rows, scene)
	return nil
}

type SceneIDDocument struct {
	ID string
}

type SceneIDConsumer struct {
	Rows []id.SceneID
}

func (c *SceneIDConsumer) Consume(raw bson.Raw) error {
	if raw == nil {
		return nil
	}

	var doc SceneIDDocument
	if err := bson.Unmarshal(raw, &doc); err != nil {
		return err
	}
	scene, err := id.SceneIDFrom(doc.ID)
	if err != nil {
		return err
	}
	c.Rows = append(c.Rows, scene)
	return nil
}

func NewScene(scene *scene.Scene) (*SceneDocument, string) {
	widgets := scene.WidgetSystem().Widgets()
	was := scene.WidgetAlignSystem()
	plugins := scene.PluginSystem().Plugins()

	widgetsDoc := make([]SceneWidgetDocument, 0, len(widgets))
	pluginsDoc := make([]ScenePluginDocument, 0, len(plugins))
	alignSysDoc := NewWidgetAlignSystem(was)

	for _, w := range widgets {
		layout := WidgetLayoutDocument{
			Extendable:      w.WidgetLayout().Extendable,
			Extended:        w.WidgetLayout().Extended,
			Floating:        w.WidgetLayout().Floating,
			DefaultLocation: (*WidgetLocationDocument)(w.WidgetLayout().DefaultLocation),
		}
		widgetsDoc = append(widgetsDoc, SceneWidgetDocument{
			ID:           w.ID().String(),
			Plugin:       w.Plugin().String(),
			Extension:    string(w.Extension()),
			Property:     w.Property().String(),
			Enabled:      w.Enabled(),
			WidgetLayout: &layout,
		})
	}

	for _, sp := range plugins {
		pluginsDoc = append(pluginsDoc, ScenePluginDocument{
			Plugin:   sp.Plugin().String(),
			Property: sp.Property().StringRef(),
		})
	}

	id := scene.ID().String()
	return &SceneDocument{
		ID:          id,
		Project:     scene.Project().String(),
		Team:        scene.Team().String(),
		RootLayer:   scene.RootLayer().String(),
		Widgets:     widgetsDoc,
		AlignSystem: *alignSysDoc,
		Plugins:     pluginsDoc,
		UpdateAt:    scene.UpdatedAt(),
		Property:    scene.Property().String(),
	}, id
}

func (d *SceneDocument) Model() (*scene.Scene, error) {
	sid, err := id.SceneIDFrom(d.ID)
	if err != nil {
		return nil, err
	}
	projectID, err := id.ProjectIDFrom(d.Project)
	if err != nil {
		return nil, err
	}
	prid, err := id.PropertyIDFrom(d.Property)
	if err != nil {
		return nil, err
	}
	tid, err := id.TeamIDFrom(d.Team)
	if err != nil {
		return nil, err
	}
	lid, err := id.LayerIDFrom(d.RootLayer)
	if err != nil {
		return nil, err
	}

	ws := make([]*scene.Widget, 0, len(d.Widgets))
	ps := make([]*scene.Plugin, 0, len(d.Plugins))
	as := d.AlignSystem

	for _, w := range d.Widgets {
		wid, err := id.WidgetIDFrom(w.ID)
		if err != nil {
			return nil, err
		}
		pid, err := id.PluginIDFrom(w.Plugin)
		if err != nil {
			return nil, err
		}
		prid, err := id.PropertyIDFrom(w.Property)
		if err != nil {
			return nil, err
		}
		wl := scene.WidgetLayout{}
		if w.WidgetLayout != nil {
			wl = scene.WidgetLayout{
				Extendable:      w.WidgetLayout.Extendable,
				Extended:        w.WidgetLayout.Extended,
				Floating:        w.WidgetLayout.Floating,
				DefaultLocation: (*scene.WidgetLocation)(w.WidgetLayout.DefaultLocation),
			}
		}
		sw, err := scene.NewWidget(
			wid,
			pid,
			id.PluginExtensionID(w.Extension),
			prid,
			w.Enabled,
			&wl,
		)
		if err != nil {
			return nil, err
		}
		ws = append(ws, sw)
	}

	for _, p := range d.Plugins {
		pid, err := id.PluginIDFrom(p.Plugin)
		if err != nil {
			return nil, err
		}
		ps = append(ps, scene.NewPlugin(pid, id.PropertyIDFromRef(p.Property)))
	}

	nas := d.AlignSystem.ToModelAlignSystem(as)

	return scene.New().
		ID(sid).
		Project(projectID).
		Team(tid).
		RootLayer(lid).
		WidgetSystem(scene.NewWidgetSystem(ws)).
		WidgetAlignSystem(nas).
		PluginSystem(scene.NewPluginSystem(ps)).
		UpdatedAt(d.UpdateAt).
		Property(prid).
		Build()
}

type SceneLockConsumer struct {
	Rows []scene.LockMode
}

type SceneLockDocument struct {
	Scene string
	Lock  string
}

func (c *SceneLockConsumer) Consume(raw bson.Raw) error {
	if raw == nil {
		return nil
	}

	var doc SceneLockDocument
	if err := bson.Unmarshal(raw, &doc); err != nil {
		return err
	}
	_, sceneLock, err := doc.Model()
	if err != nil {
		return err
	}
	c.Rows = append(c.Rows, sceneLock)
	return nil
}

func NewSceneLock(sceneID id.SceneID, lock scene.LockMode) *SceneLockDocument {
	return &SceneLockDocument{
		Scene: sceneID.String(),
		Lock:  string(lock),
	}
}

func (d *SceneLockDocument) Model() (id.SceneID, scene.LockMode, error) {
	sceneID, err := id.SceneIDFrom(d.Scene)
	if err != nil {
		return sceneID, scene.LockMode(""), err
	}
	sceneLock, ok := scene.LockMode(d.Lock).Validate()
	if !ok {
		return sceneID, sceneLock, errors.New("invalid scene lock mode")
	}
	return sceneID, sceneLock, nil
}
