package builder

var builderStrategy = map[string]func() IBuilder{
	"normal": func() IBuilder { return &normalBuilder{} },
	"igloo":  func() IBuilder { return &iglooBuilder{} },
}
