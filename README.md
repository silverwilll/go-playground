# Go Playground - Coffee Decorator Experiment

## Overview
This is a playground project for experimenting with Go and common software design patterns. Currently, it implements a **Decorator Pattern** to model a coffee ordering system where ingredients and espresso shots can be dynamically added to a base coffee.

## Technical Stack
- **Language**: Go 1.26.1
- **Module**: `playground`

## Core Architecture
- **Interface**: `Coffee` (defined in `coffee.go`) requires `Cost() float64` and `Description() string`.
- **Base Implementation**: `SimpleCoffee` (in `coffee_impl.go`) provides the foundation.
- **Decorators**: Ingredients like `OatMilk`, `WholeMilk`, `StandardEspresso`, and `Ristretto` wrap an existing `Coffee` instance to extend its behavior.

## Development Guidelines
- **Adding Ingredients**: Create a new struct that embeds or contains a `Coffee` interface. Implement `Cost()` and `Description()` by delegating to the wrapped instance and adding the ingredient's specific values.
- **Testing**: Add new demonstration cases in `main.go` to verify the behavior of new decorators.
- **Naming**: Use descriptive names for ingredients (e.g., `Syrup`, `WhippedCream`).
- **Formatting**: Always run `go fmt ./...` before committing changes.

## Commands
- **Run**: `go run .`
- **Format**: `go fmt ./...`
