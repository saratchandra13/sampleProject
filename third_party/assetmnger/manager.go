package assetmnger

import (
	"github.com/gobuffalo/packr/v2"
)

type Manager struct {
	configBox *packr.Box
}

func (m *Manager) Get(fileName string) ([]byte, error) {
	byteValue, err := m.configBox.Find(fileName)
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}

func Initialize() *Manager {
	var m = &Manager{}
	box := packr.New("configBox", "../../asset")
	m.configBox = box
	return m
}
