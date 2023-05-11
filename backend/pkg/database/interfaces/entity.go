package interfaces

import "my-app/backend/pkg/snowflake/interfaces"

type IEntity interface {
	// Set Id generator for entity
	SetSnowflake(interfaces.ISnowflake)
}
