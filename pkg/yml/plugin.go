package yml

import "github.com/apache/apisix-control-plane/pkg/mem"

func (p *Plugin) ToMem() []mem.MemModel {
	result := make([]mem.MemModel, 0)
	plugins := make([]*mem.PluginSet, 0)
	for _, s := range p.Sets {
		plugin := s.ToMem()
		plugins = append(plugins, plugin)
	}
	plugin := &mem.Plugin{
		Kind:     p.Kind,
		Selector: p.Selector,
		Sets:     plugins,
	}
	result = append(result, plugin)
	return result
}

func (s *PluginSet) ToMem() *mem.PluginSet {
	return &mem.PluginSet{
		Name: s.Name,
		Conf: s.Conf,
	}
}
