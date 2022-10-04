package components

import "hermannvincent/deliveryservice/ecs"

var PositionComponent = ecs.Component{Id: "position", Data: map[string]interface{}{"x": 0, "y": 0}}
